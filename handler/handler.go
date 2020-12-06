package handler

import (
	"go-web-basic/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Selamat Datang Kawan"))

	tmplt, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is Happening, Keep Calm", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"title":   "I'm learning Golang web",
	// 	"content": "I'm Learning Golang web in BuildWithAngga",
	// }

	// data := entity.Product{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 2}

	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 4},
		{ID: 2, Name: "Xpander", Price: 520000000, Stock: 2},
		{ID: 3, Name: "Pajero Sport", Price: 1120000000, Stock: 1},
	}

	tmplt.Execute(w, data)

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Halo dunia, saya sedang belajar golang web"))

}

func EggyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Nama saya Eggy Andika wardani"))

}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumber, err := strconv.Atoi(id)

	if err != nil || idNumber < 1 {
		http.NotFound(w, r)
		return
	}

	// fmt.Fprintf(w, "Product : %d", idNumber)

	tmplt, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is Happening, Keep Calm", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"content": idNumber,
	}
	tmplt.Execute(w, data)

}

func GetPost(w http.ResponseWriter, r *http.Request) {

	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("Ini GET"))
	case "POST":
		w.Write([]byte("Ini POSt"))

	default:
		http.Error(w, "Error is Happening", http.StatusBadRequest)
	}

}

func Form(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tmplt, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is Happening, Keep Calm", http.StatusInternalServerError)
			return
		}

		err = tmplt.Execute(w, nil)

		if err != nil {
			log.Println(err)
			http.Error(w, "Error is Happening, Keep Calm", http.StatusInternalServerError)
			return

		}
		return
	}

	http.Error(w, "Error is Happening, Keep Calm", http.StatusBadRequest)

}

func Process(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := r.ParseForm()

		if err != nil {
			log.Println(err)
			http.Error(w, "Error is Happening, Keep Calm", http.StatusInternalServerError)
			return

		}

		name := r.Form.Get("name")
		message := r.Form.Get("message")
		data := map[string]interface{}{
			"name":    name,
			"message": message,
		}

		tmplt, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is Happening, Keep Calm", http.StatusInternalServerError)
			return

		}
		err = tmplt.Execute(w, data)

		if err != nil {
			log.Println(err)
			http.Error(w, "Error is Happening, Keep Calm", http.StatusInternalServerError)
			return

		}
		// w.Write([]byte(name))
		// w.Write([]byte(message))

		return
	}

	http.Error(w, "Error is Happening, Keep Calm", http.StatusBadRequest)

}
