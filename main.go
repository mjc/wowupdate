package main

import (
	"fmt"
)

func main() {
	addonDirectories := ListAddons()
	for _, v := range addonDirectories {
		fmt.Println(GetAddon(v))
	}
}
