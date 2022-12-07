package user

import (
	"context"
	"fmt"
	"log"
	"mymachine707/config"
	"mymachine707/protogen/blogpost"
	"mymachine707/storage"
	"mymachine707/util"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userService struct {
	cfg config.Config
	stg storage.Interfaces
	blogpost.UnimplementedUserServiceServer
}

func NewAuthService(cfg config.Config, stg storage.Interfaces) *userService {
	return &userService{
		cfg: cfg,
		stg: stg,
	}
}

func (s *userService) Ping(ctx context.Context, req *blogpost.Empty) (*blogpost.Pong, error) {
	log.Println("Ping")

	return &blogpost.Pong{
		Message: "Ok",
	}, nil
}

func (s *userService) CreateUser(ctx context.Context, req *blogpost.CreateUserRequest) (*blogpost.User, error) {
	fmt.Println("<<< ---- CreateUser ---->>>")
	id := uuid.New()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, " util.HashPassword: %s", err.Error())
	}

	req.Password = hashedPassword

	err = s.stg.AddUser(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddUser: %s", err.Error())
	}

	user, err := s.stg.GetUserByID(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserByID: %s", err.Error())
	}

	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, req *blogpost.UpdateUserRequest) (*blogpost.User, error) {
	fmt.Println("<<< ---- UpdateUser ---->>>")

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, " util.HashPassword: %s", err.Error())
	}

	req.Password = hashedPassword

	err = s.stg.UpdateUser(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateUser: %s", err.Error())
	}

	user, err := s.stg.GetUserByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserByID---!: %s", err.Error())
	}

	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, req *blogpost.DeleteUserRequest) (*blogpost.User, error) {
	fmt.Println("<<< ---- DeleteUser ---- >>>")

	user, err := s.stg.GetUserByID(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserByID: %s", err.Error())
	}

	err = s.stg.DeleteUser(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteUser: %s", err.Error())
	}

	return user, nil
}

func (s *userService) GetUserList(ctx context.Context, req *blogpost.GetUserListRequest) (*blogpost.GetUserListResponse, error) {
	fmt.Println("<<< ---- GetUserList ---->>>")

	res, err := s.stg.GetUserList(int(req.Offset), int(req.Limit), req.Search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserList: %s", err.Error())
	}

	return res, nil
}
func (s *userService) GetUserById(ctx context.Context, req *blogpost.GetUserByIDRequest) (*blogpost.User, error) {
	fmt.Println("<<< ---- GetUserById ---->>>")

	user, err := s.stg.GetUserByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserByID: %s", err.Error())
	}

	return user, nil
}
