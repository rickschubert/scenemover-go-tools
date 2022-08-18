package getscenes

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/rsdoiel/fountain"
)

type Scene struct {
	// Title is the title of the scene
	Title string `json:"title"`
	// Body is the scene's content without the titel
	Body string `json:"body"`
	// Content is the full scene text without amends
	Content string `json:"content"`
	// File is the file name of the scene
	File string `json:"file"`
}

func getTitle(fullSceneContent []byte) (string, error) {
	var err error

	// If the title page, get its title instead
	titleRegex := regexp.MustCompile(`(?m)Title\:.+$`)
	title := titleRegex.FindString(string(fullSceneContent))
	if title != "" {
		return title, nil
	}

	// If a montage (no INT. or EXT.), get its title instead
	montageRegex := regexp.MustCompile(`(?m)(.+)?MONTAGE(.+)?$`)
	title = montageRegex.FindString(string(fullSceneContent))
	if title != "" {
		return title, nil
	}

	// If not the title page, use fountain parser
	fountainDoc, err := fountain.Parse(fullSceneContent)
	if err != nil {
		return "", err
	}
	if len(fountainDoc.Elements) == 0 {
		return "", fmt.Errorf("scene does not contain a title")
	}
	return fountainDoc.Elements[0].Content, err
}

func parseScene(fullSceneContent []byte, scenePath string) (Scene, error) {
	title, err := getTitle(fullSceneContent)
	if err != nil {
		return Scene{}, fmt.Errorf("could not parse scene '%s': %s", scenePath, err)
	}
	body := strings.Replace(string(fullSceneContent), title, "", 1)
	body = strings.TrimLeft(body, "\n")
	return Scene{
		Title:   title,
		Body:    body,
		Content: string(fullSceneContent),
		File:    filepath.Base(scenePath),
	}, nil
}
