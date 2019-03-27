package hotkey

import "HotKeysBackend/key"

var MockHotkeys = []*Hotkey{
	{
		Id:          [16]byte{3},
		Description: "a",
		Combination: []*key.Key{
			{Code: 1, Name: "1"},
			{Code: 2, Name: "2"},
		},
	},
	{
		Id:          [16]byte{4},
		Description: "b",
		Combination: []*key.Key{
			{Code: 1, Name: "1"},
			{Code: 3, Name: "3"},
		},
	},
	{
		Id:          [16]byte{5},
		Description: "c",
		Combination: []*key.Key{
			{Code: 4, Name: "4"},
			{Code: 5, Name: "5"},
		},
	},
}

var MockHotkeys2 = []*Hotkey{
	{
		Id:          [16]byte{6},
		Description: "d",
		Combination: []*key.Key{
			{Code: 1, Name: "1"},
			{Code: 5, Name: "5"},
		},
	},
	{
		Id:          [16]byte{7},
		Description: "e",
		Combination: []*key.Key{
			{Code: 4, Name: "4"},
			{Code: 3, Name: "3"},
		},
	},
}
