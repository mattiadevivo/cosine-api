package controllers

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"webapi/utils"

	"github.com/gofiber/fiber/v2"
)

const FILENAME string = "features"

func getFilePath() string {
	p, err := filepath.Abs(filepath.Join("files", FILENAME))
	if err != nil {
		fmt.Println("Error building path:")
		panic(err)
	}
	return p
}

func DoNothing(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// Take base64 string from "features" inside JSON body, decode it and saves it to a file
func Store(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	dst, err := base64.StdEncoding.DecodeString(data["features"])
	if err != nil {
		return err
	}

	// Ensure path exists
	filePath := getFilePath()
	if _, err := os.Stat(filepath.Dir(filePath)); os.IsNotExist(err) {
		if err := os.Mkdir(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

	}
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error while creating " + filePath)
		fmt.Println(err)
		return err
	}
	// Defer are executed in LIFO order
	defer f.Close()
	defer f.Sync()

	_, err = f.Write(dst)
	if err != nil {
		fmt.Println("Error while writing to file")
		fmt.Println(err)
		return err
	}
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// Decode base64 string from "features" inside JSON body, and compare it from previously stored features
func Compare(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	received, err := base64.StdEncoding.DecodeString(data["features"])
	if err != nil {
		return err
	}
	// Ensure path exists
	filePath := getFilePath()
	if _, err := os.Stat(filepath.Dir(filePath)); os.IsNotExist(err) {
		return err
	}
	features, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	cosine := utils.Cosine(features, received)

	return c.JSON(fiber.Map{
		"message": "success",
		"cosine":  cosine,
	})
}
