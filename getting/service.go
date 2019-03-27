package getting

import (
	"HotKeysBackend/hotkey"
	"HotKeysBackend/key"
	"HotKeysBackend/program"
)

type Service interface {
	GetPrograms() ([]*program.Program, error)
	GetProgram(name string) (*program.Program, error)
	GetHotkeysForProgram(name string) ([]*hotkey.Hotkey, error)
}

type service struct {
	programRepository program.Repository
	keyRepository     key.Repository
}

func NewService(programRepository program.Repository, keyRepository key.Repository) Service {
	return &service{
		programRepository: programRepository,
		keyRepository:     keyRepository,
	}
}

func (s *service) GetPrograms() ([]*program.Program, error) {
	return s.programRepository.GetAll()
}

func (s *service) GetProgram(name string) (*program.Program, error) {
	return s.programRepository.Get(name)
}

func (s *service) GetHotkeysForProgram(name string) ([]*hotkey.Hotkey, error) {
	return s.programRepository.GetHotkeys(name)
}
