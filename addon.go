package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Addon struct {
	title          string
	author         string
	notes          string
	version        string
	revision       string
	dependencies   []string
	optionalDeps   []string
	savedVariables []string
	x              map[string]string
	path           string
	files          []string
}

func GetAddon(dir os.FileInfo) (addon Addon, err error) {
	var files []string
	files, err = listAddonFiles(dir)
	toc := findTocInList(files)
	params, files := parseToc(toc)

	addon = Addon{
		title:          params["Title"],
		author:         params["Author"],
		notes:          params["Notes"],
		version:        params["Version"],
		dependencies:   parseDependencies(params["Dependencies"]),
		optionalDeps:   parseDependencies(params["OptionalDeps"]),
		savedVariables: parseDependencies(params["SavedVariables"]),
		revision:       params["Revision"],
		files:          files,
		path:           dir.Name(),
	}

	return addon, err
}

func parseDependencies(depList string) (deps []string) {
	split := strings.SplitN(depList, ",", 2)
	for _, v := range split {
		deps = append(deps, strings.TrimSpace(v))
	}
	return deps
}

func parseToc(path string) (toc map[string]string, files []string) {
	toc = make(map[string]string)
	f, err := os.Open(path)
	if err != nil {
		fmt.Errorf("Could not open %v", path)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if tocLineIsComment(line) && len(line) > 3 {
			key, value := tocLineParseMetaData(line)
			toc[key] = value
		} else if len(line) > 2 {
			files = append(files, line)
		}
	}
	return toc, files
}

func tocLineIsComment(line string) (isComment bool) {
	return strings.HasPrefix(line, "##")
}

func tocLineParseMetaData(line string) (key, value string) {
	trimmed := strings.TrimSpace(strings.TrimPrefix(line, "##"))
	split := strings.SplitN(trimmed, ":", 2)
	return split[0], split[1]
}

func listAddonFiles(dir os.FileInfo) (files []string, err error) {
	if !dir.IsDir() {
		return files, fmt.Errorf("%v is not a directory", dir.Name())
	}
	pattern := filepath.Join(GetWowPath(), dir.Name(), "*")
	files, err = filepath.Glob(pattern)
	return files, err
}

func findTocInList(files []string) (file string) {
	for _, value := range files {
		matches, _ := regexp.MatchString("\\.toc$", value)
		if matches {
			file = value
		}
	}
	return file
}
