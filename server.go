package main

import (
	"fmt"
	"net/http"
)

func home_page(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Go is super easy")

}
func farrukh_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Farukhchik one love")
	fmt.Println(r)
}
func farrukhclass_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hey? how are you")

}
func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/farukh", farrukh_page)
	http.HandleFunc("/fa", farrukhclass_page)
	http.ListenAndServe(":1010", nil)
}
func main() {
	handleRequest()

}
