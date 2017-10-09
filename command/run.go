package command

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"os"

)

func DirList(ext string, cmd string) []string {
	dir := GetCurDir()
	file, err := ioutil.ReadDir(GetCurDir())
	res := []string{}
	if err != nil {
		panic(err)
	}
	for _, f := range file {
		if filepath.Ext(f.Name()) == ext {
			fn := dir+"/"+f.Name()
			o := RunScript(dir, &AllPaths{Dir: fn, Command: cmd})
			res = append(res, fmt.Sprintf("File:%s, Message:%s", fn, o))
		}
	}
	return res
}

func GetCurDir()string {
	d, e := os.Getwd()
	if e != nil {
		panic(e)
	}
	return d
}

func RunScript(cDir string, p Pather) string {
	b := exec.Command(p.GetCommand(), p.GetFilePath())
	b.Dir = cDir
	s, e := b.Output()
	if e != nil{
		panic(e)
	}
	return string(s)
}
