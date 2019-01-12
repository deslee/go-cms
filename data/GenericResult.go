package data

type GenericResult struct {
	ErrorMessage *string `json:"errorMessage"`
	Success      bool    `json:"success"`
}
