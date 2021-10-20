package sample_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func routes() {
	http.HandleFunc("/sendjson", sendJSON)
}

func sendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Stuart",
		Email: "stuart@stuart.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}

func init() {
	routes()
}

func TestSendJSON(t *testing.T) {
	url := "/sendjson"
	statusCode := 200

	r := httptest.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)

	if w.Code != statusCode {
		t.Log("Nope")
		t.Error("Yeah")
	}

	// And so on
}
