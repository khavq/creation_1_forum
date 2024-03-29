package tests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/khavq/creation_1_forum/api/controllers"
	"github.com/khavq/creation_1_forum/api/models"
)

var server = controllers.Server{}
var userInstance = models.User{}
var postInstance = models.Post{}
var likeInstance = models.Like{}
var commentInstance = models.Comment{}

func TestMain(m *testing.M) {
	// UNCOMMENT THIS WHILE TESTING ON LOCAL(WITHOUT USING CIRCLE CI), BUT LEAVE IT COMMENTED IF YOU ARE USING CIRCLE CI
	//var err error
	//err = godotenv.Load(os.ExpandEnv("./../.env"))
	//if err != nil {
	//log.Fatalf("Error getting env %v\n", err)
	//}

	Database()

	os.Exit(m.Run())

}

func Database() {

	var err error

	////////////////////////////////// UNCOMMENT THIS WHILE TESTING ON LOCAL(WITHOUT USING CIRCLE CI) ///////////////////////
	//TestDbDriver := os.Getenv("TEST_DB_DRIVER")
	//if TestDbDriver == "mysql" {
	//DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_PASSWORD"), os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_NAME"))
	//server.DB, err = gorm.Open(TestDbDriver, DBURL)
	//if err != nil {
	//fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
	//log.Fatal("This is the error:", err)
	//} else {
	//fmt.Printf("We are connected to the %s database\n", TestDbDriver)
	//}
	//}
	//if TestDbDriver == "postgres" {
	//DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_NAME"), os.Getenv("TEST_DB_PASSWORD"))
	//server.DB, err = gorm.Open(TestDbDriver, DBURL)
	//if err != nil {
	//fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
	//log.Fatal("This is the error:", err)
	//} else {
	//fmt.Printf("We are connected to the %s database\n", TestDbDriver)
	//}
	//}
	/////////////////////////////////  END OF LOCAL TEST DATABASE SETUP ///////////////////////////////////////////////////

	//////////////////////////////////  COMMENT THIS WHILE TESTING ON LOCAL(WITHOUT USING CIRCLE CI)  //////////////////////
	// WE HAVE TO INPUT TESTING DATA MANUALLY BECAUSE CIRCLECI, CANNOT READ THE ".env" FILE WHICH, WE WOULD HAVE ADDED THE TEST CONFIG THERE
	// SO MANUALLY ADD THE NAME OF THE DATABASE, THE USER AND THE PASSWORD, AS SEEN BELOW:

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "khavq", "password", "127.0.0.1", "3306", "creation_1_forum_test")
	//DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "127.0.0.1", "5432", "steven", "forum_db_test", "password")
	server.DB, err = gorm.Open("mysql", DBURL)
	//server.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", "postgres")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", "postgres")
	}

	//////////////////////////////// END OF USING CIRCLE CI ////////////////////////////////////////////////////////////////
}

func refreshUserTable() error {
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneUser() (models.User, error) {

	user := models.User{
		Username: "Pet",
		Email:    "pet@example.com",
		Password: "password",
	}

	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func seedUsers() ([]models.User, error) {

	var err error
	if err != nil {
		return nil, err
	}
	users := []models.User{
		models.User{
			Username: "Steven",
			Email:    "steven@example.com",
			Password: "password",
		},
		models.User{
			Username: "Kenny",
			Email:    "kenny@example.com",
			Password: "password",
		},
	}

	for i, _ := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return []models.User{}, err
		}
	}
	return users, nil
}

