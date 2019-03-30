package checking

import (
	"HotKeysBackend/key"
	"HotKeysBackend/program"
	"HotKeysBackend/utils"
)

type Service interface {
	Check(projectId uint, hotkeyId uint, combination []*key.Key) (bool, error)
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

// TODO
func (s *service) Check(projectId uint, hotkeyId uint, combination []*key.Key) (bool, error) {
	currentHotkeys, err := s.programRepository.GetHotkeys(projectId)
	if err != nil {
		return false, utils.ErrorGetProgram
	}

	for _, element := range *currentHotkeys {
		isFound := true
		for i, currentCombination := range *element.Combination {
			if combination[i].ID != currentCombination.ID {
				isFound = false
			}
		}

		if isFound {
			return true, nil
		}
	}

	return false, nil
}
