package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", http.StripPrefix("/prueba.html/", fs))
	http.HandleFunc("/static", home)

	http.HandleFunc("/info", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Host: ", req.Host)
		fmt.Fprintln(w, "URI:", req.RequestURI)
		fmt.Fprintln(w, "Method: ", req.Method)
		fmt.Fprintln(w, "RemoteAddr:", req.RemoteAddr)
	})

	http.HandleFunc("/producto", producto)
	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/producto", 301)
	})
	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Error chungo perro", 501)
	})
	http.HandleFunc("/cabeceras", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Test", "test1")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		fmt.Fprintln(w, "{\"hola\":1 }")
	})
	//tmpl := template.Must(template.ParseFiles("./yuo/templates/go.html"))
	//http.HandleFunc("/template", func(w http.ResponseWriter, r *http.Request) {
	//	tmpl.Execute(w, struct{ Saludo string }{"pe desgraciados!!!!"})
	//})

	log.Println("escuchando.. por el puerto 8080 papá")

	http.ListenAndServe(":8080", nil)

}
func home(w http.ResponseWriter, r *http.Request) {
	html := "<html>"
	html += "<body>"
	html += "<h1>Hola papas</h1>"
	html += "</body>"
	html += "</html>"

	w.Write([]byte(html))

}

var productos []string

func producto(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	add, okForm := r.Form["add"]
	if okForm && len(add) == 1 {
		productos = append(productos, string(add[0]))
		w.Write([]byte("Producto añadido correctamente"))

		return
	}
	prod, ok := r.URL.Query()["prod"]
	if ok && len(prod) == 1 {
		pos, err := strconv.Atoi(prod[0])
		if err != nil {
			return
		}
		html := "<html>"
		html += "<body>"
		html += "<h1>Productos " + productos[pos] + "</h1>"
		html += "</body>"
		html += "</html>"

		w.Write([]byte(html))
		return
	}
	html := "<html>"
	html += "<body>"
	html += "<h1>Total Productos" + strconv.Itoa(len(productos)) + "</h1>"
	html += "</body>"
	html += "</html>"
	w.Write([]byte(html))
}
