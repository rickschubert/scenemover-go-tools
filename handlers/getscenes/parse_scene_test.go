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

func TestParseSceneWithMontageTitle(t *testing.T) {
	standardScene := `
MILITARY TRAINING MONTAGE

-- Self-defence class: Tess' trainer beckons her to punch him. She hits him in the stomach. He doesn't even flinch and instead throws her to the ground.
-- Firing range: Tess fires a handgun at a distant target. Only two out of nine shots hit at all.
-- Gym: Tess performs a set of jumping jacks and pances heavily.
-- Firing range: The shooting instructur demonstrates the firing pose. Tess imitates the pose. The instructor corrects her stance.
-- Gym: Tess struggles at the butterfly press machine.
-- Coffee shop: Tess stands in the queue for coffee. Her watch beeps with an alarm - she leaves the shop and starts doing jumping jacks right on the street. People stare at her.
-- Firing range: Tess presents the instructor with a target paper perforated nine times close to the center. The instructur looks incredulous and points to a rifle behind his shoulder.
-- Gym: Tess squats a massing rack of barbells.
-- Firing range: Tess fires a half-automatic rifle, ... a submachine gun, ... a sniper rifle, ... a fully automatic machine gun M4.
-- Self-defence class: Tess punches her trainer again in the stomach. He hunches over in pain as she moves her leg under his knee bend and throws him to the ground. She tries to elbow him in the face which he can just about so fence off before begging for a time out.

END MILITARY TRAINING MONTAGE
`
	scene, err := parseScene([]byte(standardScene), "./scenes/02_standard.fountain")
	assert.Equal(t, "MILITARY TRAINING MONTAGE", scene.Title)
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
