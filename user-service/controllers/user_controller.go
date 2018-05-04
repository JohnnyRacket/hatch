package controllers

import (
	"context"
	pb "hatch/rpc/user"
	"math/rand"
	"strconv"

	"github.com/twitchtv/twirp"
)

type Server struct{}

func (s *Server) GetUser(ctx context.Context, userId *pb.UserId) (hat *pb.User, err error) {
	if userId.Id <= 0 {
		return nil, twirp.InvalidArgumentError(strconv.Itoa(int(userId.Id)), "ID out of range")
	}

	return &pb.User{
		Id:    userId.Id,
		Email: []string{"white", "black", "brown", "red", "blue"}[rand.Intn(4)],
		Name:  []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(3)],
	}, nil
}
