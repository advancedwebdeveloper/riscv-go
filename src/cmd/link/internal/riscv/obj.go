//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package riscv

import (
	"cmd/internal/objabi"
	"cmd/internal/sys"
	"cmd/link/internal/ld"
	"fmt"
	"log"
)

func Init() (*sys.Arch, ld.Arch) {
	arch := sys.ArchRISCV

	theArch := ld.Arch{
		Funcalign: FuncAlign, /* XXX */
		Maxalign: MaxAlign,
		Minalign: MinAlign,
		Dwarfregsp: DWARFREGSP,
		Dwarfreglr: DWARFREGLR,

		Adddynrel: adddynrel,
		Archinit: archinit,
		Archreloc: archreloc,
		Archrelocvariant: archrelocvariant,
		Asmb: asmb,
		Elfreloc1: elfreloc1,
		Elfsetupplt: elfsetupplt,
		Gentext: gentext,
		Machoreloc1: machoreloc1,

		Linuxdynld: "/lib/ld.so.1",

		// TODO: FreeBSD and NetBSD have RISCV ports, but we don't support,
		// them yet.,
		Freebsddynld: "XXX",
		Netbsddynld: "XXX",

		Openbsddynld: "XXX",
		Dragonflydynld: "XXX",
		Solarisdynld: "XXX",
	}

	return arch, theArch
}

func archinit(ctxt *ld.Link) {
	// getgoextlinkenabled is based on GO_EXTLINK_ENABLED when
	// Go was built; see ../../make.bash.
	if ctxt.LinkMode == ld.LinkAuto && objabi.Getgoextlinkenabled() == "0" {
		ctxt.LinkMode = ld.LinkInternal
	}

	switch ctxt.HeadType {
	default:
		if ctxt.LinkMode == ld.LinkAuto {
			ctxt.LinkMode = ld.LinkInternal
		}
		if ctxt.LinkMode == ld.LinkExternal && objabi.Getgoextlinkenabled() != "1" {
			log.Fatalf("cannot use -linkmode=external with -H %v", ctxt.HeadType)
		}
	}

	switch ctxt.HeadType {
	default:
		ld.Exitf("unknown -H option: %v", ctxt.HeadType)

	case objabi.Hlinux: /* riscv elf */
		ld.Elfinit(ctxt)
		ld.HEADR = ld.ELFRESERVE
		if *ld.FlagTextAddr == -1 {
			*ld.FlagTextAddr = 0x10000 + int64(ld.HEADR)
		}
		if *ld.FlagDataAddr == -1 {
			*ld.FlagDataAddr = 0
		}
		if *ld.FlagRound == -1 {
			*ld.FlagRound = 0x10000
		}
	}

	if *ld.FlagDataAddr != 0 && *ld.FlagRound != 0 {
		fmt.Printf("warning: -D0x%x is ignored because of -R0x%x\n", uint64(*ld.FlagDataAddr), uint32(*ld.FlagRound))
	}
}
