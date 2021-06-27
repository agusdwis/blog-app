package seed

import (
	"log"

	"github.com/agusdwis/blog-app/models"
	"gorm.io/gorm"
)


var users = []models.User{
	models.User{
		Username: "userone",
		Email: "user1@mail.com",
		Password: "qwe123",
	},
	models.User{
		Username: "usertwo",
		Email: "user2@mail.com",
		Password: "qwe123",
	},
}

var posts = []models.Post{
	models.Post{
		Title: "Title 1",
		Content: "Content 1",
	},
	models.Post{
		Title: "Title 2",
		Content: "Content 2",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().Migrator().DropTable(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("Cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("Cannot migrate table: %v", err)
	}

	for i := range users {
		err1 := db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err1 != nil {
			log.Fatalf("Cannot seed users table: %v", err)
		}

		posts[i].AuthorID = users[i].ID

		err2 := db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err2 != nil {
			log.Fatalf("Cannot seed posts table: %v", err)
		}
	}


}
