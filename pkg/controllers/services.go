// services.go
package controllers

import (
	"backend-manager/pkg/database"
	"backend-manager/pkg/models"
	"backend-manager/pkg/response"
	"backend-manager/pkg/utils"
	"context"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

// CreateService creates a new service in the MongoDB collection
func CreateService(c echo.Context) error {
	lang := utils.GetLanguage(c)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	var service models.Service
	defer cancel()

	// Validate the request body
	if err := c.Bind(&service); err != nil {
		return response.BadRequestResponse.WriteToResponse(c, nil, lang)
	}

	// Use the validator library to validate required fields
	if validationErr := validate.Struct(&service); validationErr != nil {
		return response.ValidationErrorResponse.WriteToResponse(c, &echo.Map{"data": validationErr.Error()}, lang)
	}

	// fmt.Printf("Image Data:\n%s\n", service.Image)

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

	// Log before inserting the service
	log.Printf("Inserting service: %+v", newService)

	// Insert the new service into the MongoDB collection
	servicesCollection := database.GetServiceCollection()
	result, err := servicesCollection.InsertOne(ctx, newService)
	if err != nil {
		log.Printf("Error creating service: %v", err)
		return response.InternalServerErrorResponse.WriteToResponse(c, nil, lang)
	}

	// Log after successful insertion
	log.Printf("Service inserted successfully. ID: %v", result.InsertedID)

	return response.ServiceCreateResponse.WriteToResponse(c, &echo.Map{"service": newService}, lang)
}

// GetServices fetches services from MongoDB
func GetServices(c echo.Context) error {
	lang := utils.GetLanguage(c)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Fetch services from MongoDB with appropriate error handling
	servicesCollection := database.GetServiceCollection()

	cursor, err := servicesCollection.Find(ctx, bson.M{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return response.NotFoundResponse.WriteToResponse(c, nil, lang)
		}
		return response.InternalServerErrorResponse.WriteToResponse(c, nil, lang)
	}
	defer cursor.Close(ctx)

	var services []models.Service
	for cursor.Next(ctx) {
		var service models.Service
		if err := cursor.Decode(&service); err != nil {
			return response.InternalServerErrorResponse.WriteToResponse(c, nil, lang)
		}
		services = append(services, service)
	}
	return response.GetServicesResponse.WriteToResponse(c, &echo.Map{"services": services}, lang)
}

func GetServiceByName(c echo.Context) error {
	lang := utils.GetLanguage(c)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var service models.Service
	serviceName := c.Param("name")

	log.Printf("Service Name: %s", serviceName)

	// Fetch the service from MongoDB based on the provided name
	servicesCollection := database.GetServiceCollection()
	filter := bson.M{"name": serviceName}
	err := servicesCollection.FindOne(ctx, filter).Decode(&service)

	// userId := c.Param("userId")
	// objId, _ := primitive.ObjectIDFromHex(userId)
	// err := servicesCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&service)

	if err != nil {
		log.Printf("Error fetching service: %v", err)
		if err == mongo.ErrNoDocuments {
			return response.NotFoundResponse.WriteToResponse(c, nil, lang)
		}
		return response.InternalServerErrorResponse.WriteToResponse(c, nil, lang)
	}
	return response.GetServiceByNameResponse.WriteToResponse(c, &echo.Map{"service": service}, lang)
}

// UpdateService updates an existing service in the MongoDB collection
func UpdateService(c echo.Context) error {
	lang := utils.GetLanguage(c)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var updatedService models.Service
	serviceName := c.Param("name")

	// Fetch the existing service from MongoDB based on the provided ID
	servicesCollection := database.GetServiceCollection()
	filter := bson.M{"name": serviceName}
	err := servicesCollection.FindOne(ctx, filter).Decode(&updatedService)
	if err != nil {
		log.Printf("Error fetching service: %v", err)
		if err == mongo.ErrNoDocuments {
			return response.NotFoundResponse.WriteToResponse(c, nil, lang)
		}
		return response.InternalServerErrorResponse.WriteToResponse(c, nil, lang)
	}

	// Update the service fields based on the request body
	if err := c.Bind(&updatedService); err != nil {
		return response.BadRequestResponse.WriteToResponse(c, nil, lang)
	}

	if validationErr := validate.Struct(&updatedService); validationErr != nil {
		return response.ValidationErrorResponse.WriteToResponse(c, &echo.Map{"data": validationErr.Error()}, lang)
	}

	// Update the service in the MongoDB collection
	update := bson.M{
		"$set": bson.M{
			"name":        updatedService.Name,
			"statement":   updatedService.Statement,
			"description": updatedService.Description,
			"input":       updatedService.Input,
			"output":      updatedService.Output,
			"image":       updatedService.Image,
		},
	}

	_, err = servicesCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return response.InternalServerErrorResponse.WriteToResponse(c, nil, lang)
	}

	return response.UpdateServiceResponse.WriteToResponse(c, &echo.Map{"service": updatedService}, lang)
}

// DeleteService deletes an existing service from the MongoDB collection
func DeleteService(c echo.Context) error {
	lang := utils.GetLanguage(c)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Extract service name from the URL parameter
	serviceName := c.Param("name")

	// Fetch the service from MongoDB based on the provided name
	servicesCollection := database.GetServiceCollection()
	filter := bson.M{"name": serviceName}

	// Check if the service exists
	var existingService models.Service
	err := servicesCollection.FindOne(ctx, filter).Decode(&existingService)

	if err != nil {
		log.Printf("Error fetching service: %v", err)
		if err == mongo.ErrNoDocuments {
			return response.NotFoundResponse.WriteToResponse(c, nil, lang)
		}
		return response.InternalServerErrorResponse.WriteToResponse(c, nil, lang)
	}

	// Perform the delete operation in MongoDB
	result, err := servicesCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("Error deleting service: %v", err)
		return response.InternalServerErrorResponse.WriteToResponse(c, &echo.Map{"error": err.Error()}, lang)
	}

	// Check if any documents were deleted
	if result.DeletedCount == 0 {
		return response.NotFoundResponse.WriteToResponse(c, nil, lang)
	}
	return response.DeleteServiceResponse.WriteToResponse(c, &echo.Map{"service": existingService}, lang)
}
