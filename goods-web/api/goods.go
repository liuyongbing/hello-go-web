package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/liuyongbing/hello-go-web/goods-web/global"
	"github.com/liuyongbing/hello-go-web/goods-web/global/response"
	"github.com/liuyongbing/hello-go-web/goods-web/models"
	"github.com/liuyongbing/hello-go-web/goods-web/proto"
)

/*
HandleGrpcErrorToHttp
将 grpc 的 code 转换成 http 的状态码
*/
func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "服务不可用",
				})
			case codes.AlreadyExists:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "用户已存在",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误",
				})
			}

			return
		}
	}
}

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

/*
HandleValidatorError
表单难错误处理
*/
func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}

/*
FilterServices
发现服务(筛选)
*/
func FilterServices() (host string, port int) {
	consulInfo := global.ServerConfig.ConsulInfo
	cfg := api.DefaultConfig()
	// cfg.Address = "127.0.0.1:8500"
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// filter := `Service == "goods-web"`
	filter := fmt.Sprintf(`Service == "%s"`, global.ServerConfig.GoodsSrvInfo.Name)
	data, err := client.Agent().ServicesWithFilter(filter)
	if err != nil {
		panic(err)
	}

	userSrvHost := ""
	userSrvPort := 0

	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}

	return userSrvHost, userSrvPort
}

/*
GetGoodsList
*/
func GetGoodsList(ctx *gin.Context) {
	// 跨域的问题 - 后端解决 也可以前端来解决
	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户: %d", currentUser.ID)

	ip := global.ServerConfig.GoodsSrvInfo.Host
	port := global.ServerConfig.GoodsSrvInfo.Port

	// 拨号连接 user grpc 服务
	clientConn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[GetGoodsList] 连接 [用户服务] 失败",
			"msg", err.Error(),
		)
		HandleGrpcErrorToHttp(err, ctx)
	}

	// 生成 grpc 的 client 并调用接口
	goodsSrvClient := proto.NewGoodsClient(clientConn)

	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)

	rsp, err := goodsSrvClient.GoodsList(context.Background(), &proto.GoodsFilterRequest{
		// PriceMin:    0,
		// PriceMax:    0,
		// IsHot:       false,
		// IsNew:       false,
		// IsTab:       false,
		// TopCategory: 0,
		Pages:       int32(pnInt),
		PagePerNums: int32(pSizeInt),
		// KeyWords:    "",
		// Brand:       0,
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] 调用 [goodsSrvClient.GoodsList] 失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		user := response.GoodsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		result = append(result, user)
	}

	ctx.JSON(http.StatusOK, result)

	zap.S().Infow("获取用户列表")
}
