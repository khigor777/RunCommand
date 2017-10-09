package command

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"os"
)

func DirList(ext string, cmd string) []string {
	file, err := ioutil.ReadDir(getCurDir())
	res := []string{}
	if err != nil {
		panic(err)
	}
	for _, f := range file {
		if filepath.Ext(f.Name()) == ext {
			n := f.Name()
			o := RunScript(&AllPaths{Dir: n, Command: cmd})
			res = append(res, fmt.Sprintf("File:%s, Message:%s", n, o))
		}
	}
	return res
}

func getCurDir()string {
	d, e := os.Getwd()
	if e != nil {
		panic(e)
	}
	return d
}

func RunScript(p Pather) string {
	b, e := exec.Command(p.GetCommand(), p.GetFilePath()).Output()
	if e != nil {
		panic(e)
	}
	return string(b)
}
