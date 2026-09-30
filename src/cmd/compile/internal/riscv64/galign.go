package riscv64

import (
	"cmd/compile/internal/ssagen"
	"cmd/internal/obj/riscv"
)

func Init(arch *ssagen.ArchInfo) {
	arch.LinkArch = &riscv.LinkRISCV64

	arch.REGSP = riscv.REG_SP
	arch.MAXWIDTH = 1 << 50

	arch.Ginsnop = ginsnop
	arch.ZeroRange = zeroRange

	arch.SSAMarkMoves = ssaMarkMoves
	arch.SSAGenValue = ssaGenValue
	arch.SSAGenBlock = ssaGenBlock
	arch.LoadRegResult = loadRegResult
	arch.SpillArgReg = spillArgReg
}
