package main

import (
	"fmt"
)

func main() {
	addonDirectories := ListAddons()
	fmt.Println(GetAddon(addonDirectories[0]))
}
