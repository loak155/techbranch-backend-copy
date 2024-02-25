package usecase

import (
	"context"

	"github.com/loak155/techbranch-backend/internal/user/domain"
	"github.com/loak155/techbranch-backend/internal/user/repository"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IUserUsecase interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	ListUsers(ctx context.Context) ([]domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) (bool, error)
	DeleteUser(ctx context.Context, id int) (bool, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return domain.User{}, status.Errorf(codes.Internal, "failed to bcrypt generate password: %v", err)
	}
	newUser := domain.User{Username: user.Username, Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return domain.User{}, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}
	return newUser, nil
}

func (uu *userUsecase) GetUser(ctx context.Context, id int) (domain.User, error) {
	user, err := uu.ur.GetUser(id)
	if err != nil {
		return domain.User{}, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	return *user, nil
}

func (uu *userUsecase) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := uu.ur.GetUserByEmail(email)
	if err != nil {
		return domain.User{}, status.Errorf(codes.Internal, "failed to get user by email: %v", err)
	}
	return *user, nil
}

func (uu *userUsecase) ListUsers(ctx context.Context) ([]domain.User, error) {
	users, err := uu.ur.ListUsers()
	if err != nil {
		return []domain.User{}, status.Errorf(codes.Internal, "failed to get user list: %v", err)
	}
	return *users, nil
}

func (uu *userUsecase) UpdateUser(ctx context.Context, user domain.User) (bool, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return false, status.Errorf(codes.NotFound, "failed to bcrypt generate password: %v", err)
	}
	updatedUser := domain.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: string(hash),
	}
	if err := uu.ur.UpdateUser(&updatedUser); err != nil {
		return false, status.Errorf(codes.Internal, "failed to update user: %v", err)
	}
	return true, nil
}

func (uu *userUsecase) DeleteUser(ctx context.Context, id int) (bool, error) {
	if err := uu.ur.DeleteUser(id); err != nil {
		return false, status.Errorf(codes.Internal, "failed to delete user: %v", err)
	}
	return true, nil
}
