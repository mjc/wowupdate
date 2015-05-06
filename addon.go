package main

import (
	"bufio"
	"fmt"
	"log"
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
	updateMethod   string
}

func GetAddon(dir os.FileInfo) (addon Addon, err error) {
	var files []string
	files, err = listAddonFiles(dir)
	toc := findTocInList(files)
	params, tocFileList := parseToc(toc)

	addon = Addon{
		title:          params["Title"],
		author:         params["Author"],
		notes:          params["Notes"],
		version:        params["Version"],
		dependencies:   parseDependencies(params["Dependencies"]),
		optionalDeps:   parseDependencies(params["OptionalDeps"]),
		savedVariables: parseDependencies(params["SavedVariables"]),
		revision:       params["Revision"],
		files:          tocFileList,
		path:           dir.Name(),
		updateMethod:   getUpdateMethod(files),
	}

	return addon, err
}

func getUpdateMethod(files []string) (updateMethod string) {
	updateMethod = "curse"
	for _, value := range files {
		if isGit(value) {
			updateMethod = "git"
		} else if isSvn(value) {
			updateMethod = "svn"
		}
	}
	return updateMethod
}

func isGit(dir string) (result bool) {
	matches, _ := regexp.MatchString("^\\.git$", filepath.Base(dir))
	return matches
}

func isSvn(dir string) (result bool) {
	matches, _ := regexp.MatchString("^\\.svn$", filepath.Base(dir))
	return matches
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
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if tocLineIsComment(line) && len(line) > 3 {
			split := tocLineParseMetaData(line)
			if len(split) >= 2 {
				toc[split[0]] = split[1]
			}
		} else if len(line) > 2 {
			files = append(files, line)
		}
	}
	return toc, files
}

func tocLineIsComment(line string) (isComment bool) {
	return strings.HasPrefix(line, "##")
}

func tocLineParseMetaData(line string) (split []string) {
	trimmed := strings.TrimSpace(strings.TrimPrefix(line, "##"))
	split = strings.SplitN(trimmed, ":", 2)
	return split
}

func listAddonFiles(dir os.FileInfo) (files []string, err error) {
	if !dir.IsDir() {
		return files, fmt.Errorf("%v is not a directory", dir.Name())
	}
	pattern := filepath.Join(getWowPath(), dir.Name(), "*")
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
