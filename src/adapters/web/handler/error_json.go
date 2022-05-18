package handler

import "encoding/json"

func JsonError(msg string) []byte {
	error := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}

	result, err := json.Marshal(error)
	if err != nil {
		return []byte(err.Error())
	}

	return result
}
