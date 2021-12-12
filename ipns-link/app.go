/*
 * author - S@M
 */

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	http.HandleFunc("/", all_handle)
	http.Handle("/asset/", http.StripPrefix("/asset", http.FileServer(http.Dir("./local_asset"))))
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func all_handle(w http.ResponseWriter, req *http.Request) {

	//fmt.Println(req.Method, req.URL.Path, req.URL.RawQuery)
	//fmt.Println(req.Header)
	fmt.Println(req.RequestURI)

	switch req.URL.Path {
	case "/":
		index(w, req)
		break
	default:
		fmt.Println(req.URL.Path)
	}
}

func index(w http.ResponseWriter, req *http.Request) {

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

/*
func about(w http.ResponseWriter, req *http.Request) {

	tpl.ExecuteTemplate(w, "about.gohtml", nil)
}
*/
