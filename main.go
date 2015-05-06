package main

import (
	"fmt"
)

func main() {
	addonDirectories := ListAddons()
	for _, v := range addonDirectories {
		addon, _ := GetAddon(v)
		fmt.Println(addon.updateMethod)
	}
}
