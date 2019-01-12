package data

type ItemAsset struct {
	ItemID string `gorm:"type:text;primary_key;column:ItemId"`
	SiteID string `gorm:"type:text;primary_key;column:SiteId"`
	Order  int    `gorm:"type:integer;column:Order"`
	AuditFields
}

func (ItemAsset) TableName() string {
	return "ItemAsset"
}
