package transitionscene

import (
	"testing"

	"github.com/rickschubert/scenemover/handlers/getscenes"
	"github.com/stretchr/testify/assert"
)

func TestMoveSceneToIndex(t *testing.T) {
	testCases := []struct {
		name     string
		scenes   []getscenes.Scene
		file     string
		newIndex int
		expected []SceneAfterTransition
	}{
		{
			name: "moving first item to back",
			scenes: []getscenes.Scene{
				{
					File: "01_a.fountain",
				},
				{
					File: "02_b.fountain",
				},
				{
					File: "03_c.fountain",
				},
			},
			file:     "01_a.fountain",
			newIndex: 2,
			expected: []SceneAfterTransition{
				{
					Scene: getscenes.Scene{
						File: "01_b.fountain",
					},
					OldFile: "02_b.fountain",
				},
				{
					Scene: getscenes.Scene{
						File: "02_c.fountain",
					},
					OldFile: "03_c.fountain",
				},
				{
					Scene: getscenes.Scene{
						File: "03_a.fountain",
					},
					OldFile: "01_a.fountain",
				},
			},
		},
		{
			name: "moving last item to the front",
			scenes: []getscenes.Scene{
				{
					File: "01_a.fountain",
				},
				{
					File: "02_b.fountain",
				},
				{
					File: "03_c.fountain",
				},
			},
			file:     "03_c.fountain",
			newIndex: 0,
			expected: []SceneAfterTransition{
				{
					Scene: getscenes.Scene{
						File: "01_c.fountain",
					},
					OldFile: "03_c.fountain",
				},
				{
					Scene: getscenes.Scene{
						File: "02_a.fountain",
					},
					OldFile: "01_a.fountain",
				},
				{
					Scene: getscenes.Scene{
						File: "03_b.fountain",
					},
					OldFile: "02_b.fountain",
				},
			},
		},
		{
			name: "moving first item to the same position it already was in",
			scenes: []getscenes.Scene{
				{
					File: "01_a.fountain",
				},
				{
					File: "02_b.fountain",
				},
				{
					File: "03_c.fountain",
				},
			},
			file:     "01_a.fountain",
			newIndex: 0,
			expected: []SceneAfterTransition{
				{
					Scene: getscenes.Scene{
						File: "01_a.fountain",
					},
					OldFile: "01_a.fountain",
				},
				{
					Scene: getscenes.Scene{
						File: "02_b.fountain",
					},
					OldFile: "02_b.fountain",
				},
				{
					Scene: getscenes.Scene{
						File: "03_c.fountain",
					},
					OldFile: "03_c.fountain",
				},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			originalScenes := make([]getscenes.Scene, len(testCase.scenes))
			copy(originalScenes, testCase.scenes)
			newScenes, err := moveSceneToIndex(testCase.scenes, testCase.file, testCase.newIndex)
			assert.Equal(t, testCase.expected, newScenes)
			// Generates a copy instead of amending the source slice
			assert.Equal(t, testCase.scenes, originalScenes)
			assert.NoError(t, err)
		})
	}
}

func TestMoveSceneToIndexWithInvalidScene(t *testing.T) {
	testCases := []struct {
		name          string
		scenes        []getscenes.Scene
		file          string
		newIndex      int
		expectedError string
	}{
		{
			name: "moving first item to back",
			scenes: []getscenes.Scene{
				{
					File: "01_a.fountain",
				},
				{
					File: "02_b.fountain",
				},
				{
					File: "03_c.fountain",
				},
			},
			file:          "04_d.fountain",
			newIndex:      2,
			expectedError: "scene 04_d.fountain not found",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := moveSceneToIndex(testCase.scenes, testCase.file, testCase.newIndex)
			assert.Equal(t, testCase.expectedError, err.Error())
		})
	}
}

func TestMoveSceneToIndexWithImpossibleIndex(t *testing.T) {
	testCases := []struct {
		name          string
		scenes        []getscenes.Scene
		file          string
		newIndex      int
		expectedError string
	}{
		{
			name: "moving first item to back",
			scenes: []getscenes.Scene{
				{
					File: "01_a.fountain",
				},
				{
					File: "02_b.fountain",
				},
				{
					File: "03_c.fountain",
				},
			},
			file:          "01_a.fountain",
			newIndex:      3,
			expectedError: "impossible index 3",
		},
		{
			name: "moving first item to back",
			scenes: []getscenes.Scene{
				{
					File: "01_a.fountain",
				},
				{
					File: "02_b.fountain",
				},
				{
					File: "03_c.fountain",
				},
			},
			file:          "01_a.fountain",
			newIndex:      -23,
			expectedError: "impossible index -23",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := moveSceneToIndex(testCase.scenes, testCase.file, testCase.newIndex)
			assert.Equal(t, testCase.expectedError, err.Error())
		})
	}
}
