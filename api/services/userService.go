package services

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/funcs"
	"alfath_lms/api/interfaces"
	"alfath_lms/api/models"
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	mongo         *mongo.Database
	instructorSvc interfaces.InstructorServiceInterface
	studentSvc    interfaces.StudentServiceInterface
	redis         *redis.Client
}

func (userSvc *UserService) Inject(
	mongo *mongo.Database,
	studentService interfaces.StudentServiceInterface,
	instructorService interfaces.InstructorServiceInterface,
	redis *redis.Client,
) {
	userSvc.mongo = mongo
	userSvc.instructorSvc = instructorService
	userSvc.studentSvc = studentService
	userSvc.redis = redis
}

func (userSvc *UserService) Refresh(Data map[string]interface{}) (definitions.LoginResponse, error) {

	token, err := jwt.Parse(Data["RefreshToken"].(string), func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return "Not authorized!", nil
		}

		return []byte(os.Getenv("JWT_KEY")), nil //Parse function must return a key. remember it's called the "Keyfunc".
	})

	if err != nil {
		return definitions.LoginResponse{}, nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		storedToken, err := userSvc.redis.Get(context.Background(), "refresh_token_"+claims["email"].(string)).Result()

		fmt.Println(storedToken)

		if err != nil {
			return definitions.LoginResponse{
				Status:       500,
				Message:      err.Error(),
				Token:        "",
				RefreshToken: "",
			}, nil
		}

		if storedToken != Data["RefreshToken"].(string) {
			return definitions.LoginResponse{
				Status:       400,
				Message:      "Wrong data supplied",
				Token:        "",
				RefreshToken: "",
			}, nil
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email":     claims["email"],
			"name":      claims["name"],
			"role_name": claims["role_name"],
			"exp":       time.Now().Add(time.Minute * 60).Unix(),
		})

		parsedToken, tokenErr := token.SignedString([]byte(os.Getenv("JWT_KEY")))

		if tokenErr != nil {
			return definitions.LoginResponse{
				Status:       400,
				Message:      tokenErr.Error(),
				Token:        "",
				RefreshToken: "",
			}, nil
		}

		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email":     claims["email"],
			"name":      claims["name"],
			"role_name": claims["role_name"],
			"exp":       time.Now().Add(time.Hour * 24).Unix(),
		})

		parsedRefreshToken, tokenErr := refreshToken.SignedString([]byte(os.Getenv("JWT_KEY")))

		if tokenErr != nil {
			return definitions.LoginResponse{
				Status:       400,
				Message:      tokenErr.Error(),
				Token:        "",
				RefreshToken: "",
			}, nil
		}

		userEmail := claims["email"].(string)
		redisKey := "refresh_token_" + userEmail
		redisErr := userSvc.redis.Set(context.Background(), redisKey, parsedRefreshToken, 24*time.Hour).Err()

		if redisErr != nil {
			return definitions.LoginResponse{
				Status:       400,
				Message:      tokenErr.Error(),
				Token:        "",
				RefreshToken: "",
			}, nil
		}

		return definitions.LoginResponse{
			Status:       200,
			Message:      "Login Success",
			Token:        parsedToken,
			RefreshToken: parsedRefreshToken,
		}, nil
	} else {
		return definitions.LoginResponse{
			Status:       400,
			Message:      "Invalid refresh token",
			Token:        "",
			RefreshToken: "",
		}, nil
	}
}

