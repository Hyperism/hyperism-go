package controllers

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"
    "net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/hyperism/hyperism-go/config"
	"github.com/hyperism/hyperism-go/models"
	"github.com/hyperism/hyperism-go/utix"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

    "github.com/golang-jwt/jwt/v4"
)

func GetAllMeta(c *fiber.Ctx) error {
    metaCollection := config.MI.DB.Collection("meta")
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

    var meta []models.Meta

    filter := bson.M{}
    findOptions := options.Find()

    if s := c.Query("s"); s != "" {
        filter = bson.M{
            "$or": []bson.M{
                {
                    "owner": bson.M{
                        "$regex": primitive.Regex{
                            Pattern: s,
                            Options: "i",
                        },
                    },
                },
                {
                    "price": bson.M{
                        "$regex": primitive.Regex{
                            Pattern: s,
                            Options: "i",
                        },
                    },
                },
            },
        }
    }

    page, _ := strconv.Atoi(c.Query("page", "1"))
    limitVal, _ := strconv.Atoi(c.Query("limit", "10"))
    var limit int64 = int64(limitVal)

    total, _ := metaCollection.CountDocuments(ctx, filter)

    findOptions.SetSkip((int64(page) - 1) * limit)
    findOptions.SetLimit(limit)

    cursor, err := metaCollection.Find(ctx, filter, findOptions)
    defer cursor.Close(ctx)

    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "success": false,
            "message": "meta Not found",
            "error":   err,
        })
    }

    for cursor.Next(ctx) {
        var metadata models.Meta
        cursor.Decode(&metadata)
        meta = append(meta, metadata)
    }

    last := math.Ceil(float64(total / limit))
    if last < 1 && total > 0 {
        last = 1
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "data":      meta,
        "total":     total,
        "page":      page,
        "last_page": last,
        "limit":     limit,
    })
}

func GetMeta(c *fiber.Ctx) error {
    metaCollection := config.MI.DB.Collection("meta")
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

    var metadata models.Meta
    objId, err := primitive.ObjectIDFromHex(c.Params("id"))
    findResult := metaCollection.FindOne(ctx, bson.M{"_id": objId})
    if err := findResult.Err(); err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "success": false,
            "message": "meta Not found",
            "error":   err,
        })
    }
   
    err = findResult.Decode(&metadata)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "success": false,
            "message": "meta Not found",
            "error":   err,
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "data":    metadata,
        "success": true,
    })
}

func AddMeta(c *fiber.Ctx) error {
    metaDataCollection := config.MI.DB.Collection("meta")
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    meta := new(models.Meta)

    if err := c.BodyParser(meta); err != nil {
        log.Println(err)
        return c.Status(400).JSON(fiber.Map{
            "success": false,
            "message": "Failed to parse body",
            "error":   err,
        })
    }
    meta.MintDate = time.Now()
    
    result, err := metaDataCollection.InsertOne(ctx, meta)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "success": false,
            "message": "meta failed to insert",
            "error":   err,
        })
    }
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "data":    result,
        "success": true,
        "message": "meta inserted successfully",
    })

}

func UpdateMeta(c *fiber.Ctx) error {
    metaCollection := config.MI.DB.Collection("meta")
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    meta := new(models.Meta)

    // var userinfo models.User
	var err error
	// id := c.Params("id")

	user := c.Locals("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	fmt.Println(user)

	// userinfo, err = GetByID("_id", id)
	if err != nil {
		c.
			Status(http.StatusUnprocessableEntity).
			JSON(utix.NewJError(err))

	}

    if err := c.BodyParser(meta); err != nil {
        log.Println(err)
        return c.Status(400).JSON(fiber.Map{
            "success": false,
            "message": "Failed to parse body",
            "error":   err,
        })
    }

    objId, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "success": false,
            "message": "meta not found",
            "error":   err,
        })
    }
    // if userinfo.ID.Hex() == claims["Id"] && userinfo.ID.Hex() == claims["Issuer"] {
	// 	fmt.Println("both claims match")

	// 	if err != nil {
	// 		fmt.Println(err, " file upload ERRRRR")
	// 		return c.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able upload your attachment"}})

	// 	}
    // }
    update := bson.M{
        "$set": meta,
    }
    _, err = metaCollection.UpdateOne(ctx, bson.M{"_id": objId}, update)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "success": false,
            "message": "meta failed to update",
            "error":   err.Error(),
        })
    }
    
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "success": true,
        "message": "meta updated successfully",
    })
    
}

