package service

import (
	"go-gin-template/src/dto/request"
	"go-gin-template/src/entity"
	"go-gin-template/src/repository"
)

type (
	sysPortalUsecase struct{}

	SysPortal interface {
		GetListPortal(data request.PortalListRequest) ([]entity.SysPortal, error)
	}
)

func SysPortalService() SysPortal {
	return &sysPortalUsecase{}
}

func (u *sysPortalUsecase) GetListPortal(data request.PortalListRequest) ([]entity.SysPortal, error) {
	res, err := repository.GetListPortal(data)
	if err != nil {
		return nil, err
	}
	return res, nil
}
