package mcis

import tbcommon "github.com/cloud-barista/cm-butterfly/src/model/tumblebug/common"

type TbCustomImageInfo struct {
	AssociatedObjectList []string `json:"associatedObjectList"`
	ConnectionName       string   `json:"connectionName"`
	CreatedTime          string   `json:"createdTime"`
	CspCustomImageId     string   `json:"cspCustomImageId"`
	CspCustomImageName   string   `json:"cspCustomImageName"`
	Description          string   `json:"description"`
	GuestOS              string   `json:"guestOS"`

	ID              string                `json:"id"`
	isAutoGenerated bool                  `json:"isAutoGenerated"`
	KeyValueList    []tbcommon.TbKeyValue `json:"keyValueList"`
	Name            string                `json:"name"`
	NameSpace       string                `json:"namespace"`
	SourceVmId      string                `json:"sourceVmId"`

	Status      string `json:"status"`
	SystemLabel string `json:"systemLabel"`
}
