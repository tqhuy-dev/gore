package utilities

// T2 creates a tuple from a list of values.
func T2[T0, T1 any](a T0, b T1) Tuple2[T0, T1] {
	return Tuple2[T0, T1]{T0: a, T1: b}
}

// T3 creates a tuple from a list of values.
func T3[T0, T1, T2 any](a T0, b T1, c T2) Tuple3[T0, T1, T2] {
	return Tuple3[T0, T1, T2]{T0: a, T1: b, T2: c}
}

// T4 creates a tuple from a list of values.
func T4[T0, T1, T2, T3 any](a T0, b T1, c T2, d T3) Tuple4[T0, T1, T2, T3] {
	return Tuple4[T0, T1, T2, T3]{T0: a, T1: b, T2: c, T3: d}
}

// T5 creates a tuple from a list of values.
func T5[T0, T1, T2, T3, T4 any](a T0, b T1, c T2, d T3, e T4) Tuple5[T0, T1, T2, T3, T4] {
	return Tuple5[T0, T1, T2, T3, T4]{T0: a, T1: b, T2: c, T3: d, T4: e}
}

// T6 creates a tuple from a list of values.
func T6[T0, T1, T2, T3, T4, T5 any](a T0, b T1, c T2, d T3, e T4, f T5) Tuple6[T0, T1, T2, T3, T4, T5] {
	return Tuple6[T0, T1, T2, T3, T4, T5]{T0: a, T1: b, T2: c, T3: d, T4: e, T5: f}
}

// T7 creates a tuple from a list of values.
func T7[T0, T1, T2, T3, T4, T5, T6 any](a T0, b T1, c T2, d T3, e T4, f T5, g T6) Tuple7[T0, T1, T2, T3, T4, T5, T6] {
	return Tuple7[T0, T1, T2, T3, T4, T5, T6]{T0: a, T1: b, T2: c, T3: d, T4: e, T5: f, T6: g}
}

// T8 creates a tuple from a list of values.
func T8[T0, T1, T2, T3, T4, T5, T6, T7 any](a T0, b T1, c T2, d T3, e T4, f T5, g T6, h T7) Tuple8[T0, T1, T2, T3, T4, T5, T6, T7] {
	return Tuple8[T0, T1, T2, T3, T4, T5, T6, T7]{T0: a, T1: b, T2: c, T3: d, T4: e, T5: f, T6: g, T7: h}
}

// T9 creates a tuple from a list of values.
func T9[T0, T1, T2, T3, T4, T5, T6, T7, T8 any](a T0, b T1, c T2, d T3, e T4, f T5, g T6, h T7, i T8) Tuple9[T0, T1, T2, T3, T4, T5, T6, T7, T8] {
	return Tuple9[T0, T1, T2, T3, T4, T5, T6, T7, T8]{T0: a, T1: b, T2: c, T3: d, T4: e, T5: f, T6: g, T7: h, T8: i}
}

// Unpack2 returns values contained in tuple.
func Unpack2[T0, T1 any](tuple Tuple2[T0, T1]) (T0, T1) {
	return tuple.T0, tuple.T1
}

// Unpack3 returns values contained in tuple.
func Unpack3[T0, T1, T2 any](tuple Tuple3[T0, T1, T2]) (T0, T1, T2) {
	return tuple.T0, tuple.T1, tuple.T2
}

// Unpack4 returns values contained in tuple.
func Unpack4[T0, T1, T2, T3 any](tuple Tuple4[T0, T1, T2, T3]) (T0, T1, T2, T3) {
	return tuple.T0, tuple.T1, tuple.T2, tuple.T3
}

// Unpack5 returns values contained in tuple.
func Unpack5[T0, T1, T2, T3, T4 any](tuple Tuple5[T0, T1, T2, T3, T4]) (T0, T1, T2, T3, T4) {
	return tuple.T0, tuple.T1, tuple.T2, tuple.T3, tuple.T4
}

