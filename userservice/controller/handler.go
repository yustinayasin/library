package controller

import (
	"context"
	users "userservice/business"

	"shared/proto"
	"shared/utils"

	"google.golang.org/grpc"
)

type UserService struct {
	usecase users.UserUseCaseInterface
	proto.UnimplementedUserServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase users.UserUseCaseInterface) {
	userGrpc := &UserService{usecase: usecase}
	proto.RegisterUserServiceServer(grpcServer, userGrpc)
}

func NewUserService(uc users.UserUseCaseInterface) *UserService {
	return &UserService{
		usecase: uc,
	}
}

func (s *UserService) SignUp(ctx context.Context, req *proto.UserEditRequest) (*proto.BaseResponse, error) {
	userUseCase := ToUsecaseEdit(req)
	user, err := s.usecase.SignUp(*userUseCase)
	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponse{
		Message: "User successfully created",
		User:    FromUsecase(user),
	}, nil
}

// Login handles user login functionality
func (s *UserService) Login(ctx context.Context, req *proto.UserLoginRequest) (*proto.BaseResponse, error) {
	userUseCase := ToUsecaseLogin(req) // assuming you have a ToUseCase() method
	user, err := s.usecase.Login(*userUseCase)
	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponse{
		Message: "User successfully login",
		User:    FromUsecase(user),
	}, nil
}

func (s *UserService) EditUser(ctx context.Context, req *proto.UserEditRequest) (*proto.BaseResponse, error) {
	userId := req.GetId()
	userEdit := ToUsecaseEdit(req) // assuming ToUseCase converts to business layer object
	user, err := s.usecase.EditUser(*userEdit, int(userId))
	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponse{
		Message: "User successfully updated",
		User:    FromUsecase(user),
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *proto.UserIdRequest) (*proto.BaseResponse, error) {
	userId := req.Id
	user, err := s.usecase.DeleteUser(int(userId))
	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponse{
		Message: "User successfully deleted",
		User:    FromUsecase(user),
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *proto.UserIdRequest) (*proto.BaseResponse, error) {
	userId := req.GetId()
	user, err := s.usecase.GetUser(int(userId))
	if err != nil {
		return nil, utils.NewGrpcError(err)
	}

	return &proto.BaseResponse{
		Message: "Successfully retrieved user",
		User:    FromUsecase(user),
	}, nil
}
