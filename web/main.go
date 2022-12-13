package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"mofu/ent"
	"mofu/value"
	"net/http"
)

type IWeb interface {
	Auth(ctx context.Context, token string) *string
	CommandProcess(command string, args string, operator string) (value.MessageMakeup, error)
	ListHistory(ctx context.Context, timeBefore *int64, offset, limit int) ([]*ent.History, error)
}

func New(auth IWeb) *gin.Engine {
	g := gin.Default()

	g.Use(func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		user := auth.Auth(ctx, token)
		if user == nil {
			ctx.Status(http.StatusUnauthorized)
		} else {
			ctx.Set("user", *user)
		}
	})

	g.POST("/command/", func(context *gin.Context) {
		process(context, auth)
	})
	g.GET("/", func(context *gin.Context) {
		list(context, auth)
	})

	return g
}

type command struct {
	Command string
	Args    string
}

func process(ctx *gin.Context, functions IWeb) {
	var c command
	err := ctx.ShouldBindJSON(&c)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	msg, err := functions.CommandProcess(c.Command, c.Args, ctx.MustGet("user").(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"content": msg.ToMsg().Content()})
}

type query struct {
	TimeBefore *int64 `form:"time_before"`
	Offset     int    `form:"offset"`
	Limit      int    `form:"limit"`
}

func list(ctx *gin.Context, functions IWeb) {
	var q query
	err := ctx.ShouldBindQuery(&q)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	if q.Limit <= 0 || q.Limit >= 100 {
		q.Limit = 100
	}
	msg, err := functions.ListHistory(ctx, q.TimeBefore, q.Offset, q.Limit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, msg)
}
