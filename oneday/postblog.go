package oneday

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

// https://github.com/tebeka/selenium/blob/master/example_test.go
// https://juejin.cn/post/6949557048203804709
const (
	// 设置 ChromeDriver 的路径
	chromeDriverPath = "D:\\doc\\2023\\05-third\\chromedriver_win32\\chromedriver.exe"
	port             = 8080
)

func PostGeekTime() {
	//Start a ChromeDriver server instance
	opts := []selenium.ServiceOption{
		selenium.Output(os.Stderr), // Output debug information to STDERR.
	}
	selenium.SetDebug(true)

	service, err := selenium.NewSeleniumService(chromeDriverPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// // 设置 Chrome 的选项
	// // Connect to the WebDriver instance running locally.
	// chromeCaps := chrome.Capabilities{
	// 	Path: "",
	// 	Args: []string{
	// 		"--headless",
	// 		"--disable-gpu",
	// 		"--no-sandbox",
	// 		"--disable-dev-shm-usage",
	// 		"--disable-extensions",
	// 		"--disable-infobars",
	// 		"--disable-features=NetworkService",
	// 	},
	// }
	chromeCaps := selenium.Capabilities{"browserName": "chrome"}
	// 连接到 Chrome 浏览器
	wd, err := selenium.NewRemote(chromeCaps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wd.Quit()

	// 打开要发帖的网站
	// if err := wd.Get("https://example.com"); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // 等待页面加载完成
	// time.Sleep(5 * time.Second)

	// // 填写表单并提交
	// if err := wd.FindElement(selenium.ByCSSSelector, "#title").SendKeys("这是标题"); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// if err := wd.FindElement(selenium.ByCSSSelector, "#content").SendKeys("这是内容"); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// if err := wd.FindElement(selenium.ByCSSSelector, "#submit").Click(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // 等待页面加载完成
	// time.Sleep(5 * time.Second)

	// // 打印发帖结果
	// fmt.Println("发帖成功！")
	// Navigate to the simple playground interface.
	// Navigate to the simple playground interface.
	if err := wd.Get("http://play.golang.org/?simple=1"); err != nil {
		panic(err)
	}

	// Get a reference to the text box containing code.
	elem, err := wd.FindElement(selenium.ByCSSSelector, "#code")
	if err != nil {
		panic(err)
	}
	// Remove the boilerplate code already in the text box.
	if err := elem.Clear(); err != nil {
		panic(err)
	}

	// Enter some new code in text box.
	err = elem.SendKeys(`
        package main
        import "fmt"
        func main() {
            fmt.Println("Hello WebDriver!")
        }
    `)
	if err != nil {
		panic(err)
	}

	// Click the run button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, "#run")
	if err != nil {
		panic(err)
	}
	if err := btn.Click(); err != nil {
		panic(err)
	}

	// Wait for the program to finish running and get the output.
	outputDiv, err := wd.FindElement(selenium.ByCSSSelector, "#output")
	if err != nil {
		panic(err)
	}

	var output string
	for {
		output, err = outputDiv.Text()
		if err != nil {
			panic(err)
		}
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("%s", strings.Replace(output, "\n\n", "\n", -1))
}
