package table

import (
	"fmt"
	"troy/src/dasm"
)

const (
	ForegroundDefault = "white"
	ForegroundJump    = "yellow"

	BackgroundDefault = "black"
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
	d := fmt.Sprintf("[%s:%s:-]", ForegroundDefault, BackgroundDefault)

	switch ins.OpCode.String() {
	case "JUMPI":
		fallthrough
	case "JUMPDEST":
		return fmt.Sprintf("[%s:%s:b]", ForegroundJump, BackgroundDefault), d
	}

	return d, d
}
