package command

import (
	"strings"
	"testing"

	"path"
	"runtime"
)

func TestDirList(t *testing.T) {
	d := GetBaseDir()

	r := RunDirScript(".php", "php", d+"/testdata")
	for _, v := range r {

		if strings.Index(string(v), "Message1") == -1 {
			t.Error("It's not working")

		}
	}

}

func TestRunScript(t *testing.T) {
	//RunScript(&CurrentPath{Dir:"23", Command:"php"})
}

func GetBaseDir() string {
	_, k, _, _ := runtime.Caller(0)
	return path.Dir(k)
}
