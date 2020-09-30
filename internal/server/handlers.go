package server

import (
	"github.com/gin-gonic/gin"
)

type CreateExchangeRate struct {
	Currency1 string `json:"currency1" binding:"required"`
	Currency2 string `json:"currency2" binding:"required"`
}

type ConvertReq struct {
	CurrrencyFrom string  `json:"currencyFrom" binding:"required"`
	CurrrencyTo   string  `json:"currencyTo" binding:"required"`
	Value         float64 `json:"value" binding:"required"`
}

func (s *Server) CreateExchangeRate(c *gin.Context) {
	var input CreateExchangeRate
	c.ShouldBindJSON(&input)

	if err := s.db.CreateExchangeRate(c.Request.Context(), input.Currency1, input.Currency2); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{})
}
func (s *Server) Convert(c *gin.Context) {
	var input ConvertReq
	err := c.BindJSON(&input)
	if err != nil {
		c.Error(err)
	}

	exchangeRate, err := s.db.GetExchangeRate(c, input.CurrrencyFrom, input.CurrrencyTo)
	if err != nil || len(exchangeRate) < 1 {
		c.Error(err)
		c.JSON(400, gin.H{
			"error": "Can't find currency pair",
		})
		return
	}

	sum := exchangeRate[0].Cource * input.Value

	c.JSON(200, gin.H{
		"result": sum,
	})
}
