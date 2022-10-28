package main

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/rmcaxoul/grpc_server/calculator"
	mock_calculator "github.com/rmcaxoul/grpc_server/mocks"
)

// Parsing with a positive and negative int.
func TestParseInts_OK(t *testing.T) {
	a, b, err := parseInts("1", "-4")
	if err != nil || a != 1 || b != -4 {
		t.Errorf("parseInt with a=1, b=-4")
	}
}

// Cannot parse a float.
func TestParseInt_Err_Float(t *testing.T) {
	a, _, err := parseInts("1.5", "4")
	if err == nil {
		t.Errorf("parsed a float 1.5 to %d", a)
	}
}

// Cannot parse a string.
func TestParseInt_Err_String(t *testing.T) {
	a, _, err := parseInts("1string", "4")
	if err == nil {
		t.Errorf("parsed a string 1string to %d", a)
	}
}

// The input cannot be empty.
func TestParseInt_Err_Empty(t *testing.T) {
	a, _, err := parseInts("", "4")
	if err == nil {
		t.Errorf("parsed an empty string to %d", a)
	}
}

func TestMakeRequest_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	calcMock := mock_calculator.NewMockCalculatorClient(ctrl)

	req := &calculator.OpRequest{
		First:  1,
		Second: 2,
	}

	resp := &calculator.OpResponse{
		Result: 3,
	}

	calcMock.EXPECT().Add(gomock.Any(), gomock.Eq(req)).Return(resp, nil).Times(1)

	makeRequest(calcMock, 1, 2, "+")
}
