package controllers

import (
	"net/http"
	"strconv"

	"github.com/azura-labs/models"

	repo "github.com/azura-labs/repositories"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func PostTicket(c echo.Context) (err error) {
	var request *models.RequestTicket
	var response *models.ResponseTicket

	//auto-populate json request into struct
	request = &models.RequestTicket{}
	if err = c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.WithField("request", request).Info("post tiket request received")

	response = repo.InsertTicket(request)
	if response.Success == false {
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func PutTicket(c echo.Context) (err error) {
	var request *models.RequestTicket
	var response *models.ResponseTicket

	id := c.Param("ticketId")
	intVar, _ := strconv.Atoi(id)
	//auto-populate json request into struct
	request = &models.RequestTicket{}
	if err = c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.WithField("request", request).Info("put ticket request received")

	response = repo.UpdateTicket(request, intVar)
	if response.Success == false {
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func DeleteTicket(c echo.Context) (err error) {
	var response *models.ResponseTicket

	id := c.Param("ticketId")
	intVar, _ := strconv.Atoi(id)
	log.WithField("request", intVar).Info("delete ticket request received")
	response = repo.DeleteTicket(intVar)
	if response.Success == false {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func DetailTicket(c echo.Context) (err error) {
	var response *models.ResponseTicket

	id := c.Param("ticketId")
	intVar, _ := strconv.Atoi(id)
	log.WithField("request", intVar).Info("get detail ticket request received")
	response = repo.DetailTicket(intVar)
	if response.Success == false {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func ListTicket(c echo.Context) (err error) {
	var response *models.ResponseTickets

	limit := c.QueryParam("limit")
	intLimit, _ := strconv.Atoi(limit)
	skip := c.QueryParam("skip")
	intSkip, _ := strconv.Atoi(skip)
	q := c.QueryParam("q")
	log.WithField("request", intLimit).Info("get list ticket request received")
	response = repo.ListTicket(intLimit, intSkip, q)
	c.JSON(http.StatusOK, response)
	return
}
