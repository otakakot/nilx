package nilex_test

import (
	"reflect"
	"testing"

	"github.com/otakakot/nilex"
)

func TestNilDefInt(t *testing.T) {
	got := nilex.NilDef(nil, 1)

	if !reflect.DeepEqual(got, 1) {
		t.Errorf("NilDef() = %v, want %v", got, 1)
	}

	v := new(int)

	*v = 5

	got = nilex.NilDef(v, 1)

	if !reflect.DeepEqual(got, 5) {
		t.Errorf("NilDef() = %v, want %v", got, 5)
	}
}

func TestNilDefInt16(t *testing.T) {
	got := nilex.NilDef(nil, int16(1))

	if !reflect.DeepEqual(got, int16(1)) {
		t.Errorf("NilDef() = %v, want %v", got, int16(0))
	}

	v := new(int16)

	*v = 5

	got = nilex.NilDef(v, 1)

	if !reflect.DeepEqual(got, int16(5)) {
		t.Errorf("NilDef() = %v, want %v", got, int16(5))
	}
}

func TestNilDefInt32(t *testing.T) {
	got := nilex.NilDef(nil, int32(1))

	if !reflect.DeepEqual(got, int32(1)) {
		t.Errorf("NilDef() = %v, want %v", got, int32(1))
	}

	v := new(int32)

	*v = 5

	got = nilex.NilDef(v, 1)

	if !reflect.DeepEqual(got, int32(5)) {
		t.Errorf("NilDef() = %v, want %v", got, int32(5))
	}
}

func TestNilDefInt64(t *testing.T) {
	got := nilex.NilDef(nil, int64(1))

	if !reflect.DeepEqual(got, int64(1)) {
		t.Errorf("NilDef() = %v, want %v", got, int64(1))
	}

	v := new(int64)

	*v = 5

	got = nilex.NilDef(v, 1)

	if !reflect.DeepEqual(got, int64(5)) {
		t.Errorf("NilDef() = %v, want %v", got, int64(5))
	}
}

func TestNilDefUint(t *testing.T) {
	got := nilex.NilDef(nil, uint(1))

	if !reflect.DeepEqual(got, uint(1)) {
		t.Errorf("NilDef() = %v, want %v", got, uint(1))
	}

	v := new(uint)

	*v = 5

	got = nilex.NilDef(v, 1)

	if !reflect.DeepEqual(got, uint(5)) {
		t.Errorf("NilDef() = %v, want %v", got, uint(5))
	}
}

func TestNilDefUint16(t *testing.T) {
	got := nilex.NilDef(nil, uint16(1))

	if !reflect.DeepEqual(got, uint16(1)) {
		t.Errorf("NilDef() = %v, want %v", got, uint16(1))
	}

	v := new(uint16)

	*v = 5

	got = nilex.NilDef(v, 1)

	if !reflect.DeepEqual(got, uint16(5)) {
		t.Errorf("NilDef() = %v, want %v", got, uint16(5))
	}
}

func TestNilDefUint32(t *testing.T) {
	got := nilex.NilDef(nil, uint32(1))

	if !reflect.DeepEqual(got, uint32(1)) {
		t.Errorf("NilDef() = %v, want %v", got, uint32(1))
	}

	v := new(uint32)

	*v = 5

	got = nilex.NilDef(v, 1)

	if !reflect.DeepEqual(got, uint32(5)) {
		t.Errorf("NilDef() = %v, want %v", got, uint32(5))
	}
}

func TestNilDefUint64(t *testing.T) {
	got := nilex.NilDef(nil, uint64(1))

	if !reflect.DeepEqual(got, uint64(1)) {
		t.Errorf("NilDef() = %v, want %v", got, uint64(1))
	}

	v := new(uint64)

	*v = 5

	got = nilex.NilDef(v, 1)

	if !reflect.DeepEqual(got, uint64(5)) {
		t.Errorf("NilDef() = %v, want %v", got, uint64(5))
	}
}

func TestNilDefFloat32(t *testing.T) {
	got := nilex.NilDef(nil, float32(1))

	if !reflect.DeepEqual(got, float32(1)) {
		t.Errorf("NilDef() = %v, want %v", got, float32(1))
	}

	v := new(float32)

	*v = 5.5

	got = nilex.NilDef(v, 1)

	if !reflect.DeepEqual(got, float32(5.5)) {
		t.Errorf("NilDef() = %v, want %v", got, float32(5.5))
	}
}

