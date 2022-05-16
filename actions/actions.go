package actions

import (
	_"fmt"
	_"github.com/hyperism/hyperism-go/controllers"
	_"github.com/hyperism/hyperism-go/models"
	_"github.com/hyperism/hyperism-go/utix"
	_"net/http"
	_"strconv"

	_"github.com/gofiber/fiber/v2"
	_"github.com/golang-jwt/jwt/v4"
)

// func Updateuserdata(c *fiber.Ctx) error {

// 	var userinfo models.User
// 	var err error
// 	id := c.Params("id")

// 	user := c.Locals("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	fmt.Println(user)

// 	userinfo, err = controllers.GetByID("_id", id)
// 	if err != nil {
// 		c.
// 			Status(http.StatusUnprocessableEntity).
// 			JSON(utix.NewJError(err))

// 	}
// 	if userinfo.ID.Hex() == claims["Id"] && userinfo.ID.Hex() == claims["Issuer"] {
// 		fmt.Println("both claims match")

// 		if err != nil {
// 			fmt.Println(err, " file upload ERRRRR")
// 			return c.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able upload your attachment"}})

// 		}
// 		username := c.FormValue("username")
// 		age := c.FormValue("age")
// 		ageInt, err := strconv.Atoi(age)
// 		utix.CheckErorr(err)
// 		bio := c.FormValue("bio")
// 		phone := c.FormValue("phone")
// 		gender := c.FormValue("gender")

// 		var userdata models.User
// 		userdata.Username = username
// 		userdata.Age = int64(ageInt)
// 		userdata.Bio = bio
// 		userdata.Phone = phone
// 		userdata.Gender = gender
// 		fmt.Println(userdata)

// 		err = controllers.Update("_id", id, "username", userdata.Username)
// 		if err != nil {
// 			fmt.Println(err, " username err")
// 			return c.Status(422).JSON(fiber.Map{"errors": err})

// 		}
// 		err = controllers.Updateint("_id", id, "age", userdata.Age)
// 		if err != nil {
// 			fmt.Println(err, " username err")
// 			return c.Status(422).JSON(fiber.Map{"errors": err})

// 		}
// 		err = controllers.Update("_id", id, "phone", userdata.Phone)
// 		if err != nil {
// 			fmt.Println(err, " username err")
// 			return c.Status(422).JSON(fiber.Map{"errors": err})

// 		}
// 		err = controllers.Update("_id", id, "bio", userdata.Bio)
// 		if err != nil {
// 			fmt.Println(err, " username err")
// 			return c.Status(422).JSON(fiber.Map{"errors": err})

// 		}
// 		err = controllers.Update("_id", id, "gender", userdata.Gender)
// 		if err != nil {
// 			fmt.Println(err, " username err")
// 			return c.Status(422).JSON(fiber.Map{"errors": err})

// 		}

// 		if err != nil {
// 			fmt.Println(err, " username err")
// 			return c.Status(422).JSON(fiber.Map{"errors": err})

// 		}

// 		return c.
// 			Status(http.StatusOK).
// 			JSON(fiber.Map{"message": "file upload success", "userdata": userdata})
// 	} else {

// 		return c.
// 			Status(http.StatusBadGateway).
// 			SendString("stop tryna access other profile")

// 	}

// 	//return c.SendString("UNAUTHORIZED")

// }

// func Uploadprofilepic(c *fiber.Ctx) error {

// 	var userinfo models.User
// 	var err error
// 	id := c.Params("id")

// 	user := c.Locals("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	//fmt.Println(user)

// 	userinfo, err = controllers.GetByID("_id", id)
// 	if err != nil {
// 		c.
// 			Status(http.StatusUnprocessableEntity).
// 			JSON(utix.NewJError(err))

// 	}
// 	if userinfo.ID.Hex() == claims["Id"] && userinfo.ID.Hex() == claims["Issuer"] {
// 		fmt.Println(" they matched you got it")

// 		// upload
// 		file, err := c.FormFile("attachment")
// 		if err != nil {
// 			c.
// 				Status(http.StatusUnprocessableEntity).
// 				JSON(utix.NewJError(err))

// 		}

// 		adduploadno := userinfo.Uploadsno
// 		adduploadno++

// 		profilepiclink := fmt.Sprintf("%s%s%duserprofilepicture.jpeg", id, "upno", adduploadno)

// 		c.SaveFile(file, fmt.Sprintf("./uploads/%s", profilepiclink))

// 		if err != nil {
// 			c.
// 				Status(http.StatusUnprocessableEntity).
// 				JSON(utix.NewJError(err))

// 		}

// 		//fmt.Println(adduploadno, "this is adduplaod")

// 		controllers.Updateint("_id", id, "uploadsno", adduploadno)
// 		controllers.Update("_id", id, "profilepicturelink", profilepiclink)

// 		if err != nil {
// 			c.
// 				Status(http.StatusUnprocessableEntity).
// 				JSON(utix.NewJError(err))

// 		}
// 		return c.Status(http.StatusAccepted).
// 			JSON(fiber.Map{"message": "file upload success", "userdata": userinfo, "profilepiclink": profilepiclink})
// 	} else {

// 		return c.
// 			Status(http.StatusBadGateway).
// 			SendString("stop tryna access other profile")

// 	}

// }
