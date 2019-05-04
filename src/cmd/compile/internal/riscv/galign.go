// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package riscv

import (
	"cmd/compile/internal/gc"
	"cmd/internal/obj/riscv"
)

func Init(arch *gc.Arch) {
	arch.LinkArch = &riscv.LinkRISCV

	arch.REGSP = riscv.REG_SP
	// TODO(prattmic): all the other arches use 50 bits, even though
	// they have 48-bit vaddrs. why?
	arch.MAXWIDTH = 1 << 50

	arch.ZeroAuto = zeroAuto
	arch.ZeroRange = zerorange
	arch.Ginsnop = ginsnop

	// TODO(prattmic): other fields?

	arch.SSAMarkMoves = ssaMarkMoves
	arch.SSAGenValue = ssaGenValue
	arch.SSAGenBlock = ssaGenBlock
}