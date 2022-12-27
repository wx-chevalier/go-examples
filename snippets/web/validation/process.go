package validation

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Controller) Process(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var p Payload

	if err := decoder.Decode(&p); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := c.ValidatePayload(&p); err != nil {
		switch err.(type) {
		case Verror:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
