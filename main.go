package main

import (
	"errors"
	"net/http"
	rc "restaurant-catalog/structs"

	"github.com/gin-gonic/gin"
)

// TODO turn this to MenuItems
var menuItems = []rc.MenuItem{
	{ID: "1", Name: "Test Item", Description: "Test Description", Price: 3.99, ImageID: "abc"},
	{ID: "2", Name: "Test Item2", Description: "Test Description2", Price: 5.99, ImageID: "abd"},
	{ID: "2", Name: "Test Item3", Description: "Test Description3", Price: 8.99, ImageID: "abd"},
	{ID: "1", Name: "Test Item4", Description: "Test Description4", Price: 9.99, ImageID: "abd"},
	{ID: "1", Name: "Test Item5", Description: "Test Description5", Price: 2.99, ImageID: "abd"},
}

var restaurants = rc.Restaurant{
	// ID     string `json:"id"`
	// Name   string `json:"name"`
	// Owner  string `json:"owner"`
	// Street string `json:"street"`
	// City   string `json:"city"`
	// State  string `json:"state"`
	// Zip    string `json:"zip"`

}

func main() {
	router := gin.Default()
	// Restaurants API endpoints
	router.GET("/API/v1/Restaurants", getAllRestaurants)
	router.GET("/API/v1/Restaurants/:id", getRestaurantById)
	router.POST("/API/v1/Restaurants", saveRestaurant)

	//menu API endpoints
	router.GET("/API/v1/MenuItems", getMenuItems)
	router.GET("/API/v1/MenuItems/:id", getMenuItemsByRestaurantId)
	router.POST("/API/v1/SaveMenuItem", saveMenuItem)
	// TODO change this to environment variable
	router.Run("localhost:8080")

}

func getAllRestaurants(c *gin.Context) {
	//TODO connect to db and get all restaurants
	// TODO create restaurants struct and return a list of it here
	c.IndentedJSON(http.StatusOK, "restaurants")
}

func getRestaurantById(c *gin.Context) {

}

func getRestaurantByIdFromDb(id string) ([]rc.Restaurant, error) {
	// var restaurants rc.Restaurant
	// // TODO get menuItems from DB

	// if restaurant {
	// 	return nil, errors.New("No restaurants found")
	// }
	// return restaurant, nil
}

func saveRestaurant(c *gin.Context) {

}

func getMenuItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, menuItems)
}

func saveMenuItem(c *gin.Context) {
	var newMenuItem rc.MenuItem

	if err := c.BindJSON(&newMenuItem); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Menu item not saved"})
		return
	}
	menuItems = append(menuItems, newMenuItem)
	c.IndentedJSON(http.StatusCreated, menuItems)
}

func getMenuItemsByRestaurantId(c *gin.Context) {
	id := c.Param("id")
	menuItems, err := getMenuItemsFromDB(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No menu found"})
		return
	}
	c.IndentedJSON(http.StatusFound, menuItems)
}

func getMenuItemsFromDB(id string) ([]rc.MenuItem, error) {
	var foundMenuItems []rc.MenuItem
	// TODO get menuItems from DB
	for _, m := range menuItems {
		if m.ID == id {
			foundMenuItems = append(foundMenuItems, m)
		}
	}
	if len(foundMenuItems) == 0 {
		return nil, errors.New("No values found")
	}
	return foundMenuItems, nil
}
