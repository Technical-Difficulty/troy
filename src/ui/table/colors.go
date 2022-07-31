package table

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	log "github.com/sirupsen/logrus"
	"strings"
	"troy/src/dasm"
)

func (t *ContractTable) getInstructionOutput(ins dasm.Instruction) (output string) {
	base := t.getDefaultColorTag()
	tag := t.getOpcodeTag(ins)

	if ins.Operand != nil {
		output = fmt.Sprintf("  %s %s%v%s %#x\n",
			t.getProgramCounterOutput(ins.PC), tag, ins.OpCode, base, ins.Operand,
		)
	} else {
		output = fmt.Sprintf("  %s %s%v%s\n",
			t.getProgramCounterOutput(ins.PC), tag, ins.OpCode, base,
		)
	}

	return t.colorize(output)
}

func (t *ContractTable) getProgramCounterOutput(pc uint64) string {
	return fmt.Sprintf("%s[%05x]%s",
		t.config.Colors.Table.ProgramCounter, pc, t.getDefaultColorTag(),
	)
}

func (t *ContractTable) colorize(input string) string {
	return fmt.Sprintf("%s%s", t.getDefaultColorTag(), input)
}

// Color tags can be passed as [foreground:background:flags]
// https://github.com/rivo/tview/blob/master/doc.go#L65
func (t *ContractTable) getOpcodeTag(ins dasm.Instruction) string {
	opcode := ins.OpCode.String()

	def, ok := t.config.Colors.Instructions["default"]
	if !ok {
		log.Fatal("Failed to find default instruction color in color config")
	}

	if strings.Contains(opcode, "not defined") {
		if tags, ok := t.config.Colors.Instructions["notdefined"]; ok {
			return tags.Opcode
		}
	}

	if tags, ok := t.config.Colors.Instructions[opcode]; ok {
		return tags.Opcode
	}

	return def.Opcode
}

func (t *ContractTable) getDefaultColorTag() string {
	c := t.config.Colors.Table.Default
	return fmt.Sprintf("[%s:%s:%s]", c.Foreground, c.Background, c.Flags)
}

func (t *ContractTable) getDefaultSelectedStyle() tcell.Style {
	colors := t.config.Colors.Table.Selected

	return tcell.StyleDefault.
		Foreground(t.getColor(colors.Foreground)).
		Background(t.getColor(colors.Background))
}

func (t *ContractTable) getColor(key string) (c tcell.Color) {
	if value, ok := tcell.ColorNames[key]; ok {
		return value
	}

	log.WithField("Key", key).Fatal("Failed to find specified color")
	return
}
