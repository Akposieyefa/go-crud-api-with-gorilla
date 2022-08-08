package initializers

import (
	"encoding/json"
	"net/http"

	"github.com/akposieyefa/crud-gorilla/database"
	"github.com/akposieyefa/crud-gorilla/models"
	"github.com/gorilla/mux"
)

//get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	database.DB.Find(&users)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"users":   users,
		"message": "Users fetched successfully",
		"success": true,
	})
}

//create user
func CreateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Create(&user)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"users":   user,
		"message": "success in creating users",
		"success": true,
	})
}

//get single user
func GetSingleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	params := mux.Vars(r)
	database.DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(map[string]interface{}{
		"users":   user,
		"message": "User details feteched successfully",
		"success": true,
	})
}

//update user
func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	params := mux.Vars(r)
	database.DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Save(&user)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"users":   user,
		"message": "User details updated successfully",
		"success": true,
	})
}

//delete user
func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	params := mux.Vars(r)
	database.DB.First(&user, params["id"])
	database.DB.Delete(&user)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"users":   user,
		"message": "User details deleted successfully",
		"success": true,
	})
}
