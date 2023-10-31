package response

import (
	"time"
)

type sysPortalListResponse struct {
	PortalID     uint      `json:"portalId"`
	PortalNumber string    `json:"portalNumber"`
	PortalName   string    `json:"portalName"`
	PortalDesc   string    `json:"portalDesc"`
	PortalLink   *string   `json:"portalLink"`
	MetaTitle    string    `json:"metaTitle"`
	MetaDesc     string    `json:"metaDesc"`
	MetaTag      string    `json:"metaTag"`
	ModifiedBy   int       `json:"modifiedBy"`
	ModifiedDate time.Time `json:"modifiedDate"`
}
// func PortalListResponse(en entity.SysPortal) *sysPortalListResponse {
// 	return &sysPortalListResponse{
// 		PortalID: en.PortalID,
// 		PortalNumber: en.PortalNumber,
// 		PortalName: en.PortalName,
// 		PortalDesc: en.PortalDesc,
// 		PortalLink: en.PortalLink,
// 		MetaTitle: en.MetaTitle,
// 		MetaDesc: en.MetaDesc,
// 		MetaTag: en.MetaTag,
// 		ModifiedBy: en.ModifiedBy,
// 		ModifiedDate: en.ModifiedDate,
// 	}
// }	