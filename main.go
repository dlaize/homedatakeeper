// main.go

package main

import (
	"github.com/dlaize/homedatakeeper/database"
	"github.com/dlaize/homedatakeeper/util"
)

func main() {
	util.InitLogger()
	a := App{}
	database.Connect("5432")
	a.Initialize()
    	a.Run(":8000")
}
