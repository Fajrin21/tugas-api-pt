package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type flower struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var flowers []flower

func sendresponse(code int, message string, data interface{}, w http.ResponseWriter) {
	resp := response{
		Code:    code,
		Data:    data,
		Message: message,
	}
	dataByte, err := json.Marshal(resp)
	if err != nil {
		resp := response{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "internal server error",
		}
		dataByte, _ = json.Marshal(resp)

	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(dataByte)

}

func main() {
	http.HandleFunc("/api/v1/flowers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			sendresponse(http.StatusOK, "succes", flowers, w)
			return
		}
		if r.Method == http.MethodPost {
			dataByte, err := io.ReadAll(r.Body)
			if err != nil {
				sendresponse(http.StatusBadRequest, "bad request", nil, w)
			}
			defer r.Body.Close()
			var flower flower
			err = json.Unmarshal(dataByte, &flower)
			if err != nil {
				sendresponse(http.StatusInternalServerError, "internal server error", nil, w)
			}

			flowers = append(flowers, flower)
			sendresponse(http.StatusCreated, "success", nil, w)
			return
		}

		if r.Method == http.MethodPut {
			w.Write([]byte("ini put"))
			return
		}
		if r.Method == http.MethodDelete {
			w.Write([]byte("ini delete"))
			return
		}
		w.Write([]byte("wrong method"))
	})

	port := "8000"
	fmt.Println("server run on port", port)
	http.ListenAndServe(":"+port, nil)

}
