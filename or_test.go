package nilx_test

import (
	"testing"
	"time"

	"github.com/otakakot/nilx"
)

func TestOrInt(t *testing.T) {
	t.Parallel()

	val1 := 1

	val2 := 2

	got := nilx.Or(nil, &val1, &val2)

	if got != val1 {
		t.Errorf("Or() = %v, want %v", got, 1)
	}

	if got := nilx.Or[int](nil, nil, nil); got != int(0) {
		t.Errorf("Or() = %v, want %v", got, 0)
	}
}

func TestOrInt16(t *testing.T) {
	t.Parallel()

	val1 := int16(1)

	val2 := int16(2)

	got := nilx.Or(nil, &val1, &val2)

	if got != val1 {
		t.Errorf("Or() = %v, want %v", got, int16(1))
	}

	if got := nilx.Or[int16](nil, nil, nil); got != int16(0) {
		t.Errorf("Or() = %v, want %v", got, int16(0))
	}
}

func TestOrInt32(t *testing.T) {
	t.Parallel()

	val1 := int32(1)

	val2 := int32(2)

	got := nilx.Or(nil, &val1, &val2)

	if got != val1 {
		t.Errorf("Or() = %v, want %v", got, int32(1))
	}

	if got := nilx.Or[int32](nil, nil, nil); got != int32(0) {
		t.Errorf("Or() = %v, want %v", got, int32(0))
	}
}

func TestOrInt64(t *testing.T) {
	t.Parallel()

	val1 := int64(1)

	val2 := int64(2)

	got := nilx.Or(nil, &val1, &val2)

	if got != val1 {
		t.Errorf("Or() = %v, want %v", got, int64(1))
	}

	if got := nilx.Or[int64](nil, nil, nil); got != int64(0) {
		t.Errorf("Or() = %v, want %v", got, int64(0))
	}
}

func TestOrUint(t *testing.T) {
	t.Parallel()

	val1 := uint(1)

	val2 := uint(2)

	got := nilx.Or(nil, &val1, &val2)

	if got != val1 {
		t.Errorf("Or() = %v, want %v", got, uint(1))
	}

	if got := nilx.Or[uint](nil, nil, nil); got != uint(0) {
		t.Errorf("Or() = %v, want %v", got, uint(0))
	}
}

func TestOrUint16(t *testing.T) {
	t.Parallel()

	val1 := uint16(1)

	val2 := uint16(2)

	got := nilx.Or(nil, &val1, &val2)

	if got != val1 {
		t.Errorf("Or() = %v, want %v", got, uint16(1))
	}

	if got := nilx.Or[uint16](nil, nil, nil); got != uint16(0) {
		t.Errorf("Or() = %v, want %v", got, uint16(0))
	}
}

func TestOrUint32(t *testing.T) {
	t.Parallel()

	val1 := uint32(1)

	val2 := uint32(2)

	got := nilx.Or(nil, &val1, &val2)

	if got != val1 {
		t.Errorf("Or() = %v, want %v", got, uint32(1))
	}

	if got := nilx.Or[uint32](nil, nil, nil); got != uint32(0) {
		t.Errorf("Or() = %v, want %v", got, uint32(0))
	}
}

func TestOrUint64(t *testing.T) {
	t.Parallel()

	val1 := uint64(1)

	val2 := uint64(2)

	got := nilx.Or(nil, &val1, &val2)

	if got != val1 {
		t.Errorf("Or() = %v, want %v", got, uint64(1))
	}

	if got := nilx.Or[uint64](nil, nil, nil); got != uint64(0) {
		t.Errorf("Or() = %v, want %v", got, uint64(0))
	}
}

func TestOrFloat32(t *testing.T) {
	t.Parallel()

	val1 := float32(1)

	val2 := float32(2)

	got := nilx.Or(nil, &val1, &val2)

	if got != val1 {
		t.Errorf("Or() = %v, want %v", got, float32(1))
	}

	if got := nilx.Or[float32](nil, nil, nil); got != float32(0) {
		t.Errorf("Or() = %v, want %v", got, float32(0))
	}
}

func TestOrFloat64(t *testing.T) {
	t.Parallel()

	val1 := float64(1)

	val2 := float64(2)

	got := nilx.Or(nil, &val1, &val2)

	if got != val1 {
		t.Errorf("Or() = %v, want %v", got, float64(1))
	}

	if got := nilx.Or[float64](nil, nil, nil); got != float64(0) {
		t.Errorf("Or() = %v, want %v", got, float64(0))
	}
}

