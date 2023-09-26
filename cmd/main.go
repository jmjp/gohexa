package main

import (
	"deluze/pkg/servers"

	"github.com/joho/godotenv"
)

func main() {
	// var userService *UserService.UserService
	// db_type := "mongo"

	// switch db_type {
	// case "mongo":
	// 	mongoRepo := UserRepository.NewUserMongoRepository()
	// 	userService = UserService.New(mongoRepo)
	// 	break
	// case "pg":
	// 	pgRepo := UserRepository.NewUserPostgresRepository()
	// 	userService = UserService.New(pgRepo)
	// }

	// usr, err := userService.FindMany(nil, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(usr)
	godotenv.Load(".env")
	servers.StartupApi(":8080")
	// mongoRepo := UserRepository.NewUserMongoRepository()
	// otpRepo := OtpRepository.NewOtpMongoRepository()
	// authService := AuthService.New(mongoRepo, otpRepo)
	// err := authService.Login("joaozinhosleepy@gmail.com")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("success")
	// user, token, err := authService.Verify("joaozinhosleepy@gmail.com", "515978")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(*token)
	// fmt.Println(user)

}
