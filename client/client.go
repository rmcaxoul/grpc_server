package main

import (
	"errors"
	"flag"
	"log"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/rmcaxoul/grpc_server/calculator"
)

// Transform two strings respectively into two int32.
// Error is returned if one of the two conversions fails.
func parseInts(a string, b string) (int32, int32, error) {
	if a == "" || b == "" {
		return 0, 0, errors.New("need to input two numbers")
	}

	inta, err := strconv.ParseInt(a, 10, 32)
	if err != nil {
		return 0, 0, err
	}

	intb, err := strconv.ParseInt(b, 10, 32)
	if err != nil {
		return 0, 0, err
	}

	return int32(inta), int32(intb), nil

}

// Depending on the operator, the correct endpoint of the server will be called.
// Returns an error if the operator is incorrect or empty.
func makeRequest(c calculator.CalculatorClient, a int32, b int32, op string) error {
	switch op {
	case "+":
		res, err := c.Add(context.Background(), &calculator.OpRequest{
			First:  a,
			Second: b,
		})
		if err != nil {
			return err
		}
		log.Printf("%d + %d = %d", a, b, res.GetResult())
	case "-":
		res, err := c.Sub(context.Background(), &calculator.OpRequest{
			First:  a,
			Second: b,
		})
		if err != nil {
			return err
		}
		log.Printf("%d - %d = %d", a, b, res.GetResult())
	case "*":
		res, err := c.Mult(context.Background(), &calculator.OpRequest{
			First:  a,
			Second: b,
		})
		if err != nil {
			return err
		}
		log.Printf("%d * %d = %d", a, b, res.GetResult())
	case "/":
		if b == 0 {
			return errors.New("second number cannot be zero for division")
		}
		res, err := c.Div(context.Background(), &calculator.DivRequest{
			First:  a,
			Second: b,
		})
		if err != nil {
			return err
		}
		log.Printf("%d * %d = %f", a, b, res.GetResult())
	default:
		return errors.New("method could not be parsed; should be +,-,* or /")
	}
	return nil
}

// Initialize the RPC and analyse arguments in order to compute the operation wanted.
func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	defer conn.Close()

	c := calculator.NewCalculatorClient(conn)

	var a, b, method string
	flag.StringVar(&a, "a", "", "an int")
	flag.StringVar(&b, "b", "", "an int")
	flag.StringVar(&method, "method", "", "An operator (+,-,* or /)")
	flag.Parse()
	inta, intb, err := parseInts(a, b)
	if err != nil {
		log.Fatalf("error: %s", err)
		return
	}

	err = makeRequest(c, inta, intb, method)
	if err != nil {
		log.Fatalf("error: %s", err)
		return
	}
}
