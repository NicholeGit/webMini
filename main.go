package main

import (
	_ "github.com/NicholeGit/webMini/frontend"
	_ "github.com/NicholeGit/webMini/frontend/user"
	"github.com/NicholeGit/webMini/util"
)

func main() {
	util.Run(":80")
}
