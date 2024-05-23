package entities

type (
	Def_Roles struct {
		ID     string `gorm:"column:id; primaryKey"`
		Write  bool   `gorm:"column:write"`
		Read   bool   `gorm:"column:read"`
		Delete bool   `gorm:"column:delete"`
	}
)
