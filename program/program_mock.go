package program

import "HotKeysBackend/hotkey"

var MockPrograms = []*Program{
	{
		Id:       [16]byte{1},
		Name:     "Sketch",
		ImageURL: "image url",
		Hotkeys:  hotkey.MockHotkeys,
	},
	{
		Id:       [16]byte{2},
		Name:     "Figma",
		ImageURL: "image url2",
		Hotkeys:  hotkey.MockHotkeys2,
	},
}
