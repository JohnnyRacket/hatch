package controllers

import (
	"context"
	"errors"
	pb "hatch/rpc/user"
	"hatch/user-service/data"
	"hatch/user-service/models"
	"log"

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
func (s *UserServiceServer) RegisterUser(ctx context.Context, user *pb.NewUser) (*pb.Status, error) {
	if user.Email == "" || user.Name == "" {
		return nil, twirp.InvalidArgumentError("", "Email and User must be valid")
	}

	model := models.User{Name: user.Name, Email: user.Email}
	if !model.Validate() {
		return nil, errors.New("Malformed User")
	}

	if !s.checkUserExists(user.Email) {
		id, err := s.userRepo.AddUser(user.Email, user.Name)
		if err != nil {
			return nil, err
		}
		s.emailCodeRepo.AddEmailCode(id, uuid.New())

		return &pb.Status{Code: 200, Message: "Success"}, nil
	} else {
		return &pb.Status{Code: 303, Message: "User Already Exists"}, nil
	}

}

func (s *UserServiceServer) checkUserExists(email string) bool {
	res, err := s.userRepo.CheckUserExists(email)
	if err != nil {
		log.Fatal(err)
		return true
	}
	return res
}
