package api

import "github.com/gin-gonic/gin"

type EnordApi interface {
	RegisterDrone(*gin.Context)
	LoadDrone(*gin.Context)
	CheckDoneLoadMedication(*gin.Context)
	CheckDroneAvailability(*gin.Context)
	GetDroneBatteryLevel(*gin.Context)
}

type enord struct{}

func (e *enord) RegisterDrone(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (e *enord) LoadDrone(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (e *enord) CheckDoneLoadMedication(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (e *enord) CheckDroneAvailability(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (e *enord) GetDroneBatteryLevel(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewEnordAPI() EnordApi {
	return &enord{}
}
