package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "graduate/api/album/service/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:8001", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAlbumClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// in := &pb.ListAlbumReq{PageNum: 1, PageSize: 10}
	// r, err := c.ListAlbum(ctx, in)
	// if err != nil {
	// 	log.Fatalf("could not ListAlbum: %v", err)
	// }

	// fmt.Println("count:", r.Count)
	// for _, o := range r.Albums {
	// 	fmt.Printf("Id:%d,Title:%s,Artist:%s,CreateAt:%s\r\n", o.Id, o.Title, o.Artist, o.CreateAt)
	// }

	in := &pb.GetAlbumByIdReq{Id: 1}
	o, err := c.GetAlbumById(ctx, in)
	if err != nil {
		log.Fatalf("could not GetAlbumByIdReq: %v", err)
	}
	fmt.Printf("Id:%d,Title:%s,Artist:%s,CreateAt:%s\r\n", o.Album.Id, o.Album.Title, o.Album.Artist, o.Album.CreateAt)

}
