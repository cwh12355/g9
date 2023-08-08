package main

import (
	"encoding/json"
	"fmt"
)

type terson struct {
	Color   string  `json:"color"`
	Weath   float32 `json:"weath"`
	Convert string  `json:"convert"`
}

var response1 = []terson{
	terson{Color: "yellow", Weath: 0.431, Convert: "slim"},
	terson{Color: "blue", Weath: 8.99, Convert: "punte"},
	terson{Color: "black", Weath: 0.01, Convert: "strong"},
}

var response2 = []byte(`{
	 "Color":"yellow",
     "Weath":["23","34"]
}`)

func main() {

	var person map[string]interface{}

	err := json.Unmarshal(response2, &person)
	if err != nil {
		fmt.Println("response2 error")
	}

	fmt.Println(person["Color"].(string))
	//var tee terson
	tee := terson{Color: "cs", Weath: 0.34, Convert: "yes"}
	tee.become_person(0.23, 9.25)
}

var response3 = []byte(`{
   stripe = "cs"
   paypal = []{
	"china","russia","United status of American"
   }
`)

type response4 struct {
	stripe string
	acera  int
	band   string
}

var por map[string]interface{}

func (c *response4) covet_json() {
	json.Unmarshal(response3, &por)
}
func (t *terson) become_person(head, hand float32) {
	//var head ,hand float32
	var person map[float32]interface{}

	complany := head*0.25 + hand*0.75

	Totally_Weath := t.Weath * complany

	fmt.Println(Totally_Weath)

	sum := []float32{head, hand, complany, Totally_Weath}

	Head := person[sum[0]]
	hc := person[head]

	fmt.Printf("the float32 of head is %v,the string of head is %v", Head, hc)

	fmt.Println(response1[0].Color)

}
