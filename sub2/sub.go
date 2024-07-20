package sub2

import (
	"embed"
	"fmt"
)

//go:embed hello.txt
var f embed.FS

func Print() {
	data, _ := f.ReadFile("hello.txt")
	fmt.Println(string(data))
}
