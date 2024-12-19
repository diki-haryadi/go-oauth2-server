package authDomain

import (
	"context"
	model "github.com/diki-haryadi/go-micro-template/internal/authentication/domain/model"
	authDto "github.com/diki-haryadi/go-micro-template/internal/authentication/dto"
	auhenticationV1 "github.com/diki-haryadi/protobuf-ecomerce/oauth2_server_service/authentication/v1"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
)

type Configurator interface {
	Configure(ctx context.Context) error
}

type UseCase interface {
	AuthClient(ctx context.Context, clientID, secret string) (*model.Client, error)
	Register(ctx context.Context, dto *authDto.UserRequestDto) (*authDto.UserResponse, error)
	ChangePassword(ctx context.Context, dto *authDto.ChangePasswordRequest) (*authDto.ChangePasswordResponse, error)
	ForgotPassword(ctx context.Context, dto *authDto.ForgotPasswordRequest) (*authDto.ForgotPasswordResponse, error)
	UpdateUsername(ctx context.Context, user *model.Users, username string) error
}

type Repository interface {
	CreateClientCommon(ctx context.Context, clientID, secret, redirectURI string) (*model.Client, error)
	FindClientByClientID(ctx context.Context, clientID string) (*model.Client, error)
	FindRoleByID(ctx context.Context, id string) (*model.Role, error)
	GetScope(ctx context.Context, requestedScope string) (string, error)
	GetDefaultScope(ctx context.Context) string
	ScopeExists(ctx context.Context, requestedScope string) bool
	FindUserByUsername(ctx context.Context, username string) (*model.Users, error)
	CreateUserCommon(ctx context.Context, roleID, username, password string) (*model.Users, error)
	SetPasswordCommon(ctx context.Context, user *model.Users, password string) error
	UpdateUsernameCommon(ctx context.Context, user *model.Users, username string) error
	UpdatePassword(ctx context.Context, uuid, password string) error
}

type GrpcController interface {
	Register(ctx context.Context, req *auhenticationV1.RegisterRequest) (*auhenticationV1.RegisterResponse, error)
	ChangePassword(ctx context.Context, req *auhenticationV1.ChangePasswordRequest) (*auhenticationV1.ChangePasswordResponse, error)
	ForgotPassword(ctx context.Context, req *auhenticationV1.ForgotPasswordRequest) (*auhenticationV1.ForgotPasswordResponse, error)
	UpdateUsername(ctx context.Context, req *auhenticationV1.UpdateUsernameRequest) (*auhenticationV1.UpdateUsernameResponse, error)
}

type HttpController interface {
	Register(c echo.Context) error
	ChangePassword(c echo.Context) error
	ForgotPassword(c echo.Context) error
	UpdateUsername(ctx echo.Context) error
}

type Job interface {
	StartJobs(ctx context.Context)
}

type KafkaProducer interface {
	PublishCreateEvent(ctx context.Context, messages ...kafka.Message) error
}

type KafkaConsumer interface {
	RunConsumers(ctx context.Context)
}
