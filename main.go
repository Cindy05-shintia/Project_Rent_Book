package main

import (
	"fmt"

	config "Project_Rent_Book/Config"
	"Project_Rent_Book/entity"
)

func main() {
	conn := config.InitDB()
	SignUp := entity.SignUp{conn}

	var input int = 0
	for input != 9 {
		fmt.Println("--Menu--")
		fmt.Println("1. Sign Up")
		fmt.Println("2. Login")
		fmt.Println("3. Update Profile")
		fmt.Println("4. Delete Account")
		fmt.Println("10. Exit")
		fmt.Print("Pilih menu :")
		fmt.Scanln(&input)

		switch input {
		case 1:
			var UserBaru entity.User
			fmt.Print("Nama:")
			fmt.Scanln(&UserBaru.Nama)
			fmt.Print("Email:")
			fmt.Scanln(&UserBaru.Email)
			fmt.Print("Password:")
			fmt.Scanln(&UserBaru.Password)
			fmt.Print("Status")
			fmt.Scanln(&UserBaru.Status)
			res := SignUp.RegisterUser(UserBaru)
			if res.ID == 0 {
				fmt.Println("tidak bisa input User")
				break
			}
			fmt.Println("berhasil input user")

		case 2:
			fmt.Println("Daftar semua User")
			for _, v := range SignUp.LoginUser() {
				fmt.Println(v)
			}
		case 4:
			var ID uint
			fmt.Print("Masukkan ID yang akan dihapus: ")
			fmt.Scanln(&ID)
			SignUp.DeleteUser(ID)
		default:
			continue
		}
	}
	fmt.Println("Thank you for using our program, see next time")
}
