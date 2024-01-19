package main

import (
	"github.com/vincentnamle/CSFUpdater/update"
)

func main() {
	err := update.UpdateFolderFromUrl("http://localhost:8890/.mcl", "C:\\Users\\Vincent\\Desktop\\test\\.mcl")
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
