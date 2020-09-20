package main

import (
	"encoding/json"
	"fmt"
	"github.com/emerli/gomodrestservice"
	"log"
	"net/http"
	"os"
)

type AddRequest struct {
	Message string
}

func main() {
	var InternalData []*string

	var role = os.Getenv("SERVICE_ROLE")
	switch role {
	case "MASTER":
		log.Print("This is master")
	case "SLAVE":
		log.Print("This is slave")
	default:
		log.Fatal("SERVICE_ROLE env is mandatory!")
	}

	var addr = os.Getenv("SERVICE_ADDRESS")
	log.Print(addr)
	if addr == "" {
		log.Fatal("SERVICE_ADDRESS env is mandatory!")
	}

	var Service services.RESTService
	Service.AddPostMethod("/add", func(w http.ResponseWriter, r *http.Request) {
		var payload AddRequest
		/* bytes,_ := ioutil.ReadAll(r.Body)

		json.Unmarshal(bytes,&payload)*/

		err1 := json.NewDecoder(r.Body).Decode(&payload)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}

		InternalData = append(InternalData, &payload.Message)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		//w.Write([]byte(fmt.Sprintf(`{"id":"%d","you have sent ":"%s"}`, 1, "response post")))
		w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, "OK")))

	})
	Service.AddGetMethod("/list", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		content, err1 := json.Marshal(InternalData)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}

		w.Write(content)

	})

	Service.AddGetMethod("/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"port":"%s","role":"%s"}`, addr, role)))

	})

	Service.Start(addr)
}
