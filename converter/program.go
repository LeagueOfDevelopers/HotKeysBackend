package converter

import "HotKeysBackend/program"

type ProgramResponse struct {
	Name     string `json:"name"`
	ImageURL string `json:"url"`
}

func ConvertProgramsToResponse(programs []*program.Program) []*ProgramResponse {
	programsResponse := make([]*ProgramResponse, len(programs))
	for i, element := range programs {
		programsResponse[i] = ConvertProgramToResponse(element)
	}
	return programsResponse
}

func ConvertProgramToResponse(program *program.Program) *ProgramResponse {
	return &ProgramResponse{
		Name:     program.Name,
		ImageURL: program.ImageURL,
	}
}
