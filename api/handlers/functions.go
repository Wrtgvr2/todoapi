package handlers

import (
	"regexp"
	"strconv"
)

func GetIdFromUrl(urlPath string) (uint64, error) {
	r := regexp.MustCompile(`^\/(?:users|todos)/(\d+)(:?/todos)?`)
	var id uint64

	matches := r.FindStringSubmatch(urlPath)
	if len(matches) > 1 {
		var err error
		id, err = strconv.ParseUint(matches[1], 10, 64)
		if err != nil {
			return 0, err
		}
	} else {
		return 0, ErrInvalidUserID
	}

	return id, nil
}
