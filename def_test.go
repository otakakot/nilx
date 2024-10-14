package nilx_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/otakakot/nilx"
)

func TestDefInt(t *testing.T) {
	t.Parallel()

	got := nilx.Def(nil, 1)

	if !reflect.DeepEqual(got, 1) {
		t.Errorf("Def() = %v, want %v", got, 1)
	}

	v := new(int)

	*v = 5

	got = nilx.Def(v, 1)

	if !reflect.DeepEqual(got, 5) {
		t.Errorf("Def() = %v, want %v", got, 5)
	}
}

func TestDefInt16(t *testing.T) {
	t.Parallel()

	got := nilx.Def(nil, int16(1))

	if !reflect.DeepEqual(got, int16(1)) {
		t.Errorf("Def() = %v, want %v", got, int16(0))
	}

	v := new(int16)

	*v = 5

	got = nilx.Def(v, 1)

	if !reflect.DeepEqual(got, int16(5)) {
		t.Errorf("Def() = %v, want %v", got, int16(5))
	}
}

func TestDefInt32(t *testing.T) {
	t.Parallel()

	got := nilx.Def(nil, int32(1))

	if !reflect.DeepEqual(got, int32(1)) {
		t.Errorf("Def() = %v, want %v", got, int32(1))
	}

	v := new(int32)

	*v = 5

	got = nilx.Def(v, 1)

	if !reflect.DeepEqual(got, int32(5)) {
		t.Errorf("Def() = %v, want %v", got, int32(5))
	}
}

func TestDefInt64(t *testing.T) {
	t.Parallel()

	got := nilx.Def(nil, int64(1))

	if !reflect.DeepEqual(got, int64(1)) {
		t.Errorf("Def() = %v, want %v", got, int64(1))
	}

	v := new(int64)

	*v = 5

	got = nilx.Def(v, 1)

	if !reflect.DeepEqual(got, int64(5)) {
		t.Errorf("Def() = %v, want %v", got, int64(5))
	}
}

func TestDefUint(t *testing.T) {
	t.Parallel()

	got := nilx.Def(nil, uint(1))

	if !reflect.DeepEqual(got, uint(1)) {
		t.Errorf("Def() = %v, want %v", got, uint(1))
	}

	v := new(uint)

	*v = 5

	got = nilx.Def(v, 1)

	if !reflect.DeepEqual(got, uint(5)) {
		t.Errorf("Def() = %v, want %v", got, uint(5))
	}
}

func TestDefUint16(t *testing.T) {
	t.Parallel()

	got := nilx.Def(nil, uint16(1))

	if !reflect.DeepEqual(got, uint16(1)) {
		t.Errorf("Def() = %v, want %v", got, uint16(1))
	}

	v := new(uint16)

	*v = 5

	got = nilx.Def(v, 1)

	if !reflect.DeepEqual(got, uint16(5)) {
		t.Errorf("Def() = %v, want %v", got, uint16(5))
	}
}

func TestDefUint32(t *testing.T) {
	t.Parallel()

	got := nilx.Def(nil, uint32(1))

	if !reflect.DeepEqual(got, uint32(1)) {
		t.Errorf("Def() = %v, want %v", got, uint32(1))
	}

	v := new(uint32)

	*v = 5

	got = nilx.Def(v, 1)

	if !reflect.DeepEqual(got, uint32(5)) {
		t.Errorf("Def() = %v, want %v", got, uint32(5))
	}
}

func TestDefUint64(t *testing.T) {
	t.Parallel()

	got := nilx.Def(nil, uint64(1))

	if !reflect.DeepEqual(got, uint64(1)) {
		t.Errorf("Def() = %v, want %v", got, uint64(1))
	}

	v := new(uint64)

	*v = 5

	got = nilx.Def(v, 1)

	if !reflect.DeepEqual(got, uint64(5)) {
		t.Errorf("Def() = %v, want %v", got, uint64(5))
	}
}

func TestDefFloat32(t *testing.T) {
	t.Parallel()

	got := nilx.Def(nil, float32(1))

	if !reflect.DeepEqual(got, float32(1)) {
		t.Errorf("Def() = %v, want %v", got, float32(1))
	}

	v := new(float32)

	*v = 5.5

	got = nilx.Def(v, 1)

	if !reflect.DeepEqual(got, float32(5.5)) {
		t.Errorf("Def() = %v, want %v", got, float32(5.5))
	}
}

func TestDefFloat64(t *testing.T) {
	t.Parallel()

	got := nilx.Def(nil, float64(1))

	if !reflect.DeepEqual(got, float64(1)) {
		t.Errorf("Def() = %v, want %v", got, float64(1))
	}

	v := new(float64)

	*v = 5.5

	got = nilx.Def(v, 1)

	if !reflect.DeepEqual(got, float64(5.5)) {
		t.Errorf("Def() = %v, want %v", got, float64(5.5))
	}
}

