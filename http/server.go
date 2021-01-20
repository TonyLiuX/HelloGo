package http

import (
	"fmt"
	"net/http"
)

type IndexHandler struct {
	content string
}

func (ih *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, ih.content)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `<!doctype html>
				<META http-equiv="Content-Type" content="text/html" charset="utf-8">
				<html lang="zh-CN">
					<head>
						<title>Golang</title>
						<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0;" />
						<body>
							<div id="app">Welcome!</div>
						</body>
					</head>
				</html>`
	fmt.Fprintf(w, html)
}

func Handler() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(htmlHandler))
	mux.HandleFunc("/welcome", welcome)
	//http.Handle("/", &IndexHandler{content: "Hello world!"})
	//http.HandleFunc("/welcome", welcome)
	_ = http.ListenAndServe(":9090", mux)
}
