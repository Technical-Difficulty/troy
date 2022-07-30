package table

import (
	"fmt"
	"log"
	"troy/src/dasm"
)

func (t *InstructionTable) getInstructionOutput(ins dasm.Instruction) (output string) {
	prefix, suffix := t.getColorValues(ins)

	if ins.Operand != nil {
		output = fmt.Sprintf("  [%05x] %s%v%s %#x\n", ins.PC, prefix, ins.OpCode, suffix, ins.Operand)
	} else {
		output = fmt.Sprintf("  [%05x] %s%v%s\n", ins.PC, prefix, ins.OpCode, suffix)
	}

	return
}

// Color tags can be passed as [foreground:background:flags]
// https://github.com/rivo/tview/blob/master/doc.go#L65
func (t *InstructionTable) getColorValues(ins dasm.Instruction) (prefix string, suffix string) {
	def, ok := t.config.Colors.Instructions["default"]
	if !ok {
		log.Fatal("Failed to find default instruction colors in color config")
	}

	if tags, ok := t.config.Colors.Instructions[ins.OpCode.String()]; ok {
		return tags.Opcode.Prefix, tags.Opcode.Suffix
	}

	return def.Opcode.Prefix, def.Opcode.Suffix
}
