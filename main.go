package main

import (
	"fmt"
	"os"
)

//Print chars constants
const (
	LastItemInDirChar          = "\u2514"
	ItemNameEmphasizeChar      = "\u2500"
	ItemNameStartChar          = "\u251C"
	DirectoryVerticalSpaceChar = "\u007C"
	SpaceChar 				   = "\u0020"
)

//Directory tree print struct, have root directory and print only dirs mode
type Tree struct {
	root string
	directory bool
}

//Print file system tree entrance function
func (tree *Tree) Print() error {
	return tree.printContent(tree.root, make([]int, 0))
}

//Filter directories function
func (tree *Tree) FilterDirectories(dirs []os.FileInfo) []os.FileInfo {
	if tree.directory {
		directories := make([]os.FileInfo, 0)

		for _, file := range dirs {
			if file.IsDir() {
				directories = append(directories, file)
			}
		}
		return directories
	}
	return dirs
}

//Space builder function by slice
func SpaceBuilder(spaces []int) string {
	space := ""

	for _, v := range spaces {
		if v > 0 {
			space += DirectoryVerticalSpaceChar
		} else {
			space += SpaceChar
		}
		space += SpaceChar + SpaceChar
	}

	return space
}

//Recursive print content of directory function
func (tree * Tree)printContent(path string, spaces []int) error {
	dir, err := os.Open(path)

	if err != nil {
		return err
	}

	dirs, err := dir.Readdir(-1)

	dir.Close()

	if err != nil {
		return err
	}

	dirs = tree.FilterDirectories(dirs)

	if len(dirs) == 0 {
		return nil
	}

	lastFile := dirs[len(dirs) - 1]

	for _, file := range dirs {
		space := SpaceBuilder(spaces)

		fileSign := ItemNameStartChar

		if file == lastFile {
			fileSign = LastItemInDirChar
		}

		if tree.directory && file.IsDir() || !tree.directory {
			fmt.Println(space + fileSign + ItemNameEmphasizeChar + file.Name())
		}

		if file.IsDir() {
			space := 0

			if file != lastFile {
				space = 1
			}

			spaces = append(spaces, space)

			tree.printContent(path + string(os.PathSeparator) + file.Name(), spaces)

			if file != lastFile {
				spaces = spaces[:len(spaces) - 1]
			}

		}
	}

	return nil
}

//Get command line arguments logic
func GetArguments(args []string) (string , bool){
	root := ""
	directory := false

	//As parameters don't have any order so do some routines
	if len(args) < 2 || (len(args) < 3 && args[1] == "-d"){
		root, _ = os.Getwd()

		if !(len(args) < 2) {
			directory = true
		}
	} else if len(args) < 3 {
		root = os.Args[1]
	} else {
		directory = true
		root = args[1]

		if args[1] == "-d" {
			root = args[2]
		}
	}

	return root, directory
}

func main () {

	root, directory := GetArguments(os.Args)

	tree := Tree{root, directory}
	err := tree.Print()

	if err != nil {
		fmt.Println(err)
	}
}

