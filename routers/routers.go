package routers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io"
)

func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	io.WriteString(w,"hello world")
}

func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	name:=p.ByName("user_name")
	io.WriteString(w,name)
}

