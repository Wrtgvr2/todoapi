package errorhandler

import (
	"errors"
	"net/http"
)

func GetErrorStatusCode(w http.ResponseWriter, err error) int {
	if err == nil {
		return 0
	}
	for knownErr, status := range errorsToStatusMap {
		if errors.Is(err, knownErr) {
			return status
		}
	}
	return http.StatusInternalServerError
}
