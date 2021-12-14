package controllers

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadImageTest(c *fiber.Ctx) error {

	form, err := c.MultipartForm()
	files := form.File["image"]

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	var imageURL = []string{}

	for _, file := range files {
		// generate new uuid for image name
		uniqueId := uuid.New()

		// remove "- from imageName"

		filename := strings.Replace(uniqueId.String(), "-", "", -1)

		// extract image extension from original file filename

		fileExt := filepath.Ext(file.Filename)
		// generate image from filename and extension
		image := fmt.Sprintf("%s%s", filename, fileExt)

		// save image to ./images dir
		err = c.SaveFile(file, fmt.Sprintf("./images/%s", image))

		if err != nil {
			log.Println("image save error --> ", err)
			return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
		}

		// generate image url to serve to client using CDN

		imageLoc := fmt.Sprintf("http://localhost:8050/images/%s", image)

		imageURL = append(imageURL, imageLoc)
	}

	for _, image := range imageURL {
		fmt.Println(image)
	}
	// data := map[string]interface{}{
	// testing := c.FormValue("testing")
	// data := string(testing)

	// "imageName": image,
	// "imageUrl":  imageUrl,
	// "header":    file.Header,
	// "size":      file.Size,
	// "fileExt":   fileExt,
	// "filename":  filename,
	// }

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": imageURL})

}
