package data

type GenericResult struct {
	ErrorMessage *string `json:"errorMessage"`
	Success      bool    `json:"success"`
}

func GenericError() GenericResult {
	return GenericErrorMessage("An error has occurred")
}

func GenericSuccess() GenericResult {
	return GenericResult{
		Success: true,
	}
}

func GenericErrorMessage(errorMessage string) GenericResult {
	return GenericResult{
		ErrorMessage: &errorMessage,
		Success:      false,
	}
}