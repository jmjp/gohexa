package routes

import (
	AuthHandler "deluze/internal/adapters/handler/http"
	EmailRepository "deluze/internal/adapters/repository/email"
	OtpRepository "deluze/internal/adapters/repository/otp"
	UserRepository "deluze/internal/adapters/repository/user"
	AuthService "deluze/internal/core/services/auth"
	"deluze/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	//repositories
	userRepo := UserRepository.NewUserMongoRepository()
	otpRepo := OtpRepository.NewOtpRedisRepository()
	emailRepo := EmailRepository.NewEmailTrapRepository()

	//services
	authService := AuthService.New(userRepo, otpRepo, emailRepo)

	//handlers
	authHandler := AuthHandler.NewHttpAuthHandler(authService)

	//routes
	auth := app.Group("/auth")
	auth.Post("/", authHandler.Login)
	auth.Get("/", authHandler.Verify)
	auth.Get("/logout", middlewares.IsAuthenticated, authHandler.Logout)
}
