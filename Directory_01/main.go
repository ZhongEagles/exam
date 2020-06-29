//计算出从你生日之日起到今天一共过了多少天？（python、go、shell均可）
//运维工程师 金山WPS 2020
package main

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now()
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	birthday, err := time.ParseInLocation("2006/01/02 15:04:05", "1997/06/19 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(birthday.Format("2006/01/02 15:04:05"))
	days:=now.Sub(birthday)
	fmt.Printf("days from my birthday is: %v\n",int(days.Hours()/24))

}
