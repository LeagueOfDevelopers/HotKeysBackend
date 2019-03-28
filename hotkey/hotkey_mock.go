package hotkey

import "HotKeysBackend/key"

var MockHotkeys = &[]Hotkey{
	{
		Description: "a",
		Combination: &[]key.Key{
			{Code: 1, Name: "1"},
			{Code: 2, Name: "2"},
		},
	},
	{
		Description: "b",
		Combination: &[]key.Key{
			{Code: 1, Name: "1"},
			{Code: 3, Name: "3"},
		},
	},
	{
		Description: "c",
		Combination: &[]key.Key{
			{Code: 4, Name: "4"},
			{Code: 5, Name: "5"},
		},
	},
}

var MockHotkeys2 = &[]Hotkey{
	{
		Description: "d",
		Combination: &[]key.Key{
			{Code: 1, Name: "1"},
			{Code: 5, Name: "5"},
		},
	},
	{
		Description: "e",
		Combination: &[]key.Key{
			{Code: 4, Name: "4"},
			{Code: 3, Name: "3"},
		},
	},
}
