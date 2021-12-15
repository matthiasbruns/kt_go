package standard

import (
	"testing"
)

func TestLet(t *testing.T) {
	input := "something great"
	expected := "test"
	result := let[string, string](&input, func(t *string) string {
		return expected
	})

	if result != expected {
		t.Fatalf("result '%s' is not expected '%s", result, expected)
	}
}

type AlsoTextStruct struct {
	Value1 int
	Value2 int
}

func TestAlso(t *testing.T) {
	expected := AlsoTextStruct{
		Value1: 1337,
		Value2: 2000,
	}

	result := also[AlsoTextStruct](&AlsoTextStruct{
		Value1: 1337,
		Value2: 10,
	}, func(t *AlsoTextStruct) {
		t.Value2 = 2000
	})

	if *result != expected {
		t.Fatalf("result '%v' is not expected '%v", result, expected)
	}
}

func TestTakeIf(t *testing.T) {
	expected := "take me"

	result := takeIf[string](&expected, func(t *string) bool { return true })
	if *result != expected {
		t.Fatalf("take me '%v' is not expected '%v", result, expected)
	}

	expected = "dont take me"

	result = takeIf[string](&expected, func(t *string) bool { return false })
	if result != nil {
		t.Fatalf("dont take me '%v' is not expected '%v", result, expected)
	}
}

func TestTakeUnless(t *testing.T) {
	expected := "take me"

	result := takeUnless[string](&expected, func(t *string) bool { return false })
	if *result != expected {
		t.Fatalf("take me '%v' is not expected '%v", result, expected)
	}

	expected = "dont take me"

	result = takeUnless[string](&expected, func(t *string) bool { return true })
	if result != nil {
		t.Fatalf("dont take me '%v' is not expected '%v", result, expected)
	}
}

func TestRepeat(t *testing.T) {
	var repeated []int
	times := 10
	repeat(times, func(i int) {
		repeated = append(repeated, i)
	})

	if len(repeated) != times {
		t.Fatalf("expected repeat count %d - result %d", times, len(repeated))
	}
}
