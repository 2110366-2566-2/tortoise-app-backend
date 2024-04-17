package test

import (
    "bytes"
    // "encoding/json"
    // "fmt"
    "net/http"
    "net/http/httptest"
    "testing"
	"time"
	"context"
    // "log"


    "github.com/2110366-2566-2/tortoise-app-backend/internal/database"
    "github.com/2110366-2566-2/tortoise-app-backend/internal/storage"
    "github.com/2110366-2566-2/tortoise-app-backend/internal/services"
    "github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	// "github.com/2110366-2566-2/tortoise-app-backend/configs"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func setup() (*gin.Engine, *database.Handler) {
    //Create a new handler
    db, _, _ := database.ConnectMongo("mongodb+srv://admin:M3jgA2Kan6RgrIcR@petpal.ai2diov.mongodb.net/petpal?retryWrites=true&w=majority", "petpal", 10*time.Second)
	dbHandler := database.NewHandler(db)

	// connect to firebase
	stg, _ := storage.ConnectFirebase(context.Background(), "../configs/config.json")
	stgHandler := storage.NewHandler(stg)

    userHandler := services.NewUserHandler(dbHandler, stgHandler)
    
    // Create a new gin router
    r := gin.Default()
    r.PUT("/api/v1/user/:userID", userHandler.UpdateUser)
        
    // Create a new HTTP request
    w := httptest.NewRecorder()
    //-----------------------------
    req, _ := http.NewRequest("PUT", "/api/v1/user/661f8ce33e12e57c0c400302", bytes.NewBuffer([]byte(`
	{
		"username": "mahiru",
		"email": "mahiru@gmail.com",
		"first_name": "mahiru",
		"last_name": "shiina",
		"gender": "Female",
		"phoneNumber": "0123456789",
		"role": 2,
		"address": {
			"province": "",
			"district": "",
			"subdistrict": "",
			"postalCode": "",
			"street": "",
			"building": "",
			"houseNumber": ""
		}
	}
    `)))

    //-----------------------------
    req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTM0MzAxMTUsInJvbGUiOiJidXllciIsInVzZXJJRCI6IjY2MWY4Y2UzM2UxMmU1N2MwYzQwMDMwMiIsInVzZXJuYW1lIjoibWFoaXJ1In0.p-pI12Id1-uAzwVjmOvuyAPGK3Jy8iWj4MYeo1ouxCk")
    // Serve the HTTP request to the handler
    r.ServeHTTP(w, req)

    return r,dbHandler
}



func defaultAssert(user models.User, t *testing.T) {
    assert.Equal(t, "mahiru", user.FirstName)
    assert.Equal(t, "shiina", user.LastName)
    assert.Equal(t, "Female", user.Gender)
    assert.Equal(t, "0123456789", user.PhoneNumber)
}

func TestCorrectUpdateUser(t *testing.T) {
    r,dbHandler:= setup()
    w := httptest.NewRecorder()

    //-----------------------------
    req, _ := http.NewRequest("PUT", "/api/v1/user/661f8ce33e12e57c0c400302", bytes.NewBuffer([]byte(`
	{
		"first_name": "yamada",
		"last_name": "anna",
		"gender": "Male",
		"phoneNumber": "9876543210"
	}
    `)))
    //-----------------------------

    // req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTM0MzAxMTUsInJvbGUiOiJidXllciIsInVzZXJJRCI6IjY2MWY4Y2UzM2UxMmU1N2MwYzQwMDMwMiIsInVzZXJuYW1lIjoibWFoaXJ1In0.p-pI12Id1-uAzwVjmOvuyAPGK3Jy8iWj4MYeo1ouxCk")

    // Serve the HTTP request to the handler
    r.ServeHTTP(w, req)

    user,_ := dbHandler.GetUserByUserID(context.Background(),"661f8ce33e12e57c0c400302")
    

    // assert.NotEqual(t, http.StatusOK, w.Code)
    //     // w = w_prev
    //     // user :=  MarshalManage(w)
    // defaultAssert(*user, t)


    assert.Equal(t, http.StatusOK, w.Code)
        // user :=  MarshalManage(w)
    //-----------------------------
    assert.Equal(t, "yamada", user.FirstName)
    assert.Equal(t, "anna", user.LastName)
    assert.Equal(t, "Male", user.Gender)
    assert.Equal(t, "9876543210", user.PhoneNumber)
    //-----------------------------
    

}

func TestNoFirstNameUpdateUser(t *testing.T) {
    r,dbHandler:= setup()
    w := httptest.NewRecorder()

    //-----------------------------
    req, _ := http.NewRequest("PUT", "/api/v1/user/661f8ce33e12e57c0c400302", bytes.NewBuffer([]byte(`
	{
		"first_name": "",
		"last_name": "anna",
		"gender": "Female",
		"phoneNumber": "9876543210"
	}
    `)))
    //-----------------------------

    // req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTM0MzAxMTUsInJvbGUiOiJidXllciIsInVzZXJJRCI6IjY2MWY4Y2UzM2UxMmU1N2MwYzQwMDMwMiIsInVzZXJuYW1lIjoibWFoaXJ1In0.p-pI12Id1-uAzwVjmOvuyAPGK3Jy8iWj4MYeo1ouxCk")

    // Serve the HTTP request to the handler
    r.ServeHTTP(w, req)

    user,_ := dbHandler.GetUserByUserID(context.Background(),"661f8ce33e12e57c0c400302")
    
    assert.NotEqual(t, http.StatusOK, w.Code)
    defaultAssert(*user, t)
}

func TestNoLastNameUpdateUser(t *testing.T) {
    r,dbHandler:= setup()
    w := httptest.NewRecorder()

    //-----------------------------
    req, _ := http.NewRequest("PUT", "/api/v1/user/661f8ce33e12e57c0c400302", bytes.NewBuffer([]byte(`
	{
		"first_name": "yamada",
		"last_name": "",
		"gender": "Female",
		"phoneNumber": "9876543210"
	}
    `)))
    //-----------------------------

    // req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTM0MzAxMTUsInJvbGUiOiJidXllciIsInVzZXJJRCI6IjY2MWY4Y2UzM2UxMmU1N2MwYzQwMDMwMiIsInVzZXJuYW1lIjoibWFoaXJ1In0.p-pI12Id1-uAzwVjmOvuyAPGK3Jy8iWj4MYeo1ouxCk")

    // Serve the HTTP request to the handler
    r.ServeHTTP(w, req)

    user,_ := dbHandler.GetUserByUserID(context.Background(),"661f8ce33e12e57c0c400302")
    
    assert.NotEqual(t, http.StatusOK, w.Code)
    defaultAssert(*user, t)
}


    //func MarshalManage(w *httptest.ResponseRecorder) models.User {
    //     // Assuming w.Body.String() gives you the JSON string
    //     jsonStr := w.Body.String()
    
    //     // We'll unmarshal into this variable
    //     var result map[string]interface{}
    
    //     // Unmarshal the JSON string
    //     _ = json.Unmarshal([]byte(jsonStr), &result)
    
    //     // Now, result["data"] should hold the User data
    //     userData := result["data"].(map[string]interface{})
    
    //     // We'll unmarshal the user data into a User struct
    //     var user models.User
    //     userDataBytes, _ := json.Marshal(userData)
    //     _ = json.Unmarshal(userDataBytes, &user)
    //     log.Println("user:",user)
    //     return user
    // }