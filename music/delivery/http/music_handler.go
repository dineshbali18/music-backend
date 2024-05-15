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

	name := c.FormValue("name")
	image := c.FormValue("image")

	file, err := c.FormFile("file")

	if err != nil {
		return err
	}

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

	music.Name = name
	music.File = fileContent
	music.Image = image
	music.CreatedAt = time.Now()

	if err := d.musicUsecase.AddMusic(&music); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, music)
}

func (d *delivery) updateMusic(c echo.Context) error {
	id := c.Param("id")

	musicID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid music ID"})
	}

	var music domain.Music
	if err := c.Bind(&music); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	name := c.FormValue("name")
	image := c.FormValue("image")

	file, err := c.FormFile("file")
	if err != nil {
		if err != http.ErrMissingFile {
			return err
		}
	} else {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		fileContent, err := ioutil.ReadAll(src)
		if err != nil {
			return err
		}

		music.File = fileContent
	}

	music.ID = musicID
	music.Name = name
	music.Image = image
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
