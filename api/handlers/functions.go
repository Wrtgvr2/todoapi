package handlers

import (
	"strconv"
	"strings"
)

func GetIdFromUrl(urlPath string) (uint64, error) {
	idStr := strings.TrimPrefix(urlPath, "/users/")

	if idStr == "" {
		return uint64(0), ErrUserIdRequired
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return uint64(0), ErrInvalidUserID
	}

	return id, nil
}
