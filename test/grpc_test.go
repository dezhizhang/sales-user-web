package test

import (
	"context"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
	"log"
	"sales-user-web/proto"
	"testing"
)

func TestGrpc(t *testing.T) {
	//conn, err := grpc.Dial(
	//	"consul://127.0.0.1:8500/user_srv?wait=14s&tag=srv",
	//	grpc.WithInsecure(),
	//	grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer conn.Close()
	conn, err := grpc.Dial(
		"consul://127.0.0.1:8500/user-srv?wait=14s&tag=srv",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewUserClient(conn)
	list, err1 := client.GetUserList(context.Background(), &proto.PageInfo{
		PageSize:  1,
		PageIndex: 10,
	})
	if err1 != nil {
		panic(err1)
	}

	fmt.Println(list)
}
