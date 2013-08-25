package main

import (
	"net"
	"net/http"
	"fmt"
	"os"
	"html/template"
	"encoding/json"
	"io/ioutil"
)

const (
	port = "5001"
)

var (
	tcpPort = os.Getenv("PORT")
	index = template.Must(template.ParseFiles("templates/index.html"))
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	index.Execute(w, nil)
}

//OpenBLAS : Your OS does not support AVX instructions. OpenBLAS is using Nehalem kernels as a fallback, which may give poorer performance.

func codeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		conn, err := net.Dial("tcp", "localhost:" + tcpPort)
		data := make(map[string]string)
		data["message"] = ""
		data["error"] = ""
		defer conn.Close()
		if err != nil {
			data["error"] = err.Error()
			d, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Got an error encoding json error", err.Error())
			}
			fmt.Fprint(w, d)
			return
		}
		code := r.FormValue("code")
		conn.Write([]byte(code))
		conn.Write([]byte("\n"))
		conn.Write([]byte("exit()"))
		conn.Write([]byte("\n"))
		buf, _ := ioutil.ReadAll(conn)
		w.Header().Set("Content-Type", "application/json")
		data["message"] = string(buf)
		d, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Got an error encoding json message", err.Error())
		}
		fmt.Fprint(w, string(d))
		return
	}
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/code", codeHandler)
	http.HandleFunc("/", indexHandler)
	fmt.Println("Serving on port", port)
	http.ListenAndServe(":" + port, nil)
}

