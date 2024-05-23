package mock

import (
	"gorm.io/gorm"
	"xcylla.io/common/uid"
	"xcylla.io/config/internal/users/entities"
)

func FillSystem(systemDb *gorm.DB) {

}

func FillUserData(userDb *gorm.DB) {
	roleAdmin := entities.Def_Roles{ID: uid.NewId(), Write: true, Read: true, Delete: true}
	userDb.Create(&roleAdmin)

	userManu := entities.Def_Users{ID: uid.NewId(), Name: "Manu", Subname: "Rubio", Username: "mrubio", Email: "manuel.rubio@xmartcs.com", HashedPassword: "8d3d16633b1aea4d5e628b0efc84516048701671c824538765ee3bd8bdd9f1ad", Role_Id: roleAdmin.ID}
	userDb.Create(&userManu)
}
