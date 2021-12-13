package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zulfikarmuzakir/golang-pengaduan-masyarakat/database"
	"github.com/zulfikarmuzakir/golang-pengaduan-masyarakat/models"
)

func UpdateUser(c *fiber.Ctx) error {
	type UpdateUser struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Telp     string `json:"telp"`
		Nik      string `json:"nik"`
	}
	id := c.Params("id")
	var user models.User

	database.DB.First(&user, id)

	if user.Id == 0 {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	var updateData UpdateUser
	err := c.BodyParser(&updateData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	var updatedUser models.User
	updatedUser.Name = updateData.Name
	updatedUser.Email = updateData.Email
	updatedUser.Username = updateData.Username
	updatedUser.Telp = updateData.Telp
	updatedUser.Nik = updateData.Nik

	//update data
	database.DB.Model(&user).Updates(updatedUser)

	return c.JSON(fiber.Map{
		"status":   "success",
		"messsage": "data updated",
		"data":     user,
	})
}
