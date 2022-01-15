package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Welcome struct {
	Name string
	Time string
}

func main() {

	// Display the name "User" with the Time right now (when running server)
	welcome := Welcome{"User", time.Now().Format(time.Stamp)}

	// Parsing the template for the HTML
	templates := template.Must(template.ParseFiles("src/templates/index.html"))

	http.Handle("static",
		http.StripPrefix("static",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}

		if err := templates.ExecuteTemplate(w, "index.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Up and running on Port 8000")
	fmt.Println(http.ListenAndServe(":8000", nil))
}
