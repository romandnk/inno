package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/bogatyr285/auth-go/internal/auth/entity"
	"github.com/bogatyr285/auth-go/internal/auth/repository"
	"github.com/bogatyr285/auth-go/internal/buildinfo"
	"github.com/bogatyr285/auth-go/internal/gateway/http/gen"
	"github.com/golang-jwt/jwt/v5"
)

var ErrInternalError = errors.New("internal error")

type UserRepository interface {
	RegisterUser(ctx context.Context, u entity.UserAccount) (int, error)
	FindUserByEmail(ctx context.Context, email string) (entity.UserAccount, error)
}

type CryptoPassword interface {
	HashPassword(password string) ([]byte, error)
	ComparePasswords(fromUser, fromDB string) bool
}

type JWTManager interface {
	NewAccessToken(sub string) (string, error)
	NewRefreshToken(sub string) (string, error)
	VerifyAccessToken(tokenString string) (*jwt.Token, error)
}

type AuthUseCase struct {
	ur UserRepository
	cp CryptoPassword
	jm JWTManager
	bi buildinfo.BuildInfo
}

func NewUseCase(
	ur UserRepository,
	cp CryptoPassword,
	jm JWTManager,
	bi buildinfo.BuildInfo,
) AuthUseCase {
	return AuthUseCase{
		ur: ur,
		cp: cp,
		jm: jm,
		bi: bi,
	}
}

func (u AuthUseCase) PostLogin(ctx context.Context, request gen.PostLoginRequestObject) (gen.PostLoginResponseObject, error) {
	user, err := u.ur.FindUserByEmail(ctx, request.Body.Email)
	if err != nil {
		return gen.PostLogin500JSONResponse{
			Error: err.Error(),
		}, nil
	}

	if !u.cp.ComparePasswords(user.Password, request.Body.Password) {
		return gen.PostLogin401JSONResponse{Error: "unauth"}, nil
	}

	accessToken, err := u.jm.NewAccessToken(user.Email)
	if err != nil {
		return gen.PostLogin500JSONResponse{}, err
	}

	refreshToken, err := u.jm.NewRefreshToken(user.Email)
	if err != nil {
		return gen.PostLogin500JSONResponse{}, err
	}

	return gen.PostLogin200JSONResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u AuthUseCase) PostRegister(ctx context.Context, request gen.PostRegisterRequestObject) (gen.PostRegisterResponseObject, error) {
	hashedPassword, err := u.cp.HashPassword(request.Body.Password)
	if err != nil {
		return gen.PostRegister500JSONResponse{}, nil
	}

	user := entity.UserAccount{
		Email:    request.Body.Email,
		Password: string(hashedPassword),
	}

	id, err := u.ur.RegisterUser(ctx, user)
	if err != nil {
		if errors.Is(err, repository.ErrNicknameAlreadyExists) {
			return gen.PostRegister400JSONResponse{
				Error: fmt.Sprintf("email already exists: %s", request.Body.Email),
			}, nil
		}
		return gen.PostRegister500JSONResponse{
			Error: ErrInternalError.Error(),
		}, nil
	}
	return gen.PostRegister201JSONResponse{
		Id:    id,
		Email: request.Body.Email,
	}, nil
}

func (u AuthUseCase) GetBuildinfo(ctx context.Context, request gen.GetBuildinfoRequestObject) (gen.GetBuildinfoResponseObject, error) {
	return gen.GetBuildinfo200JSONResponse{
		Arch:       u.bi.Arch,
		BuildDate:  u.bi.BuildDate,
		CommitHash: u.bi.CommitHash,
		Compiler:   u.bi.Compiler,
		GoVersion:  u.bi.GoVersion,
		Os:         u.bi.OS,
		Version:    u.bi.Version,
	}, nil
}
