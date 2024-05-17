package http

import (
	"fmt"
	"io/ioutil"
	"musicApp/domain"
	"net/http"
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

	// musicID, err := strconv.ParseUint(id, 10, 64)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid music ID"})
	// }

	var music domain.Music
	if err := c.Bind(&music); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	fmt.Println("AAAAAAAAAAAAAA1")

	existingMusic, err := d.musicUsecase.GetMusicById(id)
	if err != nil {
		return err
	}
	fmt.Println("AAAAAAAAAAAAAA2")
	name := c.FormValue("name")
	image := c.FormValue("image")

	file, err := c.FormFile("file")
	if err != nil {
		// responding with a missing file
		if err != http.ErrMissingFile {
			music.File = existingMusic[0].File
			// return err
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

	music.ID = existingMusic[0].ID
	music.Name = name
	music.Image = image
	if len(name) == 0 {
		music.Name = existingMusic[0].Name
	}
	if len(image) == 0 {
		music.Image = existingMusic[0].Image
	}

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
