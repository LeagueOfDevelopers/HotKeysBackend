package converter

import "HotKeysBackend/hotkey"

type HotkeyResponse struct {
	ID          uint
	Description string `json:"desc"`
}

func ConvertHotkeysToResponse(hotkeys []*hotkey.Hotkey) []*HotkeyResponse {
	hotkeysResponse := make([]*HotkeyResponse, len(hotkeys))
	for i, element := range hotkeys {
		hotkeysResponse[i] = ConvertHotkeyToResponse(element)
	}
	return hotkeysResponse
}

func ConvertHotkeyToResponse(hotkey *hotkey.Hotkey) *HotkeyResponse {
	return &HotkeyResponse{
		ID:          hotkey.ID,
		Description: hotkey.Description,
	}
}
