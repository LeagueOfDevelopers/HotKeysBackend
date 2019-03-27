package storage

import (
	"HotKeysBackend/key"
	"HotKeysBackend/program"
)

type InMemoryStorage struct {
	programRepository programInMemoryRepository
	keyRepository     keyInMemoryRepository
}

type programInMemoryRepository struct {
	Programs []*program.Program
}

type keyInMemoryRepository struct {
	Keys []*key.Key
}

var programRepository = &programInMemoryRepository{
	Programs: program.MockPrograms,
}

var keyRepository = &keyInMemoryRepository{
	Keys: key.MockKeys,
}

func GetProgramInMemoryRepository() program.Repository {
	return programRepository
}

func GetKeyInMemoryRepository() key.Repository {
	return keyRepository
}

func (repository *programInMemoryRepository) GetAll() ([]*program.Program, error) {
	return repository.Programs, nil
}

func (repository *programInMemoryRepository) Get(name string) (*program.Program, error) {
	for _, element := range repository.Programs {
		if element.Name == name {
			return element, nil
		}
	}
	return nil, program.ErrorProgramNotFound
}

func (repository *keyInMemoryRepository) GetAll() ([]*key.Key, error) {
	return repository.Keys, nil
}

func (repository *keyInMemoryRepository) Get(code int) (*key.Key, error) {
	for _, element := range repository.Keys {
		if element.Code == code {
			return element, nil
		}
	}
	return nil, key.ErrorKeyNotFound
}
