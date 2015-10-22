package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type NameObject struct {
	Name string `json:"username"`
}

type GreetObject struct {
	Greet string `json:"reponse"`
}

func grabName(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var nameObject NameObject
	var greetObject GreetObject
	
	
	decoder := json.NewDecoder(r.Body)
    fmt.Println(r.Body)

    err1 := decoder.Decode(&nameObject)
    if err1 != nil {
    panic(err1)
    }

    greetObject.Greet="Hey,Hi "+nameObject.Name

    output,_ := json.Marshal(greetObject)
    fmt.Fprintf(w,string(output))
	
}

func displayName(rw http.ResponseWriter, jsonRequest *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "How are you doing?,%s!\n", p.ByName("username"))
}

func main() {
	r := httprouter.New()
	r.GET("/hello/:username", displayName)
	r.POST("/hello", grabName)
	http.ListenAndServe("localhost:8080", r)
}