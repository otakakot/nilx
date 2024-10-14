package nilx_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/otakakot/nilx"
)

func TestZeroInt(t *testing.T) {
	t.Parallel()

	got := nilx.Zero[int](nil)

	if !reflect.DeepEqual(got, 0) {
		t.Errorf("Zero() = %v, want %v", got, 0)
	}

	v := 5

	got = nilx.Zero(&v)

	if !reflect.DeepEqual(got, 5) {
		t.Errorf("Zero() = %v, want %v", got, 5)
	}
}

func TestZeroInt16(t *testing.T) {
	t.Parallel()

	got := nilx.Zero[int16](nil)

	if !reflect.DeepEqual(got, int16(0)) {
		t.Errorf("Zero() = %v, want %v", got, int16(0))
	}

	v := int16(5)

	got = nilx.Zero(&v)

	if !reflect.DeepEqual(got, int16(5)) {
		t.Errorf("Zero() = %v, want %v", got, int16(5))
	}
}

func TestZeroInt32(t *testing.T) {
	t.Parallel()

	got := nilx.Zero[int32](nil)

	if !reflect.DeepEqual(got, int32(0)) {
		t.Errorf("Zero() = %v, want %v", got, int32(0))
	}

	v := int32(5)

	got = nilx.Zero(&v)

	if !reflect.DeepEqual(got, int32(5)) {
		t.Errorf("Zero() = %v, want %v", got, int32(5))
	}
}

func TestZeroInt64(t *testing.T) {
	t.Parallel()

	got := nilx.Zero[int64](nil)

	if !reflect.DeepEqual(got, int64(0)) {
		t.Errorf("Zero() = %v, want %v", got, int64(0))
	}

	v := int64(5)

	got = nilx.Zero(&v)

	if !reflect.DeepEqual(got, int64(5)) {
		t.Errorf("Zero() = %v, want %v", got, int64(5))
	}
}

func TestZeroUint(t *testing.T) {
	t.Parallel()

	got := nilx.Zero[uint](nil)

	if !reflect.DeepEqual(got, uint(0)) {
		t.Errorf("Zero() = %v, want %v", got, uint(0))
	}

	v := uint(5)

	got = nilx.Zero(&v)

	if !reflect.DeepEqual(got, uint(5)) {
		t.Errorf("Zero() = %v, want %v", got, uint(5))
	}
}

func TestZeroUint16(t *testing.T) {
	t.Parallel()

	got := nilx.Zero[uint16](nil)

	if !reflect.DeepEqual(got, uint16(0)) {
		t.Errorf("Zero() = %v, want %v", got, uint16(0))
	}

	v := uint16(5)

	got = nilx.Zero(&v)

	if !reflect.DeepEqual(got, uint16(5)) {
		t.Errorf("Zero() = %v, want %v", got, uint16(5))
	}
}

func TestZeroUint32(t *testing.T) {
	t.Parallel()

	got := nilx.Zero[uint32](nil)

	if !reflect.DeepEqual(got, uint32(0)) {
		t.Errorf("Zero() = %v, want %v", got, uint32(0))
	}

	v := uint32(5)

	got = nilx.Zero(&v)

	if !reflect.DeepEqual(got, uint32(5)) {
		t.Errorf("Zero() = %v, want %v", got, uint32(5))
	}
}

func TestZeroUint64(t *testing.T) {
	t.Parallel()

	got := nilx.Zero[uint64](nil)

	if !reflect.DeepEqual(got, uint64(0)) {
		t.Errorf("Zero() = %v, want %v", got, uint64(0))
	}

	v := uint64(5)

	got = nilx.Zero(&v)

	if !reflect.DeepEqual(got, uint64(5)) {
		t.Errorf("Zero() = %v, want %v", got, uint64(5))
	}
}

func TestZeroFloat32(t *testing.T) {
	t.Parallel()

	got := nilx.Zero[float32](nil)

	if !reflect.DeepEqual(got, float32(0)) {
		t.Errorf("Zero() = %v, want %v", got, float32(0))
	}

	v := float32(5.5)

	got = nilx.Zero(&v)

	if !reflect.DeepEqual(got, float32(5.5)) {
		t.Errorf("Zero() = %v, want %v", got, float32(5.5))
	}
}

func TestZeroFloat64(t *testing.T) {
	t.Parallel()

	got := nilx.Zero[float64](nil)

	if !reflect.DeepEqual(got, float64(0)) {
		t.Errorf("Zero() = %v, want %v", got, float64(0))
	}

	v := float64(5.5)

	got = nilx.Zero(&v)

	if !reflect.DeepEqual(got, float64(5.5)) {
		t.Errorf("Zero() = %v, want %v", got, float64(5.5))
	}
}

