package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"./routers"
	//"io"
)

/*var (
	w http.ResponseWriter
	r *http.Request
	p httprouter.Params
)
*/


func RegisterHanders()*httprouter.Router {
	router:=httprouter.New()
	router.POST("/user",routers.CreateUser)
	router.POST("/user/:user_name",routers.Login)
	return router
}

func main() {
	r:=RegisterHanders()
	http.ListenAndServe(":9999",r)
}
