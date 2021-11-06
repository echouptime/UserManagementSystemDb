package main

import (
	"UserManagementSystem/routers"
	"fmt"
	"net/http"
)

func main() {
	Addr := ":9000"

	routers.Register()
	if err := http.ListenAndServe(Addr, nil); err != nil {
		fmt.Println(err)
	}

}
