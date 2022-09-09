package check

import (
	"fmt"
	"net/http"
)

func Error(w http.ResponseWriter, err error) bool {
	if err != nil {
		RaiseInternalError(w, err)
		return true
	}

	return false
}

func Existence(w http.ResponseWriter, ok bool, id string) bool {
	if !ok {
		http.Error(w,
			fmt.Sprintf("film with id=%s not found", id),
			http.StatusNotFound)

		return true
	}

	return false
}
