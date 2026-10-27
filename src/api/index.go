// Point d'entrée Vercel : même routeur que main.go, servi par une fonction serverless.
// Les fichiers de static/ sont servis directement par Vercel (route "filesystem").
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
	})
	mux.ServeHTTP(w, r)
}
