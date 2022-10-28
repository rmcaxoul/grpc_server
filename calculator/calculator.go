package calculator

import (
	"errors"
	"log"

	"golang.org/x/net/context"
)

type Server struct{}

// Add two numbers.
func (s *Server) Add(ctx context.Context, req *OpRequest) (*OpResponse, error) {
	log.Printf("Add method called by %s, with parameters %s", ctx, req)

	a, b := req.GetFirst(), req.GetSecond()
	res := a + b
	if (res < 0 && a > 0 && b > 0) || (res > 0 && a < 0 && b < 0) {
		return nil, errors.New("overflow, result non computable")
	}

	return &OpResponse{
		Result: res,
	}, nil
}

// Substratct two numbers.
func (s *Server) Sub(ctx context.Context, req *OpRequest) (*OpResponse, error) {
	log.Printf("Sub method called by %s, with parameters %s", ctx, req)

	a, b := req.GetFirst(), req.GetSecond()
	res := a - b
	if (res < 0 && a > 0 && b < 0) || (res > 0 && a < 0 && b > 0) {
		return nil, errors.New("overflow, result non computable")
	}

	return &OpResponse{
		Result: res,
	}, nil
}

// Multiply two numbers.
func (s *Server) Mult(ctx context.Context, req *OpRequest) (*OpResponse, error) {
	log.Printf("Mult method called by %s, with parameters %s", ctx, req)

	a, b := req.GetFirst(), req.GetSecond()
	res := a * b
	if a != 0 && res/b != a {
		return nil, errors.New("overflow, result non computable")
	}

	return &OpResponse{
		Result: res,
	}, nil
}

// Divide two numbers.
func (s *Server) Div(ctx context.Context, req *DivRequest) (*DivResponse, error) {
	log.Printf("Div method called by %s, with parameters %s", ctx, req)

	a, b := req.GetFirst(), req.GetSecond()
	res := float32(a / b)
	return &DivResponse{
		Result: res,
	}, nil
}
