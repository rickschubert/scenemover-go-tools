package transitionscene

import (
	"bytes"
	"fmt"
	"os"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rickschubert/scenemover/handlers/common"
	"github.com/rickschubert/scenemover/handlers/getscenes"
	"github.com/rickschubert/scenemover/utils"
)

type transitionScenePayload struct {
	// Work around in order to allow for zero values as valid input
	NewIndex *int   `json:"newIndex" validate:"required,numeric"`
	File     string `json:"file" validate:"required"`
}

type SceneAfterTransition struct {
	getscenes.Scene
	OldFile string `json:"oldFile"`
}

// findOriginalPositionOfSceneToMove returns the original position (0 indexed)
// of the scene that is to be moved. If the scene is not found, -1 is returned.
func findOriginalPositionOfSceneToMove(scenes []getscenes.Scene, sceneFileNameToMove string) int {
	for idx, scene := range scenes {
		if scene.File == sceneFileNameToMove {
			return idx
		}
	}
	return -1
}

func isImpossibleIndex(scenes []getscenes.Scene, newIndex int) bool {
	return newIndex < 0 || newIndex > len(scenes)-1
}

func moveSceneToIndex(scenes []getscenes.Scene, sceneFileNameToMove string, newIndex int) ([]SceneAfterTransition, error) {
	if isImpossibleIndex(scenes, newIndex) {
		return nil, fmt.Errorf("impossible index %d", newIndex)
	}
	fromIndex := findOriginalPositionOfSceneToMove(scenes, sceneFileNameToMove)
	if fromIndex == -1 {
		return nil, fmt.Errorf("scene %s not found", sceneFileNameToMove)
	}

	newScenes := make([]getscenes.Scene, len(scenes))
	copy(newScenes, scenes)

	var startIndex int
	if fromIndex < 0 {
		startIndex = len(newScenes) + fromIndex
	} else {
		startIndex = fromIndex
	}

	if startIndex >= 0 && startIndex < len(newScenes) {
		var endIndex int
		if newIndex < 0 {
			endIndex = len(newScenes) + newIndex
		} else {
			endIndex = newIndex
		}

		firstItem := splice(&newScenes, fromIndex, 1)[0]
		splice(&newScenes, endIndex, 0, firstItem)
	}

	scenedReindexed := reindexScenes(newScenes)

	return scenedReindexed, nil
}

func reindexScenes(scenes []getscenes.Scene) []SceneAfterTransition {
	var reindexedScenes []SceneAfterTransition
	leadNumberRegex := regexp.MustCompile(`^\d+_`)
	for idx, scene := range scenes {
		var lead bytes.Buffer
		newLeadingNumber := idx + 1
		if newLeadingNumber < 10 {
			lead.WriteString("0")
		}
		lead.WriteString(fmt.Sprintf("%d_", newLeadingNumber))

		newSceneFile := leadNumberRegex.ReplaceAll([]byte(scene.File), lead.Bytes())
		utils.LogInfo(fmt.Sprintf("Renaming %s to %s", scene.File, newSceneFile))
		reindexedScenes = append(reindexedScenes, SceneAfterTransition{
			Scene:   scene,
			OldFile: scene.File,
		})
		reindexedScenes[idx].File = string(newSceneFile)
	}
	return reindexedScenes
}

var Validator = validator.New()

type ValidationErrors struct {
	Field string
	Tag   string
	Value string
}

type ValidationErrorResponse struct {
	Errors  []*ValidationErrors `json:"validationErrors"`
	Message string              `json:"message"`
}

// validatePayload returns a boolean indicating whether the payload is valid or
// not. If not valid, an additional response body is returned.
func validatePayload(payload *transitionScenePayload) (success bool, errorResponse ValidationErrorResponse) {
	var errors []*ValidationErrors
	err := Validator.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el ValidationErrors
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
	}
	if len(errors) > 0 {
		return false, ValidationErrorResponse{
			Errors:  errors,
			Message: "Payload validation failed. See attached errors.",
		}
	}
	return true, ValidationErrorResponse{}
}

// renameScenes renames a list of scenes so that they match their new positions
// given a list of changes
func renameScenes(transitions []SceneAfterTransition) error {
	for _, transition := range transitions {
		if transition.OldFile == transition.File {
			continue
		}
		err := os.Rename(fmt.Sprintf("%s/%s", utils.ScenesDirectory, transition.OldFile), fmt.Sprintf("%s/%s", utils.ScenesDirectory, transition.File))
		if err != nil {
			return fmt.Errorf("failed to rename %s to %s: %w", transition.OldFile, transition.File, err)
		}
	}
	return nil
}

func Handler(context *fiber.Ctx) error {
	payload := new(transitionScenePayload)

	// Parse body into struct
	context.BodyParser(&payload)

	valid, validationErrors := validatePayload(payload)
	if !valid {
		return context.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	currentScenes := getscenes.GetScenes()
	newScenes, err := moveSceneToIndex(currentScenes, payload.File, *payload.NewIndex)
	if err != nil {
		return common.SendError(err, context)
	}

	err = renameScenes(newScenes)
	if err != nil {
		return common.SendError(err, context)
	}

	return context.JSON(newScenes)
}
