package http

import (
	"fmt"
	"io/ioutil"
	"musicApp/domain"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type delivery struct {
	musicUsecase domain.MusicUsecase
}

func NewMusicHandler(e *echo.Echo, usecase domain.MusicUsecase) {
	handler := &delivery{musicUsecase: usecase}

	// Define routes for music endpoints
	e.GET("/v1/get/music", handler.getAllMusic)
	e.POST("/v1/add/music", handler.addMusic)
	e.PUT("/v1/update/music/:id", handler.updateMusic)
	e.DELETE("/v1/delete/music/:id", handler.deleteMusic)
}

func (d *delivery) getAllMusic(c echo.Context) error {
	musicList, err := d.musicUsecase.GetAllMusic()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, musicList)
}

func (d *delivery) addMusic(c echo.Context) error {
	var music domain.Music
	// fmt.Println("aaa")
	// Read form fields
	name := c.FormValue("name")
	image := c.FormValue("image")
	// fmt.Println("name:", name)
	// Get file
	file, err := c.FormFile("file")
	// fmt.Println("file", file)
	if err != nil {
		return err
	}
	// fmt.Println("ccc")
	// Open file
	src, err := file.Open()
	if err != nil {
		return err
	}
	fmt.Println("ddd")
	defer src.Close()

	// Create a buffer to store the file content
	fileContent, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	fmt.Println("FIle", fileContent)

	// Assign values to the Music object
	music.Name = name
	music.File = fileContent
	music.Image = image
	music.CreatedAt = time.Now()

	// fmt.Println("ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ")

	// Add music to the database
	if err := d.musicUsecase.AddMusic(&music); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, music)
}

func (d *delivery) updateMusic(c echo.Context) error {
	id := c.Param("id")

	// Get the music ID from the URL parameter
	musicID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid music ID"})
	}

	// Get the music object from the form data
	var music domain.Music
	if err := c.Bind(&music); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	name := c.FormValue("name")
	music.Name = name
	// Check if a new file is uploaded
	file, err := c.FormFile("file")
	if err != nil {
		if err != http.ErrMissingFile {
			return err
		}
	} else {
		// Open the file
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Read the file content
		fileContent, err := ioutil.ReadAll(src)
		if err != nil {
			return err
		}

		// Update the file content if a new file is uploaded
		music.File = fileContent
	}

	// Set the music ID
	music.ID = musicID
	// music.Name=
	music.CreatedAt = time.Now()

	// Update the music in the database
	if err := d.musicUsecase.UpdateMusic(&music); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, music)
}

func (d *delivery) deleteMusic(c echo.Context) error {
	id := c.Param("id")
	if err := d.musicUsecase.DeleteMusic(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
