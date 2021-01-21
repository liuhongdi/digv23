package service

import (
    "github.com/liuhongdi/digv23/dao"
    "github.com/liuhongdi/digv23/global"
    "strconv"
    "github.com/go-redsync/redsync/v4"
    "github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

//购买一件商品
func BuyOneGoods(goodsId int64,buyNum int) error {
    return	dao.DecreaseOneGoodsStock(goodsId,buyNum);
}

//购买一件商品,by lock
func LockBuyOneGoods(goodsId int64,buyNum int) error {

    //fmt.Println("begin LockBuyOneGoods")

    pool := goredis.NewPool(global.RedisDb) // or, pool := redigo.NewPool(...)
    // Create an instance of redisync to be used to obtain a mutual exclusion
    // lock.
    rs := redsync.New(pool)
    // Obtain a new mutex by using the same name for all instances wanting the
    // same lock.
    mutexname := "goods_"+strconv.FormatInt(goodsId,10)
    mutex := rs.NewMutex(mutexname)
    // Obtain a lock for our given mutex. After this is successful, no one else
    // can obtain the same lock (the same mutex name) until we unlock it.
    if err := mutex.Lock(); err != nil {
        //panic(err)
        //fmt.Println("get lock error:")
        //fmt.Println(err)
        return err
    } else {
        //fmt.Println("get lock success:")
    }
    // Do your work that requires the lock.
    errdecre := dao.DecreaseOneGoodsStock(goodsId,buyNum);
    //fmt.Println(errdecre)

    // Release the lock so other processes or threads can obtain a lock.
    if ok, err := mutex.Unlock(); !ok || err != nil {
        //panic("unlock failed")
        //fmt.Println("unlock failed:")
        //fmt.Println(err)
        return err
    }

    if (errdecre!=nil){
        //fmt.Println("decrease failed:")
        //fmt.Println(errdecre)
        return errdecre
    }

    return nil
}