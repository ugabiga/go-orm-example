package internal

import (
	"encoding/json"
	"log"
)

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func PrintJSONLog(v any) {
	output, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(output))
}
