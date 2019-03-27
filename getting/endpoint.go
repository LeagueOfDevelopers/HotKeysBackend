package getting

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleGetPrograms(context *gin.Context, getter Service) {
	programs, err := getter.GetPrograms()
	if err != nil {
		sendError(context, err)
		panic("error while getting programs")
	}

	context.JSON(http.StatusOK, gin.H{
		"programs": programs,
	})
}

func HandleGetProgram(context *gin.Context, getter Service) {
	programName := context.Param("program")
	program, err := getter.GetProgram(programName)
	if err != nil {
		sendError(context, err)
		panic("error while getting program " + programName)
	}

	context.JSON(http.StatusOK, gin.H{
		"program": program,
	})
}

func sendError(context *gin.Context, err error) {
	context.JSON(http.StatusInternalServerError, gin.H{
		// TODO is it safe?
		"cause": err,
	})
}
