package vecprocess

// AsmAxpyInc is
//
//	for i := 0; i < int(n); i++ {
//		y[iy] += alpha * x[ix]
//		ix += incX
//		iy += incY
//	}
func AsmAxpyInc(alpha float32, x, y []float32, n, incX, incY, ix, iy uintptr)

func AsmAxpyPointer_Align11(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
func AsmAxpyPointer_Align16(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)

func AsmAxpyPointerLoop_Align11(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
func AsmAxpyPointerLoop_Align16(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)

func AsmAxpyUnsafe_Align11(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
func AsmAxpyUnsafe_Align16(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)