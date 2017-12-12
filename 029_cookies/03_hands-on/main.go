package main

import(
	"fmt"
	"net/http"
	"strconv"
	"io"
)

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/reset",reset)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter, req * http.Request){
	
	count := 1
	cookie,_ := req.Cookie("VisitCount")

	if cookie != nil {
			
		var err error
		count, err = strconv.Atoi(cookie.Value)
		if err == nil{
			count += 1
		}
	}

	http.SetCookie(w,&http.Cookie{
		Name: "VisitCount",
		Value: strconv.Itoa(count),
	})
	
	io.WriteString(w,fmt.Sprintf("You've visited this website %d times",count))
}

func reset(w http.ResponseWriter, req * http.Request){

	cookie,_ := req.Cookie("VisitCount")
	if cookie != nil{
		cookie.MaxAge = -1
	}
	http.SetCookie(w,cookie)
	http.Redirect(w,req,"/",http.StatusSeeOther)
}