package update

import (
	"io"
	"net/http"
	"os"
)

func downloadFile(url string, localPath string) error {
	// Envoyer une requête GET à l'URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Ouvrir un fichier pour écrire les données téléchargées
	out, err := os.Create(localPath)
	_ = err
	if err != nil {
		return err
	}
	defer out.Close()

	// Copier les données depuis le Body de la réponse vers le fichier
	_, err = io.Copy(out, resp.Body)
	return err
}

func removeFirstChar(s string) string {
	runes := []rune(s)
	if len(runes) > 0 {
		s = string(runes[1:])
	}
	return s
}
