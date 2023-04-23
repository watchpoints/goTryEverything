// dailyTask
package oneday

import (
	"fmt"

	"log"

	"runtime"

	"os/exec"

	"github.com/gen2brain/beeep"
	"github.com/watchpoints/goTryEverything/wechat"

	//"github.com/go-toast/toast"
	"github.com/jakecoffman/cron"
)

/*
*

1. 完整框架：https://github.com/ouqiang/gocron
2. 只有后端服务端的定时器 https://pkg.go.dev/github.com/robfig/cron
3. https://blog.csdn.net/sryan/article/details/50129133
4. 用法：https://github.com/eddycjy/go-gin-example/tree/master/service/article_service

*
*/
type DailyAction struct {
	Name string
}

var JobStandUpFun = func() {
	sysType := runtime.GOOS
	if sysType == "windows" {
		NotifyStadnUpWindos()
	}

}

var JobSleepNow = func() {
	ShutDownPc()

}

var JobPutDownComputerAndPhone = func() {
	PutDownComputerAndPhone()

}

var SleepTrip = func() {
	wechat.EverydaySleep()
}

func (da *DailyAction) StartCron() {

	fmt.Println("start StartCron")
	c := cron.New()

	// job_stand_up := "0 25 * * * *" //25提醒站立一次
	// c.AddFunc(job_stand_up, JobStandUpFun, "番茄时间")

	// c.AddFunc("0 30 22 * * *", JobSleepNow, "如果自己控制不了，执行22:30 强制关机")

	// c.AddFunc("0 0 20 * * *", JobPutDownComputerAndPhone, "20:00 远离手机和电脑")
	c.AddFunc("0 30 21 * * *", SleepTrip, "20:00 远离手机和电脑")

	c.Start()
}

/*
*

	消息通知弹框

https://github.com/gen2brain/beeep
https://github.com/0xAX/notificator
https://github.com/go-toast/toast
https://github.com/ying32/govcl
https://zhuanlan.zhihu.com/p/404062419
*
*/
func NotifyStadnUp() {

	err := beeep.Notify("Title", "Message body", "assets/information.png")
	if err != nil {
		panic(err)
	}
}
func NotifyStadnUpWindos() {
	/*
		notification := toast.Notification{
			AppID:   "Microsoft.Windows.Shell.RunDialog",
			Title:   "站立提醒",
			Message: "是的到了休息一下",
			Icon:    "d:\\golang_official_logo_icon_169092.png", // 文件必须存在
			Actions: []toast.Action{
				{"protocol", "马上去做", ""},
				{"protocol", "延迟", ""},
			},
		}
		err := notification.Push()

		if err != nil {
			log.Fatalln(err)
		}*/
}

// ShutDownPc cmd
func ShutDownPc() {

	if runtime.GOOS == "windows" {

		fmt.Println("ShutDownPc")

		//60*5
		arg := []string{"-s", "-t", "30"}
		cmd := exec.Command("shutdown", arg...)
		d, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(string(d))

	}

}

// 习惯：不拖堂 20点远离手机和电脑
// -s"正常关机、"-f"强制关机、"-r"重启、"-t"定时关机
// shutdown -h 休眠
func PutDownComputerAndPhone() {

	if runtime.GOOS == "windows" {

		fmt.Println("习惯：不拖堂 20点远离手机和电脑")

		cmd := exec.Command("shutdown", "-h")
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}

	}

}
