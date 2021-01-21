package service

import (
    "github.com/liuhongdi/digv23/dao"
    "github.com/liuhongdi/digv23/global"
    //goredislib "github.com/go-redis/redis/v8"
    "github.com/go-redsync/redsync/v4"
    "github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

//购买一件商品
func BuyOneGoods(goodsId int64,buyNum int) error {
    return	dao.DecreaseOneGoodsStock(goodsId,buyNum);
}

//购买一件商品,by lock
func LockBuyOneGoods(goodsId int64,buyNum int) error {
   /*
    client := goredislib.NewClient(&goredislib.Options{
        Addr: "localhost:6379",
    })
    */
    pool := goredis.NewPool(global.RedisDb) // or, pool := redigo.NewPool(...)

    // Create an instance of redisync to be used to obtain a mutual exclusion
    // lock.
    rs := redsync.New(pool)

    // Obtain a new mutex by using the same name for all instances wanting the
    // same lock.
    mutexname := "my-global-mutex"
    mutex := rs.NewMutex(mutexname)

    // Obtain a lock for our given mutex. After this is successful, no one else
    // can obtain the same lock (the same mutex name) until we unlock it.
    if err := mutex.Lock(); err != nil {
        //panic(err)
        return err
    }

    // Do your work that requires the lock.
    errdecre := dao.DecreaseOneGoodsStock(goodsId,buyNum);
    if (errdecre!=nil){
        return errdecre
    }
    // Release the lock so other processes or threads can obtain a lock.
    if ok, err := mutex.Unlock(); !ok || err != nil {
        //panic("unlock failed")
        return err
    }

    return nil
}