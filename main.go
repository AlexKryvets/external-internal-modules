package main

import (
	"fmt"
	pkg2 "github.com/AlexKryvets/external-internal-modules/internal/package"
	pkg1 "github.com/AlexKryvets/external-internal-modules/package"
)

func main() {
	fmt.Println(pkg1.GetMessage())
	fmt.Println(pkg2.GetMessage())
}
