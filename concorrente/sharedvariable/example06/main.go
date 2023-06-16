package main

import (
	"sync"
)

var (
	keys          map[string]string
	loadIconsOnce sync.Once
)

func loadKeys() {
	keys = make(map[string]string)
	keys["file01"] = "file01.txt"
	keys["file02"] = "file02.txt"
	keys["file03"] = "file03.txt"
	keys["file04"] = "file04.txt"
	keys["file05"] = "file05.txt"
	keys["file06"] = "file06.txt"
}

func main() {
	loadIconsOnce.Do(loadKeys)
}

