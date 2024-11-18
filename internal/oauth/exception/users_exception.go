package oauthException

import (
	errorList "github.com/diki-haryadi/ztools/constant/error/error_list"
	customErrors "github.com/diki-haryadi/ztools/error/custom_error"
	errorUtils "github.com/diki-haryadi/ztools/error/error_utils"
)

func CreateUsersValidationExc(err error) error {
	ve, ie := errorUtils.ValidationErrorHandler(err)
	if ie != nil {
		return ie
	}

	validationError := errorList.InternalErrorList.ValidationError
	return customErrors.NewValidationError(validationError.Msg, validationError.Code, ve)
}

func UsersBindingExc() error {
	usersBindingError := errorList.InternalErrorList.ArticleExceptions.BindingError
	return customErrors.NewBadRequestError(usersBindingError.Msg, usersBindingError.Code, nil)
}
