package table

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
	"troy/src/dasm"
)

func (t *InstructionTable) getInstructionOutput(ins dasm.Instruction) (output string) {
	base := t.getDefaultColorTag()
	tag := t.getOpcodeTag(ins)

	if ins.Operand != nil {
		output = fmt.Sprintf("  [%05x] %s%v%s %#x\n", ins.PC, tag, ins.OpCode, base, ins.Operand)
	} else {
		output = fmt.Sprintf("  [%05x] %s%v%s\n", ins.PC, tag, ins.OpCode, base)
	}

	return t.colorize(output)
}

func (t *InstructionTable) colorize(input string) string {
	return fmt.Sprintf("%s%s", t.getDefaultColorTag(), input)
}

// Color tags can be passed as [foreground:background:flags]
// https://github.com/rivo/tview/blob/master/doc.go#L65
func (t *InstructionTable) getOpcodeTag(ins dasm.Instruction) string {
	opcode := ins.OpCode.String()

	def, ok := t.config.Colors.Instructions["default"]
	if !ok {
		log.Fatal("Failed to find default instruction color in color config")
	}

	if strings.Contains(opcode, "not defined") {
		if tags, ok := t.config.Colors.Instructions["notdefined"]; ok {
			return tags.Opcode.Prefix
		}
	}

	if tags, ok := t.config.Colors.Instructions[opcode]; ok {
		return tags.Opcode.Prefix
	}

	return def.Opcode.Prefix
}

func (t *InstructionTable) getDefaultColorTag() string {
	c := t.config.Colors.Table.Default
	return fmt.Sprintf("[%s:%s:%s]", c.Foreground, c.Background, c.Flags)
}
