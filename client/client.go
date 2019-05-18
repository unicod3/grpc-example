package main

import (
	"context"
	"grpc-demo/pcalc"
	"log"

	"google.golang.org/grpc"
)

func main() {

	addr := "localhost:3000"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pcalc.NewCalculatorClient(conn)

	req := &pcalc.CalcRequest{

		X: 5,
		Y: 7,
	}

	res, err := client.Topla(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Result is %d\n", res.GetResult())

}
