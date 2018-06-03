package controllers

import (
	"context"
	"errors"
	pb "hatch/rpc/user"
	"hatch/user-service/data"
	"hatch/user-service/models"
	"log"
	"time"

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

//ValidateUser validates a users email code and returns the access & refresh tokens
func (s *UserServiceServer) ValidateUser(ctx context.Context, guid *pb.Guid) (*pb.UserToken, error) {
	id, err := uuid.Parse(guid.Value)
	if err != nil {
		return nil, err
	}
	emailCode, err := s.emailCodeRepo.GetEmailCode(id)
	if emailCode.Expiration.Before(time.Now()) {
		//TODO create access and fresh tokens and insert them
		return nil, nil
	} else {
		return nil, errors.New("Email Code is Expired")
	}
}

//DeleteUser removes a user from the system using the Id
func (s *UserServiceServer) DeleteUser(ctx context.Context, userId *pb.UserId) (*pb.Status, error) {
	id, err := uuid.Parse(userId.Id)
	if err != nil {
		return nil, err
	}
	err = s.userRepo.RemoveUser(id)

	if err != nil {
		return &pb.Status{Code: 500, Message: "Server Error Deleting User."}, nil
	} else {
		return &pb.Status{Code: 200, Message: "User Deleted."}, nil
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
