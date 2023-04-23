package oneday

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {
	// 设置 ChromeDriver 的路径
	const (
		path = "/usr/local/bin/chromedriver"
		port = 9515
	)

	// 启动 ChromeDriver
	opts := []selenium.ServiceOption{
		selenium.ChromeDriver(path),
	}
	service, err := selenium.NewChromeDriverService(path, port, opts...)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer service.Stop()

	// 设置 Chrome 的选项
	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: []string{
			"--headless",
			"--disable-gpu",
			"--no-sandbox",
			"--disable-dev-shm-usage",
			"--disable-extensions",
			"--disable-infobars",
			"--disable-features=NetworkService",
		},
	}

	// 连接到 Chrome 浏览器
	wd, err := selenium.NewRemote(chromeCaps)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wd.Quit()

	// 打开要发帖的网站
	if err := wd.Get("https://example.com"); err != nil {
		fmt.Println(err)
		return
	}

	// 等待页面加载完成
	time.Sleep(5 * time.Second)

	// 填写表单并提交
	if err := wd.FindElement(selenium.ByCSSSelector, "#title").SendKeys("这是标题"); err != nil {
		fmt.Println(err)
		return
	}
	if err := wd.FindElement(selenium.ByCSSSelector, "#content").SendKeys("这是内容"); err != nil {
		fmt.Println(err)
		return
	}
	if err := wd.FindElement(selenium.ByCSSSelector, "#submit").Click(); err != nil {
		fmt.Println(err)
		return
	}

	// 等待页面加载完成
	time.Sleep(5 * time.Second)

	// 打印发帖结果
	fmt.Println("发帖成功！")
}
