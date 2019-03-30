package storage

import (
	"HotKeysBackend/hotkey"
	"HotKeysBackend/key"
	"HotKeysBackend/program"
)

type programInMemoryRepository struct {
	Programs *[]program.Program
}

type keyInMemoryRepository struct {
	Keys *[]key.Key
}

var programRepository = &programInMemoryRepository{
	Programs: &program.MockPrograms,
}

var keyRepository = &keyInMemoryRepository{
	Keys: &key.MockKeys,
}

func GetProgramInMemoryRepository() program.Repository {
	return programRepository
}

func GetKeyInMemoryRepository() key.Repository {
	return keyRepository
}

func (repository *programInMemoryRepository) GetAll() (*[]program.Program, error) {
	return repository.Programs, nil
}

func (repository *programInMemoryRepository) Get(id uint) (*program.Program, error) {
	for _, element := range *repository.Programs {
		if element.ID == id {
			return &element, nil
		}
	}

	return nil, program.ErrorProgramNotFound
}

func (repository *programInMemoryRepository) GetHotkeys(id uint) (*[]hotkey.Hotkey, error) {
	currentProgram, err := repository.Get(id)
	if err != nil {
		return nil, program.ErrorGetProgram
	}

	return currentProgram.Hotkeys, nil
}

func (repository *keyInMemoryRepository) GetAll() (*[]key.Key, error) {
	return repository.Keys, nil
}

func (repository *keyInMemoryRepository) Get(code int) (*key.Key, error) {
	for _, element := range *repository.Keys {
		if element.Code == code {
			return &element, nil
		}
	}

	return nil, key.ErrorKeyNotFound
}
