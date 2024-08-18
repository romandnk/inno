package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/bogatyr285/auth-go/internal/auth/entity"
	storageerrors "github.com/bogatyr285/auth-go/internal/auth/repository/errors"
	"github.com/bogatyr285/auth-go/internal/buildinfo"
	"github.com/bogatyr285/auth-go/internal/gateway/http/gen"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var ErrInternalError = errors.New("internal error")

type UserRepository interface {
	RegisterUser(ctx context.Context, u entity.UserAccount) (int, error)
	FindUserByEmail(ctx context.Context, email string) (entity.UserAccount, error)
}

type SaveTokenRequest struct {
	UserId      int
	Token       string
	Fingerprint string
	ExpiresIn   time.Time
}

type TokenRepository interface {
	SaveToken(ctx context.Context, req SaveTokenRequest) error
}

type CryptoPassword interface {
	HashPassword(password string) ([]byte, error)
	ComparePasswords(fromUser, fromDB string) bool
}

type JWTManager interface {
	NewAccessToken(sub string) (string, error)
	VerifyAccessToken(tokenString string) (*jwt.Token, error)
}

type AuthUseCase struct {
	userRepository  UserRepository
	tokenRepository TokenRepository
	cryptoPassword  CryptoPassword
	jwtManager      JWTManager
	buildInfo       buildinfo.BuildInfo
}

func NewUseCase(
	userRepository UserRepository,
	tokenRepository TokenRepository,
	cryptoPassword CryptoPassword,
	jwtManager JWTManager,
	buildInfo buildinfo.BuildInfo,
) AuthUseCase {
	return AuthUseCase{
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
		cryptoPassword:  cryptoPassword,
		jwtManager:      jwtManager,
		buildInfo:       buildInfo,
	}
}

func (u AuthUseCase) PostLogin(ctx context.Context, request gen.PostLoginRequestObject) (gen.PostLoginResponseObject, error) {
	user, err := u.userRepository.FindUserByEmail(ctx, request.Body.Email)
	if err != nil {
		return gen.PostLogin500JSONResponse{
			Error: err.Error(),
		}, nil
	}

	if !u.cryptoPassword.ComparePasswords(user.Password, request.Body.Password) {
		return gen.PostLogin401JSONResponse{Error: "unauth"}, nil
	}

	if request.Params.UserAgent == "" {
		return gen.PostLogin400JSONResponse{Error: "empty user agent"}, nil
	}

	accessToken, err := u.jwtManager.NewAccessToken(user.Email)
	if err != nil {
		return gen.PostLogin500JSONResponse{Error: "error creating access token"}, err
	}

	refreshToken := uuid.NewString()

	err = u.tokenRepository.SaveToken(ctx, SaveTokenRequest{
		UserId:      user.Id,
		Token:       refreshToken,
		Fingerprint: request.Params.UserAgent,
		ExpiresIn:   time.Now().UTC().Add(24 * time.Hour),
	})
	if err != nil {
		return gen.PostLogin500JSONResponse{Error: "error saving refresh token"}, err
	}

	return gen.PostLogin200JSONResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u AuthUseCase) PostRegister(ctx context.Context, request gen.PostRegisterRequestObject) (gen.PostRegisterResponseObject, error) {
	hashedPassword, err := u.cryptoPassword.HashPassword(request.Body.Password)
	if err != nil {
		return gen.PostRegister500JSONResponse{}, nil
	}

	user := entity.UserAccount{
		Email:    request.Body.Email,
		Password: string(hashedPassword),
	}

	id, err := u.userRepository.RegisterUser(ctx, user)
	if err != nil {
		if errors.Is(err, storageerrors.ErrNicknameAlreadyExists) {
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
		Arch:       u.buildInfo.Arch,
		BuildDate:  u.buildInfo.BuildDate,
		CommitHash: u.buildInfo.CommitHash,
		Compiler:   u.buildInfo.Compiler,
		GoVersion:  u.buildInfo.GoVersion,
		Os:         u.buildInfo.OS,
		Version:    u.buildInfo.Version,
	}, nil
}
