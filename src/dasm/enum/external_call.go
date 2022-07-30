package enum

import (
	"troy/src/dasm"
)

type ExtCall struct {
	FuncSig FuncSig
	ins     dasm.Instruction
}

// Initial naive assumption that the closest PUSH4
// to the CALL is the function signature who knows if this
// is true in all cases... probably isn't.
// TODO: parse CALL target and arguments
func ExtCalls(instructions []dasm.Instruction) (out []ExtCall) {
	var lastPush4Ins dasm.Instruction

	for _, ins := range instructions {
		if ins.OpCode.String() == "PUSH4" {
			lastPush4Ins = ins
		}

		if ins.OpCode.String() == "CALL" && lastPush4Ins.PC != 0 {
			out = append(out, ExtCall{
				FuncSig: FuncSig{ins: lastPush4Ins},
				ins:     ins,
			})
			lastPush4Ins = dasm.Instruction{}
		}
	}
	return out
}
