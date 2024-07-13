package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, _ *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	_, err := io.WriteString(res,
		`<doctype html>
	<html>
		<head>
			<title>Hello World!</title>
		</head>
		<body>
			Hello World
		</body>
	</html>`,
	)
	if err != nil {
		println(err.Error())
	}
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		println(err.Error())
	}
}
