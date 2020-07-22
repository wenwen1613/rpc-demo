package simple_client

import (
	"context"
	"google.golang.org/grpc"
	"log"

	pb "oops.com/rpc-demo/proto"
)

const PORT = "9001"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)

	resp, err := client.Search(context.Background(), &pb.SearchRequest{Request: "gRpc test"})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}
	log.Printf("lee resp: %s", resp.GetResponse())

}
