package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"go-fiber-postgres/models"
	"go-fiber-postgres/storage"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}

	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}

	// Validate book data
	if book.Author == "" || book.Title == "" || book.Publisher == "" {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Invalid book data"})
		return nil
	}

	// Create the book in the database.
	err = r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not create book"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Book has been added",
	})
	return nil
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	// Delete the book from the database.
	result := r.DB.Delete(&models.Books{}, id)

	if result.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not delete book"})
		return result.Error
	}

	if result.RowsAffected == 0 {
		context.Status(http.StatusNotFound).JSON(
			&fiber.Map{"message": "Book not found"})
		return nil
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Book deleted successfully",
	})
	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}

	// Retrieve all books from the database.
	err := r.DB.Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not get books"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Books fetched successfully",
		"data":    bookModels,
	})
	return nil
}

func (r *Repository) GetBooksPaginated(context *fiber.Ctx) error {
	page, _ := strconv.Atoi(context.Query("page", "1"))
	limit, _ := strconv.Atoi(context.Query("limit", "10"))

	offset := (page - 1) * limit
	books := []models.Books{}
	err := r.DB.Offset(offset).Limit(limit).Find(&books).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not get books"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Books fetched successfully",
		"data":    books,
	})
	return nil
}

func (r *Repository) GetBookByID(context *fiber.Ctx) error {
	id := context.Params("id")
	bookModel := &models.Books{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "ID cannot be empty",
		})
		return nil
	}

	// Retrieve the book by ID from the database.
	err := r.DB.Where("id = ?", id).First(bookModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not get the book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Book ID fetched successfully",
		"data":    bookModel,
	})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_books/:id", r.GetBookByID)
	api.Get("/books", r.GetBooksPaginated)
	api.Get("/books_all", r.GetBooks)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize database configuration.
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	// Connect to the database.
	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Could not load the database")
	}

	// Migrate the books table.
	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("Could not migrate the database")
	}

	// Create a Repository instance and set up API routes.
	r := Repository{
		DB: db,
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	r.SetupRoutes(app)

	// Start the application on port 8080.
	app.Listen(":8080")
}
