package main

import (
	"net"
	"net/http"
	"fmt"
	"os"
	"html/template"
	"io/ioutil"
)

const (
	port = "5001"
)

var (
	tcpPort = os.Getenv("PORT")
	//index = template.Must(template.ParseFiles("templates/index.html"))
)
/*
func indexHandler(w http.ResponseWriter, r *http.Request) {
	index.Execute(w, nil)
}
*/

//OpenBLAS : Your OS does not support AVX instructions. OpenBLAS is using Nehalem kernels as a fallback, which may give poorer performance.

func codeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		conn, err := net.Dial("tcp", "localhost:" + tcpPort)
		defer conn.Close()
		if err != nil {
			http.Error(w, err, http.StatusInternalServerError)
		}
		data := r.FormValue("code")
		conn.Write([]byte(data))
		conn.Write([]byte("exit()"))
		buf, _ := ioutil.ReadAll(conn)
		fmt.Println(buf.String())
	}
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func main() {
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/code", codeHandler)
//	http.HandleFunc("/", indexHandler)
	fmt.Println("Serving on port", port)
	http.ListenAndServe(":" + port, nil)
}