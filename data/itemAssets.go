package data

type ItemAsset struct {
	ItemId string `db:"ItemId"`
	AssetId string `db:"AssetId"`
	Order  int    `db:"Order"`
	AuditFields
}

func (ItemAsset) TableName() string {
	return "ItemAssets"
}
