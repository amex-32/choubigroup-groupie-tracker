package pages

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

var Temp *template.Template

// Fonction custom utilisable dans les templates
func add1(x int) int {
	return x + 1
}

// BaseDir renvoie le dossier qui contient "pages" et "static". En local c'est le dossier courant ;
// sur Vercel les fichiers inclus peuvent se trouver dans un sous-dossier "src".
func BaseDir() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Erreur récupération chemin actuel : %v", err)
	}
	for _, dir := range []string{wd, filepath.Join(wd, "src")} {
		if _, err := os.Stat(filepath.Join(dir, "pages")); err == nil {
			return dir
		}
	}
	return wd
}

func Init() {
	Temp = template.New("").Funcs(template.FuncMap{
		"add1": add1,
	})

	templatesPath := filepath.Join(BaseDir(), "pages")

	err := filepath.Walk(templatesPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			_, e := Temp.ParseFiles(path)
			if e != nil {
				log.Printf("Erreur parse template %s : %v", path, e)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Erreur lors de l'initialisation des templates : %v", err)
	}

	log.Println("Templates chargés depuis :", templatesPath)
}
