package main

import (
	"dblib"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func myhandlefunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case "GET":
		if employees, err := dblib.List(); err != nil {
			fmt.Println(err)
		} else {
			barr, _ := json.Marshal(employees)
			w.Write(barr)
		}
	case "POST":
		fmt.Fprint(w, "<h1>Post Method</h1>"+r.Method)
		emp := dblib.Emp{}
		err := json.NewDecoder(r.Body).Decode(&emp)
		fmt.Println("err", err)
		fmt.Println("in post ", emp)
		if err := dblib.Create(emp); err != nil {
			fmt.Println(err)
		}
		fmt.Fprint(w, "<h1>Inserted </h1>")
	case "PUT":
		fmt.Fprint(w, "<h1>PUT Method</h1>"+r.Method)
		emp := dblib.Emp{}
		err := json.NewDecoder(r.Body).Decode(&emp)
		fmt.Println("err", err)
		fmt.Println("in put ", emp)
		if err := dblib.Update(emp); err != nil {
			fmt.Println(err)
		}
		fmt.Fprint(w, "<h1>UPDATE </h1>")
	}	
}
func myhandlefunc1(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case "DELETE":
		fmt.Fprint(w, "<h1>DELETE Method</h1>"+r.Method)
		empno := r.PathValue("empno")
		fmt.Println(empno )
		if err := dblib.Delete(empno); err != nil {
			fmt.Println(err)
		}
		fmt.Fprint(w, "<h1>DELETED </h1>")
	}
}
func main() {

	http.HandleFunc("/dept", myhandlefunc)
	http.HandleFunc("/dept/{deptno}", myhandlefunc1)
    fmt.Println("Server starting on 8080.........")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
