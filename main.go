package main

import (
	"encoding/json"
	"github.com/jmervine/exec/v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port, action, conditions := CliHandler(os.Args)

	if len(action) == 0 {
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		var push PushEvent
		if err := json.Unmarshal(body, &push); err != nil {
			log.Println(err)
		}

		sig := r.Header.Get("X-Hub-Signature")
		if conditions.Auth(body, sig) && conditions.AreMet(push) {
			if _, _, err := exec.ExecTee2(os.Stdout, os.Stderr, action[0], action[1:]...); err != nil {
				log.Println(err)
				w.WriteHeader(500)
				return
			}
		} else {
			log.Printf("conditions are not met") // todo: push.ToString()
			w.WriteHeader(403)
			return
		}

		w.WriteHeader(200)
	})

	log.Printf("Starting on *:%s\n", port)
	log.Println("---")
	log.Println("Configuration")
	log.Printf(" - Command : %v\n", action)
	log.Printf(" - Branches: %v\n", conditions.Branches)
	log.Printf(" - Secret  : %v\n", !!(len(conditions.Secret) > 0))
	log.Printf(" - Owner   : %v\n", conditions.Owner)
	log.Printf(" - Admin   : %v\n", conditions.Admin)
	log.Printf(" - Master  : %v\n", conditions.Admin)
	log.Panic(http.ListenAndServe(":"+port, nil))
}

// safeSplit handles quoting well for commands
// e.g.: "/bin/bash bash -l -c 'echo \"foo bar bah bin\"'"
func safeSplit(s string) []string {
	split := strings.Split(s, " ")

	var result []string
	var inquote string
	var block string
	for _, i := range split {
		if inquote == "" {
			if strings.HasPrefix(i, "'") || strings.HasPrefix(i, "\"") {
				inquote = string(i[0])
				block = strings.TrimPrefix(i, inquote) + " "
			} else {
				result = append(result, i)
			}
		} else {
			if !strings.HasSuffix(i, inquote) {
				block += i + " "
			} else {
				block += strings.TrimSuffix(i, inquote)
				inquote = ""
				result = append(result, block)
				block = ""
			}
		}
	}

	return result
}
