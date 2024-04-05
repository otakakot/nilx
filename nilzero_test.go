package nilex_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/otakakot/nilex"
)

func TestNilZeroInt(t *testing.T) {
	v := new(int)

	got := nilex.NilZero(v)

	if !reflect.DeepEqual(got, 0) {
		t.Errorf("NilZero() = %v, want %v", got, 0)
	}

	*v = 5

	got = nilex.NilZero(v)

	if !reflect.DeepEqual(got, 5) {
		t.Errorf("NilZero() = %v, want %v", got, 5)
	}
}

func TestNilZeroInt16(t *testing.T) {
	v := new(int16)

	got := nilex.NilZero(v)

	if !reflect.DeepEqual(got, int16(0)) {
		t.Errorf("NilZero() = %v, want %v", got, int16(0))
	}

	*v = 5

	got = nilex.NilZero(v)

	if !reflect.DeepEqual(got, int16(5)) {
		t.Errorf("NilZero() = %v, want %v", got, int16(5))
	}
}

func TestNilZeroInt32(t *testing.T) {
	v := new(int32)

	got := nilex.NilZero(v)

	if !reflect.DeepEqual(got, int32(0)) {
		t.Errorf("NilZero() = %v, want %v", got, int32(0))
	}

	*v = 5

	got = nilex.NilZero(v)

	if !reflect.DeepEqual(got, int32(5)) {
		t.Errorf("NilZero() = %v, want %v", got, int32(5))
	}
}

func TestNilZeroInt64(t *testing.T) {
	v := new(int64)

	got := nilex.NilZero(v)

	if !reflect.DeepEqual(got, int64(0)) {
		t.Errorf("NilZero() = %v, want %v", got, int64(0))
	}

	*v = 5

	got = nilex.NilZero(v)

	if !reflect.DeepEqual(got, int64(5)) {
		t.Errorf("NilZero() = %v, want %v", got, int64(5))
	}
}

func TestNilZeroUint(t *testing.T) {
	v := new(uint)

	got := nilex.NilZero(v)

	if !reflect.DeepEqual(got, uint(0)) {
		t.Errorf("NilZero() = %v, want %v", got, uint(0))
	}

	*v = 5

	got = nilex.NilZero(v)

	if !reflect.DeepEqual(got, uint(5)) {
		t.Errorf("NilZero() = %v, want %v", got, uint(5))
	}
}

func TestNilZeroUint16(t *testing.T) {
	v := new(uint16)

	got := nilex.NilZero(v)

	if !reflect.DeepEqual(got, uint16(0)) {
		t.Errorf("NilZero() = %v, want %v", got, uint16(0))
	}

	*v = 5

	got = nilex.NilZero(v)

	if !reflect.DeepEqual(got, uint16(5)) {
		t.Errorf("NilZero() = %v, want %v", got, uint16(5))
	}
}

func TestNilZeroUint32(t *testing.T) {
	v := new(uint32)

	got := nilex.NilZero(v)

	if !reflect.DeepEqual(got, uint32(0)) {
		t.Errorf("NilZero() = %v, want %v", got, uint32(0))
	}

	*v = 5

	got = nilex.NilZero(v)

	if !reflect.DeepEqual(got, uint32(5)) {
		t.Errorf("NilZero() = %v, want %v", got, uint32(5))
	}
}

func TestNilZeroUint64(t *testing.T) {
	v := new(uint64)

	got := nilex.NilZero(v)

	if !reflect.DeepEqual(got, uint64(0)) {
		t.Errorf("NilZero() = %v, want %v", got, uint64(0))
	}

	*v = 5

	got = nilex.NilZero(v)

	if !reflect.DeepEqual(got, uint64(5)) {
		t.Errorf("NilZero() = %v, want %v", got, uint64(5))
	}
}

func TestNilZeroFloat32(t *testing.T) {
	v := new(float32)

	got := nilex.NilZero(v)

	if !reflect.DeepEqual(got, float32(0)) {
		t.Errorf("NilZero() = %v, want %v", got, float32(0))
	}

	*v = 5.5

	got = nilex.NilZero(v)

	if !reflect.DeepEqual(got, float32(5.5)) {
		t.Errorf("NilZero() = %v, want %v", got, float32(5.5))
	}
}

func TestNilZeroFloat64(t *testing.T) {
	v := new(float64)

	got := nilex.NilZero(v)

	if !reflect.DeepEqual(got, float64(0)) {
		t.Errorf("NilZero() = %v, want %v", got, float64(0))
	}

	*v = 5.5

	got = nilex.NilZero(v)

	if !reflect.DeepEqual(got, float64(5.5)) {
		t.Errorf("NilZero() = %v, want %v", got, float64(5.5))
	}
}

