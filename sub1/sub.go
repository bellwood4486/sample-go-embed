package sub1

import (
	_ "embed"
	"fmt"
)

//go:embed hello.txt
var s string

func Print() {
	fmt.Println(s)
}
