package url

import (
	"github.com/vincentnamle/CSFUpdater/tree"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func UrlInTree(url, baseUrl string) (tree.Node, error) {
	resp, err := http.Get(url)

	if err != nil {
		return tree.NewEmptyNode(), err
	}

	defer resp.Body.Close()
	tokenizer := html.NewTokenizer(resp.Body)
	children := []tree.Node{}

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			// Fin de l'analyse, retournez sans erreur
			return tree.NewFolderNode(strings.Replace(url, baseUrl, "", -1), children), nil
		}

		token := tokenizer.Token()

		if tokenType != html.StartTagToken || token.Data != "a" {
			continue
		}

		for _, attr := range token.Attr {
			if attr.Key != "href" {
				continue
			}

			link := attr.Val

			if isRoot(link) {
				continue
			}

			if isFile(link) {
				children = append(children, tree.NewFileNode(link))

			} else if isDirectory(link) {
				// Enlever le / à la fin du dossier
				link = link[:len(link)-1]

				childFolderNode, err := UrlInTree(url+"/"+link, baseUrl)

				if err != nil {
					return tree.NewEmptyNode(), err
				}

				children = append(children, childFolderNode)
			}
		}
	}
}

func isRoot(link string) bool {
	return link == "/"
}

func isFile(link string) bool {
	return !isDirectory(link) && !strings.HasPrefix(link, "/")
}

func isDirectory(link string) bool {
	return strings.HasSuffix(link, "/") && !strings.HasPrefix(link, "/")
}
