package schema

type ErrorMessage struct {
	Indonesian string `json:"indonesian"`
	English    string `json:"english"`
}

type ErrorSchema struct {
	ErrorCode    string       `json:"errorCode"`
	ErrorMessage ErrorMessage `json:"errorMessage"`
}

func NewError(code, idMsg, enMsg string) ErrorSchema {
	return ErrorSchema{
		ErrorCode: code,
		ErrorMessage: ErrorMessage{
			Indonesian: idMsg,
			English:    enMsg,
		},
	}
}

func NewSuccess() ErrorSchema {
	return NewError("00", "Berhasil", "Success")
}
