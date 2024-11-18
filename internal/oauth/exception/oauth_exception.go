package oauthException

import (
	errorList "github.com/diki-haryadi/go-micro-template/pkg/constant/error/error_list"
	customErrors "github.com/diki-haryadi/go-micro-template/pkg/error/custom_error"
	errorUtils "github.com/diki-haryadi/go-micro-template/pkg/error/error_utils"
)

//var (
//	errStatusCodeMap = map[error]int{
//		ErrAuthorizationCodeNotFound:     http.StatusNotFound,
//		ErrAuthorizationCodeExpired:      http.StatusBadRequest,
//		ErrInvalidRedirectURI:            http.StatusBadRequest,
//		ErrInvalidScope:                  http.StatusBadRequest,
//		ErrInvalidUsernameOrPassword:     http.StatusBadRequest,
//		ErrRefreshTokenNotFound:          http.StatusNotFound,
//		ErrRefreshTokenExpired:           http.StatusBadRequest,
//		ErrRequestedScopeCannotBeGreater: http.StatusBadRequest,
//		ErrTokenMissing:                  http.StatusNotFound,
//		ErrTokenHintInvalid:              http.StatusBadRequest,
//		ErrAccessTokenNotFound:           http.StatusNotFound,
//		ErrRefreshTokenNotFound:          http.StatusNotFound,
//		ErrTokenMissing:                  http.StatusBadRequest,
//		ErrTokenHintInvalid:              http.StatusBadRequest,
//		ErrInvalidUsernameOrPassword:     http.StatusUnauthorized,
//	}
//)
//
//func getErrStatusCode(err error) int {
//	code, ok := errStatusCodeMap[err]
//	if ok {
//		return code
//	}
//
//	return http.StatusInternalServerError
//}

func AuthorizationCodeGrantValidationExc(err error) error {
	ve, ie := errorUtils.ValidationErrorHandler(err)
	if ie != nil {
		return ie
	}

	validationError := errorList.InternalErrorList.ValidationError
	return customErrors.NewValidationError(validationError.Msg, validationError.Code, ve)
}

func AuthorizationCodeGrantBindingExc() error {
	oauthBindingError := errorList.InternalErrorList.OauthExceptions.BindingError
	return customErrors.NewBadRequestError(oauthBindingError.Msg, oauthBindingError.Code, nil)
}

func PasswordGrantValidationExc(err error) error {
	ve, ie := errorUtils.ValidationErrorHandler(err)
	if ie != nil {
		return ie
	}

	validationError := errorList.InternalErrorList.ValidationError
	return customErrors.NewValidationError(validationError.Msg, validationError.Code, ve)
}

func PasswordGrantBindingExc() error {
	oauthBindingError := errorList.InternalErrorList.OauthExceptions.BindingError
	return customErrors.NewBadRequestError(oauthBindingError.Msg, oauthBindingError.Code, nil)
}

func GrantClientCredentialGrantValidationExc(err error) error {
	ve, ie := errorUtils.ValidationErrorHandler(err)
	if ie != nil {
		return ie
	}

	validationError := errorList.InternalErrorList.ValidationError
	return customErrors.NewValidationError(validationError.Msg, validationError.Code, ve)
}

func GrantClientCredentialGrantBindingExc() error {
	oauthBindingError := errorList.InternalErrorList.OauthExceptions.BindingError
	return customErrors.NewBadRequestError(oauthBindingError.Msg, oauthBindingError.Code, nil)
}

func RefreshTokenGrantValidationExc(err error) error {
	ve, ie := errorUtils.ValidationErrorHandler(err)
	if ie != nil {
		return ie
	}

	validationError := errorList.InternalErrorList.ValidationError
	return customErrors.NewValidationError(validationError.Msg, validationError.Code, ve)
}

func RefreshTokenGrantBindingExc() error {
	oauthBindingError := errorList.InternalErrorList.OauthExceptions.BindingError
	return customErrors.NewBadRequestError(oauthBindingError.Msg, oauthBindingError.Code, nil)
}

func IntrospectValidationExc(err error) error {
	ve, ie := errorUtils.ValidationErrorHandler(err)
	if ie != nil {
		return ie
	}

	validationError := errorList.InternalErrorList.ValidationError
	return customErrors.NewValidationError(validationError.Msg, validationError.Code, ve)
}

func IntrospectBindingExc() error {
	oauthBindingError := errorList.InternalErrorList.OauthExceptions.BindingError
	return customErrors.NewBadRequestError(oauthBindingError.Msg, oauthBindingError.Code, nil)
}
