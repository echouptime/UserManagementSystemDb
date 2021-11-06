package routers

import (
	"UserManagementSystem/controllers"
	"net/http"
)

func Register() {
	http.HandleFunc("/", controllers.BaseInformation)
	http.HandleFunc("/create/", controllers.CreateUser)
	//http.HandleFunc("/query/", controllers.QueryUser)
	http.HandleFunc("/update/", controllers.UpdateUser)
	http.HandleFunc("/delete/", controllers.DeleteUser)
}
