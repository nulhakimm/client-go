package model

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID   string `json:"id" form:"id"`
	Name string `json:"name,omitempty" form:"name"`
	Age  int    `json:"age,omitempty" form:"age"`
}

type UserModel interface {
	FindAll(ctx context.Context) (*WebResponse, error)
}

type UserModelImpl struct {
}

func NewUserModel() UserModel {
	return &UserModelImpl{}
}

func (model *UserModelImpl) FindAll(ctx context.Context) (*WebResponse, error) {

	url := "https://www.faridlan.nostracode.com/api/users"

	agent := fiber.Get(url)

	_, data, err := agent.Bytes()
	if err != nil {
		return nil, fmt.Errorf("failed to make request : %v", err)
	}

	webResponse := &WebResponse{
		Data: []User{},
	}

	jsonErr := json.Unmarshal(data, &webResponse)
	if jsonErr != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", jsonErr)
	}

	return webResponse, nil

}
