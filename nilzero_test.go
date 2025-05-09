package nilx_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/otakakot/nilx"
)

func TestNilZeroInt(t *testing.T) {
	t.Parallel()

	got := nilx.NilZero[int](nil)

	if !reflect.DeepEqual(got, 0) {
		t.Errorf("NilZero() = %v, want %v", got, 0)
	}

	val := 5

	got = nilx.NilZero(&val)

	if !reflect.DeepEqual(got, 5) {
		t.Errorf("NilZero() = %v, want %v", got, 5)
	}
}

func TestNilZeroInt16(t *testing.T) {
	t.Parallel()

	got := nilx.NilZero[int16](nil)

	if !reflect.DeepEqual(got, int16(0)) {
		t.Errorf("NilZero() = %v, want %v", got, int16(0))
	}

	val := int16(5)

	got = nilx.NilZero(&val)

	if !reflect.DeepEqual(got, int16(5)) {
		t.Errorf("NilZero() = %v, want %v", got, int16(5))
	}
}

func TestNilZeroInt32(t *testing.T) {
	t.Parallel()

	got := nilx.NilZero[int32](nil)

	if !reflect.DeepEqual(got, int32(0)) {
		t.Errorf("NilZero() = %v, want %v", got, int32(0))
	}

	val := int32(5)

	got = nilx.NilZero(&val)

	if !reflect.DeepEqual(got, int32(5)) {
		t.Errorf("NilZero() = %v, want %v", got, int32(5))
	}
}

func TestNilZeroInt64(t *testing.T) {
	t.Parallel()

	got := nilx.NilZero[int64](nil)

	if !reflect.DeepEqual(got, int64(0)) {
		t.Errorf("NilZero() = %v, want %v", got, int64(0))
	}

	val := int64(5)

	got = nilx.NilZero(&val)

	if !reflect.DeepEqual(got, int64(5)) {
		t.Errorf("NilZero() = %v, want %v", got, int64(5))
	}
}

func TestNilZeroUint(t *testing.T) {
	t.Parallel()

	got := nilx.NilZero[uint](nil)

	if !reflect.DeepEqual(got, uint(0)) {
		t.Errorf("NilZero() = %v, want %v", got, uint(0))
	}

	val := uint(5)

	got = nilx.NilZero(&val)

	if !reflect.DeepEqual(got, uint(5)) {
		t.Errorf("NilZero() = %v, want %v", got, uint(5))
	}
}

func TestNilZeroUint16(t *testing.T) {
	t.Parallel()

	got := nilx.NilZero[uint16](nil)

	if !reflect.DeepEqual(got, uint16(0)) {
		t.Errorf("NilZero() = %v, want %v", got, uint16(0))
	}

	v := uint16(5)

	got = nilx.NilZero(&v)

	if !reflect.DeepEqual(got, uint16(5)) {
		t.Errorf("NilZero() = %v, want %v", got, uint16(5))
	}
}

func TestNilZeroUint32(t *testing.T) {
	t.Parallel()

	got := nilx.NilZero[uint32](nil)

	if !reflect.DeepEqual(got, uint32(0)) {
		t.Errorf("NilZero() = %v, want %v", got, uint32(0))
	}

	val := uint32(5)

	got = nilx.NilZero(&val)

	if !reflect.DeepEqual(got, uint32(5)) {
		t.Errorf("NilZero() = %v, want %v", got, uint32(5))
	}
}

func TestNilZeroUint64(t *testing.T) {
	t.Parallel()

	got := nilx.NilZero[uint64](nil)

	if !reflect.DeepEqual(got, uint64(0)) {
		t.Errorf("NilZero() = %v, want %v", got, uint64(0))
	}

	val := uint64(5)

	got = nilx.NilZero(&val)

	if !reflect.DeepEqual(got, uint64(5)) {
		t.Errorf("NilZero() = %v, want %v", got, uint64(5))
	}
}

func TestNilZeroFloat32(t *testing.T) {
	t.Parallel()

	got := nilx.NilZero[float32](nil)

	if !reflect.DeepEqual(got, float32(0)) {
		t.Errorf("NilZero() = %v, want %v", got, float32(0))
	}

	val := float32(5.5)

	got = nilx.NilZero(&val)

	if !reflect.DeepEqual(got, float32(5.5)) {
		t.Errorf("NilZero() = %v, want %v", got, float32(5.5))
	}
}

func TestNilZeroFloat64(t *testing.T) {
	t.Parallel()

	got := nilx.NilZero[float64](nil)

	if !reflect.DeepEqual(got, float64(0)) {
		t.Errorf("NilZero() = %v, want %v", got, float64(0))
	}

	val := float64(5.5)

	got = nilx.NilZero(&val)

	if !reflect.DeepEqual(got, float64(5.5)) {
		t.Errorf("NilZero() = %v, want %v", got, float64(5.5))
	}
}