// Unpack6 returns values contained in tuple.
// Play: https://go.dev/play/p/xVP_k0kJ96W
func Unpack6[T0, T1, T2, T3, T4, T5 any](tuple Tuple6[T0, T1, T2, T3, T4, T5]) (T0, T1, T2, T3, T4, T5) {
	return tuple.T0, tuple.T1, tuple.T2, tuple.T3, tuple.T4, tuple.T5
}

// Unpack7 returns values contained in tuple.
func Unpack7[T0, T1, T2, T3, T4, T5, T6 any](tuple Tuple7[T0, T1, T2, T3, T4, T5, T6]) (T0, T1, T2, T3, T4, T5, T6) {
	return tuple.T0, tuple.T1, tuple.T2, tuple.T3, tuple.T4, tuple.T5, tuple.T6
}

// Unpack8 returns values contained in tuple.
func Unpack8[T0, T1, T2, T3, T4, T5, T6, T7 any](tuple Tuple8[T0, T1, T2, T3, T4, T5, T6, T7]) (T0, T1, T2, T3, T4, T5, T6, T7) {
	return tuple.T0, tuple.T1, tuple.T2, tuple.T3, tuple.T4, tuple.T5, tuple.T6, tuple.T7
}

// Unpack9 returns values contained in tuple.
func Unpack9[T0, T1, T2, T3, T4, T5, T6, T7, T8 any](tuple Tuple9[T0, T1, T2, T3, T4, T5, T6, T7, T8]) (T0, T1, T2, T3, T4, T5, T6, T7, T8) {
	return tuple.T0, tuple.T1, tuple.T2, tuple.T3, tuple.T4, tuple.T5, tuple.T6, tuple.T7, tuple.T8
}

func Zip2[T0, T1 any](a []T0, b []T1) []Tuple2[T0, T1] {
	size := Max([]int{len(a), len(b)})

	result := make([]Tuple2[T0, T1], 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)

		result = append(result, Tuple2[T0, T1]{
			T0: _a,
			T1: _b,
		})
	}

	return result
}

func Zip3[T0, T1, T2 any](a []T0, b []T1, c []T2) []Tuple3[T0, T1, T2] {
	size := Max([]int{len(a), len(b), len(c)})

	result := make([]Tuple3[T0, T1, T2], 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)

		result = append(result, Tuple3[T0, T1, T2]{
			T0: _a,
			T1: _b,
			T2: _c,
		})
	}

	return result
}

// Zip4 creates a slice of grouped elements, the first of which contains the first elements
// of the given arrays, the second of which contains the second elements of the given arrays, and so on.
// When collections have different size, the Tuple attributes are filled with zero value.
// Play: https://go.dev/play/p/jujaT06T6aJTp
func Zip4[T0, T1, T2, T3 any](a []T0, b []T1, c []T2, d []T3) []Tuple4[T0, T1, T2, T3] {
	size := Max([]int{len(a), len(b), len(c), len(d)})

	result := make([]Tuple4[T0, T1, T2, T3], 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)
		_d, _ := Nth(d, index)

		result = append(result, Tuple4[T0, T1, T2, T3]{
			T0: _a,
			T1: _b,
			T2: _c,
			T3: _d,
		})
	}

	return result
}

func Zip5[T0, T1, T2, T3, T4 any](a []T0, b []T1, c []T2, d []T3, e []T4) []Tuple5[T0, T1, T2, T3, T4] {
	size := Max([]int{len(a), len(b), len(c), len(d), len(e)})

	result := make([]Tuple5[T0, T1, T2, T3, T4], 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)
		_d, _ := Nth(d, index)
		_e, _ := Nth(e, index)

		result = append(result, Tuple5[T0, T1, T2, T3, T4]{
			T0: _a,
			T1: _b,
			T2: _c,
			T3: _d,
			T4: _e,
		})
	}

	return result
}

func Zip6[T0, T1, T2, T3, T4, T5 any](a []T0, b []T1, c []T2, d []T3, e []T4, f []T5) []Tuple6[T0, T1, T2, T3, T4, T5] {
	size := Max([]int{len(a), len(b), len(c), len(d), len(e), len(f)})

	result := make([]Tuple6[T0, T1, T2, T3, T4, T5], 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)
		_d, _ := Nth(d, index)
		_e, _ := Nth(e, index)
		_f, _ := Nth(f, index)

		result = append(result, Tuple6[T0, T1, T2, T3, T4, T5]{
			T0: _a,
			T1: _b,
			T2: _c,
			T3: _d,
			T4: _e,
			T5: _f,
		})
	}

	return result
}

