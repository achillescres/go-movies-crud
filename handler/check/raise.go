package check

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func RaiseInternalError(w http.ResponseWriter, err error) {
	log.Error(err)
	http.Error(w, "internal server error", http.StatusInternalServerError)
}
