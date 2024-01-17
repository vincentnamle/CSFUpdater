package update

import (
	"UpdaterProject/tree"
	"fmt"
	"os"
	"strings"
)

func RemoveInfilePath(localPath, url, subpath string, fileTree tree.Node) error {
	for _, child := range fileTree.Children {
		if child.IsFile {
			subpath2 := strings.ReplaceAll(subpath, "\\", "/")
			fmt.Println("Downloading " + url + subpath2 + "/" + child.Name + "  to " + localPath + subpath + "\\" + child.Name)

			os.Remove(localPath + subpath + "\\" + child.Name)
		} else {
			name := strings.ReplaceAll(child.Name, "/", "\\")

			fmt.Println("Creating directory " + localPath + name)

			err := os.RemoveAll(localPath + name)

			err = RemoveInfilePath(localPath, url, name, child)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
