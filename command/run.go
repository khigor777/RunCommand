package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func RunDirScript(ext string, cmd string, path string) []string {
	file, err := ioutil.ReadDir(GetCurDir(path))
	res := []string{}
	if err != nil {
		panic(err)
	}
	for _, f := range file {
		if filepath.Ext(f.Name()) == ext {
			fn := path + "/" + f.Name()

			o := RunScript(&AllPaths{AllCommand: &AllCommand{cmd}, Dir: fn})
			res = append(res, fmt.Sprintf("File:%s, Message:%s", fn, o))
		}
	}
	return res
}

func RunScript(p Pather) string {
	c := exec.Command(p.GetCommand(), p.GetFilePath())
	b, e := c.Output()
	if e != nil {
		panic(p)
	}
	return string(b)
}

func GetCurDir(path string) string {
	if path != "" {
		return path
	}
	d, e := os.Getwd()
	if e != nil {
		panic(e)
	}
	return d
}