func Zip7[T0, T1, T2, T3, T4, T5, T6 any](a []T0, b []T1, c []T2, d []T3, e []T4, f []T5, g []T6) []Tuple7[T0, T1, T2, T3, T4, T5, T6] {
	size := Max([]int{len(a), len(b), len(c), len(d), len(e), len(f), len(g)})

	result := make([]Tuple7[T0, T1, T2, T3, T4, T5, T6], 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)
		_d, _ := Nth(d, index)
		_e, _ := Nth(e, index)
		_f, _ := Nth(f, index)
		_g, _ := Nth(g, index)

		result = append(result, Tuple7[T0, T1, T2, T3, T4, T5, T6]{
			T0: _a,
			T1: _b,
			T2: _c,
			T3: _d,
			T4: _e,
			T5: _f,
			T6: _g,
		})
	}

	return result
}

func Zip8[T0, T1, T2, T3, T4, T5, T6, T7 any](a []T0, b []T1, c []T2, d []T3, e []T4, f []T5, g []T6, h []T7) []Tuple8[T0, T1, T2, T3, T4, T5, T6, T7] {
	size := Max([]int{len(a), len(b), len(c), len(d), len(e), len(f), len(g), len(h)})

	result := make([]Tuple8[T0, T1, T2, T3, T4, T5, T6, T7], 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)
		_d, _ := Nth(d, index)
		_e, _ := Nth(e, index)
		_f, _ := Nth(f, index)
		_g, _ := Nth(g, index)
		_h, _ := Nth(h, index)

		result = append(result, Tuple8[T0, T1, T2, T3, T4, T5, T6, T7]{
			T0: _a,
			T1: _b,
			T2: _c,
			T3: _d,
			T4: _e,
			T5: _f,
			T6: _g,
			T7: _h,
		})
	}

	return result
}

func Zip9[T0, T1, T2, T3, T4, T5, T6, T7, T8 any](a []T0, b []T1, c []T2, d []T3, e []T4, f []T5, g []T6, h []T7, i []T8) []Tuple9[T0, T1, T2, T3, T4, T5, T6, T7, T8] {
	size := Max([]int{len(a), len(b), len(c), len(d), len(e), len(f), len(g), len(h), len(i)})

	result := make([]Tuple9[T0, T1, T2, T3, T4, T5, T6, T7, T8], 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)
		_d, _ := Nth(d, index)
		_e, _ := Nth(e, index)
		_f, _ := Nth(f, index)
		_g, _ := Nth(g, index)
		_h, _ := Nth(h, index)
		_i, _ := Nth(i, index)

		result = append(result, Tuple9[T0, T1, T2, T3, T4, T5, T6, T7, T8]{
			T0: _a,
			T1: _b,
			T2: _c,
			T3: _d,
			T4: _e,
			T5: _f,
			T6: _g,
			T7: _h,
			T8: _i,
		})
	}

	return result
}

func ZipT1y2[T0 any, T1 any, Out any](a []T0, b []T1, iteratee func(a T0, b T1) Out) []Out {
	size := Max([]int{len(a), len(b)})

	result := make([]Out, 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)

		result = append(result, iteratee(_a, _b))
	}

	return result
}

func ZipT1y3[T0 any, T1 any, T2 any, Out any](a []T0, b []T1, c []T2, iteratee func(a T0, b T1, c T2) Out) []Out {
	size := Max([]int{len(a), len(b), len(c)})

	result := make([]Out, 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)

		result = append(result, iteratee(_a, _b, _c))
	}

	return result
}

