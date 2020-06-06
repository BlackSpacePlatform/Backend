package helper

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

)

// ExtractIdFromRequest takes as input a request object
// and extracts an id from it
func ExtractIDFromRequest(r *http.Request) uint32 {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	processedID := uint32(id)
	return processedID
}

// ProcessMalformedRequest handles aggregated errors occuring from interactions with various external services
func ProcessMalformedRequest(w http.ResponseWriter, err error) {
	var mr *MalformedRequest
	if errors.As(err, &mr) {
		http.Error(w, mr.Msg, mr.Status)
	} else {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

