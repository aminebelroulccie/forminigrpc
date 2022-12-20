package main

import (
	"context"
	"fmt"
	"grpc/proto"
	"io"
	"log"
	"time"

	//"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	channal := proto.NewNoteServiceClient(conn)
	// res, err := channal.CreateNotesUnary(context.TODO(), &proto.CreateNote{
	// 	Title:       "devops training",
	// 	Description: "it's training",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("res id: %d\n", res.Id)
	// stream,err := channal.CreateNotesClientStream(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for i := 0; i < 10; i++ {
	// 	if err := stream.Send(&proto.CreateNote{
	// 		Title:       fmt.Sprintf("first %d",i),
	// 		Description: fmt.Sprintf("first %d description",i),
	// 	}); err != nil {
	// 		log.Println(err)
	// 		continue
	// 	}
	// 	fmt.Println("still waiting for server")
	// 	time.Sleep(time.Second*3)

	// }
	// res, err := stream.CloseAndRecv()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("res : %+v\n", res)
	// stream,err := channal.CreateNotesServerStream(context.TODO(),&proto.CreateNote{
	// 	Title: "devops training",
	// 	Description: "it's devops description",
	// })
	// for {
	// 	msg, err := stream.Recv()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("msg id: %d\n",msg.Id)
	// }
	stream, err := channal.CreateNotesBdirectional(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i< 10; i++{
		if err := stream.Send(&proto.CreateNote{
					Title:       fmt.Sprintf("first %d",i),
					Description: fmt.Sprintf("first %d description",i),
				}); err != nil {
					log.Println(err)
					continue
				}
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d\n", msg.Id)
		time.Sleep(time.Second* 3)
		
	}
	
}
