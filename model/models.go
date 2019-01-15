package model

// WARNING: the struct fields must be in the same order as the columns in the database. this is just how the generated repository works.

type JSONObject = string

type Site struct {
	Id   string     `dbGen:"column:Id;Pk:1" db:"Id"`
	Name string     `dbGen:"column:Name" db:"Name"`
	Data JSONObject `dbGen:"column:Data" db:"Data"`
	AuditFields
}

func (Site) TableName() string {
	return "Sites"
}

type Asset struct {
	Id     string `dbGen:"column:Id;Pk:1" db:"Id"`
	SiteId string `dbGen:"column:SiteId" db:"SiteId"`
	State  string `dbGen:"column:State" db:"State"`
	Type   string `dbGen:"column:Type" db:"Type"`
	Data   string `dbGen:"column:Data" db:"Data"`
	AuditFields
}

func (Asset) TableName() string {
	return "Assets"
}

func (asset Asset) FileName() string {
	return "not implemented"
}

func (asset Asset) Extension() string {
	return "not implemented"
}

type AuditFields struct {
	CreatedAt     string `dbGen:"column:CreatedAt" db:"CreatedAt"`
	CreatedBy     string `dbGen:"column:CreatedBy" db:"CreatedBy"`
	LastUpdatedAt string `dbGen:"column:LastUpdatedAt" db:"LastUpdatedAt"`
	LastUpdatedBy string `dbGen:"column:LastUpdatedBy" db:"LastUpdatedBy"`
}

type Group struct {
	Id     string     `dbGen:"column:Id;Pk:1" db:"Id"`
	SiteId string     `dbGen:"column:SiteId" db:"SiteId"`
	Data   JSONObject `dbGen:"column:Data" db:"Data"`
	Name   string     `dbGen:"column:Name" db:"Name"`
	AuditFields
}

func (Group) TableName() string {
	return "Groups"
}

type Item struct {
	Id     string     `dbGen:"column:Id;Pk:1" db:"Id"`
	SiteId string     `dbGen:"column:SiteId" db:"SiteId"`
	Data   JSONObject `dbGen:"column:Data" db:"Data"`
	Type   string     `dbGen:"column:Type" db:"Type"`
	AuditFields
}

func (Item) TableName() string {
	return "Items"
}

type ItemAsset struct {
	ItemId  string `dbGen:"column:ItemId;Pk:1" db:"ItemId"`
	AssetId string `dbGen:"column:AssetId;Pk:2" db:"AssetId"`
	Order   int    `dbGen:"column:Order" db:"Order"`
	AuditFields
}

func (ItemAsset) TableName() string {
	return "ItemAssets"
}

type ItemGroup struct {
	ItemId  string `dbGen:"column:ItemId;Pk:1" db:"ItemId"`
	GroupId string `dbGen:"column:GroupId;Pk:2" db:"GroupId"`
	Order   int    `dbGen:"column:Order" db:"Order"`
	AuditFields
}

func (ItemGroup) TableName() string {
	return "ItemGroups"
}

type SiteUser struct {
	UserId string `dbGen:"column:UserId;Pk:1" db:"UserId"`
	SiteId string `dbGen:"column:SiteId;Pk:2" db:"SiteId"`
	Order  int    `dbGen:"column:Order" db:"Order"`
	AuditFields
}

func (SiteUser) TableName() string {
	return "SiteUsers"
}

type User struct {
	Id       string     `dbGen:"column:Id;Pk:1" db:"Id"`
	Email    string     `dbGen:"column:Email;needsGetter:true" db:"Email"`
	Password string     `dbGen:"column:Password" db:"Password"`
	Salt     string     `dbGen:"column:Salt" db:"Salt"`
	Data     JSONObject `dbGen:"column:Data" db:"Data"`
	AuditFields
}

func (User) TableName() string {
	return "Users"
}
