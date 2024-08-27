package usecase

import (
	"context"
	"github.com/bogatyr285/auth-go/internal/auth/entity"
	cryptopasswordmock "github.com/bogatyr285/auth-go/internal/auth/usecase/mock/crypto_password"
	userrepositorymock "github.com/bogatyr285/auth-go/internal/auth/usecase/mock/user_repository"
	"github.com/bogatyr285/auth-go/internal/buildinfo"
	"github.com/bogatyr285/auth-go/internal/gateway/http/gen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func newUserRepositoryMock(t *testing.T) (*gomock.Controller, *userrepositorymock.MockUserRepository) {
	ctrl := gomock.NewController(t)
	mock := userrepositorymock.NewMockUserRepository(ctrl)
	return ctrl, mock
}

func newCryptoPasswordMock(t *testing.T) (*gomock.Controller, *cryptopasswordmock.MockCryptoPassword) {
	ctrl := gomock.NewController(t)
	mock := cryptopasswordmock.NewMockCryptoPassword(ctrl)
	return ctrl, mock
}

func TestAuthUseCase_PostRegister_Success(t *testing.T) {
	userRepositoryCtrl, userRepositoryMock := newUserRepositoryMock(t)
	defer userRepositoryCtrl.Finish()

	cryptoPasswordCtrl, cryptoPasswordMock := newCryptoPasswordMock(t)
	defer cryptoPasswordCtrl.Finish()

	authUseCase := NewUseCase(userRepositoryMock, nil, cryptoPasswordMock, nil, buildinfo.BuildInfo{})

	age := 18
	request := gen.PostRegisterRequestObject{
		Body: &gen.PostRegisterJSONRequestBody{
			Age:      &age,
			Email:    "test@test.com",
			Password: "1234",
		},
	}

	ctx := context.Background()

	expectedHashedPassword := "hashed"
	cryptoPasswordMock.EXPECT().HashPassword(request.Body.Password).Return([]byte(expectedHashedPassword), nil)
	userRepositoryMock.EXPECT().RegisterUser(ctx, entity.UserAccount{
		Email:    "test@test.com",
		Password: expectedHashedPassword,
	}).Return(1, nil)

	actualResponse, err := authUseCase.PostRegister(ctx, request)
	require.NoError(t, err)

	expectedResponse := gen.PostRegister201JSONResponse{
		Email: "test@test.com",
		Id:    1,
	}

	assert.Equal(t, expectedResponse, actualResponse)
}
