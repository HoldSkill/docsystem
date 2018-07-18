package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/astaxie/beego"
	"github.com/holdskill/docsystem/conf"
)

//检查最新版本.
func CheckUpdate() {

	fmt.Println("System current version => ", conf.VERSION)

	os.Exit(0)

}
