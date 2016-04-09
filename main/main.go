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
			tmpl.Execute(w, nil)
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
<h1>Static Template</h1> <br />
This is static template demo page.
</body>
</html>
`