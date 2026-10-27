package pages

import (
	"embed"
	"html/template"
	"log"
)

// Les gabarits sont embarqués dans le binaire : ils fonctionnent partout (local, Vercel).
//
//go:embed *.html
var files embed.FS

var Temp *template.Template

// Fonction custom utilisable dans les templates
func add1(x int) int {
	return x + 1
}

func Init() {
	var err error
	Temp, err = template.New("").Funcs(template.FuncMap{"add1": add1}).ParseFS(files, "*.html")
	if err != nil {
		log.Fatalf("Erreur lors de l'initialisation des templates : %v", err)
	}
	log.Println("Templates chargés (embarqués)")
}
