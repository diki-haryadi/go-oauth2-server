package oauthDomain

import (
	"context"
	model "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	oauthDto "github.com/diki-haryadi/go-micro-template/internal/oauth/dto"
	oauth2 "github.com/diki-haryadi/protobuf-ecomerce/oauth2_server_service/oauth2/v1"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
)

type Configurator interface {
	Configure(ctx context.Context) error
}

type UseCase interface {
	AuthClient(ctx context.Context, clientID, secret string) (*model.Client, error)
	CreateClient(ctx context.Context, clientID, secret, redirectURI string) (*model.Client, error)
	ClientExists(ctx context.Context, clientID string) bool
	ClientCredentialsGrant(ctx context.Context, scope string, refreshToken string, client *model.Client) (*oauthDto.AccessTokenResponse, error)
	AuthorizationCodeGrant(ctx context.Context, code, redirectURI string, client *model.Client) (*oauthDto.AccessTokenResponse, error)
	PasswordGrant(ctx context.Context, username, password string, scope string, client *model.Client) (*oauthDto.AccessTokenResponse, error)
	RefreshTokenGrant(ctx context.Context, token, scope string, client *model.Client) (*oauthDto.AccessTokenResponse, error)
	IntrospectToken(ctx context.Context, token, tokenTypeHint string, client *model.Client) (*oauthDto.IntrospectResponse, error)
	Login(ctx context.Context, client *model.Client, user *model.Users, scope string) (*model.AccessToken, *model.RefreshToken, error)
	GetRefreshTokenScope(ctx context.Context, refreshToken *model.RefreshToken, requestedScope string) (string, error)
	GrantAccessToken(ctx context.Context, client *model.Client, user *model.Users, expiresIn int, scope string) (*oauthDto.AccessTokenResponse, error)
	Register(ctx context.Context, dto *oauthDto.UserRequestDto) (*oauthDto.UserResponse, error)
	ChangePassword(ctx context.Context, dto *oauthDto.ChangePasswordRequest) (*oauthDto.ChangePasswordResponse, error)
	ForgotPassword(ctx context.Context, dto *oauthDto.ForgotPasswordRequest) (*oauthDto.ForgotPasswordResponse, error)
	UpdateUsername(ctx context.Context, user *model.Users, username string) error
}

type Repository interface {
	GrantAccessToken(ctx context.Context, client *model.Client, user *model.Users, expiresIn int, scope string) (*model.AccessToken, error)
	Authenticate(ctx context.Context, token string) (*model.AccessToken, error)
	GrantAuthorizationCode(ctx context.Context, client *model.Client, user *model.Users, expiresIn int, redirectURI, scope string) (*model.AuthorizationCode, error)
	GetValidAuthorizationCode(ctx context.Context, code, redirectURI string, client *model.Client) (*model.AuthorizationCode, error)
	CreateClientCommon(ctx context.Context, clientID, secret, redirectURI string) (*model.Client, error)
	FindClientByClientID(ctx context.Context, clientID string) (*model.Client, error)
	FetchAuthorizationCodeByCode(ctx context.Context, client *model.Client, code string) (*model.AuthorizationCode, error)
	DeleteAuthorizationCode(ctx context.Context, authorizationCodeID string) error
	FetchClientByClientID(ctx context.Context, clientID string) (*model.Client, error)
	FetchUserByUserID(ctx context.Context, userID string) (*model.Users, error)
	GetOrCreateRefreshToken(ctx context.Context, client *model.Client, user *model.Users, expiresIn int, scope string) (*model.RefreshToken, error)
	GetValidRefreshToken(ctx context.Context, token string, client *model.Client) (*model.RefreshToken, error)
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
	PasswordGrant(ctx context.Context, req *oauth2.PasswordGrantRequest) (*oauth2.PasswordGrantResponse, error)
	AuthorizationCodeGrant(ctx context.Context, req *oauth2.AuthorizationCodeGrantRequest) (*oauth2.AuthorizationCodeGrantResponse, error)
	ClientCredentialsGrant(ctx context.Context, req *oauth2.ClientCredentialsGrantRequest) (*oauth2.ClientCredentialsGrantResponse, error)
	RefreshTokenGrant(ctx context.Context, req *oauth2.RefreshTokenGrantRequest) (*oauth2.RefreshTokenGrantResponse, error)
}

type HttpController interface {
	Tokens(c echo.Context) error
	Introspect(c echo.Context) error
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
