package file

import (
	"UpdaterProject/tree"
	"os"
	"strings"
)

func FilePathInTree(filePath string, yamlStruct map[string]interface{}) (tree.Node, error) {
	return filePathInTree(filePath, filePath, "", yamlStruct)
}

func filePathInTree(filePath, baseUrl, subPath string, yamlStruct map[string]interface{}) (tree.Node, error) {
	files, err := os.ReadDir(filePath)

	if err != nil {
		return tree.NewEmptyNode(), err
	}

	var children []tree.Node

	for _, file := range files {
		if file.IsDir() {

			name := strings.ReplaceAll(filePath+"\\"+file.Name(), "\\", "/")

			childNode, err := filePathInTree(name, baseUrl, file.Name(), yamlStruct)

			if err != nil {
				return tree.NewEmptyNode(), err
			}

			children = append(children, childNode)
		} else {
			children = append(children, tree.NewFileNode(file.Name()))
		}
	}

	name := strings.ReplaceAll(filePath, baseUrl, "")
	baseUrl = strings.ReplaceAll(baseUrl, "\\", "/")

	name = strings.ReplaceAll(name, baseUrl, "")

	return tree.NewFolderNode(name, children), nil
}
