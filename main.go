package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type arg struct {
	date string			//时间参数
	t int				//1:unix时间戳 2:日期格式
}

func main() {
	var A arg
	// 获取命令行参数
	lens := len(os.Args)
	switch lens {
	case 1:
		fmt.Printf("参数为空!")
		os.Exit(0)
	case 2:
		//通过数据长度识别是时间戳还是日期
		if len(os.Args[1]) == 10{
			A.date = os.Args[1]
			A.t = 1
		}else {
			A.date = os.Args[1] + " " + "00:00:00"
			A.t = 2
		}
	case 3:
		A.date = os.Args[1] + " " + os.Args[2]
		A.t = 2
	default:
		fmt.Printf("参数过多无法识别")
		os.Exit(0)
	}

	A.Transform()
}

//时间格式转换
func (a *arg)Transform()  {
	if a.t == 2 {
		//时间格式模板
		timeLayout := "20060102 15:04:05"
		LocalTime, _ := time.ParseInLocation(timeLayout, a.date, time.Local)
		UtcTime, _ := time.ParseInLocation(timeLayout, a.date, time.UTC)
		fmt.Printf("LocalTimeStamp = %d\n",LocalTime.Unix())
		fmt.Printf("UtcTimeStamp = %d\n",UtcTime.Unix())
		fmt.Printf("输入的时间为 %s\n",a.date)
		os.Exit(0)
	}
	if a.t == 1 {
		timeLayout := "20060102 15:04:05"
		ts, _ := strconv.ParseInt(a.date,10,64)
		Local := time.Unix(ts, 0).Format(timeLayout)
		Utc := time.Unix(ts, 0).UTC().Format(timeLayout)
		fmt.Printf("本地时间 %s\n",Local)
		fmt.Printf("Utc时间 %s\n",Utc)
		fmt.Printf("输入的时间戳为 %s\n",a.date)
		os.Exit(0)
	}
	return
}
