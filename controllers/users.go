package controllers

import (
	"net/http"
	"strconv"

	"github.com/azura-labs/models"

	repo "github.com/azura-labs/repositories"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func PostUser(c echo.Context) (err error) {
	var request *models.RequestUser
	var response *models.ResponseUser

	//auto-populate json request into struct
	request = &models.RequestUser{}
	if err = c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.WithField("request", request).Info("post user request received")

	response = repo.InsertUser(request)
	if response.Success == false {
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func PutUser(c echo.Context) (err error) {
	var request *models.RequestUser
	var response *models.ResponseUser

	id := c.Param("userId")
	intVar, _ := strconv.Atoi(id)
	//auto-populate json request into struct
	request = &models.RequestUser{}
	if err = c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.WithField("request", request).Info("put user request received")

	response = repo.UpdateUser(request, intVar)
	if response.Success == false {
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func DeleteUser(c echo.Context) (err error) {
	var response *models.ResponseUser

	id := c.Param("userId")
	intVar, _ := strconv.Atoi(id)
	log.WithField("request", intVar).Info("delete user request received")
	response = repo.DeleteUser(intVar)
	if response.Success == false {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func DetailUser(c echo.Context) (err error) {
	var response *models.ResponseUser

	id := c.Param("userId")
	intVar, _ := strconv.Atoi(id)
	log.WithField("request", intVar).Info("get detail user request received")
	response = repo.DetailUser(intVar)
	if response.Success == false {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func ListUser(c echo.Context) (err error) {
	var response *models.ResponseUsers

	limit := c.QueryParam("limit")
	intLimit, _ := strconv.Atoi(limit)
	skip := c.QueryParam("skip")
	intSkip, _ := strconv.Atoi(skip)
	q := c.QueryParam("q")
	log.WithField("request", intLimit).Info("get list user request received")
	response = repo.ListUser(intLimit, intSkip, q)
	c.JSON(http.StatusOK, response)
	return
}
