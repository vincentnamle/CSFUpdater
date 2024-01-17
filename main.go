package main

import (
	"UpdaterProject/file"
	"UpdaterProject/tree"
	"UpdaterProject/update"
	"UpdaterProject/url"
	"fmt"
)

func main() {
	_, err := url.UrlInTree("http://localhost:8890/", "http://localhost:8890/")

	if err != nil {
		panic(err)
	}

	localTree, err := file.FilePathInTree("C:\\Users\\Vincent\\Desktop\\test", nil)
	yamlTree, _ := tree.LoadFromYaml()

	//intersection := localTree.Intersection(yamlTree)

	if err != nil {
		panic(err)
	}

	serverTree, err := url.UrlInTree("http://localhost:8890/", "http://localhost:8890/")

	toBeAddedTree := serverTree.Difference(localTree)
	go update.AddInFilePath("C:\\Users\\Vincent\\Desktop\\test", "http://localhost:8890", "", toBeAddedTree)

	toBeRemovedTree := yamlTree.Difference(serverTree)

	fmt.Println("zzz")
	toBeRemovedTree.Print()

	go update.RemoveInfilePath("C:\\Users\\Vincent\\Desktop\\test", "http://localhost:8890", "", toBeRemovedTree)

	if err != nil {
		panic(err)
	}

	newLocalTree, err := url.UrlInTree("http://localhost:8890/", "http://localhost:8890/")

	if err != nil {
		panic(err)
	}

	fmt.Println("eee")
	newLocalTree.Print()

	err = tree.SaveInYaml(&newLocalTree)
	if err != nil {
		return
	}
}
