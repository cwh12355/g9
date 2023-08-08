package main

<<<<<<< HEAD
=======

>>>>>>> 4a2c29f (d1&&Initial commit)
import (
	"encoding/json"
	"fmt"
)
<<<<<<< HEAD
  
type description1 struct{
        incident string
		number []int
		volume int

}

type description2 struct{
	Name string  `json:"name"`
	Wide int `json:"wide"`
	Height int `json:"height"`
}

func main() {
	no_name := &description1{
		incident: "saup",
        number: []int{1,2,3,4,5},
		volume: 10,
}

guess_name,_ :=json.Marshal(no_name)

fmt.Println(string(guess_name))

no_inclient := &description2{
	Name: "saup",
	Wide: 10,
	Height: 20,
}
var dat map [string]interface{}

exec := []byte(`{"name":"saup","wide":10,"height":20}`)

if err :=json.Unmarshal(exec,&dat) err != nil {
	panic(err)
}


}
=======

func main() {
	var image  map[int]interface{}
	var dat map[string]interface{}
	 var cal = []byte(`{"apple":"down","banana":"up"}`)

	json.Unmarshal(cal,&dat)
	qq := dat["apple"].(string)
	fmt.Println(qq)

    



}
>>>>>>> 4a2c29f (d1&&Initial commit)
