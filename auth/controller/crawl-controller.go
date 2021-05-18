package controller

import (
	"context"
	"net/http"

	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/auth/crawler"

	"github.com/gin-gonic/gin"
)

var crawlControllerInstance *crawlController

type crawlController struct {
}

// CrawController ...
type CrawController interface {
	CrawlCategory(ctx *gin.Context)
}

func (c *crawlController) CrawlCategory(ctx *gin.Context) {
	config.CrawlerClient.CrawCategory(context.Background(), &crawler.Blank{})
	ctx.JSON(http.StatusOK, "success")
}

func GetCrawlerController() CrawController {
	if crawlControllerInstance == nil {
		crawlControllerInstance = &crawlController{}
	}
	return crawlControllerInstance
}