func TestZeroString(t *testing.T) {
	t.Parallel()

	got := nilx.Zero[string](nil)

	if !reflect.DeepEqual(got, "") {
		t.Errorf("Zero() = %v, want %v", got, "")
	}

	v := "test"

	got = nilx.Zero(&v)

	if !reflect.DeepEqual(got, "test") {
		t.Errorf("Zero() = %v, want %v", got, "test")
	}
}

func TestZeroBool(t *testing.T) {
	t.Parallel()

	got := nilx.Zero[bool](nil)

	if !reflect.DeepEqual(got, false) {
		t.Errorf("Zero() = %v, want %v", got, false)
	}

	v := true

	got = nilx.Zero(&v)

	if !reflect.DeepEqual(got, true) {
		t.Errorf("Zero() = %v, want %v", got, true)
	}
}

func TestZeroComplex64(t *testing.T) {
	t.Parallel()

	got := nilx.Zero[complex64](nil)

	if !reflect.DeepEqual(got, complex64(0)) {
		t.Errorf("Zero() = %v, want %v", got, complex64(0))
	}

	v := complex64(5 + 5i)

	got = nilx.Zero(&v)

	if !reflect.DeepEqual(got, complex64(5+5i)) {
		t.Errorf("Zero() = %v, want %v", got, complex64(5+5i))
	}
}

func TestZeroComplex128(t *testing.T) {
	t.Parallel()

	got := nilx.Zero[complex128](nil)

	if !reflect.DeepEqual(got, complex128(0)) {
		t.Errorf("Zero() = %v, want %v", got, complex128(0))
	}

	v := 5 + 5i

	got = nilx.Zero(&v)

	if !reflect.DeepEqual(got, complex128(5+5i)) {
		t.Errorf("Zero() = %v, want %v", got, complex128(5+5i))
	}
}

func TestZeroByte(t *testing.T) {
	t.Parallel()

	got := nilx.Zero[byte](nil)

	if !reflect.DeepEqual(got, byte(0)) {
		t.Errorf("Zero() = %v, want %v", got, byte(0))
	}

	v := byte('a')

	got = nilx.Zero(&v)

	if !reflect.DeepEqual(got, byte('a')) {
		t.Errorf("Zero() = %v, want %v", got, byte('a'))
	}
}

func TestZeroRune(t *testing.T) {
	t.Parallel()

	got := nilx.Zero[rune](nil)

	if !reflect.DeepEqual(got, rune(0)) {
		t.Errorf("Zero() = %v, want %v", got, rune(0))
	}

	v := 'a'

	got = nilx.Zero(&v)

	if !reflect.DeepEqual(got, 'a') {
		t.Errorf("Zero() = %v, want %v", got, rune('a'))
	}
}

func TestZeroStruct(t *testing.T) {
	t.Parallel()

	type TestStruct struct {
		Field string
	}

	got := nilx.Zero[TestStruct](nil)

	if !reflect.DeepEqual(got, TestStruct{}) {
		t.Errorf("Zero() = %v, want %v", got, TestStruct{})
	}

	v := TestStruct{
		Field: "test",
	}

	got = nilx.Zero(&v)

	if !reflect.DeepEqual(got, TestStruct{Field: "test"}) {
		t.Errorf("Zero() = %v, want %v", got, TestStruct{Field: "test"})
	}
}

func TestZeroArray(t *testing.T) {
	t.Parallel()

	got := nilx.Zero[[2]int](nil)

	if !reflect.DeepEqual(got, [2]int{}) {
		t.Errorf("Zero() = %v, want %v", got, [2]int{})
	}

	v := [2]int{1, 2}

	got = nilx.Zero(&v)

	if !reflect.DeepEqual(got, [2]int{1, 2}) {
		t.Errorf("Zero() = %v, want %v", got, [2]int{1, 2})
	}
}

func TestZeroTime(t *testing.T) {
	t.Parallel()

	got := nilx.Zero[time.Time](nil)

	if !reflect.DeepEqual(got, time.Time{}) {
		t.Errorf("Zero() = %v, want %v", got, time.Time{})
	}

	v := time.Now()

	got = nilx.Zero(&v)

	if !reflect.DeepEqual(got, v) {
		t.Errorf("Zero() = %v, want %v", got, v)
	}
}
