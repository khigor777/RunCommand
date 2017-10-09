package main

import (
	"fmt"
	"khigor777/runcommand/command"
	"strings"
	"flag"
)



func main()  {
	cmd := flag.String("cmd", "php", "-cmd=php")
	ext := flag.String("ext", ".php", "-ext=.php")
	file := flag.String("file", "", "file=file name with dir work only with ctype=2")
	ctype := flag.Int("ctype", 1, "ctype=1 or 2; if 1 all scripts run; if 2 only one run with param file")
	flag.Parse()
	if *ctype == 1{
		fmt.Println(strings.Join(command.DirList(*ext, *cmd), "\n"))
	}else if *ctype == 2{
		fmt.Println(command.RunScript(&command.CurrentPath{
			Dir:*file,
			Command:*cmd,
		}))

	}
}
