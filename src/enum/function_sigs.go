package enum

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"troy/src/dasm"
)

type FuncSig struct {
	Operand []byte
}

type FourByte struct {
	Count   int `json:"count"`
	Results []struct {
		TextSignature string `json:"text_signature"`
	} `json:"results"`
}

func (f *FuncSig) String() string {
	return "0x" + hex.EncodeToString(f.Operand)
}

func (f *FuncSig) Lookup() (FourByte, error) {
	response, err := http.Get("https://www.4byte.directory/api/v1/signatures/?hex_signature=" + f.String())

	if err != nil {
		return FourByte{}, err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)

		if err != nil {
			return FourByte{}, err
		}

		var res FourByte
		if err := json.Unmarshal(contents, &res); err != nil {
			return FourByte{}, err
		}

		if res.Count <= 0 {
			return FourByte{}, errors.New("Not found")
		}

		return res, nil

	}
}

// Function sigantures live between:
// PUSH1 0x04
// CALLDATASIZE
// ...
// CALLDATALOAD
// ...
// PUSH4 <FUNC_SIG>
// ...
// JUMPDEST
func FuncSigs(instructions []dasm.Instruction) (out []FuncSig) {
	var scan bool

	for _, ins := range instructions {
		if ins.OpCode.String() == "CALLDATALOAD" {
			scan = true
		}

		if scan && ins.OpCode.String() == "PUSH4" {
			out = append(out, FuncSig{
				Operand: ins.Operand,
			})
		}

		if scan && ins.OpCode.String() == "JUMPDEST" {
			break
		}
	}
	return out
}
