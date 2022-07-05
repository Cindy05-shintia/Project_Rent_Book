package entity

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint
	Nama       string
	Email      string
	Password   string
	Status     bool
	created_at time.Time
}

type SignUp struct {
	DB *gorm.DB
}

func (su *SignUp) RegisterUser(UserBaru User) User {
	err := su.DB.Create(UserBaru).Error
	if err != nil {
		log.Print(err)
		return User{}
	}
	return UserBaru
}

func (su *SignUp) LoginUser() []User {
	var ListUser = []User{}

	if err := su.DB.Find(&ListUser).Error; err != nil {
		log.Print(err)
		return nil
	}
	return ListUser
}

// func (su *SignUp) UpdateUser() user {
// 	var Update user
// 	err := su.DB.Model(&Update).Updates(user{Nama: "", Email: "", Password: "", Status: false})
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil
// 	}
// 	return Update
// }

func (su *SignUp) DeleteUser(ID_User uint) bool {
	//var HapusUser user
	postExc := su.DB.Where("ID = ?", ID_User).Delete(&User{})
	if err := postExc.Error; err != nil {
		log.Print(err)
		return false
	}
	if aff := postExc.RowsAffected; aff < 1 {
		log.Println("Tidak ada data yang dihapus")
		return false
	}
	return true
}
