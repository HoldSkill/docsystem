package commands

import (
	"fmt"
	"os"
	"github.com/holdskill/docsystem/conf"
)

//检查最新版本.
func CheckUpdate() {

	fmt.Println("System current version => ", conf.VERSION)

	os.Exit(0)

}
