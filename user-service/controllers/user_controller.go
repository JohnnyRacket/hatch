package controllers

import (
	"context"
	pb "hatch/rpc/user"
	"math/rand"
	"strconv"

	"github.com/google/uuid"

	"github.com/twitchtv/twirp"
)

type Server struct{}

func (s *Server) GetUser(ctx context.Context, userId *pb.UserId) (user *pb.User, err error) {
	if userId.Id <= 0 {
		return nil, twirp.InvalidArgumentError(strconv.Itoa(int(userId.Id)), "ID out of range")
	}

	return &pb.User{
		Id:    userId.Id,
		Email: []string{"white", "black", "brown", "red", "blue"}[rand.Intn(4)],
		Name:  []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(3)],
	}, nil
}

//Register user puts a user into the system and will go to create an email code
func (s *Server) RegisterUser(ctx context.Context, user *pb.NewUser) (status *pb.Status, err error) {
	if user.Email == "" || user.Name == "" {
		return nil, twirp.InvalidArgumentError("", "Email and User must be valid")
	}

	return &pb.Status{Status: "success"}, nil
}

func (s *Server) createMagicEmailCode(email string) (err error) {
	uuid := uuid.New()
	uuid.String()
	return nil
}

func (s *Server) checkUserExists(email string) (res bool, err error) {
	return false, nil
}
