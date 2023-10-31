package request

type PortalListRequest struct {
	Search          *string `json:"search" binding:"required"`
	Limit           *int    `json:"limit" binding:"required"`
	Offset          *int    `json:"offset" binding:"required"`
	OrderColumnName *string `json:"orderColumnName" binding:"required"`
	Ordering        *string `json:"ordering" binding:"required"`
}
