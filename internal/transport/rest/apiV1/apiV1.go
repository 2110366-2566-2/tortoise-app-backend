package apiV1

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/2110366-2566-2/tortoise-app-backend/configs"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/services"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PetController(r *gin.RouterGroup, h *database.Handler) {
	// Create a new pet handler
	petHandler := services.NewPetHandler(h)

	// Set up routes
	r.GET("/", petHandler.GetAllPets)
	r.GET("/filter", petHandler.GetFilteredPets)
	r.GET("/:petID", petHandler.GetPetByPetID)
	r.GET("/seller/:userID", petHandler.GetPetBySeller)

	r.POST("/seller/:userID", petHandler.CreatePet)
	r.PUT("/:petID", petHandler.UpdatePet)
	r.DELETE("/:petID", petHandler.DeletePet)

	// master data
	r.GET("/master", petHandler.GetMasterData)
	r.GET("/master/:category", petHandler.GetMasterDataByCategory)
	r.GET("/master/categories", petHandler.GetCategories)
}

// func AccountServices(r *gin.RouterGroup, h *database.Handler) {
// 	r.GET("/transactions-history", func(c *gin.Context) {
// 		services.GetTransactions(c, h)
// 	})
// }

func UserServices(r *gin.RouterGroup, h *database.Handler) {

	// Create a new user handler
	userHandler := services.NewUserHandler(h)

	r.GET("/:userID", userHandler.GetUserByUserID)
	r.PUT("/passwd/:userID", userHandler.UpdateUserPasswd)
	r.PUT("/:userID", userHandler.UpdateUser)
	r.DELETE("/:userID", userHandler.DeleteUser)
	// r.GET("/token/session", func(c *gin.Context) {
	//     services.GetSessionToken(c, h)
	// })
	r.POST("/recoverusername", userHandler.RecoveryUsername)
	r.POST("/checkvalidemail", userHandler.CheckMail)
	r.POST("/sentotp", userHandler.SentOTP)
	r.POST("/checkotp", userHandler.ValidateOTP)
	r.POST("/forgotpasswd", userHandler.UpdateForgotPassword)

	// Get me
	r.GET("/me", userHandler.WhoAmI)

}

func TransactionServices(r *gin.RouterGroup, h *database.Handler) {
	// Create a new transaction handler
	transactionHandler := services.NewTransactionHandler(h)

	r.GET("/", transactionHandler.GetTransactions)
	r.GET("/:transactionID", transactionHandler.GetTransactionByTransactionID)
}

func BankServices(r *gin.RouterGroup, h *database.Handler) {
	// Create a new seller handler
	sellerHandler := services.NewSellerHandler(h)

	// Set up routes
	r.POST("/:sellerID", sellerHandler.AddBankAccount)
	r.GET("/:sellerID", sellerHandler.GetBankAccount)
	r.DELETE("/:sellerID", sellerHandler.DeleteBankAccount)
}

func PaymentServices(r *gin.RouterGroup, h *database.Handler, env configs.EnvVars) {
	// Create a new buyer handler
	buyerHandler := services.NewPaymentHandler(h, env)

	// Set up routes
	r.POST("/create", buyerHandler.CreatePayment)
	r.POST("/confirm", buyerHandler.ConfirmPayment)
}

func ReportServices(r *gin.RouterGroup, h *database.Handler) {
	// Create a new report handler
	reportHandler := services.NewReportHandler(h)

	// Set up routes
	r.POST("/party",reportHandler.CreatePartyReport)
	r.POST("/system",reportHandler.CreateSystemReport)
}

func UnauthorizedRoutes(r *gin.RouterGroup, h *database.Handler) {
	// login and register
	r.POST("/login", func(c *gin.Context) {
		services.LoginHandler(c, h)
	})

	r.POST("/register", func(c *gin.Context) {
		services.RegisterHandler(c, h)
	})
}

