package actions

import (
	"io/ioutil"
	"net/http"
)

func FileServer(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		panic(err)
	}
	w.Write(file)
}
