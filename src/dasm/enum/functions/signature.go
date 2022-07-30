package functions

import (
	"encoding/hex"
	"encoding/json"
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
	sigs = f.offlineLookup(signatures)

	if sigs != nil {
		return sigs
	}

	return f.onlineLookup()
}

func (f *Signature) onlineLookup() (sigs []string) {
	response, err := http.Get("https://www.4byte.directory/api/v1/signatures/?hex_signature=" + f.String())

	// todo: add logging here if err != nil, perhaps with a verbose flag
	if err != nil {
		return nil
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil
	}

	var res FourByteResponse
	if err := json.Unmarshal(contents, &res); err != nil {
		return nil
	}

	if res.Count <= 0 {
		return nil
	}

	if err == nil {
		for _, s := range res.Results {
			sigs = append(sigs, s.TextSignature)
		}
	}

	return sigs
}

func (f *Signature) offlineLookup(signatures *map[string][]string) (sigs []string) {
	list := *signatures

	if sig, found := list[f.String()]; found {
		return sig
	}

	return nil
}
