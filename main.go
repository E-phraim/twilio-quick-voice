package main

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TwiML struct {
	XMLName xml.Name `xml:"Response"`

	Say  string `xml:",omitempty"`
	Play string `xml:",omitempty"`
}

type content struct {
	City string `json:"city" binding:"required"`
}

func main() {
	engine := gin.Default()

	engine.POST("/voice", func(c *gin.Context) {
		twiml := TwiML{Say: "Welcome to the code cave"}
		x, err := xml.Marshal(twiml)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.Header("Content-Type", "application/xml")
		c.String(http.StatusOK, string(x))
	})

	//and

	engine.POST("/voice", func(ctx *gin.Context) {
		var content content
		if err := ctx.ShouldBindJSON(&content); err != nil {
			ctx.Error(err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		response := TwiML{Say: "Never gonna give you up " + content.City}

		twiml := TwiML{Play: "https://demo.twilio.com/docs/classic.mp3"}
		x, err := xml.MarshalIndent(twiml, "", "  ")
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Header("Content-Type", "application/xml")
		ctx.String(http.StatusOK, string(x))
		ctx.JSON(http.StatusOK, gin.H{"message": response})
	})

	fmt.Println("server started running<><><><>")
	engine.Run(":3000")
}
