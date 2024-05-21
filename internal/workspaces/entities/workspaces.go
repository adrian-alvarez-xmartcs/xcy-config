package entities

type (
	Workspace struct {
		ID          uint32 `gorm:"primaryKey; autoIncrement"`
		Name        string `gorm:"column:name"`
		Description string `gorm:"column:description"`
	}
)
