package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Welcome struct {
	Sale string
	Time string
}

func main() {
	welcome := &Welcome{"Sale begins now", time.Now().Format(time.Stamp)}
	template := template.Must(template.ParseFiles("template/index.html"))
	fmt.Println("Hello, World!")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if sale := r.FormValue("sale"); sale != "" {
			welcome.Sale = sale
		}
		if err := template.ExecuteTemplate(w, "index.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println(http.ListenAndServe(":8000", nil))
	// port := flag.Int("port", -1, "specify a port")
	// flag.Parse()
	// log.Fatal(http.ListenAndServe(":"+*port, nil))
}
