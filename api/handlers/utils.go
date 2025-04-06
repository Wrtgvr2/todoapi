package handlers

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/wrtgvr/todoapi/internal/logger"
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

func HandleInternalError(w http.ResponseWriter, err error) {
	logger.LogError(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