func ReviewServices(r *gin.RouterGroup, h *database.Handler) {

	// Create a new review handler
	reviewHandler := services.NewReviewHandler(h)

	// Set up routes
	r.POST("/create", reviewHandler.CreateReview)
	r.PUT("/comment/:reviewID", reviewHandler.AddComment)
	r.GET("/:sellerID", reviewHandler.GetReviewBySeller)
}

// Services for Testing
func TestSellerServices() {
	log.Println("Seller services! ...")
}
func TestAdminServices() {
	log.Println("Admin services! ...")
}

// End of Tested Services

func SetupRoutes(r *gin.Engine, h *database.Handler, env configs.EnvVars) {
	// env, err := configs.LoadConfig()
	// if err != nil {
	// 	panic(err)
	// }
	// Set up routes
	apiV1 := r.Group("/api/v1")

	apiV1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "This is API v1.0.0")
	})

	// ================ Unauthorized routes ================

	UnauthorizedRoutes(apiV1, h)

	// ============ End of Unauthorized routes ============

	// Add JWT middleware to check the token
	apiV1.Use(jwtMiddleware(env))

	// All user can access
	userGroup := apiV1.Group("/user")
	userGroup.Use(roleMiddleware("seller", "admin", "buyer"))
	UserServices(userGroup, h)

	reviewGroup := apiV1.Group("/review")
	reviewGroup.Use(roleMiddleware("seller", "admin", "buyer"))
	ReviewServices(reviewGroup, h)

	reportGroup := apiV1.Group("/report")
	reportGroup.Use(roleMiddleware("seller", "admin", "buyer"))
	ReportServices(reportGroup, h)

	// Seller and Admin and Buyer can access

	// Get token session
	apiV1.GET("/token/session", roleMiddleware("seller", "admin", "buyer"), func(c *gin.Context) {
		userID, _ := c.Get("userID")
		username, _ := c.Get("username")
		role, _ := c.Get("role")
		c.JSON(http.StatusOK, gin.H{"userID": userID, "username": username, "role": role})
	})

	petsGroup := apiV1.Group("/pets")
	petsGroup.Use(roleMiddleware("seller", "admin", "buyer"))
	PetController(petsGroup, h)

	transactionGroup := apiV1.Group("/transactions")
	transactionGroup.Use(roleMiddleware("seller", "admin", "buyer"))
	TransactionServices(transactionGroup, h)

	// reveiwGrop := apiV1.Group("/review")
	// reveiwGrop.Use(roleMiddleware("seller", "admin", "buyer"))
	// ReviewServices(reveiwGrop, h)

	// Seller and Admin can access
	bankGroup := apiV1.Group("/bank")
	bankGroup.Use(roleMiddleware("seller", "admin"))
	BankServices(bankGroup, h)

	// Buyer and Admin can access
	paymentGroup := apiV1.Group("/payment")
	paymentGroup.Use(roleMiddleware("buyer", "admin"))
	PaymentServices(paymentGroup, h, env)

	// apiV1.Group("/seller").Use(roleMiddleware("seller", "admin")).GET("/", func(c *gin.Context) {
	// 	TestSellerServices()
	// })

	// Admin can access
	apiV1.Group("/admin").Use(roleMiddleware("admin")).GET("/", func(c *gin.Context) {
		TestAdminServices()
	})

	log.Println("Routes are set up successfully!")
}

func jwtMiddleware(env configs.EnvVars) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT secret key from the environment
		secretKey := env.JWT_SECRET

		// Create a new JWT middleware
		authMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
		})

		// Handle the JWT middleware
		err := authMiddleware.CheckJWT(c.Writer, c.Request)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Extract the token from the request
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Extract the role from the token
		claims := token.Claims.(jwt.MapClaims)
		userID, _ := primitive.ObjectIDFromHex(claims["userID"].(string))
		username := claims["username"].(string)
		role := claims["role"].(string)

		// Pass the role to the next middleware/handler
		c.Set("role", role)
		c.Set("userID", userID)
		c.Set("username", username)

		c.Next()
	}
}

func roleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
	}
}
