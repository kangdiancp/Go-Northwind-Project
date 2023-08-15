package controllers

import "codeid.northwind/services"

type ControllerManager struct {
	CategoryController
	UsersController
}

// constructor
func NewControllerManager(serviceMgr *services.ServiceManager) *ControllerManager {
	return &ControllerManager{
		*NewCategoryController(serviceMgr),
		*NewUsersController(&serviceMgr.UsersService),
	}
}
