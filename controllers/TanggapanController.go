package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/zulfikarmuzakir/golang-pengaduan-masyarakat/database"
	"github.com/zulfikarmuzakir/golang-pengaduan-masyarakat/models"
)

func ListTanggapan(c *fiber.Ctx) error {
	var listTanggapan []models.Tanggapan
	pengaduan_id := c.Params("id")

	database.DB.Where("pengaduan_id = ?", pengaduan_id).Find(&listTanggapan)

	return c.JSON(listTanggapan)
}

func CreateTanggapan(c *fiber.Ctx) error {
	var data map[string]string

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	id := c.Params("id")
	var pengaduan models.Pengaduan

	database.DB.First(&pengaduan, id)

	pengaduan_data := pengaduan

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	currentUser := user

	newTanggapan := models.Tanggapan{
		PengaduanId:   int(pengaduan_data.Id),
		Tgl_Tanggapan: data["tgl_tanggapan"],
		Tanggapan:     data["tanggapan"],
		UserId:        int(currentUser.Id),
	}

	database.DB.Create(&newTanggapan)

	return c.JSON(newTanggapan)
}

func DeleteTanggapan() {

}
