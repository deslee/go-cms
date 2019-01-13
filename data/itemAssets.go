package data

type ItemAsset struct {
	ItemId string `db:"ItemId"`
	SiteId string `db:"SiteId"`
	Order  int    `fb:"Order"`
	AuditFields
}

func (ItemAsset) TableName() string {
	return "ItemAsset"
}
