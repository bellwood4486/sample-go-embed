package sub3

import (
	"embed"
	"fmt"
)

//go:embed dir/*.txt
var f embed.FS

func Print() {
	data1, _ := f.ReadFile("dir/hello.txt")
	data2, _ := f.ReadFile("dir/world.txt")
	fmt.Println(string(data1))
	fmt.Println(string(data2))
}
