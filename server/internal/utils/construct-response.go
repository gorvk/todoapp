package utils

import (
	"encoding/json"

	"github.com/gorvk/todoapp/internal/types"
)

func ConstructResponse(isSuccess bool, result any) ([]byte, error) {
	response := types.RESPONSE_PARAMETERS{
		IsSuccess: isSuccess,
		Result:    result,
	}
	data, err := json.Marshal(response)
	return data, err
}
