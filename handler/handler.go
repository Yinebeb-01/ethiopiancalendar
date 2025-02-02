package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/yinebebt/ethiocal/bahirehasab"
	"github.com/yinebebt/ethiocal/dateconverter"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"errors"
)

func BahireHasab(ctx *gin.Context) {
	yearString, state := ctx.Params.Get("year")
	if !state {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Empty year value",
		})
		return
	}
	year, err := strconv.Atoi(yearString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "not a valid year",
		})
		return
	}

	festival, err := bahirehasab.BahireHasab(year)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": festival,
	})
}

func Ethiopian(ctx *gin.Context) {
	dateString, state := ctx.Params.Get("date")
	if !state {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "empty date",
		})
		return
	}
	var date = strings.Split(dateString, "-")
	if len(date) > 3 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "not a valid date",
		})
		return
	}

	day, _ := strconv.Atoi(date[2])
	month, _ := strconv.Atoi(date[1])
	year, _ := strconv.Atoi(date[0])
	EtDate, err := dateconverter.Ethiopian(year, month, day)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"ethiopian_date": EtDate.Format("2006-01-02"),
	})
}

func Gregorian(ctx *gin.Context) {
	dateString, state := ctx.Params.Get("date")
	if !state {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "empty date",
		})
		return
	}
	var date = strings.Split(dateString, "-")
	if len(date) > 3 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "not a valid date",
		})
		return
	}

	day, _ := strconv.Atoi(date[2])
	month, _ := strconv.Atoi(date[1])
	year, _ := strconv.Atoi(date[0])

	resDate, err := dateconverter.Gregorian(year, month, day)
	if err != nil {
		if err.Error() == "not a valid date" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": errors.New("internal server error"),
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"gregorian_date": resDate.Format("2006-01-02"),
	})
}

func Init() {
	handler := gin.Default()
	v0 := handler.Group("/v0")
	{
		v0.GET("/etog/:date", Gregorian)
		v0.GET("/gtoe/:date", Ethiopian)
		v0.GET("/bahir/:year", BahireHasab)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	log.Println("shutting down gracefully, received signal:", sig)

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}
}
