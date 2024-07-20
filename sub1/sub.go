package sub1

import (
	_ "embed"
	"fmt"
)

//go:embed hello.txt
var s string

//go:embed hello.txt
var b []byte

func Print() {
	fmt.Println(s)
	fmt.Println(string(b))
}
