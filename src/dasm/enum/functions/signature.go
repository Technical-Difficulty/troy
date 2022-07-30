package functions

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"troy/src/dasm"
)

type (
	Signature struct {
		ins dasm.Instruction
	}

	FourByteResponse struct {
		Count   int `json:"count"`
		Results []struct {
			TextSignature string `json:"text_signature"`
		} `json:"results"`
	}
)

func NewSignature(ins dasm.Instruction) Signature {
	return Signature{
		ins: ins,
	}
}

func (f *Signature) String() string {
	return "0x" + hex.EncodeToString(f.ins.Operand)
}

// todo: multi-thread this for online lookups
func (f *Signature) Lookup(signatures *map[string][]string) (sigs []string) {
	res, err := f.onlineLookup()

	// todo: add logging here if err != nil, perhaps with a verbose flag
	if err == nil {
		for _, s := range res.Results {
			sigs = append(sigs, s.TextSignature)
		}

		return sigs
	}

	return f.offlineLookup(signatures)
}

func (f *Signature) onlineLookup() (FourByteResponse, error) {
	response, err := http.Get("https://www.4byte.directory/api/v1/signatures/?hex_signature=" + f.String())

	if err != nil {
		return FourByteResponse{}, err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)

		if err != nil {
			return FourByteResponse{}, err
		}

		var res FourByteResponse
		if err := json.Unmarshal(contents, &res); err != nil {
			return FourByteResponse{}, err
		}

		if res.Count <= 0 {
			return FourByteResponse{}, errors.New("Not found")
		}

		return res, nil
	}
}

func (f *Signature) offlineLookup(signatures *map[string][]string) (sigs []string) {
	list := *signatures

	if sig, found := list[f.String()]; found {
		return sig
	}

	return nil
}
