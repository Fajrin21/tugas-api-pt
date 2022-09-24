package main

import (
	"fmt"
	"net/http"
)

type Flowers struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	http.HandleFunc("/api/v1/flowers", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success"))
	})
	port := "8000"
	fmt.Println("server run on port", port)
	http.ListenAndServe(":"+port, nil)
}
