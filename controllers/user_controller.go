package controllers

import (
	"fmt"
	"net/http"
)

func TestConnection(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	fmt.Println("Sukses bos!")
	SendSuccessResponse(w, "Sukses Connect!")
}