func ZipT1y4[T0 any, T1 any, T2 any, T3 any, Out any](a []T0, b []T1, c []T2, d []T3, iteratee func(a T0, b T1, c T2, d T3) Out) []Out {
	size := Max([]int{len(a), len(b), len(c), len(d)})

	result := make([]Out, 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)
		_d, _ := Nth(d, index)

		result = append(result, iteratee(_a, _b, _c, _d))
	}

	return result
}

func ZipT1y5[T0 any, T1 any, T2 any, T3 any, T4 any, Out any](a []T0, b []T1, c []T2, d []T3, e []T4, iteratee func(a T0, b T1, c T2, d T3, e T4) Out) []Out {
	size := Max([]int{len(a), len(b), len(c), len(d), len(e)})

	result := make([]Out, 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)
		_d, _ := Nth(d, index)
		_e, _ := Nth(e, index)

		result = append(result, iteratee(_a, _b, _c, _d, _e))
	}

	return result
}

func ZipT1y6[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, Out any](a []T0, b []T1, c []T2, d []T3, e []T4, f []T5, iteratee func(a T0, b T1, c T2, d T3, e T4, f T5) Out) []Out {
	size := Max([]int{len(a), len(b), len(c), len(d), len(e), len(f)})

	result := make([]Out, 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)
		_d, _ := Nth(d, index)
		_e, _ := Nth(e, index)
		_f, _ := Nth(f, index)

		result = append(result, iteratee(_a, _b, _c, _d, _e, _f))
	}

	return result
}

func ZipT1y7[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, Out any](a []T0, b []T1, c []T2, d []T3, e []T4, f []T5, g []T6, iteratee func(a T0, b T1, c T2, d T3, e T4, f T5, g T6) Out) []Out {
	size := Max([]int{len(a), len(b), len(c), len(d), len(e), len(f)})

	result := make([]Out, 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)
		_d, _ := Nth(d, index)
		_e, _ := Nth(e, index)
		_f, _ := Nth(f, index)
		_g, _ := Nth(g, index)

		result = append(result, iteratee(_a, _b, _c, _d, _e, _f, _g))
	}

	return result
}

func ZipT1y8[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, Out any](a []T0, b []T1, c []T2, d []T3, e []T4, f []T5, g []T6, h []T7, iteratee func(a T0, b T1, c T2, d T3, e T4, f T5, g T6, h T7) Out) []Out {
	size := Max([]int{len(a), len(b), len(c), len(d), len(e), len(f), len(g)})

	result := make([]Out, 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)
		_d, _ := Nth(d, index)
		_e, _ := Nth(e, index)
		_f, _ := Nth(f, index)
		_g, _ := Nth(g, index)
		_h, _ := Nth(h, index)

		result = append(result, iteratee(_a, _b, _c, _d, _e, _f, _g, _h))
	}

	return result
}

func ZipT1y9[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, Out any](a []T0, b []T1, c []T2, d []T3, e []T4, f []T5, g []T6, h []T7, i []T8, iteratee func(a T0, b T1, c T2, d T3, e T4, f T5, g T6, h T7, i T8) Out) []Out {
	size := Max([]int{len(a), len(b), len(c), len(d), len(e), len(f), len(g), len(h), len(i)})

	result := make([]Out, 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)
		_d, _ := Nth(d, index)
		_e, _ := Nth(e, index)
		_f, _ := Nth(f, index)
		_g, _ := Nth(g, index)
		_h, _ := Nth(h, index)
		_i, _ := Nth(i, index)

		result = append(result, iteratee(_a, _b, _c, _d, _e, _f, _g, _h, _i))
	}

	return result
}

func Unzip2[T0, T1 any](tuples []Tuple2[T0, T1]) ([]T0, []T1) {
	size := len(tuples)
	r1 := make([]T0, 0, size)
	r2 := make([]T1, 0, size)

	for i := range tuples {
		r1 = append(r1, tuples[i].T0)
		r2 = append(r2, tuples[i].T1)
	}

	return r1, r2
}

func Unzip3[T0, T1, T2 any](tuples []Tuple3[T0, T1, T2]) ([]T0, []T1, []T2) {
	size := len(tuples)
	r1 := make([]T0, 0, size)
	r2 := make([]T1, 0, size)
	r3 := make([]T2, 0, size)

	for i := range tuples {
		r1 = append(r1, tuples[i].T0)
		r2 = append(r2, tuples[i].T1)
		r3 = append(r3, tuples[i].T2)
	}

	return r1, r2, r3
}

