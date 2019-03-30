package converter

import "HotKeysBackend/program"

type ProgramResponse struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"url"`
}

func ConvertProgramsToResponse(programs *[]program.Program) *[]ProgramResponse {
	programsResponse := make([]ProgramResponse, len(*programs))
	for i, element := range *programs {
		programsResponse[i] = *ConvertProgramToResponse(&element)
	}
	return &programsResponse
}

func ConvertProgramToResponse(program *program.Program) *ProgramResponse {
	return &ProgramResponse{
		Id:       program.ID,
		Name:     program.Name,
		ImageURL: program.ImageURL,
	}
}
