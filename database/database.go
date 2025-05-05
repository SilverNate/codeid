package database

import (
	model "codeid/murid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// drop table each time we go run
	err = db.Migrator().DropTable(&model.Murid{}, &model.Pendidikan{})
	if err != nil {
		log.Fatalf("failed to drop tables: %v", err)
	}

	if err := db.AutoMigrate(&model.Murid{}, &model.Pendidikan{}); err != nil {
		log.Fatalf("failed to migrate user model: %v", err)
	}

	//publisher := model.Publisher{Name: "Penguin Books"}
	//db.Create(&publisher)
	//
	//author := model.Author{Name: "George Orwell"}
	//db.Create(&author)
	//
	//book := model.Book{
	//	Title:       "1984",
	//	AuthorID:    author.ID,
	//	PublisherID: publisher.ID,
	//}
	//db.Create(&book)
	//
	//var createdBook model.Book
	//db.Preload("Author").Preload("Publisher").First(&createdBook)
	//fmt.Printf("Book: %s, Author: %s, Publisher: %s\n", createdBook.Title, createdBook.Author.Name, createdBook.Publisher.Name)

	return db
}
