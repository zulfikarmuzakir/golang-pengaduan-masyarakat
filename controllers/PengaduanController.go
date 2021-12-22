package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/zulfikarmuzakir/golang-pengaduan-masyarakat/database"
	"github.com/zulfikarmuzakir/golang-pengaduan-masyarakat/models"
)

func IndexPengaduan(c *fiber.Ctx) error {
	var listPengaduan []models.Pengaduan

	statusQuery := c.Query("status")

	if statusQuery != "" {
		database.DB.Where("status = ?", statusQuery).Order("created_at DESC").Find(&listPengaduan)
	} else {
		database.DB.Order("created_at DESC").Find(&listPengaduan)
	}

	return c.JSON(listPengaduan)
}

func CreatePengaduan(c *fiber.Ctx) error {
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

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	currentUser := user

	newPengaduan := models.Pengaduan{
		Tgl_Pengaduan: data["tgl_pengaduan"],
		JudulLaporan:  data["judul_laporan"],
		Isi_Laporan:   data["isi_laporan"],
		Status:        "pending",
		UserId:        int(currentUser.Id),
	}

	database.DB.Create(&newPengaduan)

	return c.JSON(newPengaduan)
}

func UpdatePengaduan(c *fiber.Ctx) error {
	type updatePengaduan struct {
		Tgl_Pengaduan string `json:"tgl_pengaduan"`
		Isi_Laporan   string `json:"isi_laporan"`
		Foto          string `json:"foto"`
		Status        string `json:"status"`
	}
	id := c.Params("id")
	var pengaduan models.Pengaduan

	database.DB.First(&pengaduan, id)

	if pengaduan.Id == 0 {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "Data not found",
		})
	}

	var updateData updatePengaduan
	err := c.BodyParser(&updateData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	var updatedPengaduan models.Pengaduan
	updatedPengaduan.Tgl_Pengaduan = updateData.Tgl_Pengaduan
	updatedPengaduan.Isi_Laporan = updateData.Isi_Laporan
	updatedPengaduan.Status = updateData.Status

	//update data
	database.DB.Model(&pengaduan).Updates(updatedPengaduan)

	return c.JSON(fiber.Map{
		"status":   "success",
		"messsage": "data updated",
		"data":     pengaduan,
	})
}

func ShowPengaduan(c *fiber.Ctx) error {
	id := c.Params("id")
	var pengaduan models.Pengaduan
	var user models.User
	var tanggapan models.Tanggapan

	database.DB.First(&pengaduan, id)
	database.DB.Where("id = ?", pengaduan.UserId).First(&user)
	database.DB.Where("pengaduan_id = ?", pengaduan.Id).Order("created_at DESC").Find(&tanggapan)

	if pengaduan.Id == 0 {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "Data not found",
		})
	}

	return c.JSON(fiber.Map{
		"data_pengaduan": pengaduan,
		"user":           user,
		"tanggapan":      tanggapan,
	})
}

func DeletePengaduan(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	currentUser := user

	id := c.Params("id")
	var pengaduan models.Pengaduan

	database.DB.First(&pengaduan, id)
	if currentUser.Role == "petugas" {
		if pengaduan.Id == 0 {
			c.Status(fiber.StatusNotFound)
			return c.JSON(fiber.Map{
				"message": "data not found",
			})
		}

		database.DB.Delete(&pengaduan)

		return c.SendString("Successfully deleted")
	}

	return c.SendString("Kamu bukan petugas")

}