func Unzip4[T0, T1, T2, T3 any](tuples []Tuple4[T0, T1, T2, T3]) ([]T0, []T1, []T2, []T3) {
	size := len(tuples)
	r1 := make([]T0, 0, size)
	r2 := make([]T1, 0, size)
	r3 := make([]T2, 0, size)
	r4 := make([]T3, 0, size)

	for i := range tuples {
		r1 = append(r1, tuples[i].T0)
		r2 = append(r2, tuples[i].T1)
		r3 = append(r3, tuples[i].T2)
		r4 = append(r4, tuples[i].T3)
	}

	return r1, r2, r3, r4
}

func Unzip5[T0, T1, T2, T3, T4 any](tuples []Tuple5[T0, T1, T2, T3, T4]) ([]T0, []T1, []T2, []T3, []T4) {
	size := len(tuples)
	r1 := make([]T0, 0, size)
	r2 := make([]T1, 0, size)
	r3 := make([]T2, 0, size)
	r4 := make([]T3, 0, size)
	r5 := make([]T4, 0, size)

	for i := range tuples {
		r1 = append(r1, tuples[i].T0)
		r2 = append(r2, tuples[i].T1)
		r3 = append(r3, tuples[i].T2)
		r4 = append(r4, tuples[i].T3)
		r5 = append(r5, tuples[i].T4)
	}

	return r1, r2, r3, r4, r5
}

func Unzip6[T0, T1, T2, T3, T4, T5 any](tuples []Tuple6[T0, T1, T2, T3, T4, T5]) ([]T0, []T1, []T2, []T3, []T4, []T5) {
	size := len(tuples)
	r1 := make([]T0, 0, size)
	r2 := make([]T1, 0, size)
	r3 := make([]T2, 0, size)
	r4 := make([]T3, 0, size)
	r5 := make([]T4, 0, size)
	r6 := make([]T5, 0, size)

	for i := range tuples {
		r1 = append(r1, tuples[i].T0)
		r2 = append(r2, tuples[i].T1)
		r3 = append(r3, tuples[i].T2)
		r4 = append(r4, tuples[i].T3)
		r5 = append(r5, tuples[i].T4)
		r6 = append(r6, tuples[i].T5)
	}

	return r1, r2, r3, r4, r5, r6
}

func Unzip7[T0, T1, T2, T3, T4, T5, T6 any](tuples []Tuple7[T0, T1, T2, T3, T4, T5, T6]) ([]T0, []T1, []T2, []T3, []T4, []T5, []T6) {
	size := len(tuples)
	r1 := make([]T0, 0, size)
	r2 := make([]T1, 0, size)
	r3 := make([]T2, 0, size)
	r4 := make([]T3, 0, size)
	r5 := make([]T4, 0, size)
	r6 := make([]T5, 0, size)
	r7 := make([]T6, 0, size)

	for i := range tuples {
		r1 = append(r1, tuples[i].T0)
		r2 = append(r2, tuples[i].T1)
		r3 = append(r3, tuples[i].T2)
		r4 = append(r4, tuples[i].T3)
		r5 = append(r5, tuples[i].T4)
		r6 = append(r6, tuples[i].T5)
		r7 = append(r7, tuples[i].T6)
	}

	return r1, r2, r3, r4, r5, r6, r7
}

func Unzip8[T0, T1, T2, T3, T4, T5, T6, T7 any](tuples []Tuple8[T0, T1, T2, T3, T4, T5, T6, T7]) ([]T0, []T1, []T2, []T3, []T4, []T5, []T6, []T7) {
	size := len(tuples)
	r1 := make([]T0, 0, size)
	r2 := make([]T1, 0, size)
	r3 := make([]T2, 0, size)
	r4 := make([]T3, 0, size)
	r5 := make([]T4, 0, size)
	r6 := make([]T5, 0, size)
	r7 := make([]T6, 0, size)
	r8 := make([]T7, 0, size)

	for i := range tuples {
		r1 = append(r1, tuples[i].T0)
		r2 = append(r2, tuples[i].T1)
		r3 = append(r3, tuples[i].T2)
		r4 = append(r4, tuples[i].T3)
		r5 = append(r5, tuples[i].T4)
		r6 = append(r6, tuples[i].T5)
		r7 = append(r7, tuples[i].T6)
		r8 = append(r8, tuples[i].T7)
	}

	return r1, r2, r3, r4, r5, r6, r7, r8
}

