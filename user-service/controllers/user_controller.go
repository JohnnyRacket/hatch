package controllers

import (
	"context"
	pb "hatch/rpc/user"
	"hatch/user-service/data"

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
func (s *UserServiceServer) GetUser(ctx context.Context, userId *pb.UserId) (*pb.User, error) {
	id, err := uuid.Parse(userId.Id)
	if err != nil {
		return nil, err
	}
	user, err := s.userRepo.GetUser(id)
	//TODO get does not exist error
	return &pb.User{
		Id:    userId.Id,
		Email: user.Email,
		Name:  user.Name,
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
