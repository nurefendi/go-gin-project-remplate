package request

type PortalListRequest struct {
	Search          *string `json:"search" binding:"required"`
	Limit           *int    `json:"limit" binding:"required"`
	Offset          *int    `json:"offset" binding:"required"`
	OrderColumnName *string `json:"orderColumnName" binding:"required"`
	Ordering        *string `json:"ordering" binding:"required"`
}

type PortalRequest struct {
	PortalID     *uint      `json:"portalId"`
	PortalNumber string    `json:"portalNumber" binding:"required,max=2"`
	PortalName   string    `json:"portalName" binding:"required,max=45"`
	PortalDesc   string    `json:"portalDesc" binding:"required,max=255"`
	PortalLink   *string   `json:"portalLink" binding:"max=100"`
	MetaTitle    string    `json:"metaTitle" binding:"required,max=150"`
	MetaDesc     string    `json:"metaDesc" binding:"required,max=150"`
	MetaTag      string    `json:"metaTag" binding:"required,max=150"`
}