package chat

import (
	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *HelloRequest) (*HelloReply,
	error) {
	return &HelloReply{Message: "Hello" + " " + message.Name}, nil
}
