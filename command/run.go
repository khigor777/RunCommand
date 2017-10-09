package command

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"os"
)

func DirList(ext string, cmd string) []string {
	dir := getCurDir()
	file, err := ioutil.ReadDir(getCurDir())
	res := []string{}
	if err != nil {
		panic(err)
	}
	for _, f := range file {
		if filepath.Ext(f.Name()) == ext {
			fn := dir+"/"+f.Name()
			o := RunScript(&AllPaths{Dir: fn, Command: cmd})
			res = append(res, fmt.Sprintf("File:%s, Message:%s", fn, o))
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
