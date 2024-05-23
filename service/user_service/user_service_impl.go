package user_service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/urcane/post-test-golang/exception"
	"github.com/urcane/post-test-golang/helper"
	"github.com/urcane/post-test-golang/model/domain"
	webuser "github.com/urcane/post-test-golang/model/web"
	"github.com/urcane/post-test-golang/repository/user_repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, request webuser.RegisterRequest) webuser.UserResponse
	Login(ctx context.Context, request webuser.LoginRequest) webuser.UserResponse
}

type UserServiceImpl struct {
	UserRepository user_repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository user_repository.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request webuser.RegisterRequest) webuser.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	token, err := helper.GenerateJWT(request.Username)
	helper.PanicIfError(err)

	user := domain.User{
		Username: request.Username,
		Password: string(hashedPassword),
		Email:    request.Email,
		Token:    token,
	}

	user = service.UserRepository.Save(ctx, tx, user)

	return webuser.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    user.Token,
	}
}

func (service *UserServiceImpl) Login(ctx context.Context, request webuser.LoginRequest) webuser.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUsername(ctx, tx, request.Username)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		panic(exception.NewUnauthorizedError("invalid username or password"))
	}

	token, err := helper.GenerateJWT(user.Username)
	helper.PanicIfError(err)

	service.UserRepository.UpdateToken(ctx, tx, user.ID, token)

	return webuser.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}
}
