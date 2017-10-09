package command

import (
	"path/filepath"
	"os"
	"strings"
)

type Pather interface {
	GetFilePath()string
	GetCommand()string
}

type AllPaths struct {
	Dir string
	Command string
}

func (ap *AllPaths) GetFilePath() string{
	r, e := filepath.Abs(ap.Dir)
	if e != nil{
		panic(e)
	}
	return r
}

func (ap *AllPaths) GetCommand() string  {
	return trim(ap.Command)
}


type CurrentPath struct {
	Dir string
	Command string
}


func (cp *CurrentPath) GetFilePath() string {
	if _, err:= os.Stat(cp.Dir); err != nil{
		panic(err)
	}
	return cp.Dir
}

func (cp *CurrentPath) GetCommand()string  {
	return trim(cp.Command)
}


func trim(s string) string  {
	return strings.Trim(s, " ")
}