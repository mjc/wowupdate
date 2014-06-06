package main

import (
  "testing"
  "regexp"
)

func TestGetWowPath(t *testing.T) {
  const real_path = "/Applications/World of Warcraft/Interface/Addons"
  if path := getWowPath(); path != real_path {
    t.Errorf("WowPath = %v, want %v", path, real_path)
  }
}

func TestListAddons(t *testing.T) {
  BlizzardAddonNameRegexp := regexp.MustCompile("^(Blizzard_).+$")
  for _, v := range ListAddons() {
    if BlizzardAddonNameRegexp.MatchString(v.Name()) {
      t.Errorf("Found Blizzard addon in addon list: %v", v.Name())
    }
  }
}
