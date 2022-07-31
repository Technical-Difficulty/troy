package instructions

import (
	log "github.com/sirupsen/logrus"
	"troy/src/dasm"
	"troy/src/util"
)

type InstructionEnum struct {
	instructions dasm.InstructionSet
}

func NewInstructionEnum(instructions dasm.InstructionSet) InstructionEnum {
	return InstructionEnum{
		instructions: instructions,
	}
}

// todo: Add support for "JUMP" instructions, this seems to be based on stack values so may required dynamic analysis
func (e *InstructionEnum) EnumerateJumps() map[uint64]uint64 {
	var jumps = make(map[uint64]uint64)

	for idx, ins := range e.instructions.Array() {
		opcode := ins.OpCode.String()

		// Skip jumpdest instructions
		if opcode != "JUMPI" {
			continue
		}

		push, err := e.instructions.ReverseFindInstructionWithPrefix(idx, "PUSH")
		if err != nil {
			log.WithField("Error", err).Fatal("Couldn't find jump destination (preceding PUSH)")
		}

		dest, err := util.BytesToUInt64(push.Operand)
		if err != nil {
			log.WithField("Error", err).Fatal("Possibly found jump destination but couldn't decode it")
		}

		jumps[ins.PC] = dest
	}

	return jumps
}
