package checking

import "HotKeysBackend/key"

type Service interface {
	Check(hotkeyId uint, combination []*key.Key)
}

type service struct {
}
