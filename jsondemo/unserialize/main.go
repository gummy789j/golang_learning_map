package main

import (
	"fmt"
	"encoding/json"
)

type Monster struct {
	Name string `json:"monster_name"`
	Age int `json:"monter_age"`
	Birth string
	Sal float64
	Skill string
}

func unmarshalStruct() {
	str := "{\"monster_name\":\"Hawk\",\"monter_age\":200,\"Birth\":\"2021-07-01\",\"Sal\":45.6,\"Skill\":\"play\"}"
	
	var monster Monster

	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		panic("unmarshal fail")
	}

	fmt.Println(monster)
}

func unmarshalMap() {
	str := "{\"address\":[\"Taipei\",\"NewYork\"],\"age\":24,\"name\":\"Steven\"}"
	
	var m map[string]interface{}

	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		panic("unmarshal fail")
	}

	fmt.Println(m)
}

func unmarshalSlice() {
	str := "[{\"address\":\"Taipei\",\"age\":24,\"name\":\"Steven\"},{\"address\":\"newyork\",\"age\":27,\"name\":\"Kevin\"}]"
	
	var s []map[string]interface{}

	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		panic("unmarshal fail")
	}

	fmt.Println(s)
}

func main() {
	//unmarshalStruct()
	//Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v. 
	//If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.
	//unmarshalMap()
	unmarshalSlice()
}