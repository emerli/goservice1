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
		//send new data to all replicas.
	case "SLAVE":
		//receive data from master
	default:
		log.Fatal("SERVICE_ROLE env is mandatory!")
	}

	var rtype = os.Getenv("REPLICA_TYPE")
	switch rtype {
	case "SYNC":
		//send new data to all replicas during master data add and wait that alla replicas are syncronized.
	case "ASYNC":
		//send new data to all replicas during master data add in asyncronous mode.
	default:
		log.Fatal("REPLICA_TYPE env is mandatory!")
	}

	var addr = os.Getenv("SERVICE_ADDRESS")
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