func Unzip9[T0, T1, T2, T3, T4, T5, T6, T7, T8 any](tuples []Tuple9[T0, T1, T2, T3, T4, T5, T6, T7, T8]) ([]T0, []T1, []T2, []T3, []T4, []T5, []T6, []T7, []T8) {
	size := len(tuples)
	r1 := make([]T0, 0, size)
	r2 := make([]T1, 0, size)
	r3 := make([]T2, 0, size)
	r4 := make([]T3, 0, size)
	r5 := make([]T4, 0, size)
	r6 := make([]T5, 0, size)
	r7 := make([]T6, 0, size)
	r8 := make([]T7, 0, size)
	r9 := make([]T8, 0, size)

	for i := range tuples {
		r1 = append(r1, tuples[i].T0)
		r2 = append(r2, tuples[i].T1)
		r3 = append(r3, tuples[i].T2)
		r4 = append(r4, tuples[i].T3)
		r5 = append(r5, tuples[i].T4)
		r6 = append(r6, tuples[i].T5)
		r7 = append(r7, tuples[i].T6)
		r8 = append(r8, tuples[i].T7)
		r9 = append(r9, tuples[i].T8)
	}

	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

func UnzipT1y2[T8n any, T0 any, T1 any](items []T8n, iteratee func(T8n) (a T0, b T1)) ([]T0, []T1) {
	size := len(items)
	r1 := make([]T0, 0, size)
	r2 := make([]T1, 0, size)

	for i := range items {
		a, b := iteratee(items[i])
		r1 = append(r1, a)
		r2 = append(r2, b)
	}

	return r1, r2
}

func UnzipT1y3[T8n any, T0 any, T1 any, T2 any](items []T8n, iteratee func(T8n) (a T0, b T1, c T2)) ([]T0, []T1, []T2) {
	size := len(items)
	r1 := make([]T0, 0, size)
	r2 := make([]T1, 0, size)
	r3 := make([]T2, 0, size)

	for i := range items {
		a, b, c := iteratee(items[i])
		r1 = append(r1, a)
		r2 = append(r2, b)
		r3 = append(r3, c)
	}

	return r1, r2, r3
}

func UnzipT1y4[T8n any, T0 any, T1 any, T2 any, T3 any](items []T8n, iteratee func(T8n) (a T0, b T1, c T2, d T3)) ([]T0, []T1, []T2, []T3) {
	size := len(items)
	r1 := make([]T0, 0, size)
	r2 := make([]T1, 0, size)
	r3 := make([]T2, 0, size)
	r4 := make([]T3, 0, size)

	for i := range items {
		a, b, c, d := iteratee(items[i])
		r1 = append(r1, a)
		r2 = append(r2, b)
		r3 = append(r3, c)
		r4 = append(r4, d)
	}

	return r1, r2, r3, r4
}

func UnzipT1y5[T8n any, T0 any, T1 any, T2 any, T3 any, T4 any](items []T8n, iteratee func(T8n) (a T0, b T1, c T2, d T3, e T4)) ([]T0, []T1, []T2, []T3, []T4) {
	size := len(items)
	r1 := make([]T0, 0, size)
	r2 := make([]T1, 0, size)
	r3 := make([]T2, 0, size)
	r4 := make([]T3, 0, size)
	r5 := make([]T4, 0, size)

	for i := range items {
		a, b, c, d, e := iteratee(items[i])
		r1 = append(r1, a)
		r2 = append(r2, b)
		r3 = append(r3, c)
		r4 = append(r4, d)
		r5 = append(r5, e)
	}

	return r1, r2, r3, r4, r5
}

func UnzipT1y6[T8n any, T0 any, T1 any, T2 any, T3 any, T4 any, T5 any](items []T8n, iteratee func(T8n) (a T0, b T1, c T2, d T3, e T4, f T5)) ([]T0, []T1, []T2, []T3, []T4, []T5) {
	size := len(items)
	r1 := make([]T0, 0, size)
	r2 := make([]T1, 0, size)
	r3 := make([]T2, 0, size)
	r4 := make([]T3, 0, size)
	r5 := make([]T4, 0, size)
	r6 := make([]T5, 0, size)

	for i := range items {
		a, b, c, d, e, f := iteratee(items[i])
		r1 = append(r1, a)
		r2 = append(r2, b)
		r3 = append(r3, c)
		r4 = append(r4, d)
		r5 = append(r5, e)
		r6 = append(r6, f)
	}

	return r1, r2, r3, r4, r5, r6
}

func UnzipT1y7[T8n any, T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](items []T8n, iteratee func(T8n) (a T0, b T1, c T2, d T3, e T4, f T5, g T6)) ([]T0, []T1, []T2, []T3, []T4, []T5, []T6) {
	size := len(items)
	r1 := make([]T0, 0, size)
	r2 := make([]T1, 0, size)
	r3 := make([]T2, 0, size)
	r4 := make([]T3, 0, size)
	r5 := make([]T4, 0, size)
	r6 := make([]T5, 0, size)
	r7 := make([]T6, 0, size)

	for i := range items {
		a, b, c, d, e, f, g := iteratee(items[i])
		r1 = append(r1, a)
		r2 = append(r2, b)
		r3 = append(r3, c)
		r4 = append(r4, d)
		r5 = append(r5, e)
		r6 = append(r6, f)
		r7 = append(r7, g)
	}

	return r1, r2, r3, r4, r5, r6, r7
}

func UnzipT1y8[T8n any, T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any](items []T8n, iteratee func(T8n) (a T0, b T1, c T2, d T3, e T4, f T5, g T6, h T7)) ([]T0, []T1, []T2, []T3, []T4, []T5, []T6, []T7) {
	size := len(items)
	r1 := make([]T0, 0, size)
	r2 := make([]T1, 0, size)
	r3 := make([]T2, 0, size)
	r4 := make([]T3, 0, size)
	r5 := make([]T4, 0, size)
	r6 := make([]T5, 0, size)
	r7 := make([]T6, 0, size)
	r8 := make([]T7, 0, size)

	for i := range items {
		a, b, c, d, e, f, g, h := iteratee(items[i])
		r1 = append(r1, a)
		r2 = append(r2, b)
		r3 = append(r3, c)
		r4 = append(r4, d)
		r5 = append(r5, e)
		r6 = append(r6, f)
		r7 = append(r7, g)
		r8 = append(r8, h)
	}

	return r1, r2, r3, r4, r5, r6, r7, r8
}

func UnzipT1y9[T8n any, T0 any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any](items []T8n, iteratee func(T8n) (a T0, b T1, c T2, d T3, e T4, f T5, g T6, h T7, i T8)) ([]T0, []T1, []T2, []T3, []T4, []T5, []T6, []T7, []T8) {
	size := len(items)
	r1 := make([]T0, 0, size)
	r2 := make([]T1, 0, size)
	r3 := make([]T2, 0, size)
	r4 := make([]T3, 0, size)
	r5 := make([]T4, 0, size)
	r6 := make([]T5, 0, size)
	r7 := make([]T6, 0, size)
	r8 := make([]T7, 0, size)
	r9 := make([]T8, 0, size)

	for i := range items {
		a, b, c, d, e, f, g, h, i := iteratee(items[i])
		r1 = append(r1, a)
		r2 = append(r2, b)
		r3 = append(r3, c)
		r4 = append(r4, d)
		r5 = append(r5, e)
		r6 = append(r6, f)
		r7 = append(r7, g)
		r8 = append(r8, h)
		r9 = append(r9, i)
	}

	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}
