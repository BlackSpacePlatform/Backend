package database

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"

	"github.com/LensPlatform/BlackSpace/pkg/models"
)

// CreateUser creates a user record in the backend database
func (db *Db) CreateUser(ctx context.Context, user *models.User) (*models.UserORM, error) {
	tx := func(tx *gorm.DB) (interface{}, error) {
		var userOrm models.UserORM

		// attempt to make sure all user fields are valid
		if err := user.Validate(); err != nil {
			db.Logger.Error("invalid user fields", zap.Error(err))
			return nil, err
		}
		// check if the user presently exists in the db based on user email or username
		recordNotFound := tx.Where(models.UserORM{Email: user.Email, Username: user.Username}).First(&userOrm).RecordNotFound()
		// if the user does exist and the user account is not active, reactivate the user account
		// and update the user state in the backend database
		if !recordNotFound && !userOrm.IsActive {
			userOrm.IsActive = true
			if err := tx.Save(&userOrm).Error; err != nil {
				db.Logger.Error("failed to re-activate user", zap.Error(err))
				return nil, err
			}
			return &userOrm, nil
		} else if !recordNotFound {
			db.Logger.Error("failed to create user because user already exists")
			return nil, errors.New("user already exists")
		}

		// convert the input user field to orm
		userOrm, err := user.ToORM(ctx)
		if err != nil {
			db.Logger.Error("failed to create user", zap.Error(err))
			return nil, err
		}

		// hash and salt the user password then update the password fields with the hashed version
		// hash user password
		userOrm.Password, err = db.ValidateAndHashPassword(userOrm)
		if err != nil {
			db.Logger.Error("unable to hash and salt user password", zap.Error(err))
			return nil, err
		}

		// save the user to the database
		if err := tx.Create(&userOrm).Error; err != nil {
			db.Logger.Error("failed to create user", zap.Error(err))
			return nil, err
		}

		db.Logger.Info("user successfully created",
			zap.String("id", string(userOrm.Id)),
			zap.String("username", userOrm.Username),
			zap.String("email", userOrm.Email))

		return &userOrm, nil
	}

	result, err := db.PerformComplexTransaction(tx)
	if err != nil {
		return nil, err
	}

	createdUser := result.(*models.UserORM)
	return createdUser, nil
}

// UpdateUser updates a user account in the database
func (db *Db) UpdateUser(ctx context.Context, user *models.User) (*models.UserORM, error) {
	transaction := func(tx *gorm.DB) (interface{}, error) {
		// first validate user object has all proper fields of interest
		if err := user.Validate(); err != nil {
			return nil, err
		}

		// first and foremost we check for the existence of the user
		exists, _, err := db.GetUserIfExists(ctx, user.Id, "", "")
		if !exists {
			db.Logger.Error("failed to obtain user by id as user does not exist", zap.Error(err))
			return nil, err
		}

		// convert the user to an ORM type
		userOrm, err := user.ToORM(ctx)
		if err != nil {
			db.Logger.Error("failed to convert user object to orm type", zap.Error(err))
			return nil, err
		}

		// update the actual user in the database
		if err := tx.Save(&userOrm).Error; err != nil {
			db.Logger.Error("failed to update user", zap.Error(err))
			return nil, err
		}

		db.Logger.Info("Successfully updated user", zap.String("id", string(userOrm.Id)),
			zap.String("user name", userOrm.Username))

		return &userOrm, nil
	}

	result, err := db.PerformComplexTransaction(transaction)
	if err != nil {
		return nil, err
	}

	updatedUser := result.(*models.UserORM)
	return updatedUser, nil
}

