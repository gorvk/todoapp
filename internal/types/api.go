package types

type RESPONSE_COMMON_PARAMETERS struct {
	IsSuccess bool `json:"isSuccess"`
}

type ERROR_RESPONSE struct {
	RESPONSE_COMMON_PARAMETERS
}

type TODO_RESPONSE_PARAMETERS struct {
	RESPONSE_COMMON_PARAMETERS
}
