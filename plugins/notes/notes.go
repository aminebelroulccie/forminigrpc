package notes

import (
	"context"
	"fmt"
	"grpc/proto"
	"io"
	"time"
)

type Notes struct {
	proto.UnimplementedNoteServiceServer
}

func (p *Notes) CreateNotesUnary(ctx context.Context, req *proto.CreateNote) (*proto.CreateNoteRes, error) {
	fmt.Printf("title %s , description %s \n", req.Title, req.Description)
	return &proto.CreateNoteRes{
		Id: 1,
	}, nil
}

func (p *Notes) CreateNotesClientStream(stream proto.NoteService_CreateNotesClientStreamServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(msg)
	}
	stream.SendAndClose(&proto.CreateNoteRes{
		Id: 1,
	})
	return nil
}

func (p *Notes) CreateNotesServerStream(req *proto.CreateNote, stream proto.NoteService_CreateNotesServerStreamServer) error {
	fmt.Printf("title: %s , description: %s \n", req.Title, req.Description)
	for i := 0; i < 10; i++ {
		stream.Send(&proto.CreateNoteRes{
			Id: int32(i),
		})
		time.Sleep(time.Second * 3)
	}
	return nil
}

func (p *Notes) CreateNotesBdirectional(stream proto.NoteService_CreateNotesBdirectionalServer) error {
	for i := 0; i < 10; i++ {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("title: %s , description: %s \n", res.Title, res.Description)
		err = stream.Send(&proto.CreateNoteRes{
			Id: int32(i),
		})
		if err != nil {
			return err
		}

	}
	return nil
}
