package update

import (
	"github.com/vincentnamle/CSFUpdater/tree"
	"os"
	"strings"
)

func AddInFilePath(localPath, url, subpath string, fileTree tree.Node) error {
	for _, child := range fileTree.Children {
		if child.IsFile {
			subpath2 := strings.ReplaceAll(subpath, "\\", "/")
			//fmt.Println("Downloading " + url + subpath2 + "/" + child.Name + "  to " + localPath + subpath + "\\" + child.Name)

			err := downloadFile(
				url+subpath2+"/"+child.Name,
				localPath+subpath+"\\"+child.Name)
			if err != nil {
				return err
			}
		} else {
			name := strings.ReplaceAll(child.Name, "/", "\\")

			//fmt.Println("Creating directory " + localPath + name)
			os.Mkdir(localPath+name, 0755)

			err := AddInFilePath(localPath, url, name, child)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
