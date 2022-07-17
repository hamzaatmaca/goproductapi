package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/productapi/database"
	"github.com/productapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = database.GetCollection(database.DB, "product")

func GetProduct(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	results, err := productCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	var employees []model.Product = make([]model.Product, 0)

	// iterate the cursor and decode each item into an Employee
	if err := results.All(c.Context(), &employees); err != nil {
		return c.Status(500).SendString(err.Error())

	}
	// return employees list in JSON format
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    employees,
	})

}

func GetOneProduct(c *fiber.Ctx) error {

	params := c.Params("id")

	_id, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.D{{"_id", _id}}

	var result model.Product

	if err := productCollection.FindOne(c.Context(), filter).Decode(&result); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"data":    err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    result,
	})
}

func AddProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	payload := new(model.Product)
	defer cancel()

	if err := c.BodyParser(payload); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	result, err := productCollection.InsertOne(ctx, payload)

	if err != nil {
		println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"data":    err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    result,
	})

}

func UpdateProduct(c *fiber.Ctx) error {

	id := c.Params("id")
	productID, err := primitive.ObjectIDFromHex(id)

	payload := new(model.Product)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"data":    err.Error(),
		})
	}

	if err := c.BodyParser(payload); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	query := bson.D{{Key: "_id", Value: productID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: payload.Name},
				{Key: "age", Value: payload.Price},
				{Key: "salary", Value: payload.Brand},
			},
		},
	}

	err = productCollection.FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(404)
		}
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// return the updated employee
	payload.ID = id
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    payload,
	})

}

func DeleteProduct(c *fiber.Ctx) error {

	productId, err := primitive.ObjectIDFromHex(
		c.Params("id"),
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"data":    err.Error(),
		})
	}

	query := bson.D{{Key: "_id", Value: productId}}
	result, err := productCollection.DeleteOne(c.Context(), query)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// the employee might not exist
	if result.DeletedCount < 1 {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Product Not Found",
		})
	}

	// the record was deleted
	return c.Status(204).JSON(fiber.Map{
		"success": true,
		"data":    result,
		"message": "Product Deleted",
	})
}