func TestNilZeroString(t *testing.T) {
	v := new(string)

	got := nilex.NilZero(v)

	if !reflect.DeepEqual(got, "") {
		t.Errorf("NilZero() = %v, want %v", got, "")
	}

	*v = "test"

	got = nilex.NilZero(v)

	if !reflect.DeepEqual(got, "test") {
		t.Errorf("NilZero() = %v, want %v", got, "test")
	}
}

func TestNilZeroBool(t *testing.T) {
	v := new(bool)

	got := nilex.NilZero(v)

	if !reflect.DeepEqual(got, false) {
		t.Errorf("NilZero() = %v, want %v", got, false)
	}

	*v = true

	got = nilex.NilZero(v)

	if !reflect.DeepEqual(got, true) {
		t.Errorf("NilZero() = %v, want %v", got, true)
	}
}

func TestNilZeroComplex64(t *testing.T) {
	v := new(complex64)

	got := nilex.NilZero(v)

	if !reflect.DeepEqual(got, complex64(0)) {
		t.Errorf("NilZero() = %v, want %v", got, complex64(0))
	}

	*v = 5 + 5i

	got = nilex.NilZero(v)

	if !reflect.DeepEqual(got, complex64(5+5i)) {
		t.Errorf("NilZero() = %v, want %v", got, complex64(5+5i))
	}
}

func TestNilZeroComplex128(t *testing.T) {
	v := new(complex128)

	got := nilex.NilZero(v)

	if !reflect.DeepEqual(got, complex128(0)) {
		t.Errorf("NilZero() = %v, want %v", got, complex128(0))
	}

	*v = 5 + 5i

	got = nilex.NilZero(v)

	if !reflect.DeepEqual(got, complex128(5+5i)) {
		t.Errorf("NilZero() = %v, want %v", got, complex128(5+5i))
	}
}

func TestNilZeroByte(t *testing.T) {
	v := new(byte)

	got := nilex.NilZero(v)

	if !reflect.DeepEqual(got, byte(0)) {
		t.Errorf("NilZero() = %v, want %v", got, byte(0))
	}

	*v = byte('a')

	got = nilex.NilZero(v)

	if !reflect.DeepEqual(got, byte('a')) {
		t.Errorf("NilZero() = %v, want %v", got, byte('a'))
	}
}

func TestNilZeroRune(t *testing.T) {
	v := new(rune)

	got := nilex.NilZero(v)

	if !reflect.DeepEqual(got, rune(0)) {
		t.Errorf("NilZero() = %v, want %v", got, rune(0))
	}

	*v = 'a'

	got = nilex.NilZero(v)

	if !reflect.DeepEqual(got, 'a') {
		t.Errorf("NilZero() = %v, want %v", got, rune('a'))
	}
}

func TestNilZeroStruct(t *testing.T) {
	type TestStruct struct {
		Field string
	}

	v := new(TestStruct)

	got := nilex.NilZero(v)

	if !reflect.DeepEqual(got, TestStruct{}) {
		t.Errorf("NilZero() = %v, want %v", got, TestStruct{})
	}

	*v = TestStruct{
		Field: "test",
	}

	got = nilex.NilZero(v)

	if !reflect.DeepEqual(got, TestStruct{Field: "test"}) {
		t.Errorf("NilZero() = %v, want %v", got, TestStruct{Field: "test"})
	}
}

func TestNilZeroArray(t *testing.T) {
	v := new([2]int)

	got := nilex.NilZero(v)

	if !reflect.DeepEqual(got, [2]int{}) {
		t.Errorf("NilZero() = %v, want %v", got, [2]int{})
	}

	*v = [2]int{1, 2}

	got = nilex.NilZero(v)

	if !reflect.DeepEqual(got, [2]int{1, 2}) {
		t.Errorf("NilZero() = %v, want %v", got, [2]int{1, 2})
	}
}

func TestNilZeroTime(t *testing.T) {
	v := new(time.Time)

	got := nilex.NilZero(v)

	if !reflect.DeepEqual(got, time.Time{}) {
		t.Errorf("NilZero() = %v, want %v", got, time.Time{})
	}

	*v = time.Now()

	got = nilex.NilZero(v)

	if !reflect.DeepEqual(got, *v) {
		t.Errorf("NilZero() = %v, want %v", got, *v)
	}
}
