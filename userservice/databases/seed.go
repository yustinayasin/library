package users

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

// Seeding only for 2 roles 1. admin, 2. user
func SeedRoles(DB *gorm.DB) {
	var admin Role
	var user Role

	if err := DB.First(&admin, "id = ?", 1).Error; err != nil {
		admin = Role{Id: 1, Name: "admin"}
		if err := DB.Create(&admin).Error; err != nil {
			log.Fatalf("Failed to seed admin role: %v", err)
		}
		fmt.Println("Admin role seeded")
	}

	if err := DB.First(&user, "id = ?", 2).Error; err != nil {
		user = Role{Id: 2, Name: "user"}
		if err := DB.Create(&user).Error; err != nil {
			log.Fatalf("Failed to seed user role: %v", err)
		}
		fmt.Println("User role seeded")
	}
}
