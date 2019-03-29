package program

import "HotKeysBackend/hotkey"

var MockPrograms = []Program{
	{
		Name:     "Sketch",
		ImageURL: "image url",
		Hotkeys:  &hotkey.MockHotkeys,
	},
	{
		Name:     "Figma",
		ImageURL: "image url2",
		Hotkeys:  &hotkey.MockHotkeys2,
	},
}
