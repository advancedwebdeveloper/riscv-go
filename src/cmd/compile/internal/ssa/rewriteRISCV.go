// Code generated from gen/RISCV.rules; DO NOT EDIT.
// generated with: cd gen; go run *.go

package ssa

import "fmt"
import "math"
import "cmd/internal/obj"
import "cmd/internal/objabi"
import "cmd/compile/internal/types"

var _ = fmt.Println   // in case not otherwise used
var _ = math.MinInt8  // in case not otherwise used
var _ = obj.ANOP      // in case not otherwise used
var _ = objabi.GOROOT // in case not otherwise used
var _ = types.TypeMem // in case not otherwise used

func rewriteValueRISCV(v *Value) bool {
	switch v.Op {
	case OpAdd16:
		return rewriteValueRISCV_OpAdd16_0(v)
	case OpAdd32:
		return rewriteValueRISCV_OpAdd32_0(v)
	case OpAdd32F:
		return rewriteValueRISCV_OpAdd32F_0(v)
	case OpAdd64:
		return rewriteValueRISCV_OpAdd64_0(v)
	case OpAdd64F:
		return rewriteValueRISCV_OpAdd64F_0(v)
	case OpAdd8:
		return rewriteValueRISCV_OpAdd8_0(v)
	case OpAddPtr:
		return rewriteValueRISCV_OpAddPtr_0(v)
	case OpAddr:
		return rewriteValueRISCV_OpAddr_0(v)
	case OpAnd16:
		return rewriteValueRISCV_OpAnd16_0(v)
	case OpAnd32:
		return rewriteValueRISCV_OpAnd32_0(v)
	case OpAnd64:
		return rewriteValueRISCV_OpAnd64_0(v)
	case OpAnd8:
		return rewriteValueRISCV_OpAnd8_0(v)
	case OpAndB:
		return rewriteValueRISCV_OpAndB_0(v)
	case OpAvg64u:
		return rewriteValueRISCV_OpAvg64u_0(v)
	case OpClosureCall:
		return rewriteValueRISCV_OpClosureCall_0(v)
	case OpCom16:
		return rewriteValueRISCV_OpCom16_0(v)
	case OpCom32:
		return rewriteValueRISCV_OpCom32_0(v)
	case OpCom64:
		return rewriteValueRISCV_OpCom64_0(v)
	case OpCom8:
		return rewriteValueRISCV_OpCom8_0(v)
	case OpConst16:
		return rewriteValueRISCV_OpConst16_0(v)
	case OpConst32:
		return rewriteValueRISCV_OpConst32_0(v)
	case OpConst32F:
		return rewriteValueRISCV_OpConst32F_0(v)
	case OpConst64:
		return rewriteValueRISCV_OpConst64_0(v)
	case OpConst64F:
		return rewriteValueRISCV_OpConst64F_0(v)
	case OpConst8:
		return rewriteValueRISCV_OpConst8_0(v)
	case OpConstBool:
		return rewriteValueRISCV_OpConstBool_0(v)
	case OpConstNil:
		return rewriteValueRISCV_OpConstNil_0(v)
	case OpConvert:
		return rewriteValueRISCV_OpConvert_0(v)
	case OpCvt32Fto32:
		return rewriteValueRISCV_OpCvt32Fto32_0(v)
	case OpCvt32Fto64:
		return rewriteValueRISCV_OpCvt32Fto64_0(v)
	case OpCvt32Fto64F:
		return rewriteValueRISCV_OpCvt32Fto64F_0(v)
	case OpCvt32to32F:
		return rewriteValueRISCV_OpCvt32to32F_0(v)
	case OpCvt32to64F:
		return rewriteValueRISCV_OpCvt32to64F_0(v)
	case OpCvt64Fto32:
		return rewriteValueRISCV_OpCvt64Fto32_0(v)
	case OpCvt64Fto32F:
		return rewriteValueRISCV_OpCvt64Fto32F_0(v)
	case OpCvt64Fto64:
		return rewriteValueRISCV_OpCvt64Fto64_0(v)
	case OpCvt64to32F:
		return rewriteValueRISCV_OpCvt64to32F_0(v)
	case OpCvt64to64F:
		return rewriteValueRISCV_OpCvt64to64F_0(v)
	case OpDiv16:
		return rewriteValueRISCV_OpDiv16_0(v)
	case OpDiv16u:
		return rewriteValueRISCV_OpDiv16u_0(v)
	case OpDiv32:
		return rewriteValueRISCV_OpDiv32_0(v)
	case OpDiv32F:
		return rewriteValueRISCV_OpDiv32F_0(v)
	case OpDiv32u:
		return rewriteValueRISCV_OpDiv32u_0(v)
	case OpDiv64:
		return rewriteValueRISCV_OpDiv64_0(v)
	case OpDiv64F:
		return rewriteValueRISCV_OpDiv64F_0(v)
	case OpDiv64u:
		return rewriteValueRISCV_OpDiv64u_0(v)
	case OpDiv8:
		return rewriteValueRISCV_OpDiv8_0(v)
	case OpDiv8u:
		return rewriteValueRISCV_OpDiv8u_0(v)
	case OpEq16:
		return rewriteValueRISCV_OpEq16_0(v)
	case OpEq32:
		return rewriteValueRISCV_OpEq32_0(v)
	case OpEq32F:
		return rewriteValueRISCV_OpEq32F_0(v)
	case OpEq64:
		return rewriteValueRISCV_OpEq64_0(v)
	case OpEq64F:
		return rewriteValueRISCV_OpEq64F_0(v)
	case OpEq8:
		return rewriteValueRISCV_OpEq8_0(v)
	case OpEqB:
		return rewriteValueRISCV_OpEqB_0(v)
	case OpEqPtr:
		return rewriteValueRISCV_OpEqPtr_0(v)
	case OpGeq16:
		return rewriteValueRISCV_OpGeq16_0(v)
	case OpGeq16U:
		return rewriteValueRISCV_OpGeq16U_0(v)
	case OpGeq32:
		return rewriteValueRISCV_OpGeq32_0(v)
	case OpGeq32F:
		return rewriteValueRISCV_OpGeq32F_0(v)
	case OpGeq32U:
		return rewriteValueRISCV_OpGeq32U_0(v)
	case OpGeq64:
		return rewriteValueRISCV_OpGeq64_0(v)
	case OpGeq64F:
		return rewriteValueRISCV_OpGeq64F_0(v)
	case OpGeq64U:
		return rewriteValueRISCV_OpGeq64U_0(v)
	case OpGeq8:
		return rewriteValueRISCV_OpGeq8_0(v)
	case OpGeq8U:
		return rewriteValueRISCV_OpGeq8U_0(v)
	case OpGetCallerPC:
		return rewriteValueRISCV_OpGetCallerPC_0(v)
	case OpGetCallerSP:
		return rewriteValueRISCV_OpGetCallerSP_0(v)
	case OpGetClosurePtr:
		return rewriteValueRISCV_OpGetClosurePtr_0(v)
	case OpGreater16:
		return rewriteValueRISCV_OpGreater16_0(v)
	case OpGreater16U:
		return rewriteValueRISCV_OpGreater16U_0(v)
	case OpGreater32:
		return rewriteValueRISCV_OpGreater32_0(v)
	case OpGreater32F:
		return rewriteValueRISCV_OpGreater32F_0(v)
	case OpGreater32U:
		return rewriteValueRISCV_OpGreater32U_0(v)
	case OpGreater64:
		return rewriteValueRISCV_OpGreater64_0(v)
	case OpGreater64F:
		return rewriteValueRISCV_OpGreater64F_0(v)
	case OpGreater64U:
		return rewriteValueRISCV_OpGreater64U_0(v)
	case OpGreater8:
		return rewriteValueRISCV_OpGreater8_0(v)
	case OpGreater8U:
		return rewriteValueRISCV_OpGreater8U_0(v)
	case OpHmul32:
		return rewriteValueRISCV_OpHmul32_0(v)
	case OpHmul32u:
		return rewriteValueRISCV_OpHmul32u_0(v)
	case OpHmul64:
		return rewriteValueRISCV_OpHmul64_0(v)
	case OpHmul64u:
		return rewriteValueRISCV_OpHmul64u_0(v)
	case OpInterCall:
		return rewriteValueRISCV_OpInterCall_0(v)
	case OpIsInBounds:
		return rewriteValueRISCV_OpIsInBounds_0(v)
	case OpIsNonNil:
		return rewriteValueRISCV_OpIsNonNil_0(v)
	case OpIsSliceInBounds:
		return rewriteValueRISCV_OpIsSliceInBounds_0(v)
	case OpLeq16:
		return rewriteValueRISCV_OpLeq16_0(v)
	case OpLeq16U:
		return rewriteValueRISCV_OpLeq16U_0(v)
	case OpLeq32:
		return rewriteValueRISCV_OpLeq32_0(v)
	case OpLeq32F:
		return rewriteValueRISCV_OpLeq32F_0(v)
	case OpLeq32U:
		return rewriteValueRISCV_OpLeq32U_0(v)
	case OpLeq64:
		return rewriteValueRISCV_OpLeq64_0(v)
	case OpLeq64F:
		return rewriteValueRISCV_OpLeq64F_0(v)
	case OpLeq64U:
		return rewriteValueRISCV_OpLeq64U_0(v)
	case OpLeq8:
		return rewriteValueRISCV_OpLeq8_0(v)
	case OpLeq8U:
		return rewriteValueRISCV_OpLeq8U_0(v)
	case OpLess16:
		return rewriteValueRISCV_OpLess16_0(v)
	case OpLess16U:
		return rewriteValueRISCV_OpLess16U_0(v)
	case OpLess32:
		return rewriteValueRISCV_OpLess32_0(v)
	case OpLess32F:
		return rewriteValueRISCV_OpLess32F_0(v)
	case OpLess32U:
		return rewriteValueRISCV_OpLess32U_0(v)
	case OpLess64:
		return rewriteValueRISCV_OpLess64_0(v)
	case OpLess64F:
		return rewriteValueRISCV_OpLess64F_0(v)
	case OpLess64U:
		return rewriteValueRISCV_OpLess64U_0(v)
	case OpLess8:
		return rewriteValueRISCV_OpLess8_0(v)
	case OpLess8U:
		return rewriteValueRISCV_OpLess8U_0(v)
	case OpLoad:
		return rewriteValueRISCV_OpLoad_0(v)
	case OpLocalAddr:
		return rewriteValueRISCV_OpLocalAddr_0(v)
	case OpLsh16x16:
		return rewriteValueRISCV_OpLsh16x16_0(v)
	case OpLsh16x32:
		return rewriteValueRISCV_OpLsh16x32_0(v)
	case OpLsh16x64:
		return rewriteValueRISCV_OpLsh16x64_0(v)
	case OpLsh16x8:
		return rewriteValueRISCV_OpLsh16x8_0(v)
	case OpLsh32x16:
		return rewriteValueRISCV_OpLsh32x16_0(v)
	case OpLsh32x32:
		return rewriteValueRISCV_OpLsh32x32_0(v)
	case OpLsh32x64:
		return rewriteValueRISCV_OpLsh32x64_0(v)
	case OpLsh32x8:
		return rewriteValueRISCV_OpLsh32x8_0(v)
	case OpLsh64x16:
		return rewriteValueRISCV_OpLsh64x16_0(v)
	case OpLsh64x32:
		return rewriteValueRISCV_OpLsh64x32_0(v)
	case OpLsh64x64:
		return rewriteValueRISCV_OpLsh64x64_0(v)
	case OpLsh64x8:
		return rewriteValueRISCV_OpLsh64x8_0(v)
	case OpLsh8x16:
		return rewriteValueRISCV_OpLsh8x16_0(v)
	case OpLsh8x32:
		return rewriteValueRISCV_OpLsh8x32_0(v)
	case OpLsh8x64:
		return rewriteValueRISCV_OpLsh8x64_0(v)
	case OpLsh8x8:
		return rewriteValueRISCV_OpLsh8x8_0(v)
	case OpMod16:
		return rewriteValueRISCV_OpMod16_0(v)
	case OpMod16u:
		return rewriteValueRISCV_OpMod16u_0(v)
	case OpMod32:
		return rewriteValueRISCV_OpMod32_0(v)
	case OpMod32u:
		return rewriteValueRISCV_OpMod32u_0(v)
	case OpMod64:
		return rewriteValueRISCV_OpMod64_0(v)
	case OpMod64u:
		return rewriteValueRISCV_OpMod64u_0(v)
	case OpMod8:
		return rewriteValueRISCV_OpMod8_0(v)
	case OpMod8u:
		return rewriteValueRISCV_OpMod8u_0(v)
	case OpMove:
		return rewriteValueRISCV_OpMove_0(v)
	case OpMul16:
		return rewriteValueRISCV_OpMul16_0(v)
	case OpMul32:
		return rewriteValueRISCV_OpMul32_0(v)
	case OpMul32F:
		return rewriteValueRISCV_OpMul32F_0(v)
	case OpMul64:
		return rewriteValueRISCV_OpMul64_0(v)
	case OpMul64F:
		return rewriteValueRISCV_OpMul64F_0(v)
	case OpMul8:
		return rewriteValueRISCV_OpMul8_0(v)
	case OpNeg16:
		return rewriteValueRISCV_OpNeg16_0(v)
	case OpNeg32:
		return rewriteValueRISCV_OpNeg32_0(v)
	case OpNeg32F:
		return rewriteValueRISCV_OpNeg32F_0(v)
	case OpNeg64:
		return rewriteValueRISCV_OpNeg64_0(v)
	case OpNeg64F:
		return rewriteValueRISCV_OpNeg64F_0(v)
	case OpNeg8:
		return rewriteValueRISCV_OpNeg8_0(v)
	case OpNeq16:
		return rewriteValueRISCV_OpNeq16_0(v)
	case OpNeq32:
		return rewriteValueRISCV_OpNeq32_0(v)
	case OpNeq32F:
		return rewriteValueRISCV_OpNeq32F_0(v)
	case OpNeq64:
		return rewriteValueRISCV_OpNeq64_0(v)
	case OpNeq64F:
		return rewriteValueRISCV_OpNeq64F_0(v)
	case OpNeq8:
		return rewriteValueRISCV_OpNeq8_0(v)
	case OpNeqB:
		return rewriteValueRISCV_OpNeqB_0(v)
	case OpNeqPtr:
		return rewriteValueRISCV_OpNeqPtr_0(v)
	case OpNilCheck:
		return rewriteValueRISCV_OpNilCheck_0(v)
	case OpNot:
		return rewriteValueRISCV_OpNot_0(v)
	case OpOffPtr:
		return rewriteValueRISCV_OpOffPtr_0(v)
	case OpOr16:
		return rewriteValueRISCV_OpOr16_0(v)
	case OpOr32:
		return rewriteValueRISCV_OpOr32_0(v)
	case OpOr64:
		return rewriteValueRISCV_OpOr64_0(v)
	case OpOr8:
		return rewriteValueRISCV_OpOr8_0(v)
	case OpOrB:
		return rewriteValueRISCV_OpOrB_0(v)
	case OpRISCVADDI:
		return rewriteValueRISCV_OpRISCVADDI_0(v)
	case OpRISCVMOVBUload:
		return rewriteValueRISCV_OpRISCVMOVBUload_0(v)
	case OpRISCVMOVBload:
		return rewriteValueRISCV_OpRISCVMOVBload_0(v)
	case OpRISCVMOVBstore:
		return rewriteValueRISCV_OpRISCVMOVBstore_0(v)
	case OpRISCVMOVDconst:
		return rewriteValueRISCV_OpRISCVMOVDconst_0(v)
	case OpRISCVMOVDload:
		return rewriteValueRISCV_OpRISCVMOVDload_0(v)
	case OpRISCVMOVDstore:
		return rewriteValueRISCV_OpRISCVMOVDstore_0(v)
	case OpRISCVMOVHUload:
		return rewriteValueRISCV_OpRISCVMOVHUload_0(v)
	case OpRISCVMOVHload:
		return rewriteValueRISCV_OpRISCVMOVHload_0(v)
	case OpRISCVMOVHstore:
		return rewriteValueRISCV_OpRISCVMOVHstore_0(v)
	case OpRISCVMOVWUload:
		return rewriteValueRISCV_OpRISCVMOVWUload_0(v)
	case OpRISCVMOVWload:
		return rewriteValueRISCV_OpRISCVMOVWload_0(v)
	case OpRISCVMOVWstore:
		return rewriteValueRISCV_OpRISCVMOVWstore_0(v)
	case OpRound32F:
		return rewriteValueRISCV_OpRound32F_0(v)
	case OpRound64F:
		return rewriteValueRISCV_OpRound64F_0(v)
	case OpRsh16Ux16:
		return rewriteValueRISCV_OpRsh16Ux16_0(v)
	case OpRsh16Ux32:
		return rewriteValueRISCV_OpRsh16Ux32_0(v)
	case OpRsh16Ux64:
		return rewriteValueRISCV_OpRsh16Ux64_0(v)
	case OpRsh16Ux8:
		return rewriteValueRISCV_OpRsh16Ux8_0(v)
	case OpRsh16x16:
		return rewriteValueRISCV_OpRsh16x16_0(v)
	case OpRsh16x32:
		return rewriteValueRISCV_OpRsh16x32_0(v)
	case OpRsh16x64:
		return rewriteValueRISCV_OpRsh16x64_0(v)
	case OpRsh16x8:
		return rewriteValueRISCV_OpRsh16x8_0(v)
	case OpRsh32Ux16:
		return rewriteValueRISCV_OpRsh32Ux16_0(v)
	case OpRsh32Ux32:
		return rewriteValueRISCV_OpRsh32Ux32_0(v)
	case OpRsh32Ux64:
		return rewriteValueRISCV_OpRsh32Ux64_0(v)
	case OpRsh32Ux8:
		return rewriteValueRISCV_OpRsh32Ux8_0(v)
	case OpRsh32x16:
		return rewriteValueRISCV_OpRsh32x16_0(v)
	case OpRsh32x32:
		return rewriteValueRISCV_OpRsh32x32_0(v)
	case OpRsh32x64:
		return rewriteValueRISCV_OpRsh32x64_0(v)
	case OpRsh32x8:
		return rewriteValueRISCV_OpRsh32x8_0(v)
	case OpRsh64Ux16:
		return rewriteValueRISCV_OpRsh64Ux16_0(v)
	case OpRsh64Ux32:
		return rewriteValueRISCV_OpRsh64Ux32_0(v)
	case OpRsh64Ux64:
		return rewriteValueRISCV_OpRsh64Ux64_0(v)
	case OpRsh64Ux8:
		return rewriteValueRISCV_OpRsh64Ux8_0(v)
	case OpRsh64x16:
		return rewriteValueRISCV_OpRsh64x16_0(v)
	case OpRsh64x32:
		return rewriteValueRISCV_OpRsh64x32_0(v)
	case OpRsh64x64:
		return rewriteValueRISCV_OpRsh64x64_0(v)
	case OpRsh64x8:
		return rewriteValueRISCV_OpRsh64x8_0(v)
	case OpRsh8Ux16:
		return rewriteValueRISCV_OpRsh8Ux16_0(v)
	case OpRsh8Ux32:
		return rewriteValueRISCV_OpRsh8Ux32_0(v)
	case OpRsh8Ux64:
		return rewriteValueRISCV_OpRsh8Ux64_0(v)
	case OpRsh8Ux8:
		return rewriteValueRISCV_OpRsh8Ux8_0(v)
	case OpRsh8x16:
		return rewriteValueRISCV_OpRsh8x16_0(v)
	case OpRsh8x32:
		return rewriteValueRISCV_OpRsh8x32_0(v)
	case OpRsh8x64:
		return rewriteValueRISCV_OpRsh8x64_0(v)
	case OpRsh8x8:
		return rewriteValueRISCV_OpRsh8x8_0(v)
	case OpSignExt16to32:
		return rewriteValueRISCV_OpSignExt16to32_0(v)
	case OpSignExt16to64:
		return rewriteValueRISCV_OpSignExt16to64_0(v)
	case OpSignExt32to64:
		return rewriteValueRISCV_OpSignExt32to64_0(v)
	case OpSignExt8to16:
		return rewriteValueRISCV_OpSignExt8to16_0(v)
	case OpSignExt8to32:
		return rewriteValueRISCV_OpSignExt8to32_0(v)
	case OpSignExt8to64:
		return rewriteValueRISCV_OpSignExt8to64_0(v)
	case OpSlicemask:
		return rewriteValueRISCV_OpSlicemask_0(v)
	case OpSqrt:
		return rewriteValueRISCV_OpSqrt_0(v)
	case OpStaticCall:
		return rewriteValueRISCV_OpStaticCall_0(v)
	case OpStore:
		return rewriteValueRISCV_OpStore_0(v)
	case OpSub16:
		return rewriteValueRISCV_OpSub16_0(v)
	case OpSub32:
		return rewriteValueRISCV_OpSub32_0(v)
	case OpSub32F:
		return rewriteValueRISCV_OpSub32F_0(v)
	case OpSub64:
		return rewriteValueRISCV_OpSub64_0(v)
	case OpSub64F:
		return rewriteValueRISCV_OpSub64F_0(v)
	case OpSub8:
		return rewriteValueRISCV_OpSub8_0(v)
	case OpSubPtr:
		return rewriteValueRISCV_OpSubPtr_0(v)
	case OpTrunc16to8:
		return rewriteValueRISCV_OpTrunc16to8_0(v)
	case OpTrunc32to16:
		return rewriteValueRISCV_OpTrunc32to16_0(v)
	case OpTrunc32to8:
		return rewriteValueRISCV_OpTrunc32to8_0(v)
	case OpTrunc64to16:
		return rewriteValueRISCV_OpTrunc64to16_0(v)
	case OpTrunc64to32:
		return rewriteValueRISCV_OpTrunc64to32_0(v)
	case OpTrunc64to8:
		return rewriteValueRISCV_OpTrunc64to8_0(v)
	case OpWB:
		return rewriteValueRISCV_OpWB_0(v)
	case OpXor16:
		return rewriteValueRISCV_OpXor16_0(v)
	case OpXor32:
		return rewriteValueRISCV_OpXor32_0(v)
	case OpXor64:
		return rewriteValueRISCV_OpXor64_0(v)
	case OpXor8:
		return rewriteValueRISCV_OpXor8_0(v)
	case OpZero:
		return rewriteValueRISCV_OpZero_0(v)
	case OpZeroExt16to32:
		return rewriteValueRISCV_OpZeroExt16to32_0(v)
	case OpZeroExt16to64:
		return rewriteValueRISCV_OpZeroExt16to64_0(v)
	case OpZeroExt32to64:
		return rewriteValueRISCV_OpZeroExt32to64_0(v)
	case OpZeroExt8to16:
		return rewriteValueRISCV_OpZeroExt8to16_0(v)
	case OpZeroExt8to32:
		return rewriteValueRISCV_OpZeroExt8to32_0(v)
	case OpZeroExt8to64:
		return rewriteValueRISCV_OpZeroExt8to64_0(v)
	}
	return false
}
func rewriteValueRISCV_OpAdd16_0(v *Value) bool {
	// match: (Add16 x y)
	// cond:
	// result: (ADD x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpAdd32_0(v *Value) bool {
	// match: (Add32 x y)
	// cond:
	// result: (ADD x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpAdd32F_0(v *Value) bool {
	// match: (Add32F x y)
	// cond:
	// result: (FADDS x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFADDS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpAdd64_0(v *Value) bool {
	// match: (Add64 x y)
	// cond:
	// result: (ADD x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpAdd64F_0(v *Value) bool {
	// match: (Add64F x y)
	// cond:
	// result: (FADDD x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFADDD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpAdd8_0(v *Value) bool {
	// match: (Add8 x y)
	// cond:
	// result: (ADD x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpAddPtr_0(v *Value) bool {
	// match: (AddPtr x y)
	// cond:
	// result: (ADD x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpAddr_0(v *Value) bool {
	// match: (Addr {sym} base)
	// cond:
	// result: (MOVaddr {sym} base)
	for {
		sym := v.Aux
		base := v.Args[0]
		v.reset(OpRISCVMOVaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueRISCV_OpAnd16_0(v *Value) bool {
	// match: (And16 x y)
	// cond:
	// result: (AND x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpAnd32_0(v *Value) bool {
	// match: (And32 x y)
	// cond:
	// result: (AND x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpAnd64_0(v *Value) bool {
	// match: (And64 x y)
	// cond:
	// result: (AND x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpAnd8_0(v *Value) bool {
	// match: (And8 x y)
	// cond:
	// result: (AND x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpAndB_0(v *Value) bool {
	// match: (AndB x y)
	// cond:
	// result: (AND x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpAvg64u_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Avg64u <t> x y)
	// cond:
	// result: (ADD (ADD <t> (SRLI <t> [1] x) (SRLI <t> [1] y)) (ANDI <t> [1] (AND <t> x y)))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVADD)
		v0 := b.NewValue0(v.Pos, OpRISCVADD, t)
		v1 := b.NewValue0(v.Pos, OpRISCVSRLI, t)
		v1.AuxInt = 1
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpRISCVSRLI, t)
		v2.AuxInt = 1
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpRISCVANDI, t)
		v3.AuxInt = 1
		v4 := b.NewValue0(v.Pos, OpRISCVAND, t)
		v4.AddArg(x)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueRISCV_OpClosureCall_0(v *Value) bool {
	// match: (ClosureCall [argwid] entry closure mem)
	// cond:
	// result: (CALLclosure [argwid] entry closure mem)
	for {
		argwid := v.AuxInt
		_ = v.Args[2]
		entry := v.Args[0]
		closure := v.Args[1]
		mem := v.Args[2]
		v.reset(OpRISCVCALLclosure)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(closure)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueRISCV_OpCom16_0(v *Value) bool {
	// match: (Com16 x)
	// cond:
	// result: (XORI [int64(-1)] x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVXORI)
		v.AuxInt = int64(-1)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpCom32_0(v *Value) bool {
	// match: (Com32 x)
	// cond:
	// result: (XORI [int64(-1)] x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVXORI)
		v.AuxInt = int64(-1)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpCom64_0(v *Value) bool {
	// match: (Com64 x)
	// cond:
	// result: (XORI [int64(-1)] x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVXORI)
		v.AuxInt = int64(-1)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpCom8_0(v *Value) bool {
	// match: (Com8 x)
	// cond:
	// result: (XORI [int64(-1)] x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVXORI)
		v.AuxInt = int64(-1)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpConst16_0(v *Value) bool {
	// match: (Const16 [val])
	// cond:
	// result: (MOVHconst [val])
	for {
		val := v.AuxInt
		v.reset(OpRISCVMOVHconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueRISCV_OpConst32_0(v *Value) bool {
	// match: (Const32 [val])
	// cond:
	// result: (MOVWconst [val])
	for {
		val := v.AuxInt
		v.reset(OpRISCVMOVWconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueRISCV_OpConst32F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Const32F [val])
	// cond:
	// result: (FMVSX (MOVSconst [val]))
	for {
		val := v.AuxInt
		v.reset(OpRISCVFMVSX)
		v0 := b.NewValue0(v.Pos, OpRISCVMOVSconst, typ.Float32)
		v0.AuxInt = val
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpConst64_0(v *Value) bool {
	// match: (Const64 [val])
	// cond:
	// result: (MOVDconst [val])
	for {
		val := v.AuxInt
		v.reset(OpRISCVMOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueRISCV_OpConst64F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Const64F [val])
	// cond:
	// result: (FMVDX (MOVDconst [val]))
	for {
		val := v.AuxInt
		v.reset(OpRISCVFMVDX)
		v0 := b.NewValue0(v.Pos, OpRISCVMOVDconst, typ.UInt64)
		v0.AuxInt = val
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpConst8_0(v *Value) bool {
	// match: (Const8 [val])
	// cond:
	// result: (MOVBconst [val])
	for {
		val := v.AuxInt
		v.reset(OpRISCVMOVBconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueRISCV_OpConstBool_0(v *Value) bool {
	// match: (ConstBool [b])
	// cond:
	// result: (MOVBconst [b])
	for {
		b := v.AuxInt
		v.reset(OpRISCVMOVBconst)
		v.AuxInt = b
		return true
	}
}
func rewriteValueRISCV_OpConstNil_0(v *Value) bool {
	// match: (ConstNil)
	// cond:
	// result: (MOVDconst [0])
	for {
		v.reset(OpRISCVMOVDconst)
		v.AuxInt = 0
		return true
	}
}
func rewriteValueRISCV_OpConvert_0(v *Value) bool {
	// match: (Convert x mem)
	// cond:
	// result: (MOVconvert x mem)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		mem := v.Args[1]
		v.reset(OpRISCVMOVconvert)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueRISCV_OpCvt32Fto32_0(v *Value) bool {
	// match: (Cvt32Fto32 x)
	// cond:
	// result: (FCVTWS x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVFCVTWS)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpCvt32Fto64_0(v *Value) bool {
	// match: (Cvt32Fto64 x)
	// cond:
	// result: (FCVTLS x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVFCVTLS)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpCvt32Fto64F_0(v *Value) bool {
	// match: (Cvt32Fto64F x)
	// cond:
	// result: (FCVTDS x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVFCVTDS)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpCvt32to32F_0(v *Value) bool {
	// match: (Cvt32to32F x)
	// cond:
	// result: (FCVTSW x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVFCVTSW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpCvt32to64F_0(v *Value) bool {
	// match: (Cvt32to64F x)
	// cond:
	// result: (FCVTDW x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVFCVTDW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpCvt64Fto32_0(v *Value) bool {
	// match: (Cvt64Fto32 x)
	// cond:
	// result: (FCVTWD x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVFCVTWD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpCvt64Fto32F_0(v *Value) bool {
	// match: (Cvt64Fto32F x)
	// cond:
	// result: (FCVTSD x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVFCVTSD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpCvt64Fto64_0(v *Value) bool {
	// match: (Cvt64Fto64 x)
	// cond:
	// result: (FCVTLD x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVFCVTLD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpCvt64to32F_0(v *Value) bool {
	// match: (Cvt64to32F x)
	// cond:
	// result: (FCVTSL x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVFCVTSL)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpCvt64to64F_0(v *Value) bool {
	// match: (Cvt64to64F x)
	// cond:
	// result: (FCVTDL x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVFCVTDL)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpDiv16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Div16 x y)
	// cond:
	// result: (DIVW (SignExt16to32 x) (SignExt16to32 y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVDIVW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpDiv16u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Div16u x y)
	// cond:
	// result: (DIVUW (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVDIVUW)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpDiv32_0(v *Value) bool {
	// match: (Div32 x y)
	// cond:
	// result: (DIVW x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVDIVW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpDiv32F_0(v *Value) bool {
	// match: (Div32F x y)
	// cond:
	// result: (FDIVS x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFDIVS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpDiv32u_0(v *Value) bool {
	// match: (Div32u x y)
	// cond:
	// result: (DIVUW x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVDIVUW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpDiv64_0(v *Value) bool {
	// match: (Div64 x y)
	// cond:
	// result: (DIV x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVDIV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpDiv64F_0(v *Value) bool {
	// match: (Div64F x y)
	// cond:
	// result: (FDIVD x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFDIVD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpDiv64u_0(v *Value) bool {
	// match: (Div64u x y)
	// cond:
	// result: (DIVU x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVDIVU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpDiv8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Div8 x y)
	// cond:
	// result: (DIVW (SignExt8to32 x) (SignExt8to32 y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVDIVW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpDiv8u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Div8u x y)
	// cond:
	// result: (DIVUW (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVDIVUW)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpEq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Eq16 x y)
	// cond:
	// result: (SEQZ (ZeroExt16to64 (SUB <x.Type> x y)))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSEQZ)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpRISCVSUB, x.Type)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpEq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Eq32 x y)
	// cond:
	// result: (SEQZ (ZeroExt32to64 (SUB <x.Type> x y)))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSEQZ)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpRISCVSUB, x.Type)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpEq32F_0(v *Value) bool {
	// match: (Eq32F x y)
	// cond:
	// result: (FEQS x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFEQS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpEq64_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Eq64 x y)
	// cond:
	// result: (SEQZ (SUB <x.Type> x y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSEQZ)
		v0 := b.NewValue0(v.Pos, OpRISCVSUB, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpEq64F_0(v *Value) bool {
	// match: (Eq64F x y)
	// cond:
	// result: (FEQD x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFEQD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpEq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Eq8 x y)
	// cond:
	// result: (SEQZ (ZeroExt8to64 (SUB <x.Type> x y)))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSEQZ)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpRISCVSUB, x.Type)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpEqB_0(v *Value) bool {
	// match: (EqB x y)
	// cond:
	// result: (Eq8 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpEq8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpEqPtr_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (EqPtr x y)
	// cond:
	// result: (SEQZ (SUB <x.Type> x y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSEQZ)
		v0 := b.NewValue0(v.Pos, OpRISCVSUB, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpGeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Geq16 x y)
	// cond:
	// result: (Not (Less16 x y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess16, typ.Bool)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpGeq16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Geq16U x y)
	// cond:
	// result: (Not (Less16U x y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess16U, typ.Bool)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpGeq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Geq32 x y)
	// cond:
	// result: (Not (Less32 x y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess32, typ.Bool)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpGeq32F_0(v *Value) bool {
	// match: (Geq32F x y)
	// cond:
	// result: (FLES y x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFLES)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpGeq32U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Geq32U x y)
	// cond:
	// result: (Not (Less32U x y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess32U, typ.Bool)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpGeq64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Geq64 x y)
	// cond:
	// result: (Not (Less64 x y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess64, typ.Bool)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpGeq64F_0(v *Value) bool {
	// match: (Geq64F x y)
	// cond:
	// result: (FLED y x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFLED)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpGeq64U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Geq64U x y)
	// cond:
	// result: (Not (Less64U x y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess64U, typ.Bool)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpGeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Geq8 x y)
	// cond:
	// result: (Not (Less8 x y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess8, typ.Bool)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpGeq8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Geq8U x y)
	// cond:
	// result: (Not (Less8U x y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess8U, typ.Bool)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpGetCallerPC_0(v *Value) bool {
	// match: (GetCallerPC)
	// cond:
	// result: (LoweredGetCallerPC)
	for {
		v.reset(OpRISCVLoweredGetCallerPC)
		return true
	}
}
func rewriteValueRISCV_OpGetCallerSP_0(v *Value) bool {
	// match: (GetCallerSP)
	// cond:
	// result: (LoweredGetCallerSP)
	for {
		v.reset(OpRISCVLoweredGetCallerSP)
		return true
	}
}
func rewriteValueRISCV_OpGetClosurePtr_0(v *Value) bool {
	// match: (GetClosurePtr)
	// cond:
	// result: (LoweredGetClosurePtr)
	for {
		v.reset(OpRISCVLoweredGetClosurePtr)
		return true
	}
}
func rewriteValueRISCV_OpGreater16_0(v *Value) bool {
	// match: (Greater16 x y)
	// cond:
	// result: (Less16 y x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLess16)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpGreater16U_0(v *Value) bool {
	// match: (Greater16U x y)
	// cond:
	// result: (Less16U y x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLess16U)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpGreater32_0(v *Value) bool {
	// match: (Greater32 x y)
	// cond:
	// result: (Less32 y x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLess32)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpGreater32F_0(v *Value) bool {
	// match: (Greater32F x y)
	// cond:
	// result: (FLTS y x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFLTS)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpGreater32U_0(v *Value) bool {
	// match: (Greater32U x y)
	// cond:
	// result: (Less32U y x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLess32U)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpGreater64_0(v *Value) bool {
	// match: (Greater64 x y)
	// cond:
	// result: (Less64 y x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLess64)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpGreater64F_0(v *Value) bool {
	// match: (Greater64F x y)
	// cond:
	// result: (FLTD y x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFLTD)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpGreater64U_0(v *Value) bool {
	// match: (Greater64U x y)
	// cond:
	// result: (Less64U y x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLess64U)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpGreater8_0(v *Value) bool {
	// match: (Greater8 x y)
	// cond:
	// result: (Less8 y x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLess8)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpGreater8U_0(v *Value) bool {
	// match: (Greater8U x y)
	// cond:
	// result: (Less8U y x)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLess8U)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpHmul32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Hmul32 x y)
	// cond:
	// result: (SRAI [32] (MUL (SignExt32to64 x) (SignExt32to64 y)))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSRAI)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpRISCVMUL, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpHmul32u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Hmul32u x y)
	// cond:
	// result: (SRLI [32] (MUL (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSRLI)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpRISCVMUL, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpHmul64_0(v *Value) bool {
	// match: (Hmul64 x y)
	// cond:
	// result: (MULH x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVMULH)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpHmul64u_0(v *Value) bool {
	// match: (Hmul64u x y)
	// cond:
	// result: (MULHU x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVMULHU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpInterCall_0(v *Value) bool {
	// match: (InterCall [argwid] entry mem)
	// cond:
	// result: (CALLinter [argwid] entry mem)
	for {
		argwid := v.AuxInt
		_ = v.Args[1]
		entry := v.Args[0]
		mem := v.Args[1]
		v.reset(OpRISCVCALLinter)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueRISCV_OpIsInBounds_0(v *Value) bool {
	// match: (IsInBounds idx len)
	// cond:
	// result: (Less64U idx len)
	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpLess64U)
		v.AddArg(idx)
		v.AddArg(len)
		return true
	}
}
func rewriteValueRISCV_OpIsNonNil_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (IsNonNil p)
	// cond:
	// result: (NeqPtr (MOVDconst) p)
	for {
		p := v.Args[0]
		v.reset(OpNeqPtr)
		v0 := b.NewValue0(v.Pos, OpRISCVMOVDconst, typ.UInt64)
		v.AddArg(v0)
		v.AddArg(p)
		return true
	}
}
func rewriteValueRISCV_OpIsSliceInBounds_0(v *Value) bool {
	// match: (IsSliceInBounds idx len)
	// cond:
	// result: (Leq64U idx len)
	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpLeq64U)
		v.AddArg(idx)
		v.AddArg(len)
		return true
	}
}
func rewriteValueRISCV_OpLeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Leq16 x y)
	// cond:
	// result: (Not (Less16 y x))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess16, typ.Bool)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpLeq16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Leq16U x y)
	// cond:
	// result: (Not (Less16U y x))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess16U, typ.Bool)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpLeq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Leq32 x y)
	// cond:
	// result: (Not (Less32 y x))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess32, typ.Bool)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpLeq32F_0(v *Value) bool {
	// match: (Leq32F x y)
	// cond:
	// result: (FLES x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFLES)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpLeq32U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Leq32U x y)
	// cond:
	// result: (Not (Less32U y x))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess32U, typ.Bool)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpLeq64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Leq64 x y)
	// cond:
	// result: (Not (Less64 y x))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess64, typ.Bool)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpLeq64F_0(v *Value) bool {
	// match: (Leq64F x y)
	// cond:
	// result: (FLED x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFLED)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpLeq64U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Leq64U x y)
	// cond:
	// result: (Not (Less64U y x))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess64U, typ.Bool)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpLeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Leq8 x y)
	// cond:
	// result: (Not (Less8 y x))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess8, typ.Bool)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpLeq8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Leq8U x y)
	// cond:
	// result: (Not (Less8U y x))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpLess8U, typ.Bool)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpLess16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Less16 x y)
	// cond:
	// result: (SLT (SignExt16to64 x) (SignExt16to64 y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSLT)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLess16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Less16U x y)
	// cond:
	// result: (SLTU (ZeroExt16to64 x) (ZeroExt16to64 y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSLTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLess32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Less32 x y)
	// cond:
	// result: (SLT (SignExt32to64 x) (SignExt32to64 y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSLT)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLess32F_0(v *Value) bool {
	// match: (Less32F x y)
	// cond:
	// result: (FLTS x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFLTS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpLess32U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Less32U x y)
	// cond:
	// result: (SLTU (ZeroExt32to64 x) (ZeroExt32to64 y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSLTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLess64_0(v *Value) bool {
	// match: (Less64 x y)
	// cond:
	// result: (SLT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSLT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpLess64F_0(v *Value) bool {
	// match: (Less64F x y)
	// cond:
	// result: (FLTD x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFLTD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpLess64U_0(v *Value) bool {
	// match: (Less64U x y)
	// cond:
	// result: (SLTU x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSLTU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpLess8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Less8 x y)
	// cond:
	// result: (SLT (SignExt8to64 x) (SignExt8to64 y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSLT)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLess8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Less8U x y)
	// cond:
	// result: (SLTU (ZeroExt8to64 x) (ZeroExt8to64 y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSLTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLoad_0(v *Value) bool {
	// match: (Load <t> ptr mem)
	// cond: t.IsBoolean()
	// result: (MOVBUload ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsBoolean()) {
			break
		}
		v.reset(OpRISCVMOVBUload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: ( is8BitInt(t) && isSigned(t))
	// result: (MOVBload ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is8BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpRISCVMOVBload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: ( is8BitInt(t) && !isSigned(t))
	// result: (MOVBUload ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is8BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpRISCVMOVBUload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is16BitInt(t) && isSigned(t))
	// result: (MOVHload ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is16BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpRISCVMOVHload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is16BitInt(t) && !isSigned(t))
	// result: (MOVHUload ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is16BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpRISCVMOVHUload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is32BitInt(t) && isSigned(t))
	// result: (MOVWload ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is32BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpRISCVMOVWload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is32BitInt(t) && !isSigned(t))
	// result: (MOVWUload ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is32BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpRISCVMOVWUload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (MOVDload ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpRISCVMOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitFloat(t)
	// result: (FMOVWload ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is32BitFloat(t)) {
			break
		}
		v.reset(OpRISCVFMOVWload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is64BitFloat(t)
	// result: (FMOVDload ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is64BitFloat(t)) {
			break
		}
		v.reset(OpRISCVFMOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueRISCV_OpLocalAddr_0(v *Value) bool {
	// match: (LocalAddr {sym} base _)
	// cond:
	// result: (MOVaddr {sym} base)
	for {
		sym := v.Aux
		_ = v.Args[1]
		base := v.Args[0]
		v.reset(OpRISCVMOVaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueRISCV_OpLsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh16x16 <t> x y)
	// cond:
	// result: (AND (SLL <t> x y) (Neg16 <t> (SLTIU <t> [64] (ZeroExt16to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg16, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh16x32 <t> x y)
	// cond:
	// result: (AND (SLL <t> x y) (Neg16 <t> (SLTIU <t> [64] (ZeroExt32to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg16, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLsh16x64_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Lsh16x64 <t> x y)
	// cond:
	// result: (AND (SLL <t> x y) (Neg16 <t> (SLTIU <t> [64] y)))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg16, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh16x8 <t> x y)
	// cond:
	// result: (AND (SLL <t> x y) (Neg16 <t> (SLTIU <t> [64] (ZeroExt8to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg16, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh32x16 <t> x y)
	// cond:
	// result: (AND (SLL <t> x y) (Neg32 <t> (SLTIU <t> [64] (ZeroExt16to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg32, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh32x32 <t> x y)
	// cond:
	// result: (AND (SLL <t> x y) (Neg32 <t> (SLTIU <t> [64] (ZeroExt32to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg32, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLsh32x64_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Lsh32x64 <t> x y)
	// cond:
	// result: (AND (SLL <t> x y) (Neg32 <t> (SLTIU <t> [64] y)))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg32, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh32x8 <t> x y)
	// cond:
	// result: (AND (SLL <t> x y) (Neg32 <t> (SLTIU <t> [64] (ZeroExt8to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg32, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLsh64x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh64x16 <t> x y)
	// cond:
	// result: (AND (SLL <t> x y) (Neg64 <t> (SLTIU <t> [64] (ZeroExt16to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg64, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLsh64x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh64x32 <t> x y)
	// cond:
	// result: (AND (SLL <t> x y) (Neg64 <t> (SLTIU <t> [64] (ZeroExt32to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg64, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLsh64x64_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Lsh64x64 <t> x y)
	// cond:
	// result: (AND (SLL <t> x y) (Neg64 <t> (SLTIU <t> [64] y)))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg64, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLsh64x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh64x8 <t> x y)
	// cond:
	// result: (AND (SLL <t> x y) (Neg64 <t> (SLTIU <t> [64] (ZeroExt8to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg64, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh8x16 <t> x y)
	// cond:
	// result: (AND (SLL <t> x y) (Neg8 <t> (SLTIU <t> [64] (ZeroExt16to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg8, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh8x32 <t> x y)
	// cond:
	// result: (AND (SLL <t> x y) (Neg8 <t> (SLTIU <t> [64] (ZeroExt32to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg8, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLsh8x64_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Lsh8x64 <t> x y)
	// cond:
	// result: (AND (SLL <t> x y) (Neg8 <t> (SLTIU <t> [64] y)))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg8, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpLsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh8x8 <t> x y)
	// cond:
	// result: (AND (SLL <t> x y) (Neg8 <t> (SLTIU <t> [64] (ZeroExt8to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg8, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpMod16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Mod16 x y)
	// cond:
	// result: (REMW (SignExt16to32 x) (SignExt16to32 y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVREMW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpMod16u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Mod16u x y)
	// cond:
	// result: (REMUW (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVREMUW)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpMod32_0(v *Value) bool {
	// match: (Mod32 x y)
	// cond:
	// result: (REMW x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVREMW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpMod32u_0(v *Value) bool {
	// match: (Mod32u x y)
	// cond:
	// result: (REMUW x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVREMUW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpMod64_0(v *Value) bool {
	// match: (Mod64 x y)
	// cond:
	// result: (REM x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVREM)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpMod64u_0(v *Value) bool {
	// match: (Mod64u x y)
	// cond:
	// result: (REMU x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVREMU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpMod8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Mod8 x y)
	// cond:
	// result: (REMW (SignExt8to32 x) (SignExt8to32 y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVREMW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpMod8u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Mod8u x y)
	// cond:
	// result: (REMUW (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVREMUW)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpMove_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Move [0] _ _ mem)
	// cond:
	// result: mem
	for {
		if v.AuxInt != 0 {
			break
		}
		_ = v.Args[2]
		mem := v.Args[2]
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}
	// match: (Move [1] dst src mem)
	// cond:
	// result: (MOVBstore dst (MOVBload src mem) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpRISCVMOVBstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpRISCVMOVBload, typ.Int8)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [2] dst src mem)
	// cond:
	// result: (MOVHstore dst (MOVHload src mem) mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpRISCVMOVHstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpRISCVMOVHload, typ.Int16)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [4] dst src mem)
	// cond:
	// result: (MOVWstore dst (MOVWload src mem) mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpRISCVMOVWstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpRISCVMOVWload, typ.Int32)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [8] dst src mem)
	// cond:
	// result: (MOVDstore dst (MOVDload src mem) mem)
	for {
		if v.AuxInt != 8 {
			break
		}
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpRISCVMOVDstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpRISCVMOVDload, typ.Int64)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [s] {t} dst src mem)
	// cond:
	// result: (LoweredMove [t.(*types.Type).Alignment()] dst src (ADDI <src.Type> [s-moveSize(t.(*types.Type).Alignment(), config)] src) mem)
	for {
		s := v.AuxInt
		t := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpRISCVLoweredMove)
		v.AuxInt = t.(*types.Type).Alignment()
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpRISCVADDI, src.Type)
		v0.AuxInt = s - moveSize(t.(*types.Type).Alignment(), config)
		v0.AddArg(src)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueRISCV_OpMul16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Mul16 x y)
	// cond:
	// result: (MULW (SignExt16to32 x) (SignExt16to32 y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVMULW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpMul32_0(v *Value) bool {
	// match: (Mul32 x y)
	// cond:
	// result: (MULW x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVMULW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpMul32F_0(v *Value) bool {
	// match: (Mul32F x y)
	// cond:
	// result: (FMULS x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFMULS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpMul64_0(v *Value) bool {
	// match: (Mul64 x y)
	// cond:
	// result: (MUL x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpMul64F_0(v *Value) bool {
	// match: (Mul64F x y)
	// cond:
	// result: (FMULD x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpMul8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Mul8 x y)
	// cond:
	// result: (MULW (SignExt8to32 x) (SignExt8to32 y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVMULW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpNeg16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Neg16 x)
	// cond:
	// result: (SUB (MOVHconst) x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVSUB)
		v0 := b.NewValue0(v.Pos, OpRISCVMOVHconst, typ.UInt16)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpNeg32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Neg32 x)
	// cond:
	// result: (SUB (MOVWconst) x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVSUB)
		v0 := b.NewValue0(v.Pos, OpRISCVMOVWconst, typ.UInt32)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpNeg32F_0(v *Value) bool {
	// match: (Neg32F x)
	// cond:
	// result: (FNEGS x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVFNEGS)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpNeg64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Neg64 x)
	// cond:
	// result: (SUB (MOVDconst) x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVSUB)
		v0 := b.NewValue0(v.Pos, OpRISCVMOVDconst, typ.UInt64)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpNeg64F_0(v *Value) bool {
	// match: (Neg64F x)
	// cond:
	// result: (FNEGD x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVFNEGD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpNeg8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Neg8 x)
	// cond:
	// result: (SUB (MOVBconst) x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVSUB)
		v0 := b.NewValue0(v.Pos, OpRISCVMOVBconst, typ.UInt8)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpNeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Neq16 x y)
	// cond:
	// result: (SNEZ (ZeroExt16to64 (SUB <x.Type> x y)))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSNEZ)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpRISCVSUB, x.Type)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpNeq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Neq32 x y)
	// cond:
	// result: (SNEZ (ZeroExt32to64 (SUB <x.Type> x y)))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSNEZ)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpRISCVSUB, x.Type)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpNeq32F_0(v *Value) bool {
	// match: (Neq32F x y)
	// cond:
	// result: (FNES x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFNES)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpNeq64_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Neq64 x y)
	// cond:
	// result: (SNEZ (SUB <x.Type> x y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSNEZ)
		v0 := b.NewValue0(v.Pos, OpRISCVSUB, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpNeq64F_0(v *Value) bool {
	// match: (Neq64F x y)
	// cond:
	// result: (FNED x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFNED)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpNeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Neq8 x y)
	// cond:
	// result: (SNEZ (ZeroExt8to64 (SUB <x.Type> x y)))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSNEZ)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpRISCVSUB, x.Type)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpNeqB_0(v *Value) bool {
	// match: (NeqB x y)
	// cond:
	// result: (Neq8 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpNeq8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpNeqPtr_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (NeqPtr x y)
	// cond:
	// result: (SNEZ (SUB <x.Type> x y))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSNEZ)
		v0 := b.NewValue0(v.Pos, OpRISCVSUB, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpNilCheck_0(v *Value) bool {
	// match: (NilCheck ptr mem)
	// cond:
	// result: (LoweredNilCheck ptr mem)
	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpRISCVLoweredNilCheck)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueRISCV_OpNot_0(v *Value) bool {
	// match: (Not x)
	// cond:
	// result: (XORI [1] x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVXORI)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpOffPtr_0(v *Value) bool {
	// match: (OffPtr [off] ptr:(SP))
	// cond:
	// result: (MOVaddr [off] ptr)
	for {
		off := v.AuxInt
		ptr := v.Args[0]
		if ptr.Op != OpSP {
			break
		}
		v.reset(OpRISCVMOVaddr)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
	// match: (OffPtr [off] ptr)
	// cond:
	// result: (ADDI [off] ptr)
	for {
		off := v.AuxInt
		ptr := v.Args[0]
		v.reset(OpRISCVADDI)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueRISCV_OpOr16_0(v *Value) bool {
	// match: (Or16 x y)
	// cond:
	// result: (OR x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpOr32_0(v *Value) bool {
	// match: (Or32 x y)
	// cond:
	// result: (OR x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpOr64_0(v *Value) bool {
	// match: (Or64 x y)
	// cond:
	// result: (OR x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpOr8_0(v *Value) bool {
	// match: (Or8 x y)
	// cond:
	// result: (OR x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpOrB_0(v *Value) bool {
	// match: (OrB x y)
	// cond:
	// result: (OR x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpRISCVADDI_0(v *Value) bool {
	// match: (ADDI [c] (MOVaddr [d] {s} x))
	// cond: is32Bit(c+d)
	// result: (MOVaddr [c+d] {s} x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVMOVaddr {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpRISCVMOVaddr)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		return true
	}
	// match: (ADDI [0] x)
	// cond:
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueRISCV_OpRISCVMOVBUload_0(v *Value) bool {
	// match: (MOVBUload [off1] {sym1} (MOVaddr [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVBUload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVMOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCVMOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBUload [off1] {sym} (ADDI [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBUload [off1+off2] {sym} base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCVMOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueRISCV_OpRISCVMOVBload_0(v *Value) bool {
	// match: (MOVBload [off1] {sym1} (MOVaddr [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVBload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVMOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCVMOVBload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [off1] {sym} (ADDI [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBload [off1+off2] {sym} base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCVMOVBload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueRISCV_OpRISCVMOVBstore_0(v *Value) bool {
	// match: (MOVBstore [off1] {sym1} (MOVaddr [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVBstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVMOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCVMOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off1] {sym} (ADDI [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBstore [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCVMOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueRISCV_OpRISCVMOVDconst_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (MOVDconst <t> [c])
	// cond: !is32Bit(c) && int32(c) < 0
	// result: (ADD (SLLI <t> [32] (MOVDconst [c>>32+1])) (MOVDconst [int64(int32(c))]))
	for {
		t := v.Type
		c := v.AuxInt
		if !(!is32Bit(c) && int32(c) < 0) {
			break
		}
		v.reset(OpRISCVADD)
		v0 := b.NewValue0(v.Pos, OpRISCVSLLI, t)
		v0.AuxInt = 32
		v1 := b.NewValue0(v.Pos, OpRISCVMOVDconst, typ.UInt64)
		v1.AuxInt = c>>32 + 1
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpRISCVMOVDconst, typ.UInt64)
		v2.AuxInt = int64(int32(c))
		v.AddArg(v2)
		return true
	}
	// match: (MOVDconst <t> [c])
	// cond: !is32Bit(c) && int32(c) >= 0
	// result: (ADD (SLLI <t> [32] (MOVDconst [c>>32+0])) (MOVDconst [int64(int32(c))]))
	for {
		t := v.Type
		c := v.AuxInt
		if !(!is32Bit(c) && int32(c) >= 0) {
			break
		}
		v.reset(OpRISCVADD)
		v0 := b.NewValue0(v.Pos, OpRISCVSLLI, t)
		v0.AuxInt = 32
		v1 := b.NewValue0(v.Pos, OpRISCVMOVDconst, typ.UInt64)
		v1.AuxInt = c>>32 + 0
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpRISCVMOVDconst, typ.UInt64)
		v2.AuxInt = int64(int32(c))
		v.AddArg(v2)
		return true
	}
	return false
}
func rewriteValueRISCV_OpRISCVMOVDload_0(v *Value) bool {
	// match: (MOVDload [off1] {sym1} (MOVaddr [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVDload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVMOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCVMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDload [off1] {sym} (ADDI [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVDload [off1+off2] {sym} base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCVMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueRISCV_OpRISCVMOVDstore_0(v *Value) bool {
	// match: (MOVDstore [off1] {sym1} (MOVaddr [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVDstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVMOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCVMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [off1] {sym} (ADDI [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVDstore [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCVMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueRISCV_OpRISCVMOVHUload_0(v *Value) bool {
	// match: (MOVHUload [off1] {sym1} (MOVaddr [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVHUload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVMOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCVMOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHUload [off1] {sym} (ADDI [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVHUload [off1+off2] {sym} base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCVMOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueRISCV_OpRISCVMOVHload_0(v *Value) bool {
	// match: (MOVHload [off1] {sym1} (MOVaddr [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVHload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVMOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCVMOVHload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHload [off1] {sym} (ADDI [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVHload [off1+off2] {sym} base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCVMOVHload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueRISCV_OpRISCVMOVHstore_0(v *Value) bool {
	// match: (MOVHstore [off1] {sym1} (MOVaddr [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVHstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVMOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCVMOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off1] {sym} (ADDI [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVHstore [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCVMOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueRISCV_OpRISCVMOVWUload_0(v *Value) bool {
	// match: (MOVWUload [off1] {sym1} (MOVaddr [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVWUload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVMOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCVMOVWUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWUload [off1] {sym} (ADDI [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWUload [off1+off2] {sym} base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCVMOVWUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueRISCV_OpRISCVMOVWload_0(v *Value) bool {
	// match: (MOVWload [off1] {sym1} (MOVaddr [off2] {sym2} base) mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVWload [off1+off2] {mergeSym(sym1,sym2)} base mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVMOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCVMOVWload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off1] {sym} (ADDI [off2] base) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWload [off1+off2] {sym} base mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCVMOVWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueRISCV_OpRISCVMOVWstore_0(v *Value) bool {
	// match: (MOVWstore [off1] {sym1} (MOVaddr [off2] {sym2} base) val mem)
	// cond: is32Bit(off1+off2) && canMergeSym(sym1, sym2)
	// result: (MOVWstore [off1+off2] {mergeSym(sym1,sym2)} base val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVMOVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpRISCVMOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off1] {sym} (ADDI [off2] base) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWstore [off1+off2] {sym} base val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpRISCVADDI {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpRISCVMOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueRISCV_OpRound32F_0(v *Value) bool {
	// match: (Round32F x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpRound64F_0(v *Value) bool {
	// match: (Round64F x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpRsh16Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh16Ux16 <t> x y)
	// cond:
	// result: (AND (SRL <t> (ZeroExt16to64 x) y) (Neg16 <t> (SLTIU <t> [64] (ZeroExt16to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpNeg16, t)
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueRISCV_OpRsh16Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh16Ux32 <t> x y)
	// cond:
	// result: (AND (SRL <t> (ZeroExt16to64 x) y) (Neg16 <t> (SLTIU <t> [64] (ZeroExt32to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpNeg16, t)
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueRISCV_OpRsh16Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh16Ux64 <t> x y)
	// cond:
	// result: (AND (SRL <t> (ZeroExt16to64 x) y) (Neg16 <t> (SLTIU <t> [64] y)))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpNeg16, t)
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v3.AuxInt = 64
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueRISCV_OpRsh16Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh16Ux8 <t> x y)
	// cond:
	// result: (AND (SRL <t> (ZeroExt16to64 x) y) (Neg16 <t> (SLTIU <t> [64] (ZeroExt8to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpNeg16, t)
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueRISCV_OpRsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh16x16 <t> x y)
	// cond:
	// result: (SRA <t> (SignExt16to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt16to64 y)))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpRISCVOR, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpRISCVADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, y.Type)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpRsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh16x32 <t> x y)
	// cond:
	// result: (SRA <t> (SignExt16to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt32to64 y)))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpRISCVOR, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpRISCVADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, y.Type)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpRsh16x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh16x64 <t> x y)
	// cond:
	// result: (SRA <t> (SignExt16to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpRISCVOR, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpRISCVADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, y.Type)
		v3.AuxInt = 64
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpRsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh16x8 <t> x y)
	// cond:
	// result: (SRA <t> (SignExt16to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt8to64 y)))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpRISCVOR, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpRISCVADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, y.Type)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpRsh32Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh32Ux16 <t> x y)
	// cond:
	// result: (AND (SRL <t> (ZeroExt32to64 x) y) (Neg32 <t> (SLTIU <t> [64] (ZeroExt16to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpNeg32, t)
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueRISCV_OpRsh32Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh32Ux32 <t> x y)
	// cond:
	// result: (AND (SRL <t> (ZeroExt32to64 x) y) (Neg32 <t> (SLTIU <t> [64] (ZeroExt32to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpNeg32, t)
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueRISCV_OpRsh32Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh32Ux64 <t> x y)
	// cond:
	// result: (AND (SRL <t> (ZeroExt32to64 x) y) (Neg32 <t> (SLTIU <t> [64] y)))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpNeg32, t)
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v3.AuxInt = 64
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueRISCV_OpRsh32Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh32Ux8 <t> x y)
	// cond:
	// result: (AND (SRL <t> (ZeroExt32to64 x) y) (Neg32 <t> (SLTIU <t> [64] (ZeroExt8to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpNeg32, t)
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueRISCV_OpRsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh32x16 <t> x y)
	// cond:
	// result: (SRA <t> (SignExt32to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt16to64 y)))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpRISCVOR, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpRISCVADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, y.Type)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpRsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh32x32 <t> x y)
	// cond:
	// result: (SRA <t> (SignExt32to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt32to64 y)))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpRISCVOR, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpRISCVADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, y.Type)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpRsh32x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh32x64 <t> x y)
	// cond:
	// result: (SRA <t> (SignExt32to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpRISCVOR, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpRISCVADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, y.Type)
		v3.AuxInt = 64
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpRsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh32x8 <t> x y)
	// cond:
	// result: (SRA <t> (SignExt32to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt8to64 y)))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpRISCVOR, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpRISCVADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, y.Type)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpRsh64Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh64Ux16 <t> x y)
	// cond:
	// result: (AND (SRL <t> x y) (Neg64 <t> (SLTIU <t> [64] (ZeroExt16to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSRL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg64, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpRsh64Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh64Ux32 <t> x y)
	// cond:
	// result: (AND (SRL <t> x y) (Neg64 <t> (SLTIU <t> [64] (ZeroExt32to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSRL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg64, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpRsh64Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Rsh64Ux64 <t> x y)
	// cond:
	// result: (AND (SRL <t> x y) (Neg64 <t> (SLTIU <t> [64] y)))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSRL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg64, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpRsh64Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh64Ux8 <t> x y)
	// cond:
	// result: (AND (SRL <t> x y) (Neg64 <t> (SLTIU <t> [64] (ZeroExt8to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSRL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpNeg64, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpRsh64x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh64x16 <t> x y)
	// cond:
	// result: (SRA <t> x (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt16to64 y)))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSRA)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpRISCVOR, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpRISCVADDI, y.Type)
		v1.AuxInt = -1
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, y.Type)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpRsh64x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh64x32 <t> x y)
	// cond:
	// result: (SRA <t> x (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt32to64 y)))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSRA)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpRISCVOR, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpRISCVADDI, y.Type)
		v1.AuxInt = -1
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, y.Type)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpRsh64x64_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Rsh64x64 <t> x y)
	// cond:
	// result: (SRA <t> x (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSRA)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpRISCVOR, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpRISCVADDI, y.Type)
		v1.AuxInt = -1
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, y.Type)
		v2.AuxInt = 64
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpRsh64x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh64x8 <t> x y)
	// cond:
	// result: (SRA <t> x (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt8to64 y)))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSRA)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpRISCVOR, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpRISCVADDI, y.Type)
		v1.AuxInt = -1
		v2 := b.NewValue0(v.Pos, OpRISCVSLTIU, y.Type)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpRsh8Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh8Ux16 <t> x y)
	// cond:
	// result: (AND (SRL <t> (ZeroExt8to64 x) y) (Neg8 <t> (SLTIU <t> [64] (ZeroExt16to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpNeg8, t)
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueRISCV_OpRsh8Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh8Ux32 <t> x y)
	// cond:
	// result: (AND (SRL <t> (ZeroExt8to64 x) y) (Neg8 <t> (SLTIU <t> [64] (ZeroExt32to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpNeg8, t)
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueRISCV_OpRsh8Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh8Ux64 <t> x y)
	// cond:
	// result: (AND (SRL <t> (ZeroExt8to64 x) y) (Neg8 <t> (SLTIU <t> [64] y)))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpNeg8, t)
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v3.AuxInt = 64
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueRISCV_OpRsh8Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh8Ux8 <t> x y)
	// cond:
	// result: (AND (SRL <t> (ZeroExt8to64 x) y) (Neg8 <t> (SLTIU <t> [64] (ZeroExt8to64 y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVAND)
		v0 := b.NewValue0(v.Pos, OpRISCVSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpNeg8, t)
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, t)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueRISCV_OpRsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh8x16 <t> x y)
	// cond:
	// result: (SRA <t> (SignExt8to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt16to64 y)))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpRISCVOR, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpRISCVADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, y.Type)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpRsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh8x32 <t> x y)
	// cond:
	// result: (SRA <t> (SignExt8to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt32to64 y)))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpRISCVOR, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpRISCVADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, y.Type)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpRsh8x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh8x64 <t> x y)
	// cond:
	// result: (SRA <t> (SignExt8to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] y))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpRISCVOR, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpRISCVADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, y.Type)
		v3.AuxInt = 64
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpRsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh8x8 <t> x y)
	// cond:
	// result: (SRA <t> (SignExt8to64 x) (OR <y.Type> y (ADDI <y.Type> [-1] (SLTIU <y.Type> [64] (ZeroExt8to64 y)))))
	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSRA)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpRISCVOR, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpRISCVADDI, y.Type)
		v2.AuxInt = -1
		v3 := b.NewValue0(v.Pos, OpRISCVSLTIU, y.Type)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpSignExt16to32_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (SignExt16to32 <t> x)
	// cond:
	// result: (SRAI [48] (SLLI <t> [48] x))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpRISCVSRAI)
		v.AuxInt = 48
		v0 := b.NewValue0(v.Pos, OpRISCVSLLI, t)
		v0.AuxInt = 48
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpSignExt16to64_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (SignExt16to64 <t> x)
	// cond:
	// result: (SRAI [48] (SLLI <t> [48] x))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpRISCVSRAI)
		v.AuxInt = 48
		v0 := b.NewValue0(v.Pos, OpRISCVSLLI, t)
		v0.AuxInt = 48
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpSignExt32to64_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (SignExt32to64 <t> x)
	// cond:
	// result: (SRAI [32] (SLLI <t> [32] x))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpRISCVSRAI)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpRISCVSLLI, t)
		v0.AuxInt = 32
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpSignExt8to16_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (SignExt8to16 <t> x)
	// cond:
	// result: (SRAI [56] (SLLI <t> [56] x))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpRISCVSRAI)
		v.AuxInt = 56
		v0 := b.NewValue0(v.Pos, OpRISCVSLLI, t)
		v0.AuxInt = 56
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpSignExt8to32_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (SignExt8to32 <t> x)
	// cond:
	// result: (SRAI [56] (SLLI <t> [56] x))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpRISCVSRAI)
		v.AuxInt = 56
		v0 := b.NewValue0(v.Pos, OpRISCVSLLI, t)
		v0.AuxInt = 56
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpSignExt8to64_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (SignExt8to64 <t> x)
	// cond:
	// result: (SRAI [56] (SLLI <t> [56] x))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpRISCVSRAI)
		v.AuxInt = 56
		v0 := b.NewValue0(v.Pos, OpRISCVSLLI, t)
		v0.AuxInt = 56
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpSlicemask_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Slicemask <t> x)
	// cond:
	// result: (XOR (MOVDconst [-1]) (SRA <t> (SUB <t> x (MOVDconst [1])) (MOVDconst [63])))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpRISCVXOR)
		v0 := b.NewValue0(v.Pos, OpRISCVMOVDconst, typ.UInt64)
		v0.AuxInt = -1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpRISCVSRA, t)
		v2 := b.NewValue0(v.Pos, OpRISCVSUB, t)
		v2.AddArg(x)
		v3 := b.NewValue0(v.Pos, OpRISCVMOVDconst, typ.UInt64)
		v3.AuxInt = 1
		v2.AddArg(v3)
		v1.AddArg(v2)
		v4 := b.NewValue0(v.Pos, OpRISCVMOVDconst, typ.UInt64)
		v4.AuxInt = 63
		v1.AddArg(v4)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueRISCV_OpSqrt_0(v *Value) bool {
	// match: (Sqrt x)
	// cond:
	// result: (FSQRTD x)
	for {
		x := v.Args[0]
		v.reset(OpRISCVFSQRTD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpStaticCall_0(v *Value) bool {
	// match: (StaticCall [argwid] {target} mem)
	// cond:
	// result: (CALLstatic [argwid] {target} mem)
	for {
		argwid := v.AuxInt
		target := v.Aux
		mem := v.Args[0]
		v.reset(OpRISCVCALLstatic)
		v.AuxInt = argwid
		v.Aux = target
		v.AddArg(mem)
		return true
	}
}
func rewriteValueRISCV_OpStore_0(v *Value) bool {
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 1
	// result: (MOVBstore ptr val mem)
	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size() == 1) {
			break
		}
		v.reset(OpRISCVMOVBstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 2
	// result: (MOVHstore ptr val mem)
	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size() == 2) {
			break
		}
		v.reset(OpRISCVMOVHstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4 && !is32BitFloat(val.Type)
	// result: (MOVWstore ptr val mem)
	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size() == 4 && !is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpRISCVMOVWstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 8 && !is64BitFloat(val.Type)
	// result: (MOVDstore ptr val mem)
	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size() == 8 && !is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpRISCVMOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4 && is32BitFloat(val.Type)
	// result: (FMOVWstore ptr val mem)
	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size() == 4 && is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpRISCVFMOVWstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 8 && is64BitFloat(val.Type)
	// result: (FMOVDstore ptr val mem)
	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size() == 8 && is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpRISCVFMOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueRISCV_OpSub16_0(v *Value) bool {
	// match: (Sub16 x y)
	// cond:
	// result: (SUB x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpSub32_0(v *Value) bool {
	// match: (Sub32 x y)
	// cond:
	// result: (SUB x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpSub32F_0(v *Value) bool {
	// match: (Sub32F x y)
	// cond:
	// result: (FSUBS x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFSUBS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpSub64_0(v *Value) bool {
	// match: (Sub64 x y)
	// cond:
	// result: (SUB x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpSub64F_0(v *Value) bool {
	// match: (Sub64F x y)
	// cond:
	// result: (FSUBD x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVFSUBD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpSub8_0(v *Value) bool {
	// match: (Sub8 x y)
	// cond:
	// result: (SUB x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpSubPtr_0(v *Value) bool {
	// match: (SubPtr x y)
	// cond:
	// result: (SUB x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpTrunc16to8_0(v *Value) bool {
	// match: (Trunc16to8 x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpTrunc32to16_0(v *Value) bool {
	// match: (Trunc32to16 x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpTrunc32to8_0(v *Value) bool {
	// match: (Trunc32to8 x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpTrunc64to16_0(v *Value) bool {
	// match: (Trunc64to16 x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpTrunc64to32_0(v *Value) bool {
	// match: (Trunc64to32 x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpTrunc64to8_0(v *Value) bool {
	// match: (Trunc64to8 x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueRISCV_OpWB_0(v *Value) bool {
	// match: (WB {fn} destptr srcptr mem)
	// cond:
	// result: (LoweredWB {fn} destptr srcptr mem)
	for {
		fn := v.Aux
		_ = v.Args[2]
		destptr := v.Args[0]
		srcptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpRISCVLoweredWB)
		v.Aux = fn
		v.AddArg(destptr)
		v.AddArg(srcptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueRISCV_OpXor16_0(v *Value) bool {
	// match: (Xor16 x y)
	// cond:
	// result: (XOR x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpXor32_0(v *Value) bool {
	// match: (Xor32 x y)
	// cond:
	// result: (XOR x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpXor64_0(v *Value) bool {
	// match: (Xor64 x y)
	// cond:
	// result: (XOR x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpXor8_0(v *Value) bool {
	// match: (Xor8 x y)
	// cond:
	// result: (XOR x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRISCVXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueRISCV_OpZero_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Zero [0] _ mem)
	// cond:
	// result: mem
	for {
		if v.AuxInt != 0 {
			break
		}
		_ = v.Args[1]
		mem := v.Args[1]
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}
	// match: (Zero [1] ptr mem)
	// cond:
	// result: (MOVBstore ptr (MOVBconst) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpRISCVMOVBstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpRISCVMOVBconst, typ.UInt8)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [2] ptr mem)
	// cond:
	// result: (MOVHstore ptr (MOVHconst) mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpRISCVMOVHstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpRISCVMOVHconst, typ.UInt16)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [4] ptr mem)
	// cond:
	// result: (MOVWstore ptr (MOVWconst) mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpRISCVMOVWstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpRISCVMOVWconst, typ.UInt32)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [8] ptr mem)
	// cond:
	// result: (MOVDstore ptr (MOVDconst) mem)
	for {
		if v.AuxInt != 8 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpRISCVMOVDstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpRISCVMOVDconst, typ.UInt64)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// cond:
	// result: (LoweredZero [t.(*types.Type).Alignment()] ptr (ADD <ptr.Type> ptr (MOVDconst [s-moveSize(t.(*types.Type).Alignment(), config)])) mem)
	for {
		s := v.AuxInt
		t := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpRISCVLoweredZero)
		v.AuxInt = t.(*types.Type).Alignment()
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpRISCVADD, ptr.Type)
		v0.AddArg(ptr)
		v1 := b.NewValue0(v.Pos, OpRISCVMOVDconst, typ.UInt64)
		v1.AuxInt = s - moveSize(t.(*types.Type).Alignment(), config)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueRISCV_OpZeroExt16to32_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (ZeroExt16to32 <t> x)
	// cond:
	// result: (SRLI [48] (SLLI <t> [48] x))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpRISCVSRLI)
		v.AuxInt = 48
		v0 := b.NewValue0(v.Pos, OpRISCVSLLI, t)
		v0.AuxInt = 48
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpZeroExt16to64_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (ZeroExt16to64 <t> x)
	// cond:
	// result: (SRLI [48] (SLLI <t> [48] x))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpRISCVSRLI)
		v.AuxInt = 48
		v0 := b.NewValue0(v.Pos, OpRISCVSLLI, t)
		v0.AuxInt = 48
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpZeroExt32to64_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (ZeroExt32to64 <t> x)
	// cond:
	// result: (SRLI [32] (SLLI <t> [32] x))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpRISCVSRLI)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpRISCVSLLI, t)
		v0.AuxInt = 32
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpZeroExt8to16_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (ZeroExt8to16 <t> x)
	// cond:
	// result: (SRLI [56] (SLLI <t> [56] x))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpRISCVSRLI)
		v.AuxInt = 56
		v0 := b.NewValue0(v.Pos, OpRISCVSLLI, t)
		v0.AuxInt = 56
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpZeroExt8to32_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (ZeroExt8to32 <t> x)
	// cond:
	// result: (SRLI [56] (SLLI <t> [56] x))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpRISCVSRLI)
		v.AuxInt = 56
		v0 := b.NewValue0(v.Pos, OpRISCVSLLI, t)
		v0.AuxInt = 56
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueRISCV_OpZeroExt8to64_0(v *Value) bool {
	b := v.Block
	_ = b
	// match: (ZeroExt8to64 <t> x)
	// cond:
	// result: (SRLI [56] (SLLI <t> [56] x))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpRISCVSRLI)
		v.AuxInt = 56
		v0 := b.NewValue0(v.Pos, OpRISCVSLLI, t)
		v0.AuxInt = 56
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteBlockRISCV(b *Block) bool {
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	typ := &config.Types
	_ = typ
	switch b.Kind {
	case BlockIf:
		// match: (If cond yes no)
		// cond:
		// result: (BNE cond yes no)
		for {
			v := b.Control
			_ = v
			cond := b.Control
			b.Kind = BlockRISCVBNE
			b.SetControl(cond)
			b.Aux = nil
			return true
		}
	}
	return false
}
