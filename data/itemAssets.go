package data

type ItemAssets struct {
	ItemID string `gorm:"type:text;primary_key;column=ItemId"`
	SiteID string `gorm:"type:text;primary_key;column=SiteId"`
	Order  int    `gorm:"type:integer;column=Order"`
	AuditFields
}
