package controllers

import (
	"test-golang/database"
	"test-golang/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func CreateNote(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	note := models.Note{
		Title:    data["title"],
		Category: data["category"],
		Details:  data["details"],
		UserID:   user.ID,
	}

	database.DB.Create(&note)

	return c.JSON(note)
}

func GetNotes(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var notes []models.Note

	database.DB.Where("user_id = ?", claims.Issuer).Find(&notes)

	if len(notes) == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "No notes created",
		})
	}

	return c.JSON(notes)
}

func GetNoteById(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	note_id := c.Params("NoteID")

	var note models.Note

	database.DB.Where("id = ? AND user_id = ?", note_id, claims.Issuer).Find(&note)

	if note.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Note not found",
		})
	}

	return c.JSON(note)
}

func UpdateNoteById(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	note_id := c.Params("NoteID")

	var note models.Note

	database.DB.Where("id = ? AND user_id = ?", note_id, claims.Issuer).Find(&note)

	if note.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Note not found",
		})
	}

	note.Title = data["title"]
	note.Category = data["category"]
	note.Details = data["details"]
	database.DB.Save(&note)

	return c.JSON(note)
}

func DeleteNote(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}


	note_id := c.Params("NoteID")
	var note models.Note

	database.DB.Where("id = ?", note_id).Find(&note)

	if note.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Note not found",
		})
	}

	//database.DB.Unscoped().Delete(&note) // Permanent deletion
	database.DB.Delete(&note) // Soft deletion 

	return c.JSON(fiber.Map{
		"message": "Note deleted successfully",
	})
}