package repository

import (
	"go-gin-template/src/config"
	"go-gin-template/src/dto/request"
	"go-gin-template/src/entity"
)

func GetListPortal(request request.PortalListRequest) ([]entity.SysPortal, error) {
	conn := config.DBConn
	var sysPortal []entity.SysPortal

	// Define the order by clause
    if *request.OrderColumnName != "" {
		order := *request.OrderColumnName + " ASC"
		if *request.Ordering != "" {
			order = *request.OrderColumnName + " " + *request.Ordering
		}
        conn = conn.Order(order)
    }

	// Define the search query
    if *request.Search != "" {
        conn = conn.Where("portal_name LIKE ?", "%"+*request.Search+"%")
    }
	err := conn.Offset(*request.Offset).Limit(*request.Limit).Find(&sysPortal).Error
	if err != nil {
		return nil, err
	}
	return sysPortal, nil
}