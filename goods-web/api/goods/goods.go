package goods

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/liuyongbing/hello-go-web/goods-web/api"
	"github.com/liuyongbing/hello-go-web/goods-web/global"
	"github.com/liuyongbing/hello-go-web/goods-web/proto"
)

func Pong(ctx *gin.Context) {
	// TODO: gRPC of Pong
	r, err := global.GoodsSrvClient.SayHello(context.Background(), &proto.HelloRequest{
		Name: "Web request of pong",
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":      "Pong",
		"grpc_msg": r.Message,
		"time":     time.Now(),
	})
}

// List
func List(ctx *gin.Context) {
	fmt.Println("商品列表")
	// 商品的列表 pmin=abc, spring cloud, go-micro
	request := &proto.GoodsFilterRequest{}

	priceMin := ctx.DefaultQuery("pmin", "0")
	priceMinInt, _ := strconv.Atoi(priceMin)
	request.PriceMin = int32(priceMinInt)

	priceMax := ctx.DefaultQuery("pmax", "0")
	priceMaxInt, _ := strconv.Atoi(priceMax)
	request.PriceMax = int32(priceMaxInt)

	isHot := ctx.DefaultQuery("ih", "0")
	if isHot == "1" {
		request.IsHot = true
	}
	isNew := ctx.DefaultQuery("in", "0")
	if isNew == "1" {
		request.IsNew = true
	}

	isTab := ctx.DefaultQuery("it", "0")
	if isTab == "1" {
		request.IsTab = true
	}

	categoryId := ctx.DefaultQuery("c", "0")
	categoryIdInt, _ := strconv.Atoi(categoryId)
	request.TopCategory = int32(categoryIdInt)

	pages := ctx.DefaultQuery("p", "0")
	pagesInt, _ := strconv.Atoi(pages)
	request.Pages = int32(pagesInt)

	perNums := ctx.DefaultQuery("pnum", "0")
	perNumsInt, _ := strconv.Atoi(perNums)
	request.PagePerNums = int32(perNumsInt)

	keywords := ctx.DefaultQuery("q", "")
	request.KeyWords = keywords

	brandId := ctx.DefaultQuery("b", "0")
	brandIdInt, _ := strconv.Atoi(brandId)
	request.Brand = int32(brandIdInt)

	//请求商品的service服务、负载均衡
	//parent, _ := ctx.Get("parentSpan")
	//opentracing.ContextWithSpan(context.Background(), parent.(opentracing.Span))

	// e, b := sentinel.Entry("goods-list", sentinel.WithTrafficType(base.Inbound))
	// if b != nil {
	// 	ctx.JSON(http.StatusTooManyRequests, gin.H{
	// 		"msg": "请求过于频繁，请稍后重试",
	// 	})
	// 	return
	// }
	// r, err := global.GoodsSrvClient.GoodsList(context.WithValue(context.Background(), "ginContext", ctx), request)
	r, err := global.GoodsSrvClient.GoodsList(context.Background(), request)
	if err != nil {
		zap.S().Errorw("[goods.List] 查询 [gRPC.GoodsList] 失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	// e.Exit()
	reMap := map[string]interface{}{
		"total": r.Total,
	}

	goodsList := make([]interface{}, 0)
	for _, value := range r.Data {
		goodsList = append(goodsList, map[string]interface{}{
			"id":          value.Id,
			"name":        value.Name,
			"goods_brief": value.GoodsBrief,
			"desc":        value.GoodsDesc,
			"ship_free":   value.ShipFree,
			"images":      value.Images,
			"desc_images": value.DescImages,
			"front_image": value.GoodsFrontImage,
			"shop_price":  value.ShopPrice,
			"category": map[string]interface{}{
				"id":   value.Category.Id,
				"name": value.Category.Name,
			},
			"brand": map[string]interface{}{
				"id":   value.Brand.Id,
				"name": value.Brand.Name,
				"logo": value.Brand.Logo,
			},
			"is_hot":  value.IsHot,
			"is_new":  value.IsNew,
			"on_sale": value.OnSale,
		})
	}
	reMap["data"] = goodsList

	ctx.JSON(http.StatusOK, reMap)
}