// DeleteUser deletes a user account from the backend database
func (db *Db) DeleteUser(ctx context.Context, userID uint32) error {
	tx := func(tx *gorm.DB) error {
		var userOrm models.UserORM
		if userID == 0 {
			db.Logger.Error("invalid user id")
			return errors.New("invalid user id")
		}

		exist, _, err := db.GetUserIfExists(ctx, userID, "", "")
		if !exist {
			db.Logger.Error("failed to obtain user by id. user does not exist", zap.Error(err))
			return err
		}

		if err = tx.Where(models.UserORM{Id: userID}).Delete(&userOrm).Error; err != nil {
			db.Logger.Error("failed to successfully delete user account", zap.Error(err))
			return err
		}

		db.Logger.Info("successfully deleted user account", zap.String("userId", string(userID)))
		return nil
	}
	return db.PerformTransaction(tx)
}

// GetUserIfExists checks that a given user exists based on user id
func (db *Db) GetUserIfExists(ctx context.Context, userID uint32, username, email string) (bool, *models.UserORM, error) {
	transaction := func(tx *gorm.DB) (interface{}, error) {
		// check if the user exists based on username, email and id
		var user models.UserORM

		if userID != 0 && username != "" && email != "" {
			if recordNotFound := tx.Where(&models.UserORM{Id: userID, Email: email, Username: username}).Find(&user).RecordNotFound(); recordNotFound {
				return nil, errors.New("user does not exits")
			}
		} else if userID != 0 && username == "" && email == "" {
			// user name and email is empty so obtain user by querying for user id
			if recordNotFound := tx.Where(&models.UserORM{Id: userID}).Find(&user).RecordNotFound(); recordNotFound {
				return nil, errors.New("user does not exits")
			}
		} else if userID == 0 && username != "" && email != "" {
			// user name and email is empty so obtain user by querying for user id
			if recordNotFound := tx.Where(&models.UserORM{Email: email, Username: username}).Find(&user).RecordNotFound(); recordNotFound {
				return nil, errors.New("user does not exits")
			}
		} else if userID == 0 && username != "" && email == "" {
			// user name and email is empty so obtain user by querying for user id
			if recordNotFound := tx.Where(&models.UserORM{Username: username}).Find(&user).RecordNotFound(); recordNotFound {
				return nil, errors.New("user does not exits")
			}
		} else if userID == 0 && username == "" && email != "" {
			// user name and email is empty so obtain user by querying for user id
			if recordNotFound := tx.Where(&models.UserORM{Email: email}).Find(&user).RecordNotFound(); recordNotFound {
				return nil, errors.New("user does not exits")
			}
		} else {
			return nil, errors.New("Invalid input parameters")
		}

		// Convert the userORM to user object and perform validation checks
		userObj, err := user.ToPB(ctx)
		if err != nil {
			return nil, err
		}

		// Actually perform user field validation
		if err := userObj.Validate(); err != nil {
			return nil, err
		}

		return &user, nil
	}

	output, err := db.PerformComplexTransaction(transaction)
	if err != nil {
		return false, nil, err
	}

	userAccount := output.(*models.UserORM)
	return true, userAccount, nil
}

// GetUser queries the database and obtains a user record by id
func (db *Db) GetUser(ctx context.Context, userID uint32) (*models.User, error) {
	var userOrm models.UserORM

	if userID == 0 {
		db.Logger.Error("invalid user id")
		return nil, errors.New("invalid user id")
	}

	if recordNotFound := db.Engine.Where(models.UserORM{Id: userID}).First(&userOrm).RecordNotFound(); recordNotFound {
		db.Logger.Error("user does not exist", zap.String("id", string(userID)))
		return nil, errors.New("user does not exist")
	}

	// convert the obtained user ORM object to a user object and validate all fields are there
	userObj, err := userOrm.ToPB(ctx)
	if err != nil {
		db.Logger.Error("failed to convert fields to protobuf format")
		return nil, err
	}

	// perform field validation
	if err = userObj.Validate(); err != nil {
		db.Logger.Error("field validation failed")
		return nil, err
	}

	db.Logger.Info("user successfully obtained user by id",
		zap.String("id", string(userOrm.Id)),
		zap.String("username", userOrm.Username),
		zap.String("email", userOrm.Email))

	return &userObj, nil
}
