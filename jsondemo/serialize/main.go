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

func testStruct() {
	monster := Monster{
		Name : "Hawk",
		Age : 200,
		Birth : "2021-07-01",
		Sal : 45.6,
		Skill : "play",
	}
	
	data, err := json.Marshal(&monster)
	if err != nil {
		panic("serialize fail")
	}
	fmt.Println(string(data))
	
}

func testMap() {
	
	m := make(map[string]interface{})


	m["name"] = "Steven"
	m["age"] = 24
	m["address"] = []string{"Taipei","NewYork"}

	data, err := json.Marshal(m)
	if err != nil {
		panic("map searialize fail")
	}

	fmt.Println(string(data))
	
}

func testSlice() {
	//s := make([]map[string]interface{},0)
	var s []map[string]interface{}

	m1 := make(map[string]interface{})

	m1["name"] = "Steven"
	m1["age"] = 24
	m1["address"] = "Taipei"

	s = append(s,m1)
	
	m2 := make(map[string]interface{})

	m2["name"] = "Kevin"
	m2["age"] = 27
	m2["address"] = "newyork"

	s = append(s,m2)

	data, err := json.Marshal(s) // marshal就是serialize的意思
	if err != nil {
		panic("map searialize fail")
	}

	fmt.Println(string(data))

}
 
//基本數據類型數據化沒啥意義
func testFloat() {
	var f float64 = 234.6
	data, err := json.Marshal(f)
	if err != nil {
		panic("map searialize fail")
	}

	fmt.Println(string(data))
}

func main() {
	//testStruct() //{"monster_name":"Hawk","monter_age":200,"Birth":"2021-07-01","Sal":45.6,"Skill":"play"}
	//testMap() //{"address":["Taipei","NewYork"],"age":24,"name":"Steven"}
	testSlice() // [{"address":"Taipei","age":24,"name":"Steven"},{"address":"newyork","age":27,"name":"Kevin"}]
	//testFloat()
}