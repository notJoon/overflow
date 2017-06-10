package overflow

import (
	"math"
	"testing"
)
import "fmt"

func TestOperations32(t *testing.T) {
	table := []struct {
		f      string
		a, b   int32
		ok     bool
		result int32
	}{
		{"+", math.MaxInt32, 1, false, 0},
		{"+", math.MaxInt32, -1, true, math.MaxInt32 - 1},
		{"+", 0, 0, true, 0},
		{"+", 1234, 5678, true, 6912},
		{"+", math.MaxInt32, math.MaxInt32, false, 0},
		{"+", math.MinInt32 + 100, 101, true, math.MinInt32 + 201},

		{"-", math.MaxInt32, 1, true, math.MaxInt32 - 1},
		{"-", math.MaxInt32, -1, false, 0},
		{"-", 0, math.MinInt32, false, 0},
		{"-", 0, 0, true, 0},
		{"-", 1234, 5678, true, -4444},
		{"-", math.MaxInt32, math.MaxInt32, true, 0},
		{"-", math.MinInt32 + 100, 101, false, 0},

		{"*", math.MaxInt32, 1, true, math.MaxInt32},
		{"*", math.MaxInt32, -1, true, math.MinInt32 + 1},
		{"*", 0, 0, true, 0},
		{"*", 1234, 5678, true, 7006652},
		{"*", math.MaxInt32, 2, false, 0},
		{"*", 0xFFFF, 0xFFFF, false, 0},

		{"/", math.MaxInt32, 1, true, math.MaxInt32},
		{"/", math.MaxInt32, -1, true, math.MinInt32 + 1},
		{"/", 0, 0, false, 0},
		{"/", 1234, 5678, true, 0},
		{"/", 5678, 1234, true, 4},
		{"/", math.MaxInt32, 2, true, 1073741823},
		{"/", 0xFFFF, 0xFFFF, true, 1},
		{"/", math.MinInt32, math.MaxInt32, true, -1},
		{"/", math.MinInt32, -1, false, 0},
	}

	symbolicName := func(n int32) string {
		if n == math.MinInt32 {
			return "math.MinInt32"
		}
		if n == math.MaxInt32 {
			return "math.MaxInt32"
		}
		return fmt.Sprint(n)
	}

	for _, entry := range table {
		var f func(int32, int32) (int32, bool)
		switch entry.f {
		case "+":
			f = Add32
		case "-":
			f = Sub32
		case "*":
			f = Mul32
		case "/":
			f = Div32
		}
		r, ok := f(entry.a, entry.b)
		if r != entry.result || ok != entry.ok {
			t.Errorf("(%v %v %v) -> (%v, %v), expected (%v, %v)",
				symbolicName(entry.a),
				entry.f,
				symbolicName(entry.b),
				r, ok,
				symbolicName(entry.result),
				entry.ok)

		}
	}
}

