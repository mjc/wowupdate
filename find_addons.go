package main

import (
	"io/ioutil"
	"os"
	"regexp"
)

func GetWowPath() string {
	path := "/Applications/World of Warcraft/Interface/Addons"
	return path
}

func ListAddons() (addons []os.FileInfo) {
	addonNameRegexp := regexp.MustCompile("(Blizzard_).*$")
	fileInfos, _ := ioutil.ReadDir(GetWowPath())
	for i := range fileInfos {
		if !addonNameRegexp.MatchString(fileInfos[i].Name()) {
			addons = append(addons, fileInfos[i])
		}
	}
	return addons
}
