package controllers

import (
	"learn/myapp/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllBooks(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var books []models.Book
		result := *db.Raw("SELECT * from books").Scan(&books); if result.Error != nil {
			return c.String(200, result.Error.Error())
		}

		return c.JSON(200, books)
	}
}

func GetBook(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		var book models.Book
		result := *db.Raw("SELECT * from books WHERE id = ?", id).Scan(&book); if result.Error != nil {
			return c.String(200, result.Error.Error())
		}

		return c.JSON(200, book)
	}
}

func AddBook(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		book := new(models.Book)
		if err := c.Bind(book); err != nil {
			return err
		}

		result := db.Create(&book); if result.Error != nil {
			return c.String(200, result.Error.Error())
		}

		return c.JSON(200, book)
	}
}

func UpdateBook(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		book := new(models.Book)
		if err := c.Bind(book); err != nil {
			return err
		}

		result := db.Model(&models.Book{}).Where("id = ?", id).Updates(&book); if result.Error != nil {
			return c.String(200, result.Error.Error())
		}

		return c.JSON(200, book)
	}
}

func DeleteBook(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		result := db.Delete(&models.Book{}, id); if result.Error != nil {
			return c.String(200, result.Error.Error())
		}

		return c.String(200, "Book deleted")
	}
}