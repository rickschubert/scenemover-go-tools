package getscenes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseStandardScene(t *testing.T) {
	standardScene := `
INT. SWIMMING POOL - CONTINUOUS

Tess bursts out of the water in the middle of a crowded swimming pool. She crosses two lanes to reach the side. An eldery lady back crawls past Tess looking bewildered.

Tess heaves herself out of the water. The POOL BOY rushes to help.

POOL BOY
Omigod, I didn't see you fall in, I'm so sorry!

TESS
Don't worry about it.

He fishes one of her shoes out of the water and hands it over like a fine bottle of wine. She drops it on the side and shakes off her other shoe as well.

TESS
Not in trend this year anyway.
`
	scene, err := parseScene([]byte(standardScene), "./scenes/02_standard.fountain")
	assert.Equal(t, "INT. SWIMMING POOL - CONTINUOUS", scene.Title)
	assert.Equal(t, "02_standard.fountain", scene.File)
	assert.NoError(t, err)
}

func TestParseSceneWithoutTitle(t *testing.T) {
	sceneWithNoTitle := `
Tess bursts out of the water in the middle of a crowded swimming pool. She crosses two lanes to reach the side. An eldery lady back crawls past Tess looking bewildered.

Tess heaves herself out of the water. The POOL BOY rushes to help.

POOL BOY
Omigod, I didn't see you fall in, I'm so sorry!

TESS
Don't worry about it.

He fishes one of her shoes out of the water and hands it over like a fine bottle of wine. She drops it on the side and shakes off her other shoe as well.

TESS
Not in trend this year anyway.
`
	scene, err := parseScene([]byte(sceneWithNoTitle), "./scenes/03_without_title.fountain")
	assert.Equal(t, "", scene.Title)
	assert.EqualError(t, err, "could not parse scene './scenes/03_without_title.fountain': scene does not contain a title")
}

func TestParseSceneWithTitlePage(t *testing.T) {
	sceneWithNoTitle := `
	Title: THE TIME SURGEON
	Credit: By
	Author: Rick Schubert
	Contact:
	79 Ricardo Street
	E14 6EQ London
	+44 749 0947 880
	rickschubert@gmx.de
`
	scene, err := parseScene([]byte(sceneWithNoTitle), "./scenes/01_title_page.fountain")
	assert.Equal(t, "Title: THE TIME SURGEON", scene.Title)
	assert.Equal(t, "01_title_page.fountain", scene.File)
	assert.NoError(t, err)
}