func TestDefString(t *testing.T) {
	t.Parallel()

	got := nilx.Def(nil, "def")

	if !reflect.DeepEqual(got, "def") {
		t.Errorf("Def() = %v, want %v", got, "def")
	}

	v := new(string)

	*v = "test"

	got = nilx.Def(v, "def")

	if !reflect.DeepEqual(got, "test") {
		t.Errorf("Def() = %v, want %v", got, "test")
	}
}

func TestDefBool(t *testing.T) {
	t.Parallel()

	got := nilx.Def(nil, true)

	if !reflect.DeepEqual(got, true) {
		t.Errorf("Def() = %v, want %v", got, true)
	}

	v := new(bool)

	*v = true

	got = nilx.Def(v, false)

	if !reflect.DeepEqual(got, true) {
		t.Errorf("Def() = %v, want %v", got, true)
	}
}

func TestDefComplex64(t *testing.T) {
	t.Parallel()

	got := nilx.Def(nil, complex64(1))

	if !reflect.DeepEqual(got, complex64(1)) {
		t.Errorf("Def() = %v, want %v", got, complex64(1))
	}

	v := new(complex64)

	*v = 5 + 5i

	got = nilx.Def(v, complex64(1))

	if !reflect.DeepEqual(got, complex64(5+5i)) {
		t.Errorf("Def() = %v, want %v", got, complex64(5+5i))
	}
}

func TestDefComplex128(t *testing.T) {
	t.Parallel()

	got := nilx.Def(nil, complex128(1))

	if !reflect.DeepEqual(got, complex128(1)) {
		t.Errorf("Def() = %v, want %v", got, complex128(1))
	}

	v := new(complex128)

	*v = 5 + 5i

	got = nilx.Def(v, complex128(1))

	if !reflect.DeepEqual(got, complex128(5+5i)) {
		t.Errorf("Def() = %v, want %v", got, complex128(5+5i))
	}
}

func TestDefByte(t *testing.T) {
	t.Parallel()

	got := nilx.Def(nil, byte(1))

	if !reflect.DeepEqual(got, byte(1)) {
		t.Errorf("Def() = %v, want %v", got, byte(1))
	}

	v := new(byte)

	*v = byte('a')

	got = nilx.Def(v, byte(1))

	if !reflect.DeepEqual(got, byte('a')) {
		t.Errorf("Def() = %v, want %v", got, byte('a'))
	}
}

func TestDefRune(t *testing.T) {
	t.Parallel()

	got := nilx.Def(nil, rune(1))

	if !reflect.DeepEqual(got, rune(1)) {
		t.Errorf("Def() = %v, want %v", got, rune(1))
	}

	v := new(rune)

	*v = 'a'

	got = nilx.Def(v, rune(1))

	if !reflect.DeepEqual(got, 'a') {
		t.Errorf("Def() = %v, want %v", got, rune('a'))
	}
}

func TestDefStruct(t *testing.T) {
	t.Parallel()

	type TestStruct struct {
		Field string
	}

	got := nilx.Def(nil, TestStruct{Field: "def"})

	if !reflect.DeepEqual(got, TestStruct{Field: "def"}) {
		t.Errorf("Def() = %v, want %v", got, TestStruct{Field: "def"})
	}

	v := new(TestStruct)

	*v = TestStruct{
		Field: "test",
	}

	got = nilx.Def(v, TestStruct{Field: "def"})

	if !reflect.DeepEqual(got, TestStruct{Field: "test"}) {
		t.Errorf("Def() = %v, want %v", got, TestStruct{Field: "test"})
	}
}

func TestDefArray(t *testing.T) {
	t.Parallel()

	got := nilx.Def(nil, [2]int{1, 1})

	if !reflect.DeepEqual(got, [2]int{1, 1}) {
		t.Errorf("Def() = %v, want %v", got, [2]int{1, 1})
	}

	v := new([2]int)

	*v = [2]int{1, 2}

	got = nilx.Def(v, [2]int{1, 1})

	if !reflect.DeepEqual(got, [2]int{1, 2}) {
		t.Errorf("Def() = %v, want %v", got, [2]int{1, 2})
	}
}

func TestDefTime(t *testing.T) {
	t.Parallel()

	now := time.Now()

	got := nilx.Def(nil, now)

	if !reflect.DeepEqual(got, now) {
		t.Errorf("Def() = %v, want %v", got, now)
	}

	v := new(time.Time)

	*v = time.Now()

	got = nilx.Def(v, now)

	if !reflect.DeepEqual(got, *v) {
		t.Errorf("Def() = %v, want %v", got, *v)
	}
}
