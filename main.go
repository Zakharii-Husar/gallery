package main

import (
	"gallery/data"
)

func main() {
    data.Init()
    defer data.DB.Close()
    r := SetupRouter()
    r.Run()
}