func (userSvc *UserService) Login(Data map[string]interface{}) (definitions.LoginResponse, error) {
	filter := bson.D{{"email", Data["Email"].(string)}}
	searchResult := userSvc.mongo.Collection("users").FindOne(context.TODO(), filter)
	if searchResult.Err() == mongo.ErrNoDocuments {
		return definitions.LoginResponse{
			Status:       400,
			Message:      "Wrong username or password",
			Token:        "",
			RefreshToken: "",
		}, nil
	} else if searchResult.Err() != nil {
		return definitions.LoginResponse{
			Status:       500,
			Message:      "there's an error in processing your request. Please try again later",
			Token:        "",
			RefreshToken: "",
		}, nil
	} else {
		var existingUser models.User
		searchResult.Decode(&existingUser)
		if existingUser.IsDeleted {
			return definitions.LoginResponse{
				Status:       400,
				Message:      "User doesn't exist",
				Token:        "",
				RefreshToken: "",
			}, nil
		}

		if existingUser.Password != funcs.HashStringToSHA256(Data["Password"].(string)) {
			return definitions.LoginResponse{
				Status:       400,
				Message:      "Wrong username or password",
				Token:        "",
				RefreshToken: "",
			}, nil
		} else {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"email":     Data["Email"],
				"role_name": existingUser.Role.Name,
				"name":      existingUser.Name,
				"exp":       time.Now().Add(time.Minute * 60).Unix(),
			})

			parsedToken, tokenErr := token.SignedString([]byte(os.Getenv("JWT_KEY")))

			if tokenErr != nil {
				return definitions.LoginResponse{
					Status:       400,
					Message:      tokenErr.Error(),
					Token:        "",
					RefreshToken: "",
				}, nil
			}

			refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"email":     Data["Email"],
				"role_name": existingUser.Role.Name,
				"name":      existingUser.Name,
				"exp":       time.Now().Add(time.Hour * 24).Unix(),
			})

			parsedRefreshToken, tokenErr := refreshToken.SignedString([]byte(os.Getenv("JWT_KEY")))

			if tokenErr != nil {
				return definitions.LoginResponse{
					Status:       400,
					Message:      tokenErr.Error(),
					Token:        "",
					RefreshToken: "",
				}, nil
			}

			redisKey := "refresh_token_" + Data["Email"].(string)
			err := userSvc.redis.Set(context.Background(), redisKey, parsedRefreshToken, 24*time.Hour).Err()

			if err != nil {
				return definitions.LoginResponse{
					Status:       400,
					Message:      tokenErr.Error(),
					Token:        "",
					RefreshToken: "",
				}, nil
			}

			return definitions.LoginResponse{
				Status:       200,
				Message:      "Login Success",
				Token:        parsedToken,
				RefreshToken: parsedRefreshToken,
			}, nil
		}
	}
}

func (userSvc *UserService) LoginAdmin(Data map[string]interface{}) (definitions.LoginResponse, error) {
	filter := bson.D{{"email", Data["Email"].(string)}}
	searchResult := userSvc.mongo.Collection("users").FindOne(context.TODO(), filter)
	if searchResult.Err() == mongo.ErrNoDocuments {
		return definitions.LoginResponse{
			Status:       400,
			Message:      "Wrong username or password",
			Token:        "",
			RefreshToken: "",
		}, nil
	} else if searchResult.Err() != nil {
		return definitions.LoginResponse{
			Status:       500,
			Message:      "there's an error in processing your request. Please try again later",
			Token:        "",
			RefreshToken: "",
		}, nil
	} else {
		var existingUser models.User
		searchResult.Decode(&existingUser)
		if existingUser.IsDeleted {
			return definitions.LoginResponse{
				Status:       400,
				Message:      "User doesn't exist",
				Token:        "",
				RefreshToken: "",
			}, nil
		}

		if existingUser.Password != funcs.HashStringToSHA256(Data["Password"].(string)) {
			return definitions.LoginResponse{
				Status:       400,
				Message:      "Wrong username or password",
				Token:        "",
				RefreshToken: "",
			}, nil
		} else {
			if existingUser.Role.Name != "administrator" {
				return definitions.LoginResponse{
					Status:       400,
					Message:      "Please sign in with another user!",
					Token:        "",
					RefreshToken: "",
				}, nil
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"email":     Data["Email"],
				"name":      existingUser.Name,
				"role_name": existingUser.Role.Name,
				"exp":       time.Now().Add(time.Minute * 60).Unix(),
			})

			parsedToken, tokenErr := token.SignedString([]byte(os.Getenv("JWT_KEY")))

			if tokenErr != nil {
				return definitions.LoginResponse{
					Status:       400,
					Message:      tokenErr.Error(),
					Token:        "",
					RefreshToken: "",
				}, nil
			}

			refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"email":     Data["Email"],
				"name":      existingUser.Name,
				"role_name": existingUser.Role.Name,
				"exp":       time.Now().Add(time.Hour * 24).Unix(),
			})

			parsedRefreshToken, tokenErr := refreshToken.SignedString([]byte(os.Getenv("JWT_KEY")))

			if tokenErr != nil {
				return definitions.LoginResponse{
					Status:       400,
					Message:      tokenErr.Error(),
					Token:        "",
					RefreshToken: "",
				}, nil
			}

			redisKey := "refresh_token_" + Data["Email"].(string)
			err := userSvc.redis.Set(context.Background(), redisKey, parsedRefreshToken, 24*time.Hour).Err()

			if err != nil {
				return definitions.LoginResponse{
					Status:       400,
					Message:      tokenErr.Error(),
					Token:        "",
					RefreshToken: "",
				}, nil
			}

			return definitions.LoginResponse{
				Status:       200,
				Message:      "Login Success",
				Token:        parsedToken,
				RefreshToken: parsedRefreshToken,
			}, nil
		}
	}
}

