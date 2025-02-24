package main

import (
	"html/template"
	"net/http"
	"os"
)

func loadTemplate(filename string) (*template.Template, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}

	tpl, err := template.New("main").Parse(string(data))
	if err != nil {
		return nil, err
	}

	return tpl, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	tpl, err := loadTemplate("template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
