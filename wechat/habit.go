package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"gitee.com/wang_cyi/TryEverything/promise"
	"gitee.com/wang_cyi/TryEverything/public"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

//WeCHatHandler
type WeCHatHandler struct {
	LifeMap map[string]*promise.LifeData
}

//SendMsg SendMsg
func (wx *WeCHatHandler) SendMsg(msg string) {
	if len(msg) == 0 {
		return
	}
	log.Println(msg)
	public.SendMailBy163(msg)
	if true {
		if len(msg) > 140 {

		}
		public.WeiboDailyThinking(msg)
	}

}
func (wx *WeCHatHandler) dealTxtMsg(msg string) string {
	var ret string
	var plan promise.Plan
	if strings.Contains(msg, "复盘") == true {
		//log.Println("ask_deal", msg)
		wx.SendMsg(msg)

	} else if strings.Contains(msg, "起床打卡") == true || strings.Contains(msg, "6") == true {

		msg := " 起床：现在北京时间是:"
		msg += public.GetBeijingTime()
		msg += "努力一下，下次就是6点了."
		ret = msg
		wx.SendMsg(msg)

	} else if strings.Contains(msg, "地铁打卡") == true || strings.Contains(msg, "18") == true {

		msg := "365天每天地铁15分钟阅读计划:"
		msg += "\r\n"
		msg += " 已经 "
		msg += public.GetDayFrom20200101()
		msg += "天. 今天打卡时间:"
		msg += public.GetBeijingTime()
		msg += "\r\n"
		ret = msg
		wx.SendMsg(msg)

	} else if strings.Contains(msg, "吃饭") == true || strings.Contains(msg, "eat") == true {
		fmt.Println("eat plan")
		eatData, ok := wx.LifeMap["eat"]
		if ok == false {

			ret = "eatMap is null"
			log.Println(ret)
			return ret
		}
		if eatData == nil {
			ret = "eatMap is null"
			log.Println(ret)
			return ret
		}

		ret = plan.GetEatPromeData(eatData)
		fmt.Println(ret)

		wx.SendMsg(ret)

	} else if strings.Contains(msg, "放下手机") == true || strings.Contains(msg, "sleep") == true {
		fmt.Println("eat plan")
		sleepData, ok := wx.LifeMap["sleep"]
		if ok == false {

			ret = "sleepData is null"
			log.Println(ret)
			return ret
		}
		if sleepData == nil {
			ret = "sleepData is null"
			log.Println(ret)
			return ret
		}

		ret = plan.GetOnTimeSleepPromeData(sleepData)
		fmt.Println(ret)

		wx.SendMsg(ret)
	} else if strings.Contains(msg, "阅读") == true || strings.Contains(msg, "book") == true || strings.Contains(msg, "18") == true {

		bookData, ok := wx.LifeMap["book"]
		log.Println(msg)
		if ok == false {

			ret = "bookData is null"
			log.Println(ret)
			return ret
		}
		if bookData == nil {
			ret = "bookData is null"
			log.Println(ret)
			return ret
		}

		ret = plan.GetReadBookData(bookData)
		fmt.Println(ret)

		wx.SendMsg(ret)
	} else if strings.Contains(msg, "驼背") == true || strings.Contains(msg, "坐正") == true || strings.Contains(msg, "money") == true || strings.Contains(msg, "坐姿") == true || strings.Contains(msg, "疲劳") == true || strings.Contains(msg, "头疼") == true || strings.Contains(msg, "高低肩") == true {

		moneyData, ok := wx.LifeMap["money"]
		log.Println(msg)
		if ok == false {

			ret = "moneyData is null"
			log.Println(ret)
			return ret
		}
		if moneyData == nil {
			ret = "moneyData is null"
			log.Println(ret)
			return ret
		}

		ret = plan.GetMoneyData(moneyData)
		ret += public.UpdateSite()
		fmt.Println(ret)

		wx.SendMsg(ret)
	} else if strings.Contains(msg, "action") == true {
		//插入msyql 一天一小时记录 msg
		if len(msg) > 5 {
			public.DayMysql(msg)
		}

	} else if strings.Contains(msg, "22") == true || strings.Contains(msg, "压力") == true || strings.Contains(msg, "手机") == true || strings.Contains(msg, "早睡早起") == true || strings.Contains(msg, "防沉迷不解锁") == true || strings.Contains(msg, "上床睡觉") == true || strings.Contains(msg, "sex") == true || strings.Contains(msg, "不行") == true || strings.Contains(msg, "小说") == true || strings.Contains(msg, "人妻") == true || strings.Contains(msg, "周末") == true || strings.Contains(msg, "动漫") == true || strings.Contains(msg, "综艺") == true || strings.Contains(msg, "电视剧") == true || strings.Contains(msg, "美杜莎") == true || strings.Contains(msg, "视频") == true || strings.Contains(msg, "小说") == true || strings.Contains(msg, "自我救赎") == true || strings.Contains(msg, "床") == true || strings.Contains(msg, "腾讯视频") || strings.Contains(msg, "不解锁") || strings.Contains(msg, "浏览器") == true {

		//更新DB,解决因为每次重启丢失数据，因此不再使用该功能问题
		msg := public.UpdateTV()
		wx.SendMsg(msg) //发送邮件
		ret = msg

	} else if strings.Contains(msg, "活动") == true || strings.Contains(msg, "运动") == true || strings.Contains(msg, "跑步") == true || strings.Contains(msg, "keep") == true || strings.Contains(msg, "21") == true || strings.Contains(msg, "健康") == true || strings.Contains(msg, "跑步") == true || strings.Contains(msg, "椅子") == true || strings.Contains(msg, "户外") == true || strings.Contains(msg, "5公里") == true {
		//跑步
		msg := public.RunFuture()
		wx.SendMsg(msg) //发送邮件
		ret = msg
	} else if len(msg) > 20 {
		log.Println("没有格式", msg)
		public.SendMailBy163(msg)
		if true {
			if len(msg) > 140 {
				log.Println(" more than 140 ", len(msg))

			}
			public.WeiboDailyThinking(msg)
		}

	}

	return ret
}
func (wx *WeCHatHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	//fmt.Println("wechat..............")
	//配置微信参数
	//	config := &wechat.Config{
	//		AppID:          "your app id",
	//		AppSecret:      "your app secret",
	//		Token:          "your token",
	//		EncodingAESKey: "your encoding aes key",
	//	}
	//wc := wechat.NewWechat(config)
	// configtest := &wechat.Config{
	// 	AppID:          "wx019a4baceb35d648",
	// 	AppSecret:      "e6ba621f18e1eacb2c289d0d736e40e0",
	// 	Token:          "testtest",
	// 	EncodingAESKey: "test",
	// }

	configtest := &wechat.Config{
		AppID:          "wx94e03776e64ee600",
		AppSecret:      "a0af4a82bdb4a2d17af867a321da4c83",
		Token:          "testtest",
		EncodingAESKey: "Quk4OepZpDVkhWnevUNnOXfoynwT1gg83cYdBpaZXYP",
	}
	wc := wechat.NewWechat(configtest)

	// 传入request和responseWriter
	server := wc.GetServer(req, rw)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

		//回复消息：演示回复用户发送的消息
		////文本消息
		if msg.MsgType == message.MsgTypeText {
			text := message.NewText(msg.Content)
			data := wx.dealTxtMsg(msg.Content)
			if len(data) > 0 {
				log.Println(" dealTxtMsg ret ", data)
				text = message.NewText(data)
			}

			//bug1 text 不能为nil
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
		} else if msg.MsgType == message.MsgTypeVoice {
			//语音消息
		} else if msg.MsgType == message.MsgTypeImage {
			//图片消息
		}
		return nil

	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}
