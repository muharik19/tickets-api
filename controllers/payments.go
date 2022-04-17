package controllers

import (
	"net/http"
	"strconv"

	"github.com/azura-labs/models"

	repo "github.com/azura-labs/repositories"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func PostPayment(c echo.Context) (err error) {
	var request *models.RequestPayment
	var response *models.ResponsePayment

	//auto-populate json request into struct
	request = &models.RequestPayment{}
	if err = c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.WithField("request", request).Info("post payment request received")

	response = repo.InsertPayment(request)
	if response.Success == false {
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func DetailPayment(c echo.Context) (err error) {
	var response *models.ResponsePayment

	id := c.Param("paymentId")
	intVar, _ := strconv.Atoi(id)
	log.WithField("request", intVar).Info("get detail payment request received")
	response = repo.DetailPayment(intVar)
	if response.Success == false {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func ListPayment(c echo.Context) (err error) {
	var response *models.ResponsePayments

	limit := c.QueryParam("limit")
	intLimit, _ := strconv.Atoi(limit)
	skip := c.QueryParam("skip")
	intSkip, _ := strconv.Atoi(skip)
	q := c.QueryParam("q")
	log.WithField("request", intLimit).Info("get list payment request received")
	response = repo.ListPayment(intLimit, intSkip, q)
	c.JSON(http.StatusOK, response)
	return
}