func TestOrString(t *testing.T) {
	t.Parallel()

	val1 := "test1"

	val2 := "test2"

	got := nilx.Or(nil, &val1, &val2)

	if got != val1 {
		t.Errorf("Or() = %v, want %v", got, "test1")
	}

	if got := nilx.Or[string](nil, nil, nil); got != "" {
		t.Errorf("Or() = %v, want %v", got, "")
	}
}

func TestOrBool(t *testing.T) {
	t.Parallel()

	val1 := true

	val2 := false

	got := nilx.Or(nil, &val1, &val2)

	if got != val1 {
		t.Errorf("Or() = %v, want %v", got, true)
	}

	if got := nilx.Or[bool](nil, nil, nil); got != false {
		t.Errorf("Or() = %v, want %v", got, false)
	}
}

func TestOrComplex64(t *testing.T) {
	t.Parallel()

	val1 := complex64(1)

	val2 := complex64(2)

	got := nilx.Or(nil, &val1, &val2)

	if got != val1 {
		t.Errorf("Or() = %v, want %v", got, complex64(1))
	}

	if got := nilx.Or[complex64](nil, nil, nil); got != complex64(0) {
		t.Errorf("Or() = %v, want %v", got, complex64(0))
	}
}

func TestOrComplex128(t *testing.T) {
	t.Parallel()

	val1 := complex128(1)

	val2 := complex128(2)

	got := nilx.Or(nil, &val1, &val2)

	if got != val1 {
		t.Errorf("Or() = %v, want %v", got, complex128(1))
	}

	if got := nilx.Or[complex128](nil, nil, nil); got != complex128(0) {
		t.Errorf("Or() = %v, want %v", got, complex128(0))
	}
}

func TestOrByte(t *testing.T) {
	t.Parallel()

	val1 := byte(1)

	val2 := byte(2)

	got := nilx.Or(nil, &val1, &val2)

	if got != val1 {
		t.Errorf("Or() = %v, want %v", got, byte(1))
	}

	if got := nilx.Or[byte](nil, nil, nil); got != byte(0) {
		t.Errorf("Or() = %v, want %v", got, byte(0))
	}
}

func TestOrRune(t *testing.T) {
	t.Parallel()

	val1 := rune(1)

	val2 := rune(2)

	got := nilx.Or(nil, &val1, &val2)

	if got != val1 {
		t.Errorf("Or() = %v, want %v", got, rune(1))
	}

	if got := nilx.Or[rune](nil, nil, nil); got != rune(0) {
		t.Errorf("Or() = %v, want %v", got, rune(0))
	}
}

func TestOrStruct(t *testing.T) {
	t.Parallel()

	type TestStruct struct {
		Name string
	}

	val1 := TestStruct{
		Name: "test1",
	}

	val2 := TestStruct{
		Name: "test2",
	}

	got := nilx.Or(nil, &val1, &val2)

	if got != val1 {
		t.Errorf("Or() = %v, want %v", got, TestStruct{Name: "test1"})
	}

	if got := nilx.Or[TestStruct](nil, nil, nil); got != (TestStruct{}) {
		t.Errorf("Or() = %v, want %v", got, TestStruct{})
	}
}

func TestOrArray(t *testing.T) {
	t.Parallel()

	val1 := [2]int{1, 1}

	val2 := [2]int{2, 2}

	got := nilx.Or(nil, &val1, &val2)

	if got != val1 {
		t.Errorf("Or() = %v, want %v", got, [2]int{1, 1})
	}

	if got := nilx.Or[[2]int](nil, nil, nil); got != [2]int{} {
		t.Errorf("Or() = %v, want %v", got, [2]int{})
	}
}

func TestOrTime(t *testing.T) {
	t.Parallel()

	val1 := time.Now()

	val2 := time.Now().Add(1)

	got := nilx.Or(nil, &val1, &val2)

	if got != val1 {
		t.Errorf("Or() = %v, want %v", got, val1)
	}

	if got := nilx.Or[time.Time](nil, nil, nil); !got.IsZero() {
		t.Errorf("Or() = %v, want %v", got, time.Time{})
	}
}
