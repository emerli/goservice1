package main

import (
	"encoding/json"
	"fmt"
	"github.com/emerli/gomodrestservice"
	"log"
	"net/http"
)

type Request1 struct {
	//Id *string `json:"id"`
	Id int
	V1 *string
}

func main() {
	var service services.RESTService
	service.AddPostMethod("/add", func(w http.ResponseWriter, r *http.Request) {
		var payload Request1
		err1 := json.NewDecoder(r.Body).Decode(&payload)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}

		log.Println(payload)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"id":"%d","you have sent ":"%s"}`, 1, "response post")))

	})
	service.AddGetMethod("/test2", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"id":"%d","you have sent ":"%s"}`, 1, "response get")))

	})
	service.Start(":10000")
}
