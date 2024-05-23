package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"xcylla.io/common/log"
	"xcylla.io/config/internal/users"
	"xcylla.io/config/internal/users/entities"
	"xcylla.io/config/pkg/database"
)

type repository struct {
	db  *gorm.DB
	log log.Logger
}

func NewUsersRepository(db *database.Database) users.Repository {
	return &repository{
		db:  database.UserDatabase,
		log: log.NewLogger("UsersRepository"),
	}
}

func (r *repository) SetDatabase(database *gorm.DB) {
	if database == nil {
		r.log.Warning("Attempting to establish a null database")
	}
	r.db = database
}

func (r *repository) Atomic(ctx context.Context, repo func(tx users.Repository) error) error {
	txConn := r.db.Begin()
	if txConn.Error != nil {
		return txConn.Error
	}

	newRepository := &repository{db: txConn}

	err := repo(newRepository)
	if err != nil {
		return err
	}

	if newRepository.db.Error != nil {
		return newRepository.db.Error
	}

	return nil
}

func (r *repository) Login(ctx context.Context, user, hashedPw string) (*entities.Def_Users, error) {
	r.log.Trace("Attempting to log in user from repository")

	var userRecord entities.Def_Users

	tx := r.db.Joins("Role").Where("username = ?", user).First(&userRecord)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			r.log.Error("User not found")
			return nil, errors.New("user not found")
		}
		r.log.Error("Error fetching user record")
		return nil, tx.Error
	}

	if userRecord.HashedPassword != hashedPw {
		r.log.Error("Incorrect password")
		return nil, errors.New("incorrect password")
	}

	r.log.Trace("User successfully authenticated")
	return &userRecord, nil
}
