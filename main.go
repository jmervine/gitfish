package main

import (
	"encoding/json"
	"github.com/jmervine/exec/v2"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	action, conditions := CliHandler(os.Args)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var push PushEvent
		if err := decoder.Decode(&push); err != nil {
			log.Println(err)
		}

		if conditions.Auth(r) && conditions.AreMet(push) {
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

	log.Println("Starting on *:8888")
	log.Panic(http.ListenAndServe(":8888", nil))
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
