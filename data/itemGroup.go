package data

type ItemGroup struct {
	ItemId  string `db:"ItemId"`
	GroupId string `db:"GroupId"`
	Order   int    `db:"Order"`
	AuditFields
}

func (ItemGroup) TableName() string {
	return "ItemGroups"
}
