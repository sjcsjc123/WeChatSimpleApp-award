package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var Num = 10
var AlreadyBiu []int
var TotalAward float64

func main() {
	go updateNum()
	r := gin.Default()
	r.GET("/start", start)
	r.GET("/reStart", reStart)
	r.GET("/update", updateTotalAward)
	err := r.Run(":4000")
	if err != nil {
		panic(err)
	}
}

func updateTotalAward(c *gin.Context) {
	query := c.Query("award")
	award, err := strconv.ParseFloat(query, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "解析错误",
		})
		return
	}
	TotalAward = award
	c.JSON(200, gin.H{
		"msg": "更改成功",
	})
}

func start(c *gin.Context) {
	if Num == 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg":        "今日抽奖次数已用完",
			"totalAward": TotalAward,
		})
		return
	}
	Num--
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10000)
	res := float64(num) / 10000 * 100
	fmt.Println(res)
	selectIndex := 0
	switch {
	case res >= 80:
		selectIndex = 7
	case 60 <= res && res < 80:
		selectIndex = 5
	case 30 <= res && res < 60:
		selectIndex = 3
	case 5 <= res && res < 30:
		selectIndex = 1
	case 3 <= res && res < 5:
		selectIndex = 6
	case 1 <= res && res < 3:
		selectIndex = 4
	case 0.3 <= res && res < 1:
		selectIndex = 2
	case res < 0.3:
		selectIndex = 0
	}
	fmt.Println(selectIndex)
	if selectIndex == 0 || selectIndex == 2 || selectIndex == 4 || selectIndex == 6 || selectIndex == 8 {
		AlreadyBiu = append(AlreadyBiu, selectIndex)
		switch selectIndex {
		case 0:
			TotalAward = TotalAward + 52
			fmt.Println("抽中了52块钱")
			c.JSON(http.StatusOK, gin.H{
				"msg":         "恭喜你抽中了52元红包",
				"totalAward":  TotalAward,
				"indexSelect": selectIndex,
			})
		case 1:
			TotalAward = TotalAward + 0.52
			fmt.Println("抽中了0.52块钱")
			c.JSON(http.StatusOK, gin.H{
				"msg":         "恭喜你抽中了0.52元红包",
				"totalAward":  TotalAward,
				"indexSelect": selectIndex,
			})
		case 2:
			TotalAward = TotalAward + 21.21
			fmt.Println("抽中了21.21块钱")
			c.JSON(http.StatusOK, gin.H{
				"msg":         "恭喜你抽中了21.21元红包",
				"totalAward":  TotalAward,
				"indexSelect": selectIndex,
			})
		case 3:
			TotalAward = TotalAward + 1.31
			fmt.Println("抽中了1.31块钱")
			c.JSON(http.StatusOK, gin.H{
				"msg":         "恭喜你抽中了1.31元红包",
				"totalAward":  TotalAward,
				"indexSelect": selectIndex,
			})
		case 4:
			TotalAward = TotalAward + 13.14
			fmt.Println("抽中了13.14块钱")
			c.JSON(http.StatusOK, gin.H{
				"msg":         "恭喜你抽中了13.14元红包",
				"totalAward":  TotalAward,
				"indexSelect": selectIndex,
			})
		case 5:
			TotalAward = TotalAward + 2.12
			fmt.Println("抽中了2.12块钱")
			c.JSON(http.StatusOK, gin.H{
				"msg":         "恭喜你抽中了2.12元红包",
				"totalAward":  TotalAward,
				"indexSelect": selectIndex,
			})
		case 6:
			TotalAward = TotalAward + 5.2
			fmt.Println("抽中了5.2块钱")
			c.JSON(http.StatusOK, gin.H{
				"msg":         "恭喜你抽中了5.2元红包",
				"totalAward":  TotalAward,
				"indexSelect": selectIndex,
			})
		}
		Num++
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":         "再接再厉哦",
			"totalAward":  TotalAward,
			"indexSelect": selectIndex,
		})
	}
}

//手动重设抽奖次数
func reStart(c *gin.Context) {
	Num = 10
	c.JSON(200, gin.H{
		"msg": "重设抽奖次数成功",
	})
}

func updateNum() {
	for {
		now := time.Now()                                                                    //获取当前时间，放到now里面，要给next用
		next := now.Add(time.Hour * 24)                                                      //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location()) //获取下一个凌晨的日期
		fmt.Println(next)
		t := time.NewTimer(next.Sub(now)) //计算当前时间到凌晨的时间间隔，设置一个定时器
		<-t.C
		Num = 10
	}
}
