package controllers

import (
	"context"
	"fiber-mongo-api/configs"
	"fiber-mongo-api/models"
	"fiber-mongo-api/responses"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var fishCollection *mongo.Collection = configs.GetCollection(configs.DB, "fishes")
var validate = validator.New()

func CreateFish(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var fish models.Fish
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&fish); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.FishResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// ⛔️ GET /fish?id=1
	// ✅ GET /fish/1

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&fish); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.FishResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newFish := models.Fish{
		Id:            primitive.NewObjectID(),
		Name:          fish.Name,
		CookingMethod: fish.CookingMethod,
		OkToEat:       fish.OkToEat,
		MostDelicious: fish.MostDelicious,
	}

	result, err := fishCollection.InsertOne(ctx, newFish)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.FishResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.FishResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetFishes(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var fishes []models.Fish
	defer cancel()

	results, err := fishCollection.Find(ctx, bson.M{})
	fmt.Println(results)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.FishResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var fishToAppend models.Fish
		if err = results.Decode(&fishToAppend); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.FishResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		fishes = append(fishes, fishToAppend)
	}

	return c.Status(http.StatusOK).JSON(
		responses.FishResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": fishes}},
	)
}

func GetAFish(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	fishId := c.Params("fishId")
	var OneFish models.Fish
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(fishId)

	err := fishCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&OneFish)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.FishResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.FishResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": OneFish}})
}
