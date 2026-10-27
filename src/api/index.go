// Point d'entrée Vercel : même routeur que main.go, servi par une fonction serverless.
package handler

import (
	"net/http"
	"sync"

	initTemp "groupie/pages"
	"groupie/router"
)

var (
	once sync.Once
	mux  *http.ServeMux
)

func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(func() {
		initTemp.Init()
		mux = router.New()
		mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(initTemp.BaseDir()+"/static"))))
	})
	mux.ServeHTTP(w, r)
}
