package controllers

import (
	"context"
	pb "hatch/rpc/user"
	"hatch/user-service/data"
	"math/rand"
	"strconv"

	"github.com/google/uuid"

	"github.com/twitchtv/twirp"
)

//UserServiceServer implements the twirp user service server
type UserServiceServer struct {
	userRepo      data.UserRepository
	emailCodeRepo data.EmailRepository
}

//NewUserServiceServer vends you a new user service server for twrip
func NewUserServiceServer(userRepo data.UserRepository, emailCodeRepo data.EmailRepository) *UserServiceServer {
	return &UserServiceServer{userRepo: userRepo, emailCodeRepo: emailCodeRepo}
}

//GetUser gets you a user by id
func (s *UserServiceServer) GetUser(ctx context.Context, userId *pb.UserId) (user *pb.User, err error) {
	if userId.Id <= 0 {
		return nil, twirp.InvalidArgumentError(strconv.Itoa(int(userId.Id)), "ID out of range")
	}

	return &pb.User{
		Id:    userId.Id,
		Email: []string{"white", "black", "brown", "red", "blue"}[rand.Intn(4)],
		Name:  []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(3)],
	}, nil
}

//RegisterUser puts a user into the system and will go to create an email code
func (s *UserServiceServer) RegisterUser(ctx context.Context, user *pb.NewUser) (status *pb.Status, err error) {
	if user.Email == "" || user.Name == "" {
		return nil, twirp.InvalidArgumentError("", "Email and User must be valid")
	}

	return &pb.Status{Status: "success"}, nil
}

func (s *UserServiceServer) createMagicEmailCode(email string) (code string) {
	uuid := uuid.New()
	return uuid.String()
}

func (s *UserServiceServer) checkUserExists(email string) (res bool, err error) {
	return false, nil
}
