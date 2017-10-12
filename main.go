package main

import (
	"flag"
	"fmt"
	"khigor777/runcommand/command"
	"strings"
)

func main() {
	cmd := flag.String("cmd", "php", "-cmd=php")
	ext := flag.String("ext", ".php", "-ext=.php")
	dir := flag.String("dir", command.GetCurDir(""), "-dir=C:/log")
	file := flag.String("file", "", "file=file name with dir work only with ctype=2")
	flag.Usage = func() {
		fmt.Println(
			"Example#1: RunScript.exe -cmd=php -ext=.php -dir=C:/  :Run all php files in current Dir \n" +
				"Example#2: RunScript.exe -cmd=php -ext=.php -file=C:/test.php  :Run this one php file ")
	}

	flag.Parse()

	if *file == "" {
		fmt.Println(strings.Join(command.RunDirScript(*ext, *cmd, *dir), "\n"))
	} else {
		fmt.Println(command.RunScript(&command.CurrentPath{
			AllCommand: &command.AllCommand{*cmd},
			Dir:        *file,
		}))
	}
}