func TestNilDefFloat64(t *testing.T) {
	got := nilex.NilDef(nil, float64(1))

	if !reflect.DeepEqual(got, float64(1)) {
		t.Errorf("NilDef() = %v, want %v", got, float64(1))
	}

	v := new(float64)

	*v = 5.5

	got = nilex.NilDef(v, 1)

	if !reflect.DeepEqual(got, float64(5.5)) {
		t.Errorf("NilDef() = %v, want %v", got, float64(5.5))
	}
}

func TestNilDefString(t *testing.T) {
	got := nilex.NilDef(nil, "def")

	if !reflect.DeepEqual(got, "def") {
		t.Errorf("NilDef() = %v, want %v", got, "def")
	}

	v := new(string)

	*v = "test"

	got = nilex.NilDef(v, "def")

	if !reflect.DeepEqual(got, "test") {
		t.Errorf("NilDef() = %v, want %v", got, "test")
	}
}

func TestNilDefBool(t *testing.T) {
	got := nilex.NilDef(nil, true)

	if !reflect.DeepEqual(got, true) {
		t.Errorf("NilDef() = %v, want %v", got, true)
	}

	v := new(bool)

	*v = true

	got = nilex.NilDef(v, false)

	if !reflect.DeepEqual(got, true) {
		t.Errorf("NilDef() = %v, want %v", got, true)
	}
}

func TestNilDefComplex64(t *testing.T) {
	got := nilex.NilDef(nil, complex64(1))

	if !reflect.DeepEqual(got, complex64(1)) {
		t.Errorf("NilDef() = %v, want %v", got, complex64(1))
	}

	v := new(complex64)

	*v = 5 + 5i

	got = nilex.NilDef(v, complex64(1))

	if !reflect.DeepEqual(got, complex64(5+5i)) {
		t.Errorf("NilDef() = %v, want %v", got, complex64(5+5i))
	}
}

func TestNilDefComplex128(t *testing.T) {
	got := nilex.NilDef(nil, complex128(1))

	if !reflect.DeepEqual(got, complex128(1)) {
		t.Errorf("NilDef() = %v, want %v", got, complex128(1))
	}

	v := new(complex128)

	*v = 5 + 5i

	got = nilex.NilDef(v, complex128(1))

	if !reflect.DeepEqual(got, complex128(5+5i)) {
		t.Errorf("NilDef() = %v, want %v", got, complex128(5+5i))
	}
}

func TestNilDefByte(t *testing.T) {
	got := nilex.NilDef(nil, byte(1))

	if !reflect.DeepEqual(got, byte(1)) {
		t.Errorf("NilDef() = %v, want %v", got, byte(1))
	}

	v := new(byte)

	*v = byte('a')

	got = nilex.NilDef(v, byte(1))

	if !reflect.DeepEqual(got, byte('a')) {
		t.Errorf("NilDef() = %v, want %v", got, byte('a'))
	}
}

func TestNilDefRune(t *testing.T) {
	got := nilex.NilDef(nil, rune(1))

	if !reflect.DeepEqual(got, rune(1)) {
		t.Errorf("NilDef() = %v, want %v", got, rune(1))
	}

	v := new(rune)

	*v = 'a'

	got = nilex.NilDef(v, rune(1))

	if !reflect.DeepEqual(got, 'a') {
		t.Errorf("NilDef() = %v, want %v", got, rune('a'))
	}
}

func TestNilDefStruct(t *testing.T) {
	type TestStruct struct {
		Field string
	}

	got := nilex.NilDef(nil, TestStruct{Field: "def"})

	if !reflect.DeepEqual(got, TestStruct{Field: "def"}) {
		t.Errorf("NilDef() = %v, want %v", got, TestStruct{Field: "def"})
	}

	v := new(TestStruct)

	*v = TestStruct{
		Field: "test",
	}

	got = nilex.NilDef(v, TestStruct{Field: "def"})

	if !reflect.DeepEqual(got, TestStruct{Field: "test"}) {
		t.Errorf("NilDef() = %v, want %v", got, TestStruct{Field: "test"})
	}
}

func TestNilDefArray(t *testing.T) {
	got := nilex.NilDef(nil, [2]int{1, 1})

	if !reflect.DeepEqual(got, [2]int{1, 1}) {
		t.Errorf("NilDef() = %v, want %v", got, [2]int{1, 1})
	}

	v := new([2]int)

	*v = [2]int{1, 2}

	got = nilex.NilDef(v, [2]int{1, 1})

	if !reflect.DeepEqual(got, [2]int{1, 2}) {
		t.Errorf("NilDef() = %v, want %v", got, [2]int{1, 2})
	}
}
