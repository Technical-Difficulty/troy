package enum

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/vm"
	log "github.com/sirupsen/logrus"
)

var notableOps = map[byte]string{
	0xf0: "CREATE",
	0xf1: "CALL",
	0xf4: "DELEGATECALL",
	0xf5: "CREATE2",
	0xff: "SELFDESTRUCT",
	0x63: "PUSH4",
}

func EvaluateInstr(pc uint64, opcode vm.OpCode, arg []byte) {
	if _, ok := notableOps[byte(opcode)]; ok {
		// TODO: parse op codes to highlight anything notable and then pipe it
		// into some nice terminal UI. Should be able to do more advanced lookups
		// too like if the opcode is CALL look for future SSTORE. We can also
		// look at the Arg of CALL to highlight interesting function sigs like
		// approve, mint, transfer etc.
		var label string

		if arg != nil && 0 < len(arg) {
			label = fmt.Sprintf("%05x: %v %#x\n", pc, opcode, arg)
		} else {
			label = fmt.Sprintf("%05x: %v\n", pc, opcode)
		}

		log.Info(fmt.Sprintf("Noteable opcode found: %v", label))
	}
}
