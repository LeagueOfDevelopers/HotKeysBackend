package checking

import (
	"HotKeysBackend/key"
	"HotKeysBackend/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CheckHotkeyRequest struct {
	ProgramId   uint       `json:"programId"`
	HotkeyId    uint       `json:"hotkeyId"`
	Combination []*key.Key `json:"combination"`
}

func HandleCheckHotkey(context *gin.Context, checker Service) {
	paramProgramId, err := strconv.ParseUint(context.Param(utils.ProgramId), 10, 64)
	if err != nil {
		utils.SendError(context, err)
		panic(utils.ErrorWrongProgramIdFormat)
	}
	programId := uint(paramProgramId)

	paramHotkeyId, err := strconv.ParseUint(context.Param(utils.HotkeyId), 10, 64)
	if err != nil {
		utils.SendError(context, err)
		panic(utils.ErrorWrongHotkeyIdFormat)
	}
	hotkeyId := uint(paramHotkeyId)

	checkHotkeyRequest := &CheckHotkeyRequest{
		ProgramId: programId,
		HotkeyId:  hotkeyId,
		//Combination:
	}

	err = context.BindJSON(checkHotkeyRequest)
	if err != nil {
		utils.SendError(context, err)
		panic(utils.ErrorWrongBodyFormat)
	}

}
