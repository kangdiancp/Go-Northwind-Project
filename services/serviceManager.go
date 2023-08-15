package services

import "codeid.northwind/repositories"

type ServiceManager struct {
	CategoryService
	//ProductService
	UsersService
}

// constructor
func NewServiceManager(repoMgr *repositories.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		CategoryService: *NewCategoryService(repoMgr),
		UsersService:    *NewUsersService(repoMgr),
		//ProductService: *NewProductService(&repoMgr.ProductCategory),
	}
}
