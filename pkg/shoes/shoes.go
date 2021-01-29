package shoes

import (
	"html/template"
	"net/http"
)

func Adidasi(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("html/shoes/adidasi.html"))
	t.Execute(w, nil)
}
func Nike(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("html/shoes/nike.html"))
	t.Execute(w, nil)
}
func Vans(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("html/shoes/vans.html"))
	t.Execute(w, nil)
}
func AirJodan(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("html/shoes/airjodan.html"))
	t.Execute(w, nil)
}
