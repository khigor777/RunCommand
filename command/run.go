package command

import (
	"io/ioutil"
	"path/filepath"
	"fmt"
	"os/exec"
)

func DirList(ext string, cmd string)[]string {
	file, err := ioutil.ReadDir("./")
	res := []string{}
	if err != nil{
		panic(err)
	}
	for _, f := range file{
		if filepath.Ext(f.Name()) == ".php" {
			n:= f.Name()
			o := RunScript(&AllPaths{Dir:n, Command:cmd})
			res = append(res, fmt.Sprintf("File:%s, Message:%s", n, o))
		}
	}
	return res
}


func RunScript(p Pather) string  {
	b, e := exec.Command(p.GetCommand(), p.GetFilePath()).Output()
	if e != nil{
		panic(e)
	}
	return string(b)
}
