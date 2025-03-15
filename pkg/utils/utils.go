package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x any) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	json.Unmarshal(body, x)
}

