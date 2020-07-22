package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"

	pb "oops.com/rpc-demo/proto"
)

const PORT = "9002"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	defer conn.Close()

	client := pb.NewStreamServiceClient(conn)
	err = printLists(client, &pb.StreamRequest{
		Pt: &pb.StreamPoint{
			Name:  "gRPC Stream Client: List",
			Value: 2020,
		},
	})
	if err != nil {
		log.Fatalf("printLists err: %v", err)
	}

	err = printRecord(client, &pb.StreamRequest{
		Pt: &pb.StreamPoint{
			Name:  "gRPC Stream Client: Record",
			Value: 2020,
		},
	})
	if err != nil {
		log.Fatalf("printRecord err: %v", err)
	}

	err = printRoute(client, &pb.StreamRequest{
		Pt: &pb.StreamPoint{
			Name:  "gRPC Stream Client: Route",
			Value: 2020,
		},
	})
	if err != nil {
		log.Fatalf("printRoute err: %v", err)
	}

}

func printLists(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.List(context.Background(), r)
	if err != nil {
		return err
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp pj.name: %s, pj.value: %d", resp.Pt.Name, resp.Pt.Value)
	}
	return nil
}

func printRecord(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	return nil
}

func printRoute(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	return nil
}
