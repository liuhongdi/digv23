package dao

import (
	"errors"
	"fmt"
	"github.com/liuhongdi/digv23/global"
	"github.com/liuhongdi/digv23/model"
	"gorm.io/gorm"
)

//decrease stock
func DecreaseOneGoodsStock(goodsId int64,buyNum int) error {
	fmt.Println("DecreaseOneGoodsStock begin")
    //查询商品信息
	goodsOne:=&model.Goods{}
	err := global.DBLink.Where("goodsId=?",goodsId).First(&goodsOne).Error
	//fmt.Println(goodsOne)
	if (err != nil) {
		return err
	}
	//得到库存
	stock := goodsOne.Stock
	fmt.Println("当前库存:",stock)
    //fmt.Println(stock)
	if (stock < buyNum || stock <= 0) {
		return errors.New("库存不足")
	}

	//减库存
	result := global.DBLink.Debug().Table("goods").Where("goodsId = ? ", goodsId,buyNum).Update("stock", gorm.Expr("stock - ?", buyNum))
	if (result.Error != nil) {
		return result.Error
	} else {
		fmt.Println("成功减库存一次")
		return nil
	}
}