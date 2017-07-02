// main.go

package main

import (
	"github.com/dlaize/homedatakeeper/database"
	"github.com/dlaize/homedatakeeper/util"
)

func main() {
	util.InitLogger()
	a := App{}
	database.Initialize()
	a.Initialize()
	a.Run(":8000")
}
