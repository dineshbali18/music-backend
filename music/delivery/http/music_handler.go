package http

import (
	"fmt"
	"musicApp/domain"
	"net/http"
	"strconv"

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
	if err := c.Bind(&music); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := d.musicUsecase.AddMusic(&music); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, music)
}

func (d *delivery) updateMusic(c echo.Context) error {
	id := c.Param("id")
	var music domain.Music
	if err := c.Bind(&music); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var err error
	music.ID, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Error in the music Id while updating")
	}
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
