package apiV1

import (
	"log"
	"net/http"

	"github.com/2110366-2566-2/tortoise-app-backend/configs"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/storage"
	"github.com/gin-gonic/gin"
)

// Services for Testing
func TestSellerServices() {
	log.Println("Seller services! ...")
}

func TestAdminServices() {
	log.Println("Admin services! ...")
}

// End of Tested Services

func SetupRoutes(r *gin.Engine, dbH *database.Handler, stgH *storage.Handler, env configs.EnvVars) {
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

	UnauthorizedRoutes(apiV1, dbH, stgH)

	// ============ End of Unauthorized routes ============

	// Add JWT middleware to check the token
	apiV1.Use(jwtMiddleware(env))

	// All user can access
	userGroup := apiV1.Group("/user")
	userGroup.Use(roleMiddleware("seller", "admin", "buyer"))
	UserServices(userGroup, dbH, stgH)

	reviewGroup := apiV1.Group("/review")
	reviewGroup.Use(roleMiddleware("seller", "admin", "buyer"))
	ReviewServices(reviewGroup, dbH)

	reportGroup := apiV1.Group("/report")
	reportGroup.Use(roleMiddleware("seller", "admin", "buyer"))
	ReportServices(reportGroup, dbH)

	// Seller and Admin and Buyer can access

	// Get token session
	apiV1.GET("/token/session", roleMiddleware("seller", "admin", "buyer"), func(c *gin.Context) {
		userID, _ := c.Get("userID")
		username, _ := c.Get("username")
		role, _ := c.Get("role")
		c.JSON(http.StatusOK, gin.H{"userID": userID, "username": username, "role": role})
	})

	petsGroup := apiV1.Group("/pets")
	// petsGroup.Use(roleMiddleware("seller", "admin", "buyer"))
	PetController(petsGroup, dbH, stgH)

	transactionGroup := apiV1.Group("/transactions")
	transactionGroup.Use(roleMiddleware("seller", "admin", "buyer"))
	TransactionServices(transactionGroup, dbH)

	// reveiwGrop := apiV1.Group("/review")
	// reveiwGrop.Use(roleMiddleware("seller", "admin", "buyer"))
	// ReviewServices(reveiwGrop,dbH)

	// Seller and Admin can access
	bankGroup := apiV1.Group("/bank")
	bankGroup.Use(roleMiddleware("seller", "admin"))
	BankServices(bankGroup, dbH)

	// Buyer and Admin can access
	paymentGroup := apiV1.Group("/payment")
	paymentGroup.Use(roleMiddleware("buyer", "admin"))
	PaymentServices(paymentGroup, dbH, env)

	// apiV1.Group("/seller").Use(roleMiddleware("seller", "admin")).GET("/", func(c *gin.Context) {
	// 	TestSellerServices()
	// })

	// Admin can access
	adminGrqoup := apiV1.Group("/admin")
	adminGrqoup.Use(roleMiddleware("admin"))
	AdminServices(adminGrqoup, dbH)

	// apiV1.Group("/admin").Use(roleMiddleware("admin")).GET("/", func(c *gin.Context) {
	// 	TestAdminServices()
	// })

	log.Println("Routes are set up successfully!")
}
