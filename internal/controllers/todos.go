package controllers

import (
	"fmt"
	"net/http"
)

func AddTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("todo added")
}
