package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"log"
)

func DirList(ext string, cmd string, path string) []string {
	file, err := ioutil.ReadDir(GetCurDir())
	res := []string{}
	if err != nil {
		panic(err)
	}
	for _, f := range file {
		if filepath.Ext(f.Name()) == ext {
			fn := f.Name()
			o := RunScript(&AllPaths{ AllCommand:&AllCommand{cmd},  Dir:fn})
			res = append(res, fmt.Sprintf("File:%s, Message:%s", fn, o))
		}
	}
	return res
}

func GetCurDir() string {
	d, e := os.Getwd()
	if e != nil {
		panic(e)
	}
	return d
}

func RunScript(p Pather) string {
	c := exec.Command(p.GetCommand(), p.GetFilePath())
	b, e := c.Output()
	if e != nil {
		log.Fatal(p)
	}
	return string(b)
}
