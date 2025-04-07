package errdefs

import (
	"errors"
	"net/http"
)

func GetErrorStatusCode(err error) int {
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
