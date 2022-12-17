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
	UpdateHistoryFlag(id string, sent, nsfw int, operator string) (value.IMessage, error)
	ListHistory(ctx context.Context, timeBefore *int64, offset, limit int) ([]*ent.History, error)
	QueryAuthor(ctx context.Context, name string) ([]*ent.Author, error)
	QueryAuthorMedia(ctx context.Context, id string, limit, offset int) ([]*ent.History, error)
}

func New(auth IWeb) *gin.Engine {
	g := gin.Default()

	g.Use(func(ctx *gin.Context) {
		token := ctx.Query("token")
		user := auth.Auth(ctx, token)
		if user == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		} else {
			ctx.Set("user", *user)
		}
	})

	g.GET("/command/", func(context *gin.Context) {
		process(context, auth)
	})
	g.GET("/", func(context *gin.Context) {
		list(context, auth)
	})
	g.GET("/author/", func(context *gin.Context) {
		author(context, auth)
	})
	g.GET("/author/:user/", func(context *gin.Context) {
		authorMedia(context, auth)
	})

	return g
}

type command struct {
	ID      string `form:"id"`
	Approve bool   `form:"approve"`
	NSFW    bool   `form:"nsfw"`
}

func process(ctx *gin.Context, functions IWeb) {
	var c command
	err := ctx.ShouldBindQuery(&c)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	base := value.Controlled
	nsfw := value.No
	if c.Approve {
		base = base | value.Decided | value.Approved
	}
	if c.NSFW {
		nsfw = value.Nsfw
	}
	msg, err := functions.UpdateHistoryFlag(c.ID, base, nsfw, ctx.MustGet("user").(string))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"content": msg.Content()})
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
	println(*q.TimeBefore)
	msg, err := functions.ListHistory(ctx, q.TimeBefore, q.Offset, q.Limit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, msg)
}

func author(ctx *gin.Context, functions IWeb) {
	name := ctx.Param("username")

	author, err := functions.QueryAuthor(ctx, name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, author)
}

func authorMedia(ctx *gin.Context, functions IWeb) {
	name := ctx.Query("user")
	var q query
	err := ctx.ShouldBindQuery(&q)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	if q.Limit <= 0 || q.Limit >= 100 {
		q.Limit = 100
	}
	media, err := functions.QueryAuthorMedia(ctx, name, q.Limit, q.Offset)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, media)
}
