package main

import (
	"errors"
	"fmt"
	"sesi7-belajar-gorm/database"
	"sesi7-belajar-gorm/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()

	// createUser("ahanifpratama@gmail.com")
	// getUserById(1)
	// updateUserById(1, "ahpratama21@gmail.com")
	// createProduct(1, "YLO", "YYY")
	// getUserWithProduct()
	// deleteProductById(1)
}


func createUser(email string) {
	db := database.GetDB()

	User := models.User {
		Email: email,
	}

	err := db.Create(&User).Error
	
	if err != nil {
		fmt.Println("Error Creating User Data:", err)
		return
	}

	fmt.Println("New User Data:", User)
}

func getUserById(id uint) {
	db := database.GetDB()

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User Data Not Found")
			return
		}
		print("Error finding user:", err)
	}

	fmt.Printf("User Data: %+v \n", user)
}

func updateUserById(id uint, email string) {
	db := database.GetDB()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error

	if err != nil {
		fmt.Println("Error updating user data:", err)
		return
	}

	fmt.Printf("Update user's email: %+v \n", user.Email)
}

func createProduct(userId uint, brand string, name string) {
	db := database.GetDB()

	Product := models.Product {
		UserID: userId,
		Brand: brand,
		Name: name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error Creating Product Data:", err.Error())
		return
	}

	fmt.Println("New Product Data:", Product)
}

func getUserWithProduct() {
	db := database.GetDB()

	users := models.User{}

	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user datas with products:", err.Error())
		return
	}

	fmt.Println("User Datas with Products")
	fmt.Printf("%+v", users)
}

func deleteProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}

	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error Deleting Product:", err.Error())
		return
	}

	fmt.Printf("Product with id %d has been successfully deleted", id)
}