package calculator

import (
	"context"
	"math"
	"testing"
)

var s = Server{}

// Adds two numbers correctly.
func TestAdd_OK(t *testing.T) {
	res, err := s.Add(context.Background(), &OpRequest{
		First:  12,
		Second: -25,
	})

	if res.Result != 12+(-25) || err != nil {
		t.Errorf("Add(12+-25)=%d; want -37", res.Result)
	}
}

// Returns an error as the result has an overflow.
func TestAdd_Err_Overflow(t *testing.T) {
	res, err := s.Add(context.Background(), &OpRequest{
		First:  math.MaxInt32,
		Second: 1,
	})
	if res != nil && err.Error() != "overflow, result non computable" {
		t.Errorf("Add(MaxInt32 + 1)=%d", res.Result)
	}
}

// Subtracts two numbers correctly.
func TestSub_OK(t *testing.T) {
	res, err := s.Sub(context.Background(), &OpRequest{
		First:  12,
		Second: -25,
	})

	if res.Result != 12+25 || err != nil {
		t.Errorf("Sub(12--25)=%d; want 13", res.Result)
	}
}

// Returns an error as the result has an overflow.
func TestSub_Err_Overflow(t *testing.T) {
	res, err := s.Sub(context.Background(), &OpRequest{
		First:  math.MinInt32,
		Second: 1,
	})
	if res != nil && err.Error() != "overflow, result non computable" {
		t.Errorf("Sub(MinInt32 - 1)=%d", res.Result)
	}
}

// Multiplies two numbers correctly.
func TestMult_OK(t *testing.T) {
	res, err := s.Mult(context.Background(), &OpRequest{
		First:  12,
		Second: -25,
	})
	println(res.Result)
	if res.Result != 12*-25 || err != nil {
		t.Errorf("Mult(12*-25)=%d; want -300", res.Result)
	}
}

// Returns an error as the result has an overflow.
func TestMult_Err_Overflow(t *testing.T) {
	res, err := s.Mult(context.Background(), &OpRequest{
		First:  math.MaxInt32,
		Second: 2,
	})
	if res != nil && err.Error() != "overflow, result non computable" {
		t.Errorf("Mult(MaxInt32 *2)=%d", res.Result)
	}
}

func TestDiv_OK(t *testing.T) {
	res, err := s.Div(context.Background(), &DivRequest{
		First:  12,
		Second: 7,
	})
	println(res.Result)
	if res.Result != 12/7 || err != nil {
		t.Errorf("Div(12/7)=%f; want 1.71428571429", res.Result)
	}
}
