package main

import (
	"log"
)

func main() {
	addonDirectories := ListAddons()
	for _, v := range addonDirectories {
		addon, _ := GetAddon(v)
		if addon.updateMethod == "git" {
			err := GitUpdate(addon)
			if err != nil {
				log.Fatal(err)
			}
		} else if addon.updateMethod == "svn" {
			err := SvnUpdate(addon)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
