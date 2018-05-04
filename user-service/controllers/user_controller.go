package controllers

import (
	"context"
	pb "hatch/rpc/user"

	"github.com/twitchtv/twirp"
)

type Server struct{}

func (s *Server) GetUser(ctx context.Context, id *pb.UserId) (hat *pb.User, err error) {
	if id.Id <= 0 {
		return nil, twirp.InvalidArgumentError("ID out of range")
	}
	var user pb.User
	return nil, user
}
