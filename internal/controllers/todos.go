package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAllTodo(w http.ResponseWriter, r *http.Request) {
	a := &map[string]int{
		"result": 0,
	}
	res, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(res)
}
