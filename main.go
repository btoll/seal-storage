package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var t string = `
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>{{ .NodeName }}</title>
<style>
body {
    background-color: black;
}
div {
    background-color: white;
    height: 140px;
    width: 600px;

    position: absolute;
    top:0;
    bottom: 0;
    left: 0;
    right: 0;

    margin: auto;
    padding: 5px;
}
p {
    font-weight: bold;
    font-size: 1.5em;
    margin: 0;
    padding: 0;
}
</style>
</head>
<body>
    <div>
        <p>Pod: {{ .PodName }}</p>
        <p>Node: {{ .NodeName }}</p>
        <p>Namespace: {{ .Namespace }}</p>
        <p>IP: {{ .IP }}</p>
    </div>
</body>
</html>`

func f(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	tpl := template.Must(template.New("index").Parse(t))
	tpl.Execute(w, struct {
		PodName   string
		NodeName  string
		Namespace string
		IP        string
	}{
		PodName:   os.Getenv("POD_NAME"),
		NodeName:  os.Getenv("NODE_NAME"),
		Namespace: os.Getenv("POD_NAMESPACE"),
		IP:        os.Getenv("POD_IP"),
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", f)
	log.Fatal(http.ListenAndServe(":3001", mux))
	fmt.Println("Server listening on port 3001")
}
