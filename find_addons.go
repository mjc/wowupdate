package main

import (
  "os"
  "io/ioutil"
  "regexp"
)

func getWowPath() string {
	wowpath := ""
	switch runtime.GOOS {
	case "darwin":
		wowpath = "/Applications/World of Warcraft"
	case "windows":
		wowpath = "X:/Games/WoW"
	default:
		log.Fatal("No location for '%s' OS", runtime.GOOS)
	}
	return path.Join(wowpath, "Interface", "Addons")
}

func ListAddons() (addons []os.FileInfo) {
  addonNameRegexp := regexp.MustCompile("(Blizzard_).*$")
  fileInfos, _ := ioutil.ReadDir(getWowPath())
  for i := range fileInfos {
    if fileInfos[i].IsDir() && !addonNameRegexp.MatchString(fileInfos[i].Name()) {
      addons = append(addons,fileInfos[i])
    }
  }
  return addons
}
