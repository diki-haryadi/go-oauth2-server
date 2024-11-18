package oauthDto

import (
	"fmt"
	oauthModel "github.com/diki-haryadi/go-micro-template/internal/oauth/domain/model"
	"github.com/diki-haryadi/go-micro-template/pkg"
	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"time"
)

type GrantAuthorizationCodeRequestDto struct {
	GrantType   string `json:"grant_type"`
	Code        string `json:"code"`
	RedirectUri string `json:"redirect_uri"`
	ClientID    string `json:"client_id"`
}

func (g *GrantAuthorizationCodeRequestDto) ToModel(clientID uuid.UUID) *oauthModel.Client {
	return &oauthModel.Client{
		Common: oauthModel.Common{
			ID: clientID,
		},
	}
}

func (caDto *GrantAuthorizationCodeRequestDto) ValidateAuthorizationCodeDto() error {
	return validator.ValidateStruct(caDto,
		validator.Field(
			&caDto.GrantType,
			validator.Required,
		),
		validator.Field(
			&caDto.Code,
			validator.Required,
		),
		validator.Field(
			&caDto.RedirectUri,
			validator.Required,
		),
	)
}

type GrantPasswordRequestDto struct {
	GrantType string `json:"grant_type"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Scope     string `json:"scope"`
}

func (g *GrantPasswordRequestDto) ToModel(clientID uuid.UUID) *oauthModel.Client {
	return &oauthModel.Client{
		Common: oauthModel.Common{
			ID: clientID,
		},
	}
}

func (caDto *GrantPasswordRequestDto) ValidatePasswordDto() error {
	return validator.ValidateStruct(caDto,
		validator.Field(
			&caDto.GrantType,
			validator.Required,
		),
		validator.Field(
			&caDto.Username,
			validator.Required,
		),
		validator.Field(
			&caDto.Password,
			validator.Required,
		),
		validator.Field(
			&caDto.Scope,
			validator.Required,
		),
	)
}

type GrantClientCredentialsRequestDto struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`
}

func (g *GrantClientCredentialsRequestDto) ToModel(clientID uuid.UUID) *oauthModel.Client {
	return &oauthModel.Client{
		Common: oauthModel.Common{
			ID: clientID,
		},
	}
}

func (caDto *GrantClientCredentialsRequestDto) ValidateClientCredentialsDto() error {
	return validator.ValidateStruct(caDto,
		validator.Field(
			&caDto.GrantType,
			validator.Required,
		),
		validator.Field(
			&caDto.Scope,
			validator.Required,
		),
	)
}

type RefreshTokenRequestDto struct {
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func (g *RefreshTokenRequestDto) ToModel(clientID uuid.UUID) *oauthModel.Client {
	return &oauthModel.Client{
		Common: oauthModel.Common{
			ID: clientID,
		},
	}
}

func (caDto *RefreshTokenRequestDto) ValidateRefreshTokenDto() error {
	return validator.ValidateStruct(caDto,
		validator.Field(
			&caDto.GrantType,
			validator.Required,
		),
		validator.Field(
			&caDto.RefreshToken,
			validator.Required,
		),
		validator.Field(
			&caDto.Scope,
			validator.Required,
		),
	)
}

type IntrospectRequestDto struct {
	Token         string `json:"token"`
	TokenTypeHint string `json:"token_type_hint"`
}

func (g *IntrospectRequestDto) ToModel(clientID uuid.UUID) *oauthModel.Client {
	return &oauthModel.Client{
		Common: oauthModel.Common{
			ID: clientID,
		},
	}
}

func (caDto *IntrospectRequestDto) ValidateIntrospectDto() error {
	return validator.ValidateStruct(caDto,
		validator.Field(
			&caDto.Token,
			validator.Required,
		),
		validator.Field(
			&caDto.TokenTypeHint,
			validator.Required,
		),
	)
}

func NewOauthRefreshToken(client *oauthModel.Client, user *oauthModel.Users, expiresIn int, scope string) *oauthModel.RefreshToken {
	refreshToken := &oauthModel.RefreshToken{
		Common: oauthModel.Common{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		ClientID:  pkg.StringOrNull(fmt.Sprint(client.ID)),
		Token:     uuid.New().String(),
		ExpiresAt: time.Now().UTC().Add(time.Duration(expiresIn) * time.Second),
		Scope:     scope,
	}
	if user != nil {
		refreshToken.UserID = pkg.StringOrNull(fmt.Sprint(user.ID))
	}
	return refreshToken
}

func NewOauthAccessToken(client *oauthModel.Client, user *oauthModel.Users, expiresIn int, scope string) *oauthModel.AccessToken {
	accessToken := &oauthModel.AccessToken{
		Common: oauthModel.Common{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		ClientID:  pkg.StringOrNull(fmt.Sprint(client.ID)),
		Token:     uuid.New().String(),
		ExpiresAt: time.Now().UTC().Add(time.Duration(expiresIn) * time.Second),
		Scope:     scope,
	}
	if user != nil {
		accessToken.UserID = pkg.StringOrNull(fmt.Sprint(user.ID))
	}
	return accessToken
}

func NewOauthAuthorizationCode(client *oauthModel.Client, user *oauthModel.Users, expiresIn int, redirectURI, scope string) *oauthModel.AuthorizationCode {
	return &oauthModel.AuthorizationCode{
		Common: oauthModel.Common{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		ClientID:    pkg.StringOrNull(fmt.Sprint(client.ID)),
		UserID:      pkg.StringOrNull(fmt.Sprint(user.ID)),
		Code:        uuid.New().String(),
		ExpiresAt:   time.Now().UTC().Add(time.Duration(expiresIn) * time.Second),
		RedirectURI: pkg.StringOrNull(redirectURI),
		Scope:       scope,
	}
}

// AccessTokenResponse ...
type AccessTokenResponse struct {
	UserID       string `json:"user_id,omitempty"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// IntrospectResponse ...
type IntrospectResponse struct {
	Active    bool   `json:"active"`
	Scope     string `json:"scope,omitempty"`
	ClientID  string `json:"client_id,omitempty"`
	Username  string `json:"username,omitempty"`
	TokenType string `json:"token_type,omitempty"`
	ExpiresAt int    `json:"exp,omitempty"`
}

// NewAccessTokenResponse ...
func NewAccessTokenResponse(accessToken *oauthModel.AccessToken, refreshToken *oauthModel.RefreshToken, lifetime int, theTokenType string) (*AccessTokenResponse, error) {
	response := &AccessTokenResponse{
		AccessToken: accessToken.Token,
		ExpiresIn:   lifetime,
		TokenType:   theTokenType,
		Scope:       accessToken.Scope,
	}
	if accessToken.UserID.Valid {
		response.UserID = accessToken.UserID.String
	}
	if refreshToken != nil {
		response.RefreshToken = refreshToken.Token
	}
	return response, nil
}
