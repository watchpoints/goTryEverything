package wechat

import (
	"fmt"
	"log"
	"net/http"

	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

// WeCHatHandler
type WeCHatHandler struct {
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
			log.Println(text)
			//if len(text) > 0 {
			//	log.Println(" dealTxtMsg ret ", text)
			//	text = message.NewText(text)
			//}
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
