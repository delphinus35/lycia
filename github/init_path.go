package github

import (
	"fmt"
	"github.com/delphinus35/lycia/util"
	"os"
	"path"
)

func InitPath(pathStr string) (err error) {
	dir, _ := path.Split(pathStr)
	stat, err := os.Stat(dir)
	if err != nil || !stat.IsDir() {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			err = util.LyciaError(fmt.Sprintf("cannot mkdir: '%s'", dir))
			return
		}
	}
	return
}
