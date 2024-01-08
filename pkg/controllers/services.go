// services.go
package controllers

import (
	"backend-manager/pkg/database"
	"backend-manager/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateService creates a new service in the MongoDB collection
func CreateService(c echo.Context) error {
	// Parse the incoming request body into a Service struct

	// Use a proper context with timeout and cancellation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	var service models.Service
	defer cancel()

	//validate the request body
	if err := c.Bind(&service); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	fmt.Printf("Image Data:\n%s\n", service.Image)

	// Generate a new ObjectID
	newService := models.Service{
		ID:          primitive.NewObjectID(),
		Name:        service.Name,
		Statement:   service.Statement,
		Description: service.Description,
		Input:       service.Input,
		Output:      service.Output,
		Image:       service.Image,
	}

	// Log the request body JSON data
	requestData, err := json.MarshalIndent(newService, "", "  ")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to marshal JSON"})
	}
	fmt.Printf("Request Body JSON Data:\n%s\n", requestData)

	// Log before inserting the service
	log.Printf("Inserting service: %+v", newService)

	// Insert the new service into the MongoDB collection
	servicesCollection := database.GetServiceCollection()
	result, err := servicesCollection.InsertOne(ctx, newService)
	if err != nil {
		log.Printf("Error creating service: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create service", "details": err.Error()})
	}

	// Log after successful insertion
	log.Printf("Service inserted successfully. ID: %v", result.InsertedID)

	// Return the ID of the newly created service
	return c.JSON(http.StatusCreated, map[string]interface{}{"id": result.InsertedID})
}

func GetServices(c echo.Context) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Fetch services from MongoDB with appropriate error handling
	servicesCollection := database.GetServiceCollection()

	cursor, err := servicesCollection.Find(ctx, bson.M{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "No services found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch services"})
	}
	defer cursor.Close(ctx)

	var services []models.Service
	for cursor.Next(ctx) {
		var service models.Service
		if err := cursor.Decode(&service); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to decode service"})
		}
		services = append(services, service)
	}

	return c.JSON(http.StatusOK, services)
}
