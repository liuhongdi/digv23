package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv23/global"
	"github.com/liuhongdi/digv23/service"
)

type GoodsController struct{}
func NewGoodsController() GoodsController {
	return GoodsController{}
}
//购买一件商品
func (g *GoodsController) BuyOne(c *gin.Context) {
	result := global.NewResult(c)

    var goodsId int64 = 1
    buyNum :=1
	err := service.BuyOneGoods(goodsId,buyNum);
	if err != nil {
		result.Error(404,"数据查询错误")
	} else {
		result.Success("减库存成功");
	}
	return
}
//购买一件商品,by lock
func (g *GoodsController) LockBuyOne(c *gin.Context) {
	result := global.NewResult(c)

	var goodsId int64 = 1
	buyNum :=1
	err := service.LockBuyOneGoods(goodsId,buyNum);
	if err != nil {
		result.Error(404,"数据查询错误")
	} else {
		result.Success("减库存成功");
	}
	return
}