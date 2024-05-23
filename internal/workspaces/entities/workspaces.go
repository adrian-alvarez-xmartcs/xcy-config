package entities

type (
	Def_Workspace struct {
		ID          uint32 `gorm:"column:id; primaryKey; autoIncrement"`
		Name        string `gorm:"column:name"`
		Description string `gorm:"column:description"`
	}
)
