package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
	Hobbies               []string
}

func (u User) getallInfo() string {
	return fmt.Sprintf("User name is %s, he age is %d. He has money %d.", u.Name, u.Age, u.Money)
}
func (u *User) setName(newName string) {
	u.Name = newName

}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"Football", "Skate", "Dance"}}
	//bob.setName("Alex")
	//fmt.Fprintf(page, bob.getallInfo())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}
func contacts_page(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Contacts page")
}
func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":1010", nil)
}
func main() {
	//var bob User = ...

	handleRequest()

}
