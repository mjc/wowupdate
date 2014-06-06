package main

import (
  "os"
  "io/ioutil"
  "regexp"
)

func getWowPath() (string) {
  path := "/Applications/World of Warcraft/Interface/Addons"
  return path
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
