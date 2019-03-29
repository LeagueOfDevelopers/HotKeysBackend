package getting

import (
	"HotKeysBackend/hotkey"
	"HotKeysBackend/key"
	"HotKeysBackend/program"
)

type Service interface {
	GetPrograms() (*[]program.Program, error)
	GetProgram(id uint) (*program.Program, error)
	GetHotkeysForProgram(id uint) (*[]hotkey.Hotkey, error)
}

type service struct {
	programRepository program.Repository
	keyRepository     key.Repository
}

func CreateService(programRepository program.Repository, keyRepository key.Repository) Service {
	return &service{
		programRepository: programRepository,
		keyRepository:     keyRepository,
	}
}

func (s *service) GetPrograms() (*[]program.Program, error) {
	return s.programRepository.GetAll()
}

func (s *service) GetProgram(id uint) (*program.Program, error) {
	return s.programRepository.Get(id)
}

func (s *service) GetHotkeysForProgram(id uint) (*[]hotkey.Hotkey, error) {
	return s.programRepository.GetHotkeys(id)
}
