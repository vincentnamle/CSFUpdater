package main

import (
	"UpdaterProject/file"
	"UpdaterProject/tree"
	"UpdaterProject/update"
	"UpdaterProject/url"
	"sync"
)

func UpdateFolderFromUrl(link, localPath string) error {
	localTreeChan := make(chan tree.Node)
	yamlTreeChan := make(chan tree.Node)
	serverTreeChan := make(chan tree.Node)
	errors := make(chan error)
	var wg sync.WaitGroup

	wg.Add(3) // Il y a 2 goroutines à attendre

	go func() {
		defer wg.Done()
		result, err := file.FilePathInTree(localPath, nil)
		localTreeChan <- result
		errors <- err
	}()

	go func() {
		defer wg.Done()
		result, err := tree.LoadFromYaml()
		yamlTreeChan <- result
		errors <- err
	}()

	go func() {
		defer wg.Done()
		result, err := url.UrlInTree(link, link)
		serverTreeChan <- result
		errors <- err
	}()

	// Fermer les canaux une fois que toutes les goroutines ont terminé
	go func() {
		wg.Wait()
		close(localTreeChan)
		close(yamlTreeChan)
		close(serverTreeChan)
		close(errors)
	}()

	if len(errors) > 0 {
		for i := range errors {
			if i != nil {
				return i
			}
		}
	}

	localTree := <-localTreeChan
	yamlTree := <-yamlTreeChan
	serverTree := <-serverTreeChan

	toBeAddedTree := serverTree.Difference(localTree)
	err := update.AddInFilePath(localPath, link, "", toBeAddedTree)

	if err != nil {
		return err
	}

	toBeRemovedTree := yamlTree.Difference(serverTree)

	err = update.RemoveInfilePath(localPath, link, "", toBeRemovedTree)

	if err != nil {
		return err
	}

	if err != nil {
		panic(err)
	}

	newLocalTree, err := url.UrlInTree(link, link)

	if err != nil {
		panic(err)
	}

	err = tree.SaveInYaml(&newLocalTree)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := UpdateFolderFromUrl("http://localhost:8890/.mcl", "C:\\Users\\Vincent\\Desktop\\test\\.mcl")
	if err != nil {
		return
	}
}

/*
func UpdateFolderFromUrl(link, localPath string) error {
	localTree, err := file.FilePathInTree(localPath, nil)
	yamlTree, _ := tree.LoadFromYaml()

	//intersection := localTree.Intersection(yamlTree)

	if err != nil {
		panic(err)
	}

	serverTree, err := url.UrlInTree(link, link)

	toBeAddedTree := serverTree.Difference(localTree)
	err = update.AddInFilePath(localPath, link, "", toBeAddedTree)

	if err != nil {
		return err
	}

	toBeRemovedTree := yamlTree.Difference(serverTree)

	err = update.RemoveInfilePath(localPath, link, "", toBeRemovedTree)

	if err != nil {
		return err
	}

	if err != nil {
		panic(err)
	}

	newLocalTree, err := url.UrlInTree(link, link)

	if err != nil {
		panic(err)
	}

	err = tree.SaveInYaml(&newLocalTree)

	if err != nil {
		return err
	}

	return nil
}
*/
