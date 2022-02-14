package getscenes

import (
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
	"github.com/rickschubert/scenemover/utils"
)

func GetScenes() []Scene {
	scenePaths := utils.GetScenePaths()
	scenes := make([]Scene, 0, len(scenePaths))
	for _, scenePath := range scenePaths {
		content, err := ioutil.ReadFile(scenePath)
		if err != nil {
			utils.LogInfo("Error reading scene at location %s: %s", scenePath, err)
		}
		parsedScene, err := parseScene(content, scenePath)
		if err != nil {
			panic(err)
		}
		scenes = append(scenes, parsedScene)
	}
	return scenes
}

func Handler(context *fiber.Ctx) error {
	return context.JSON(GetScenes())
}
