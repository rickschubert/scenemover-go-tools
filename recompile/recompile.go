package recompile

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/rickschubert/scenemover/utils"
)

const scriptLocation = "./script.fountain"

func recompile() {
	scenes := utils.GetScenePaths()

	var fullScript bytes.Buffer

	for _, scene := range scenes {
		content, err := os.ReadFile(scene)
		if err != nil {
			utils.LogInfo("Error reading scene at location %s: %s", scene, err)
		}
		// Add scene
		fullScript.Write(content)
		// Add empty line
		fullScript.Write([]byte("\n"))
	}

	ioutil.WriteFile(scriptLocation, fullScript.Bytes(), 0644)
	log.Println("Finished recompiling")
}

func logDebugEvents(event fsnotify.Event) {
	log.Println("event:", event)
}

func enableWatching(watcher *fsnotify.Watcher) {
	err := watcher.Add(utils.ScenesDirectory)
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Add(scriptLocation)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Enabled watching")
}

func disableWatching(watcher *fsnotify.Watcher) {
	err := watcher.Remove(utils.ScenesDirectory)
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Remove(scriptLocation)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Disabled watching")
}

func LaunchWatcher() {
	utils.LogInfo("Watching %s directory for changes.", utils.ScenesDirectory)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				logDebugEvents(event)
				disableWatching(watcher)
				if event.Name == filepath.Base(scriptLocation) {
					log.Println("You changed the script - let's do something about it!")
				} else {
					recompile()
				}
				enableWatching(watcher)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	enableWatching(watcher)
	<-done
}
