package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type objects struct {
	Global_id int64 `json:"global_id"`
}

func main() {
	var v objects
	var idTotal int64

	file, err := os.Open("data-20190514T0100.json")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	dec := json.NewDecoder(file)

	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	for dec.More() {
		err := dec.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		idTotal += v.Global_id
		//		fmt.Printf("%v: %v\n", v.Global_id, v.Global_id)
	}
	fmt.Printf("%v", idTotal)
}
