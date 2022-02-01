package migrations

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/SeedlngTeam/Core/helpers"
	entity "github.com/SeedlngTeam/Core/internal/entities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func GetEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error Loading .env file!!")
	}
	return os.Getenv(key)
}

func GetEnvVariablePort(key string) int {
	value := GetEnvVariable(key)
	i, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Port Number is incorrectly returning as a string")
	}
	return i
}

func connectDB() *gorm.DB {
	url := GetEnvVariable("DB_URL")
	port := GetEnvVariablePort("DB_PORT")
	user := GetEnvVariable("DB_USER")
	pass := GetEnvVariable("DB_PASS")
	dbname := GetEnvVariable("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		url, port, user, pass, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	helpers.HandleErr(err)
	return db
}

// func createUsers() {
// 	db := connectDB()

// 	// Test User Accounts
// 	users := &[3]entity.User{
// 		{Username: "Bob", Email: "bob@seedlng.com"},
// 		{Username: "Marietta", Email: "marietss@seedlng.com"},
// 		{Username: "Young", Email: "young@seedlng.com"},
// 	}

// 	for i := 0; i < len(users); i++ {

// 		tempPassword := helpers.HashAndSalt([]byte(users[i].Username))
// 		user := &entity.User{Username: users[i].Username, Password: tempPassword, Email: users[i].Email}
// 		db.Create(&user)

// 		account := &entity.Account{Type: "Checking", Name: "Personal", UserID: user.ID, Balance: uint(5000)}
// 		db.Create(&account)
// 	}

// 	defer db.Close()
// 	// End of Test
// }

func createUsers(user *entity.User, acc *entity.Account) {
	db := connectDB()
	saltedPass := helpers.HashAndSalt([]byte(user.Password))
	newUser := &entity.User{Username: user.Username, Password: saltedPass, Email: user.Email, Phone: user.Phone}
	db.Create(&newUser)
	createAccounts(db, int(newUser.ID), acc.Type, acc.Name, acc.Balance)
	defer db.Close()
}

func createAccounts(db *gorm.DB, userid int, accType string, acctName string, balance uint) {
	account := &entity.Account{Type: accType, Name: acctName, UserID: uint(userid), Balance: balance}
	db.Create(&account)
}

func Migrate() {
	db := connectDB()
	db.AutoMigrate(&entity.User{}, &entity.Account{})
	defer db.Close()

	createUsers()
}
