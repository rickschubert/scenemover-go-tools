package utils

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func GetScenePaths() []string {
	var scenes []string
	files, err := ioutil.ReadDir(ScenesDirectory)
	if err != nil {
		LogInfo("Error reading %s directory: %s", ScenesDirectory, err)
	}
	fountainFileRegex := regexp.MustCompile(`.+\.fountain$`)
	for _, f := range files {
		isFountainFile := fountainFileRegex.Match([]byte(f.Name()))
		if !f.IsDir() && isFountainFile {
			scenes = append(scenes, fmt.Sprintf("%s/%s", ScenesDirectory, f.Name()))
		}
	}
	return scenes
}
