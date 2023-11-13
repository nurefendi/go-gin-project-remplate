package controller

import (
	"go-gin-template/src/constant"
	"go-gin-template/src/dto/request"
	"go-gin-template/src/helper"
	"go-gin-template/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

var sysPortalService = service.SysPortalService()

func LisPortal(c *gin.Context) {
	constant.RequestId = uuid.New().String()
	limit := helper.ParseStringToInt(c.Query("limit"))
    offset := helper.ParseStringToInt(c.Query("offset"))
    search := c.Query("search")
    orderColumnName := c.Query("orderColumnName")
    ordering := c.Query("ordering")

	if limit == 0 {
		limit = 10
	}

	input := request.PortalListRequest{
		Limit: &limit,
		Offset: &offset,
		Search: &search,
		OrderColumnName: &orderColumnName,
		Ordering: &ordering,
	}

	result, err := sysPortalService.GetListPortal(input)
	if err != nil {
		log.Error(constant.RequestId, " error : ", err.Error())
		c.JSON(http.StatusUnprocessableEntity, helper.ResponseWithJson(constant.Failed, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseWithJson(constant.Success, "Get Data Success", gin.H{
		"data": result,
		"limit": limit,
		"offset": offset,
	}))
}

func DeletePortal(c *gin.Context) {
	constant.RequestId = uuid.New().String()
	id := c.Param("id")
	err := sysPortalService.Delete(uint(helper.ParseStringToInt(id)))
	if err != nil {
		log.Error(constant.RequestId, " error : ", err.Error())
		c.JSON(http.StatusUnprocessableEntity, helper.ResponseWithJson(constant.Failed, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, helper.ResponseWithJson(constant.Success, "Delete Data Success", nil))
}

func StorePortal(c *gin.Context) {
	constant.RequestId = uuid.New().String()
}

func UpdatePortal(c *gin.Context) {
	constant.RequestId = uuid.New().String()
	// id := c.Param("id")
	var payload request.PortalRequest
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		log.Error(constant.RequestId, " error : ", err.Error())
		c.JSON(http.StatusUnprocessableEntity, helper.ResponseValidationError(constant.Failed, err))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseWithJson(constant.Success, "Get Data Success", nil))
}