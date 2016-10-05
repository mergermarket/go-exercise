package misterPointy

import "testing"

func TestAdd(t *testing.T) {
	result := add(3, 5)
	expected := 8

	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestAddPointers(t *testing.T) {
	a := 5
	b := 8
	result := addPointers(&a, &b)
	expected := 13

	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestAddReturnPointer(t *testing.T) {
	resultPointer := addReturnPointer(8, 13)
	result := *resultPointer
	expected := 21

	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestAddPointersReturnPointer(t *testing.T) {
	a := 13
	b := 21
	result := *addPointersReturnPointer(&a, &b)
	expected := 34

	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestAddModifyArgument(t *testing.T) {
	a := 21
	b := 34
	expected := 55
	addModifyArgument(a, &b)

	if b != expected {
		t.Error("Expected ", expected, " but got ", b)
	}
}

func TestAddPointerToPointer(t *testing.T) {
	a := 34
	b := 55
	aPointer := &a
	bPointer := &b
	expected := 89
	result := addPointerToPointer(&aPointer, &bPointer)

	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}
