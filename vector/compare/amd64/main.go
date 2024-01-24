package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"math/bits"
	"os"
	"regexp"
	"strings"

	. "github.com/mmcloughlin/avo/build"
	"github.com/mmcloughlin/avo/ir"
	. "github.com/mmcloughlin/avo/operand"
	"github.com/mmcloughlin/avo/reg"
)

var testhelp = flag.String("testhelp", "", "test helpers")

func main() {
	const variants = 6
	alignments := []int{0, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	// const variants = 1
	// alignments := []int{0}

	emitAlignments := func(emit func(variant, align int)) {
		for _, align := range alignments {
			for v := 0; v < variants; v++ {
				emit(v, align)
			}
		}
	}

	emitAlignments(AxpyPointer)
	emitAlignments(AxpyPointerLoop)
	emitAlignments(AxpyPointerLoopX)
	emitAlignments(AxpyUnsafeX)
	emitAlignments(func(variant, align int) { AxpyUnsafeXUnroll(variant, align, 4) })
	emitAlignments(func(variant, align int) { AxpyUnsafeXUnroll(variant, align, 8) })
	emitAlignments(func(variant, align int) { AxpyUnsafeXInterleaveUnroll(variant, align, 4) })
	emitAlignments(func(variant, align int) { AxpyUnsafeXInterleaveUnroll(variant, align, 8) })
	emitAlignments(func(variant, align int) { AxpyPointerLoopXUnroll(variant, align, 4) })
	emitAlignments(func(variant, align int) { AxpyPointerLoopXUnroll(variant, align, 8) })
	emitAlignments(func(variant, align int) { AxpyPointerLoopXInterleaveUnroll(variant, align, 4) })
	emitAlignments(func(variant, align int) { AxpyPointerLoopXInterleaveUnroll(variant, align, 8) })

	Generate()

	if *testhelp != "" {
		generateTestHelp("axpy_amd64.go", *testhelp)
	}
}

func generateTestHelp(stubs, out string) {
	data, err := os.ReadFile(stubs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fns := []string{}

	rx := regexp.MustCompile("func ([a-zA-Z0-9_]+)\\(")
	for _, match := range rx.FindAllStringSubmatch(string(data), -1) {
		fns = append(fns, match[1])
	}

	var b bytes.Buffer
	pf := func(format string, args ...interface{}) {
		fmt.Fprintf(&b, format, args...)
	}

	pf("// Code generated by command. DO NOT EDIT.\n\n")
	pf("package compare\n\n")
	pf("type amdAxpyDecl struct {\n")
	pf("	name string\n")
	pf("	fn   func(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)\n")
	pf("}\n\n")

	pf("var amdAxpyDecls = []amdAxpyDecl{\n")
	for _, fn := range fns {
		pf("	{name: %q, fn: %v},\n", strings.TrimPrefix(fn, "Amd"), fn)
	}
	pf("}\n")

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		fmt.Fprintln(os.Stderr, b.Bytes())
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.WriteFile(out, formatted, 0755)
}

func AxpyPointer(variant, align int) {
	TEXT(fmt.Sprintf("AmdAxpyPointer_V%vA%v", variant, align), NOSPLIT, "func(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)")

	alpha := Load(Param("alpha"), XMM())

	xs := Mem{Base: Load(Param("xs"), GP64())}
	incx := Load(Param("incx"), GP64())

	ys := Mem{Base: Load(Param("ys"), GP64())}
	incy := Load(Param("incy"), GP64())

	n := Load(Param("n"), GP64())

	end := n
	SHLQ(U8(0x2), end)
	IMULQ(incx, end)
	ADDQ(xs.Base, end)
	JMP(LabelRef("check_limit"))

	MISALIGN(align)
	Label("loop")
	{
		tmp := XMM()
		MOVSS(xs, tmp)
		MULSS(alpha, tmp)
		ADDSS(ys, tmp)
		MOVSS(tmp, ys)

		LEAQ(xs.Idx(incx, 4), xs.Base)
		LEAQ(ys.Idx(incy, 4), ys.Base)

		Label("check_limit")

		CMPQ(end, xs.Base)
		JHI(LabelRef("loop"))
	}

	RET()
}

func AxpyPointerLoop(variant, align int) {
	TEXT(fmt.Sprintf("AmdAxpyPointerLoop_V%vA%v", variant, align), NOSPLIT, "func(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)")

	alpha := Load(Param("alpha"), XMM())

	xs := Mem{Base: Load(Param("xs"), GP64())}
	incx := Load(Param("incx"), GP64())

	ys := Mem{Base: Load(Param("ys"), GP64())}
	incy := Load(Param("incy"), GP64())

	n := Load(Param("n"), GP64())
	counter := GP64()
	XORQ(counter, counter)

	JMP(LabelRef("check_limit"))

	MISALIGN(align)
	Label("loop")
	{
		tmp := XMM()
		MOVSS(xs, tmp)
		MULSS(alpha, tmp)
		ADDSS(ys, tmp)
		MOVSS(tmp, ys)

		INCQ(counter)

		LEAQ(xs.Idx(incx, 4), xs.Base)
		LEAQ(ys.Idx(incy, 4), ys.Base)

		Label("check_limit")

		CMPQ(n, counter)
		JHI(LabelRef("loop"))
	}

	RET()
}

func AxpyPointerLoopX(variant, align int) {
	TEXT(fmt.Sprintf("AmdAxpyPointerLoopX_V%vA%v", variant, align), NOSPLIT, "func(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)")

	alpha := Load(Param("alpha"), XMM())

	xs := Mem{Base: Load(Param("xs"), GP64())}
	incx := Load(Param("incx"), GP64())

	ys := Mem{Base: Load(Param("ys"), GP64())}
	incy := Load(Param("incy"), GP64())

	n := Load(Param("n"), GP64())

	JMP(LabelRef("check_limit"))

	MISALIGN(align)
	Label("loop")
	{
		tmp := XMM()
		MOVSS(xs, tmp)
		MULSS(alpha, tmp)
		ADDSS(ys, tmp)
		MOVSS(tmp, ys)

		DECQ(n)

		LEAQ(xs.Idx(incx, 4), xs.Base)
		LEAQ(ys.Idx(incy, 4), ys.Base)

		Label("check_limit")

		CMPQ(n, U8(0))
		JHI(LabelRef("loop"))
	}

	RET()
}

func log2(v int) int {
	if v&(v-1) != 0 {
		panic("not a power of two")
	}
	return bits.TrailingZeros(uint(v))
}

func AxpyPointerLoopXUnroll(variant, align, unroll int) {
	TEXT(fmt.Sprintf("AmdAxpyPointerLoopX_V%vA%vU%v", variant, align, unroll), NOSPLIT, "func(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)")

	alpha := Load(Param("alpha"), XMM())

	xs := Mem{Base: Load(Param("xs"), GP64())}
	incx := Load(Param("incx"), GP64())

	ys := Mem{Base: Load(Param("ys"), GP64())}
	incy := Load(Param("incy"), GP64())

	n := Load(Param("n"), GP64())

	JMP(LabelRef("check_limit_unroll"))

	MISALIGN(align)
	Label("loop_unroll")
	{
		for u := 0; u < unroll; u++ {
			tmp := XMM()

			MOVSS(xs, tmp)
			MULSS(alpha, tmp)
			ADDSS(ys, tmp)
			MOVSS(tmp, ys)

			LEAQ(xs.Idx(incx, 4), xs.Base)
			LEAQ(ys.Idx(incy, 4), ys.Base)
		}

		SUBQ(Imm(uint64(unroll)), n)

		Label("check_limit_unroll")

		CMPQ(n, U8(unroll))
		JHS(LabelRef("loop_unroll"))
	}

	JMP(LabelRef("check_limit"))
	Label("loop")
	{
		tmp := XMM()
		MOVSS(xs, tmp)
		MULSS(alpha, tmp)
		ADDSS(ys, tmp)
		MOVSS(tmp, ys)

		DECQ(n)

		LEAQ(xs.Idx(incx, 4), xs.Base)
		LEAQ(ys.Idx(incy, 4), ys.Base)

		Label("check_limit")

		CMPQ(n, U8(0))
		JHI(LabelRef("loop"))
	}

	RET()
}

func AxpyPointerLoopXInterleaveUnroll(variant, align, unroll int) {
	TEXT(fmt.Sprintf("AmdAxpyPointerLoopXInterleave_V%vA%vU%v", variant, align, unroll), NOSPLIT, "func(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)")

	alpha := Load(Param("alpha"), XMM())

	xs := Mem{Base: Load(Param("xs"), GP64())}
	incx := Load(Param("incx"), GP64())
	incxunroll := GP64()
	MOVQ(incx, incxunroll)
	SHLQ(U8(log2(4*unroll)), incxunroll)

	ys := Mem{Base: Load(Param("ys"), GP64())}
	incy := Load(Param("incy"), GP64())
	incyunroll := GP64()
	MOVQ(incy, incyunroll)
	SHLQ(U8(log2(4*unroll)), incyunroll)

	n := Load(Param("n"), GP64())

	JMP(LabelRef("check_limit_unroll"))

	MISALIGN(align)
	Label("loop_unroll")
	{
		tmp := make([]reg.VecVirtual, unroll)

		for u := range tmp {
			tmp[u] = XMM()
		}

		for u := 0; u < unroll; u++ {
			MOVSS(xs, tmp[u])
			LEAQ(xs.Idx(incx, 4), xs.Base)
		}
		for u := 0; u < unroll; u++ {
			MULSS(alpha, tmp[u])
		}
		for u := 0; u < unroll; u++ {
			ADDSS(ys, tmp[u])
			MOVSS(tmp[u], ys)
			LEAQ(ys.Idx(incy, 4), ys.Base)
		}

		SUBQ(Imm(uint64(unroll)), n)

		Label("check_limit_unroll")

		CMPQ(n, U8(unroll))
		JHS(LabelRef("loop_unroll"))
	}

	JMP(LabelRef("check_limit"))
	Label("loop")
	{
		tmp := XMM()
		MOVSS(xs, tmp)
		MULSS(alpha, tmp)
		ADDSS(ys, tmp)
		MOVSS(tmp, ys)

		DECQ(n)

		LEAQ(xs.Idx(incx, 4), xs.Base)
		LEAQ(ys.Idx(incy, 4), ys.Base)

		Label("check_limit")

		CMPQ(n, U8(0))
		JHI(LabelRef("loop"))
	}

	RET()
}

func AxpyUnsafeX(variant, align int) {
	TEXT(fmt.Sprintf("AmdAxpyUnsafeX_V%vA%v", variant, align), NOSPLIT, "func(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)")

	alpha := Load(Param("alpha"), XMM())

	xs := Mem{Base: Load(Param("xs"), GP64())}
	incx := Load(Param("incx"), GP64())

	ys := Mem{Base: Load(Param("ys"), GP64())}
	incy := Load(Param("incy"), GP64())

	n := Load(Param("n"), GP64())

	xi, yi := GP64(), GP64()
	XORQ(xi, xi)
	XORQ(yi, yi)

	JMP(LabelRef("check_limit"))

	MISALIGN(align)
	Label("loop")
	{
		tmp := XMM()
		MOVSS(xs.Idx(xi, 4), tmp)
		MULSS(alpha, tmp)
		ADDSS(ys.Idx(yi, 4), tmp)
		MOVSS(tmp, ys.Idx(yi, 4))

		DECQ(n)
		ADDQ(incx, xi)
		ADDQ(incy, yi)

		Label("check_limit")

		CMPQ(n, U8(0))
		JHI(LabelRef("loop"))
	}

	RET()
}

func AxpyUnsafeXUnroll(variant, align, unroll int) {
	TEXT(fmt.Sprintf("AmdAxpyUnsafeX_V%vA%vR%v", variant, align, unroll), NOSPLIT, "func(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)")

	alpha := Load(Param("alpha"), XMM())

	xs := Mem{Base: Load(Param("xs"), GP64())}
	incx := Load(Param("incx"), GP64())

	ys := Mem{Base: Load(Param("ys"), GP64())}
	incy := Load(Param("incy"), GP64())

	n := Load(Param("n"), GP64())

	xi, yi := GP64(), GP64()
	XORQ(xi, xi)
	XORQ(yi, yi)

	JMP(LabelRef("check_limit_unroll"))

	MISALIGN(align)
	Label("loop_unroll")
	{
		for u := 0; u < unroll; u++ {
			tmp := XMM()

			xat := Mem{Base: xs.Base, Index: xi, Scale: 4, Disp: 0}
			yat := Mem{Base: ys.Base, Index: yi, Scale: 4, Disp: 0}
			MOVSS(xat, tmp)
			MULSS(alpha, tmp)
			ADDSS(yat, tmp)
			MOVSS(tmp, yat)

			ADDQ(incx, xi)
			ADDQ(incy, yi)
		}

		SUBQ(Imm(uint64(unroll)), n)

		Label("check_limit_unroll")

		CMPQ(n, U8(unroll))
		JHI(LabelRef("loop_unroll"))
	}

	JMP(LabelRef("check_limit"))
	Label("loop")
	{
		tmp := XMM()
		MOVSS(xs.Idx(xi, 4), tmp)
		MULSS(alpha, tmp)
		ADDSS(ys.Idx(yi, 4), tmp)
		MOVSS(tmp, ys.Idx(yi, 4))

		DECQ(n)
		ADDQ(incx, xi)
		ADDQ(incy, yi)

		Label("check_limit")

		CMPQ(n, U8(0))
		JHI(LabelRef("loop"))
	}

	RET()
}

func AxpyUnsafeXInterleaveUnroll(variant, align, unroll int) {
	TEXT(fmt.Sprintf("AmdAxpyUnsafeXInterleave_V%vA%vR%v", variant, align, unroll), NOSPLIT, "func(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)")

	alpha := Load(Param("alpha"), XMM())

	xs := Mem{Base: Load(Param("xs"), GP64())}
	incx := Load(Param("incx"), GP64())

	ys := Mem{Base: Load(Param("ys"), GP64())}
	incy := Load(Param("incy"), GP64())

	n := Load(Param("n"), GP64())

	xi, yi := GP64(), GP64()
	XORQ(xi, xi)
	XORQ(yi, yi)

	JMP(LabelRef("check_limit_unroll"))

	MISALIGN(align)
	Label("loop_unroll")
	{
		tmp := make([]reg.VecVirtual, unroll)
		for u := range tmp {
			tmp[u] = XMM()
		}

		for u := 0; u < unroll; u++ {
			MOVSS(xs.Idx(xi, 4), tmp[u])
			ADDQ(incx, xi)
		}
		for u := 0; u < unroll; u++ {
			MULSS(alpha, tmp[u])
		}
		for u := 0; u < unroll; u++ {
			ADDSS(ys.Idx(yi, 4), tmp[u])
			MOVSS(tmp[u], ys.Idx(yi, 4))
			ADDQ(incy, yi)
		}

		SUBQ(Imm(uint64(unroll)), n)

		Label("check_limit_unroll")

		CMPQ(n, U8(unroll))
		JHS(LabelRef("loop_unroll"))
	}

	JMP(LabelRef("check_limit"))
	Label("loop")
	{
		tmp := XMM()
		MOVSS(xs.Idx(xi, 4), tmp)
		MULSS(alpha, tmp)
		ADDSS(ys.Idx(yi, 4), tmp)
		MOVSS(tmp, ys.Idx(yi, 4))

		DECQ(n)
		ADDQ(incx, xi)
		ADDQ(incy, yi)

		Label("check_limit")

		CMPQ(n, U8(0))
		JHI(LabelRef("loop"))
	}

	RET()
}

func MISALIGN(n int) {
	if n == 0 {
		return
	}

	nearestPowerOf2 := 8
	for n >= nearestPowerOf2*2 {
		nearestPowerOf2 *= 2
	}
	if nearestPowerOf2 >= 8 {
		Instruction(&ir.Instruction{
			Opcode:   "PCALIGN",
			Operands: []Op{Imm(uint64(nearestPowerOf2))},
		})
		n -= nearestPowerOf2
	}

	for i := 0; i < n; i++ {
		NOP()
	}
}
