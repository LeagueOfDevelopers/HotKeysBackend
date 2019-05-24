package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var ErrorWrongBodyFormat = errors.New("wrong body format")

var ErrorGetProgram = errors.New("error while getting program")
var ErrorProgramNotFound = errors.New("program not found")
var ErrorWrongProgramIdFormat = errors.New("program id wrong format")

var ErrorGetHotkeys = errors.New("error while getting hotkeys")
var ErrorWrongHotkeyIdFormat = errors.New("hotkey id wrong format")

var ErrorKeyNotFound = errors.New("key not found")

var ErrorChecking = errors.New("error while checking")

func SendError(context *gin.Context, err error) {
	context.JSON(http.StatusInternalServerError, gin.H{
		// TODO is it safe?
		"error": err.Error(),
	})
}
