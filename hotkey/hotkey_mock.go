package hotkey

import "HotKeysBackend/key"

var MockHotkeys = []Hotkey{
	{
		Description: "a",
		Combination: &key.MockKeys,
	},
	{
		Description: "b",
		Combination: &key.MockKeys,
	},
	{
		Description: "c",
		Combination: &key.MockKeys,
	},
}

var MockHotkeys2 = []Hotkey{
	{
		Description: "d",
		Combination: &key.MockKeys,
	},
	{
		Description: "e",
		Combination: &key.MockKeys,
	},
}
