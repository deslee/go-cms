package data

import "fmt"

type GenericResult struct {
	ErrorMessage *string `json:"errorMessage"`
	Success      bool    `json:"success"`
}

func GenericSuccess() GenericResult {
	return GenericResult{
		Success: true,
	}
}

func ErrorGenericResult(errorMessage string) GenericResult {
	return GenericResult{
		ErrorMessage: &errorMessage,
		Success:      false,
	}
}

func ErrorUserResult(errorMessage string) UserResult {
	return UserResult{GenericResult: ErrorGenericResult(errorMessage)}
}
func ErrorLoginResult(errorMessage string) LoginResult {
	return LoginResult{GenericResult: ErrorGenericResult(errorMessage)}
}
func ErrorSiteResult(errorMessage string) SiteResult {
	return SiteResult{GenericResult: ErrorGenericResult(errorMessage)}
}
func ErrorItemResult(errorMessage string) ItemResult {
	return ItemResult{GenericResult: ErrorGenericResult(errorMessage)}
}

func UnexpectedErrorGenericResult(err error) GenericResult {
	return ErrorGenericResult(fmt.Sprintf("Unexpected error %s", err))
}
func UnexpectedErrorUserResult(err error) UserResult {
	return UserResult{GenericResult: UnexpectedErrorGenericResult(err)}
}
func UnexpectedErrorLoginResult(err error) LoginResult {
	return LoginResult{GenericResult: UnexpectedErrorGenericResult(err)}
}
func UnexpectedErrorSiteResult(err error) SiteResult {
	return SiteResult{GenericResult: UnexpectedErrorGenericResult(err)}
}
func UnexpectedErrorItemResult(err error) ItemResult {
	return ItemResult{GenericResult: UnexpectedErrorGenericResult(err)}
}
