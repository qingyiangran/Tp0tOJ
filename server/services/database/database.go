package database

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"server/entity"
	"time"
)

var db *gorm.DB

func init() {
	prefix, _ := os.Getwd()
	dbPath := prefix + "/resources/data.db"
	test, err := os.Lstat(dbPath)
	if os.IsExist(err) {
		_, err := os.Create(dbPath)
		if err != nil {
			log.Panicln(err, test)
			return
		}
	} else if err != nil {
		if err != nil {
			log.Panicln(err, test)
			return
		}
	}
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Panicln("DB connect error", err.Error())
	}
	err = db.AutoMigrate(&entity.Bulletin{}, &entity.Challenge{}, &entity.Replica{}, &entity.ReplicaAlloc{}, &entity.ResetToken{}, &entity.Submit{}, &entity.User{})
	if err != nil {
		log.Panicln("DB connect error", err.Error())
		return
	}
}

func GetAllBulletin() ([]entity.Bulletin, error) {
	var allBulletin []entity.Bulletin
	result := db.Table("bulletins").Find(&allBulletin)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if result.Error != nil {
		return nil, result.Error
	}
	return allBulletin, nil
}
func AddBulletin(title string, content string, topping bool) error {
	newBulletin := entity.Bulletin{Title: title, Content: content, Topping: topping}
	result := db.Create(&newBulletin)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func FindBulletinByTitle(title string) ([]entity.Bulletin, error) {
	var bulletins []entity.Bulletin
	result := db.Table("bulletins").Where("title = ?", title).Find(&bulletins)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if result.Error != nil {
		return nil, result.Error
	}
	return bulletins, nil
}
func CheckMailExistence(mail string) (bool, error) {
	result := db.Table("users").Where("mail = ?", mail).First(&entity.User{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	} else if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func FindChallengeByState(state string) ([]entity.Challenge, error) {
	var challenges []entity.Challenge
	result := db.Table("challenges").Where("state = ?", state).Find(&challenges)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if result.Error != nil {
		return nil, result.Error
	}
	return challenges, nil
}
func FindChallengeById(id uint64) (*entity.Challenge, error) {
	var challenge entity.Challenge
	result := db.Table("challenges").Where("id = ?", id).Find(&challenge)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if result.Error != nil {
		return nil, result.Error
	}
	return &challenge, nil
}

// AddUser support role[admin|member|team] state[banned|disabled|normal]
func AddUser(name string, password string, mail string, role string, state string) error {
	newUser := entity.User{Name: name, Password: password, Mail: mail, Role: role, State: state, JoinTime: time.Now(), Score: 0}
	result := db.Create(&newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindReplicaAllocByUserId(userId uint64) ([]entity.ReplicaAlloc, error) {
	var replicaAllocs []entity.ReplicaAlloc
	result := db.Table("replica_allocs").Where("user_id").Find(&replicaAllocs)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if result.Error != nil {
		return nil, result.Error
	}
	return replicaAllocs, nil
}

func FindReplicaByChallengeId(challengeId uint64) ([]entity.Replica, error) {
	var replicas []entity.Replica
	result := db.Table("replicas").Where("challenge_id").Find(&replicas)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if result.Error != nil {
		return nil, result.Error
	}
	return replicas, nil
}

func FindResetTokenByUserId(userId uint64) (*entity.ResetToken, error) {
	var resetToken entity.ResetToken
	result := db.Table("reset_tokens").Where("user_id").Find(&resetToken)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if result.Error != nil {
		return nil, result.Error
	}
	return &resetToken, nil
}