func refreshUserAndPostTable() error {

	err := server.DB.DropTableIfExists(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneUserAndOnePost() (models.User, models.Post, error) {

	user := models.User{
		Username: "Sam",
		Email:    "sam@example.com",
		Password: "password",
	}
	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.User{}, models.Post{}, err
	}
	post := models.Post{
		Title:    "This is the title sam",
		Content:  "This is the content sam",
		AuthorID: user.ID,
	}
	err = server.DB.Model(&models.Post{}).Create(&post).Error
	if err != nil {
		return models.User{}, models.Post{}, err
	}
	return user, post, nil
}

func seedUsersAndPosts() ([]models.User, []models.Post, error) {

	var err error

	if err != nil {
		return []models.User{}, []models.Post{}, err
	}
	var users = []models.User{
		models.User{
			Username: "Steven",
			Email:    "steven@example.com",
			Password: "password",
		},
		models.User{
			Username: "Magu",
			Email:    "magu@example.com",
			Password: "password",
		},
	}
	var posts = []models.Post{
		models.Post{
			Title:   "Title 1",
			Content: "Hello world 1",
		},
		models.Post{
			Title:   "Title 2",
			Content: "Hello world 2",
		},
	}

	for i, _ := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = server.DB.Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
	return users, posts, nil
}

func refreshUserPostAndLikeTable() error {
	err := server.DB.DropTableIfExists(&models.User{}, &models.Post{}, &models.Like{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Like{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed user, post and like tables")
	return nil
}

func seedUsersPostsAndLikes() (models.Post, []models.User, []models.Like, error) {
	// The idea here is: two users can like one post
	var err error
	var users = []models.User{
		models.User{
			Username: "Steven",
			Email:    "steven@example.com",
			Password: "password",
		},
		models.User{
			Username: "Magu",
			Email:    "magu@example.com",
			Password: "password",
		},
	}
	post := models.Post{
		Title:   "This is the title",
		Content: "This is the content",
	}
	err = server.DB.Model(&models.Post{}).Create(&post).Error
	if err != nil {
		log.Fatalf("cannot seed post table: %v", err)
	}
	var likes = []models.Like{
		models.Like{
			UserID: 1,
			PostID: post.ID,
		},
		models.Like{
			UserID: 2,
			PostID: post.ID,
		},
	}
	for i, _ := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		err = server.DB.Model(&models.Like{}).Create(&likes[i]).Error
		if err != nil {
			log.Fatalf("cannot seed likes table: %v", err)
		}
	}
	return post, users, likes, nil
}

func refreshUserPostAndCommentTable() error {
	err := server.DB.DropTableIfExists(&models.User{}, &models.Post{}, &models.Comment{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed user, post and comment tables")
	return nil
}

func seedUsersPostsAndComments() (models.Post, []models.User, []models.Comment, error) {
	// The idea here is: two users can comment one post
	var err error
	var users = []models.User{
		models.User{
			Username: "Steven",
			Email:    "steven@example.com",
			Password: "password",
		},
		models.User{
			Username: "Magu",
			Email:    "magu@example.com",
			Password: "password",
		},
	}
	post := models.Post{
		Title:   "This is the title",
		Content: "This is the content",
	}
	err = server.DB.Model(&models.Post{}).Create(&post).Error
	if err != nil {
		log.Fatalf("cannot seed post table: %v", err)
	}
	var comments = []models.Comment{
		models.Comment{
			Body:   "user 1 made this comment",
			UserID: 1,
			PostID: post.ID,
		},
		models.Comment{
			Body:   "user 2 made this comment",
			UserID: 2,
			PostID: post.ID,
		},
	}
	for i, _ := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		err = server.DB.Model(&models.Like{}).Create(&comments[i]).Error
		if err != nil {
			log.Fatalf("cannot seed comments table: %v", err)
		}
	}
	return post, users, comments, nil
}

func refreshUserAndResetPasswordTable() error {
	err := server.DB.DropTableIfExists(&models.User{}, &models.ResetPassword{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}, &models.ResetPassword{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed user and resetpassword tables")
	return nil
}

// Seed the reset password table with the token
func seedResetPassword() (models.ResetPassword, error) {

	resetDetails := models.ResetPassword{
		Token: "awesometoken",
		Email: "pet@example.com",
	}
	err := server.DB.Model(&models.ResetPassword{}).Create(&resetDetails).Error
	if err != nil {
		return models.ResetPassword{}, err
	}
	return resetDetails, nil
}