func TestNilZeroString(t *testing.T) {
	t.Parallel()

	got := nilx.NilZero[string](nil)

	if !reflect.DeepEqual(got, "") {
		t.Errorf("NilZero() = %v, want %v", got, "")
	}

	val := "test"

	got = nilx.NilZero(&val)

	if !reflect.DeepEqual(got, "test") {
		t.Errorf("NilZero() = %v, want %v", got, "test")
	}
}

func TestNilZeroBool(t *testing.T) {
	t.Parallel()

	got := nilx.NilZero[bool](nil)

	if !reflect.DeepEqual(got, false) {
		t.Errorf("NilZero() = %v, want %v", got, false)
	}

	val := true

	got = nilx.NilZero(&val)

	if !reflect.DeepEqual(got, true) {
		t.Errorf("NilZero() = %v, want %v", got, val)
	}
}

func TestNilZeroComplex64(t *testing.T) {
	t.Parallel()

	got := nilx.NilZero[complex64](nil)

	if !reflect.DeepEqual(got, complex64(0)) {
		t.Errorf("NilZero() = %v, want %v", got, complex64(0))
	}

	val := complex64(5 + 5i)

	got = nilx.NilZero(&val)

	if !reflect.DeepEqual(got, complex64(5+5i)) {
		t.Errorf("NilZero() = %v, want %v", got, val)
	}
}

func TestNilZeroComplex128(t *testing.T) {
	t.Parallel()

	got := nilx.NilZero[complex128](nil)

	if !reflect.DeepEqual(got, complex128(0)) {
		t.Errorf("NilZero() = %v, want %v", got, complex128(0))
	}

	val := 5 + 5i

	got = nilx.NilZero(&val)

	if !reflect.DeepEqual(got, complex128(5+5i)) {
		t.Errorf("NilZero() = %v, want %v", got, complex128(5+5i))
	}
}

func TestNilZeroByte(t *testing.T) {
	t.Parallel()

	got := nilx.NilZero[byte](nil)

	if !reflect.DeepEqual(got, byte(0)) {
		t.Errorf("NilZero() = %v, want %v", got, byte(0))
	}

	val := byte('a')

	got = nilx.NilZero(&val)

	if !reflect.DeepEqual(got, byte('a')) {
		t.Errorf("NilZero() = %v, want %v", got, byte('a'))
	}
}

func TestNilZeroRune(t *testing.T) {
	t.Parallel()

	got := nilx.NilZero[rune](nil)

	if !reflect.DeepEqual(got, rune(0)) {
		t.Errorf("NilZero() = %v, want %v", got, rune(0))
	}

	val := 'a'

	got = nilx.NilZero(&val)

	if !reflect.DeepEqual(got, 'a') {
		t.Errorf("NilZero() = %v, want %v", got, 'a')
	}
}

func TestNilZeroStruct(t *testing.T) {
	t.Parallel()

	type TestStruct struct {
		Field string
	}

	got := nilx.NilZero[TestStruct](nil)

	if !reflect.DeepEqual(got, TestStruct{}) {
		t.Errorf("NilZero() = %v, want %v", got, TestStruct{})
	}

	val := TestStruct{
		Field: "test",
	}

	got = nilx.NilZero(&val)

	if !reflect.DeepEqual(got, TestStruct{Field: "test"}) {
		t.Errorf("NilZero() = %v, want %v", got, TestStruct{Field: "test"})
	}
}

func TestNilZeroArray(t *testing.T) {
	t.Parallel()

	got := nilx.NilZero[[2]int](nil)

	if !reflect.DeepEqual(got, [2]int{}) {
		t.Errorf("NilZero() = %v, want %v", got, [2]int{})
	}

	val := [2]int{1, 2}

	got = nilx.NilZero(&val)

	if !reflect.DeepEqual(got, [2]int{1, 2}) {
		t.Errorf("NilZero() = %v, want %v", got, [2]int{1, 2})
	}
}

func TestNilZeroTime(t *testing.T) {
	t.Parallel()

	got := nilx.NilZero[time.Time](nil)

	if !reflect.DeepEqual(got, time.Time{}) {
		t.Errorf("NilZero() = %v, want %v", got, time.Time{})
	}

	val := time.Now()

	got = nilx.NilZero(&val)

	if !reflect.DeepEqual(got, val) {
		t.Errorf("NilZero() = %v, want %v", got, val)
	}
}
