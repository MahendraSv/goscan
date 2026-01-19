package scanner

import (
	"strings"
	"testing"
)

func TestReadString(t *testing.T) {
	sc := NewFromReader(strings.NewReader("Mahendra\n"))

	result := Read[string](sc, "")

	if result != "Mahendra" {
		t.Fatalf("expected 'Mahendra', got '%s'", result)
	}
}

func TestReadInt(t *testing.T) {
	sc := NewFromReader(strings.NewReader("30\n"))

	result := Read[int](sc, "")

	if result != 30 {
		t.Fatalf("expected 30, got %d", result)
	}
}

func TestReadIntRetry(t *testing.T) {
	sc := NewFromReader(strings.NewReader("abc\n42\n"))

	result := Read[int](sc, "")

	if result != 42 {
		t.Fatalf("expected 42 after retry, got %d", result)
	}
}

func TestReadFloat(t *testing.T) {
	sc := NewFromReader(strings.NewReader("5.75\n"))

	result := Read[float64](sc, "")

	if result != 5.75 {
		t.Fatalf("expected 5.75, got %f", result)
	}
}

func TestReadBoolTrueVariants(t *testing.T) {
	tests := []string{"true", "yes", "y", "1"}

	for _, tt := range tests {
		sc := NewFromReader(strings.NewReader(tt + "\n"))
		result := Read[bool](sc, "")

		if result != true {
			t.Fatalf("expected true for input '%s'", tt)
		}
	}
}

func TestReadBoolFalseVariants(t *testing.T) {
	tests := []string{"false", "no", "n", "0"}

	for _, tt := range tests {
		sc := NewFromReader(strings.NewReader(tt + "\n"))
		result := Read[bool](sc, "")

		if result != false {
			t.Fatalf("expected false for input '%s'", tt)
		}
	}
}

func TestReadBoolRetry(t *testing.T) {
	sc := NewFromReader(strings.NewReader("maybe\nyes\n"))

	result := Read[bool](sc, "")

	if result != true {
		t.Fatalf("expected true after retry, got %v", result)
	}
}

func TestMultipleReads(t *testing.T) {
	input := "Mahendra\n30\ntrue\n"
	sc := NewFromReader(strings.NewReader(input))

	name := Read[string](sc, "")
	age := Read[int](sc, "")
	active := Read[bool](sc, "")

	if name != "Mahendra" || age != 30 || active != true {
		t.Fatalf("unexpected values: %s %d %v", name, age, active)
	}
}
