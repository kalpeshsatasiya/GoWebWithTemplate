package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content Type", "text/html")
		tmpl, err := template.New("New").Parse(doc)
		if err == nil {
			//context := Context{"Kalpesh", "Satasiya", "/tech-coder"}
			tmpl.Execute(w, req.URL.Path)
		}
	})

	http.ListenAndServe(":8070", nil)
}

const doc = `
<!DOCTYPE html>
<html>
<head>
<title> Static Template </title>
</head>
<body>
This is static template demo page.

	{{if eq . "/tech-coder" }}
		<h1> Hey, Google Go!! </h1>
	{{else}}
		<h1> Else condition - {{.}}</h1>
	{{end}}
</body>
</html>
`

type Context struct {
	Name string
	SurName string
	Path string
}