package entities

type (
	Def_Users struct {
		ID             string `gorm:"column:id; primaryKey"`
		Username       string `gorm:"column:username; unique"`
		Name           string `gorm:"column:name"`
		Subname        string `gorm:"column:subname"`
		Email          string `gorm:"column:email"`
		HashedPassword string `gorm:"column:hashed_password"`

		Role_Id string    `gorm:"column:role_id"`
		Role    Def_Roles `gorm:"foreignKey:role_id; reference:id"`
	}
)
