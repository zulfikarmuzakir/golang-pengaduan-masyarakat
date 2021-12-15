package controllers

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/zulfikarmuzakir/golang-pengaduan-masyarakat/database"
	"github.com/zulfikarmuzakir/golang-pengaduan-masyarakat/models"
)

func IndexPengaduan(c *fiber.Ctx) error {
	var listPengaduan []models.Pengaduan
	statusQuery := c.Query("status")
	if statusQuery == "pending" {
		database.DB.Where("status = ?", statusQuery).Order("created_at DESC").Find(&listPengaduan)
	} else if statusQuery == "accepted" {
		database.DB.Where("status = ?", statusQuery).Order("created_at DESC").Find(&listPengaduan)
	} else {
		database.DB.Order("created_at DESC").Find(&listPengaduan)
	}

	return c.JSON(listPengaduan)
}

func CreatePengaduan(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	// if err := c.BodyParser(&data); err != nil {
	// return err
	// }

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

	tglPengaduan := c.FormValue("tgl_pengaduan")
	judulLaporan := c.FormValue("judul_laporan")
	isiLaporan := c.FormValue("isi_laporan")

	var data = map[string]string{"tgl_pengaduan": tglPengaduan, "judul_laporan": judulLaporan, "isi_laporan": isiLaporan}

	// Image uploads
	form, err := c.MultipartForm()
	files := form.File["image"]

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	var imageURL = []string{}

	for _, file := range files {
		uniqueId := uuid.New()
		filename := strings.Replace(uniqueId.String(), "-", "", -1)
		fileExt := filepath.Ext(file.Filename)
		image := fmt.Sprintf("%s%s", filename, fileExt)
		err = c.SaveFile(file, fmt.Sprintf("./images/%s", image))
		if err != nil {
			log.Println("image save error --> ", err)
			return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
		}
		imageLoc := fmt.Sprintf("http://localhost:8050/images/%s", image)

		imageURL = append(imageURL, imageLoc)
	}

	newPengaduan := models.Pengaduan{
		Tgl_Pengaduan: data["tgl_pengaduan"],
		JudulLaporan:  data["judul_laporan"],
		Isi_Laporan:   data["isi_laporan"],
		// Foto:          imageURL[],
		Status: "pending",
		UserId: int(currentUser.Id),
	}

	database.DB.Create(&newPengaduan)

	for _, image := range imageURL {
		imageUpload := models.Image{
			PengaduanID: int(newPengaduan.Id),
			Image:       image,
		}

		database.DB.Create(&imageUpload)
	}

	var listImagePengaduan []models.Image

	database.DB.Where("pengaduan_id = ?", newPengaduan.Id).Find(&listImagePengaduan)

	return c.JSON(fiber.Map{
		"data":       newPengaduan,
		"image_data": listImagePengaduan,
	})
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

	database.DB.First(&pengaduan, id)

	var imagePengaduan []models.Image
	database.DB.Where("pengaduan_id = ?", id).Find(&imagePengaduan)

	if pengaduan.Id == 0 {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "Data not found",
		})
	}

	return c.JSON(fiber.Map{
		"data":       pengaduan,
		"image_data": imagePengaduan,
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
