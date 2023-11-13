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
		Delete(portalId uint) error
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

func (p *sysPortalUsecase) Delete(portalId uint) error {
	err := repository.DeletePortal(portalId)
	if err != nil {
		return err
	}
	return nil
}