func DeleteMeta(c *fiber.Ctx) error {
    metaCollection := config.MI.DB.Collection("meta")
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

    objId, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "success": false,
            "message": "meta not found",
            "error":   err,
        })
    }
    _, err = metaCollection.DeleteOne(ctx, bson.M{"_id": objId})
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "success": false,
            "message": "meta failed to delete",
            "error":   err,
        })
    }
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "success": true,
        "message": "meta deleted successfully",
    })
}

func Save(user *models.User) error { //   save to db
    userCollection := config.MI.DB.Collection("user")
	_, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ADDED NEW USER ", user.Email)
	return err
}

func GetByEmail(email string) (models.User, error) { // get by email
	var result models.User
	//var userlogin models.User
    userCollection := config.MI.DB.Collection("user")

	err := userCollection.FindOne(context.Background(), bson.D{{"email", email}}).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(result, " this is the result bb")
	return result, err
}

func GetByKey(key string, value string) (models.User, error) {

	filter := bson.D{{key, value}}
	var res models.User

    userCollection := config.MI.DB.Collection("user")

	err := userCollection.FindOne(context.Background(), filter).Decode(&res)

	return res, err

}

func GetUserDataByKey(key string, value string) (models.User, error) {
	var res models.User
	filter := bson.D{{key, value}}

    userCollection := config.MI.DB.Collection("user")

	result := userCollection.FindOne(context.Background(), filter)
	result.Decode(res)

	err := userCollection.FindOne(context.Background(), filter).Decode(&res)
	//fmt.Println(resultdoc.Userdata, "THIS IS RESultdoc BABY")
	return res, err

}

func GetAll() []models.User {
    userCollection := config.MI.DB.Collection("user")

	cursor, err := userCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Panic(err)
	}
	var docs []models.User

	for cursor.Next(context.Background()) {
		var single models.User
		err := cursor.Decode(&single)
		if err != nil {
			log.Panic(err)
		}
		docs = append(docs, single)
	}

	return docs

}

// func Delete(id string) (*mongo.DeleteResult, error) {

// 	_id, err1 := primitive.ObjectIDFromHex(id)
// 	if err1 != nil {
// 		panic(err1)
// 	}

// 	opts := options.Delete().SetCollation(&options.Collation{})

// 	res, err := userCollection.DeleteOne(context.Background(), bson.D{{"_id", _id}}, opts)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	return res, err
// }

// func Update(key string, value string, updatekey string, userdata string) error {
// 	if key == "_id" {
// 		_id, err1 := primitive.ObjectIDFromHex(value)
// 		if err1 != nil {
// 			utix.CheckErorr(err1)
// 		}
// 		filter := bson.D{{key, _id}}
// 		update := bson.D{{"$set", bson.D{{updatekey, userdata}}}}
// 		_, e := userCollection.UpdateOne(context.Background(), filter, update)
// 		if e != nil {
// 			return e
// 		}
// 		fmt.Println("update sucesss")
// 		fmt.Println("fired from ID sub conditionals")

// 		return nil

// 	}

// 	filter := bson.D{{key, value}}

// 	update := bson.D{{"$set", bson.D{{updatekey, userdata}}}}

// 	_, e := userCollection.UpdateOne(context.Background(), filter, update)

// 	if e != nil {
// 		return e
// 	}
// 	fmt.Println("update sucesss")

// 	return nil

// }

func Updateint(key string, value string, updatekey string, userdata int64) error {
    userCollection := config.MI.DB.Collection("user")
	if key == "_id" {
		_id, err1 := primitive.ObjectIDFromHex(value)
		if err1 != nil {
			utix.CheckErorr(err1)
		}
		filter := bson.D{{key, _id}}
		update := bson.D{{"$set", bson.D{{updatekey, userdata}}}}
		_, e := userCollection.UpdateOne(context.Background(), filter, update)
		if e != nil {
			return e
		}
		fmt.Println("update sucesss")
		fmt.Println("fired from ID sub conditionals")

		return nil

	}

	filter := bson.D{{key, value}}

	update := bson.D{{"$set", bson.D{{updatekey, userdata}}}}

	_, e := userCollection.UpdateOne(context.Background(), filter, update)

	if e != nil {
		return e
	}
	fmt.Println("update sucesss")

	return nil

}

// func Close() error {
    
// 	err := userClient.Disconnect(context.Background())
// 	fmt.Println("db closed")
// 	utix.CheckErorr(err)
// 	return err
// }

func GetByID(key string, value string) (models.User, error) {
    userCollection := config.MI.DB.Collection("user")
	_id, err1 := primitive.ObjectIDFromHex(value)
	utix.CheckErorr(err1)
	filter := bson.D{{key, _id}}
	var res models.User

	err := userCollection.FindOne(context.Background(), filter).Decode(&res)

	return res, err

}