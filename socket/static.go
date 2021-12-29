package socket

import "net/http"

func MainStatic() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":9999", nil)
}