func TestOperations64(t *testing.T) {
	table := []struct {
		f      string
		a, b   int64
		ok     bool
		result int64
	}{
		{"+", math.MaxInt64, 1, false, 0},
		{"+", math.MaxInt64, -1, true, math.MaxInt64 - 1},
		{"+", 0, 0, true, 0},
		{"+", 1234, 5678, true, 6912},
		{"+", math.MaxInt64, math.MaxInt64, false, 0},
		{"+", math.MinInt64 + 100, 101, true, math.MinInt64 + 201},

		{"-", math.MaxInt64, 1, true, math.MaxInt64 - 1},
		{"-", math.MaxInt64, -1, false, 0},
		{"-", 0, math.MinInt64, false, 0},
		{"-", 0, 0, true, 0},
		{"-", 1234, 5678, true, -4444},
		{"-", math.MaxInt64, math.MaxInt64, true, 0},
		{"-", math.MinInt64 + 100, 101, false, 0},

		{"*", math.MaxInt64, 1, true, math.MaxInt64},
		{"*", math.MaxInt64, -1, true, math.MinInt64 + 1},
		{"*", 0, 0, true, 0},
		{"*", 1234, 5678, true, 7006652},
		{"*", math.MaxInt64, 2, false, 0},
		{"*", 0xFFFFFFFF, 0xFFFFFFFF, false, 0},

		{"/", math.MaxInt64, 1, true, math.MaxInt64},
		{"/", math.MaxInt64, -1, true, math.MinInt64 + 1},
		{"/", 0, 0, false, 0},
		{"/", 1234, 5678, true, 0},
		{"/", 5678, 1234, true, 4},
		{"/", math.MaxInt64, 2, true, 4611686018427387903},
		{"/", 0xFFFF, 0xFFFF, true, 1},
		{"/", math.MinInt64, math.MaxInt64, true, -1},
		{"/", math.MinInt64, -1, false, 0},
	}

	symbolicName := func(n int64) string {
		if n == math.MinInt64 {
			return "math.MinInt64"
		}
		if n == math.MaxInt64 {
			return "math.MaxInt64"
		}
		return fmt.Sprint(n)
	}

	for _, entry := range table {
		var f func(int64, int64) (int64, bool)
		switch entry.f {
		case "+":
			f = Add64
		case "-":
			f = Sub64
		case "*":
			f = Mul64
		case "/":
			f = Div64
		}
		r, ok := f(entry.a, entry.b)
		if r != entry.result || ok != entry.ok {
			t.Errorf("(%v %v %v) -> (%v, %v), expected (%v, %v)",
				symbolicName(entry.a),
				entry.f,
				symbolicName(entry.b),
				r, ok,
				symbolicName(entry.result),
				entry.ok)

		}
	}
}

func TestOperationsInt(t *testing.T) {
	table := []struct {
		f      string
		a, b   int
		ok     bool
		result int
	}{
		{"+", 0, 0, true, 0},
		{"+", 1234, 5678, true, 6912},

		{"-", 0, 0, true, 0},
		{"-", 1234, 5678, true, -4444},

		{"*", 0, 0, true, 0},
		{"*", 1234, 5678, true, 7006652},

		{"/", math.MaxInt32, 1, true, math.MaxInt32},
		{"/", math.MaxInt32, -1, true, math.MinInt32 + 1},
		{"/", 0, 0, false, 0},
		{"/", 1234, 5678, true, 0},
		{"/", 5678, 1234, true, 4},
		{"/", math.MaxInt32, 2, true, 1073741823},
		{"/", 0xFFFF, 0xFFFF, true, 1},
		{"/", math.MinInt32, math.MaxInt32, true, -1},
	}

	symbolicName := func(n int) string {
		if n == math.MinInt32 {
			return "math.MinInt32"
		}
		if n == math.MaxInt32 {
			return "math.MaxInt32"
		}
		return fmt.Sprint(n)
	}

	for _, entry := range table {
		var f func(int, int) (int, bool)
		switch entry.f {
		case "+":
			f = Add
		case "-":
			f = Sub
		case "*":
			f = Mul
		case "/":
			f = Div
		}
		r, ok := f(entry.a, entry.b)
		if r != entry.result || ok != entry.ok {
			t.Errorf("(%v %v %v) -> (%v, %v), expected (%v, %v)",
				symbolicName(entry.a),
				entry.f,
				symbolicName(entry.b),
				r, ok,
				symbolicName(entry.result),
				entry.ok)

		}
	}
}

func TestQuotient(t *testing.T) {
	q, r, ok := Quotient(100, 3)
	if r != 1 || q != 33 || !ok {
		t.Errorf("expected 100/3 => 33, r=1")
	}
	if _, _, ok = Quotient(1, 0); ok {
		t.Error("unexpected lack of failure")
	}
}

func TestAdditionInt(t *testing.T) {
	fmt.Printf("maxint32 = %v\n", math.MaxInt32)
	fmt.Printf("minint32 = %v\n", math.MinInt32)
	fmt.Printf("maxint64 = %v\n", math.MaxInt64)
	fmt.Printf("minint64 = %v\n", math.MinInt64)
}

func Test64(t *testing.T) {
	fmt.Println("64bit:", _is64Bit())
}