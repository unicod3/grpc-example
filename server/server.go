package main

import (
	"context"
	"grpc-demo/pcalc"
	"log"
	"net"

	"google.golang.org/grpc"
)

type calcServer struct{}

func (cs *calcServer) Topla(ctx context.Context, req *pcalc.CalcRequest) (*pcalc.CalcResponse, error) {

	log.Println(req)

	x := req.GetX()
	y := req.GetY()

	result := x + y

	res := &pcalc.CalcResponse{
		Result: result,
	}

	return res, nil
}

func main() {

	addr := "localhost:3000"

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal(err)
	}

	c := &calcServer{}

	gs := grpc.NewServer()

	pcalc.RegisterCalculatorServer(gs, c)

	log.Printf("Listening at %s\n", addr)

	if err := gs.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
