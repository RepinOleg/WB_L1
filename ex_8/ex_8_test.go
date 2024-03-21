package main

import (
	"math"
	"testing"
)

func TestSetBit(t *testing.T) {
	pos := 1
	val := 1
	var number int64 = 32
	SetBit(pos, val, &number)
	expected := int64(34)
	if number != expected {
		t.Errorf("Result was incorrect, got: %d want: %d.", number, expected)
	}
}

func TestSetBit2(t *testing.T) {
	pos := 3
	val := 0
	var number int64 = 8
	SetBit(pos, val, &number)
	expected := int64(0)
	if number != expected {
		t.Errorf("Result was incorrect, got: %d want: %d.", number, expected)
	}
}

func TestSetBit3(t *testing.T) {
	pos := 0
	val := 1
	var number int64 = 0
	SetBit(pos, val, &number)
	expected := int64(1)
	if number != expected {
		t.Errorf("Result was incorrect, got: %d want: %d.", number, expected)
	}
}

func TestSetBit4(t *testing.T) {
	pos := 62
	val := 0
	var number int64 = math.MaxInt64
	SetBit(pos, val, &number)
	expected := int64(math.MaxInt64 >> 1)
	if number != expected {
		t.Errorf("Result was incorrect, got: %d want: %d.", number, expected)
	}
}
