package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var t, command, path string
var dir string = "c:/"

func main() {

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {

		txt := sc.Text()
		command = txt
		for i, v := range txt {
			if string(v) == string(" ") {
				command = string(txt[0:i])
				path = string(txt[i+1:])
			}
		}

		if command == string("cd") && len(path) != 0 {
			dir = dir + path + "/"

		}
		if command == string("cd") && len(path) == 0 {
			dir = path + "/"

		}
		if command == string("cd..") {
			dir = dir[0:len(path)] + "/"

		}

		if command == string("dir") {

			filesFromDir, err := ioutil.ReadDir(dir)
			if err != nil {
				fmt.Println(err)
			}

			for _, file := range filesFromDir {
				if file.IsDir() {
					t = "Directory: "
				} else {
					t = "File: "
				}

				fmt.Print(t)
				fmt.Printf("%s, size: %d\n", file.Name(), file.Size())
			}
		}
		path = ""
		command = ""
	}
}
