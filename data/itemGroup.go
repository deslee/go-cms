package data

type ItemGroup struct {
	ItemID  string `gorm:"type:text;primary_key;column:ItemId"`
	GroupID string `gorm:"type:text;primary_key;column:GroupId"`
	Order   int    `gorm:"type:integer;column:Order"`
	AuditFields
}

func (ItemGroup) TableName() string {
	return "ItemGroups"
}