func (userSvc *UserService) Update(Email string, Updates []bson.E) (definitions.GenericAPIMessage, error) {
	filter := bson.D{{"email", Email}}
	searchResult := userSvc.mongo.Collection("users").FindOne(context.TODO(), filter)
	if searchResult.Err() == mongo.ErrNoDocuments {
		return definitions.GenericAPIMessage{
			Status:  400,
			Message: "User not found",
		}, nil
	} else if searchResult.Err() != nil {
		return definitions.GenericAPIMessage{
			Status:  500,
			Message: searchResult.Err().Error(),
		}, nil
	} else {
		_, err := userSvc.mongo.Collection("users").UpdateOne(context.TODO(), filter, bson.D{{"$set", Updates}})

		if err != nil {
			return definitions.GenericAPIMessage{
				Status:  500,
				Message: err.Error(),
			}, nil
		}

		return definitions.GenericAPIMessage{
			Status:  200,
			Message: "User is successfully updated",
		}, nil
	}
}

func (userSvc *UserService) Delete(Email string) (definitions.GenericAPIMessage, error) {
	filter := bson.D{{"email", Email}}
	searchResult := userSvc.mongo.Collection("users").FindOne(context.TODO(), filter)
	if searchResult.Err() == mongo.ErrNoDocuments {
		return definitions.GenericAPIMessage{
			Status:  400,
			Message: "User not found",
		}, nil
	} else if searchResult.Err() != nil {
		return definitions.GenericAPIMessage{
			Status:  500,
			Message: searchResult.Err().Error(),
		}, nil
	} else {
		_, err := userSvc.mongo.Collection("users").UpdateOne(context.TODO(), filter, bson.D{{"$set", bson.D{{"is_deleted", true}}}})

		if err != nil {
			return definitions.GenericAPIMessage{
				Status:  500,
				Message: err.Error(),
			}, nil
		}

		return definitions.GenericAPIMessage{
			Status:  200,
			Message: "User is successfully deleted",
		}, nil
	}
}

func (userSvc *UserService) Create(User models.User, Role string) (definitions.GenericMongoCreationMessage, error) {
	filter := bson.D{{"email", User.Email}}
	searchResult := userSvc.mongo.Collection("users").FindOne(context.TODO(), filter)
	if searchResult.Err() == mongo.ErrNoDocuments {
		filter = bson.D{{"name", Role}}
		roleSearch := userSvc.mongo.Collection("roles").FindOne(context.TODO(), filter)

		if roleSearch.Err() == mongo.ErrNoDocuments {
			return definitions.GenericMongoCreationMessage{}, errors.New("role doesn't exist")
		} else if roleSearch.Err() != nil {
			return definitions.GenericMongoCreationMessage{}, roleSearch.Err()
		}
		var role models.Role
		roleSearch.Decode(&role)
		User.Role = role
		User.IsDeleted = false

		insertResult, err := userSvc.mongo.Collection("users").InsertOne(context.TODO(), User)

		if err != nil {
			return definitions.GenericMongoCreationMessage{}, nil
		}

		if Role == "instructor" {
			instructor := &models.Instructor{
				Name:        User.Name,
				Email:       User.Email,
				MobilePhone: User.MobilePhone,
				CreatedAt:   time.Now(),
			}
			_, err := userSvc.instructorSvc.CreateInstructor(*instructor)

			if err != nil {
				return definitions.GenericMongoCreationMessage{}, nil
			}
		} else if Role == "student" {
			student := &models.Student{
				Name:        User.Name,
				Email:       User.Email,
				MobilePhone: User.MobilePhone,
				CreatedAt:   time.Now(),
			}
			_, err := userSvc.studentSvc.CreateStudent(*student)

			if err != nil {
				return definitions.GenericMongoCreationMessage{}, nil
			}
		}

		return definitions.GenericMongoCreationMessage{
			Status:     200,
			InstanceID: insertResult.InsertedID.(primitive.ObjectID),
		}, nil
	} else if searchResult.Err() != nil {
		return definitions.GenericMongoCreationMessage{}, searchResult.Err()
	} else {
		return definitions.GenericMongoCreationMessage{}, errors.New("user with that email already exists")
	}

}
