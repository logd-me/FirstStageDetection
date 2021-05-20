/*
@Title : transfer
@Description :
@Author : 谭靖渝
@Update : 2021/5/12 15:52
*/
package RedisServe

import (
	"ZhiHu/Model"
	"ZhiHu/cache"
	"github.com/garyburd/redigo/redis"
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
)
//发表文章的时候就要有redis的初始化
var (
	wg    sync.WaitGroup
	m     sync.Mutex
)

func Transfer(i int,store[]string) {
	cache.Conn.Do("SELECT", 1)
	ch := make(chan int)
	r2, _ := redis.Values(cache.Conn.Do("keys", "*"))
	for _, v := range r2 {
		store = append(store, string(v.([]uint8)))
	}
	if len(store)==0{
		logrus.Info("没有文章数据")
	}else {
		go CreateInx(ch,store)
		for  {
			judge := <-ch
			if judge==len(store)-1 {
				wg.Add(1)
				go Get(judge,store)
				break
			}
			wg.Add(1)
			go Get(judge,store)
		}
		wg.Wait()
	}
}


func CreateInx(ch chan int,store[]string) {
	for i := 0; i < len(store); i++ {
		ch <- i
	}
	close(ch)
}

func Get(index int,store[]string)  {
	defer wg.Done()
	m.Lock()
	cache.Conn.Do("SELECT", 1)
	r,_:=redis.Values(cache.Conn.Do("HMGet",store[index],"view","like","u_like"))
	view, _ := strconv.Atoi(string(r[0].([]uint8)))
	like, _ := strconv.Atoi(string(r[1].([]uint8)))
	ulike, _ := strconv.Atoi(string(r[2].([]uint8)))
	Model.DB.Model(&Model.Article{}).Where("c_id = ?",store[index]).Update(map[string]interface{}{"view":view,"like":like,"u_like":ulike})
	m.Unlock()
}