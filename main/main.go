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
			context := Context{"Kalpesh", "Satasiya"}
			tmpl.Execute(w, context)
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
<h1>Hello {{.Name}}, {{.SurName}}</h1> <br />
This is static template demo page.
</body>
</html>
`

type Context struct {
	Name string
	SurName string
}