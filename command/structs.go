package command

import (
	"os"
	"path/filepath"
	"strings"
)

type Pather interface {
	GetFilePath() string
	GetCommand() string
}

type AllCommand struct {
	Command string
}

func (ac *AllCommand) GetCommand() string {
	return trim(ac.Command)
}

type AllPaths struct {
	*AllCommand
	Dir string
}

func (ap *AllPaths) GetFilePath() string {
	r, e := filepath.Abs(ap.Dir)
	if e != nil {
		panic(e)
	}
	return r
}

type CurrentPath struct {
	*AllCommand
	Dir string
}

func (cp *CurrentPath) GetFilePath() string {
	if _, err := os.Stat(cp.Dir); err != nil {
		panic(err)
	}
	return cp.Dir
}

func trim(s string) string {
	return strings.Trim(s, " ")
}
