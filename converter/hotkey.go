package converter

import "HotKeysBackend/hotkey"

type HotkeyResponse struct {
	ID          uint
	Description string `json:"desc"`
}

func ConvertHotkeyToResponse(hotkey *hotkey.Hotkey) *HotkeyResponse {
	return &HotkeyResponse{
		ID:          hotkey.ID,
		Description: hotkey.Description,
	}
}
