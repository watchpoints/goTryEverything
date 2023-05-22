### 奇迹提问5  python程序 支持微信公共号

ModuleNotFoundError: No module named 'Crypto'
pip install  pycryptodome

https://github.com/wechatpy/wechatpy/blob/master/examples/echo/app.py


 git clone git@github.com:wangcy6/chatgpt-on-wechat.git
cd chatgpt-on-wechat/
pip3 install -r requirements.txt
cp config-template.json config.json
python3 app.py

4. Railway部署 (✅推荐)
Railway每月提供5刀和最多500小时的免费额度。

进入 Railway。
点击 Deploy Now 按钮。
设置环境变量来重载程序运行的参数，例如open_ai_api_key, character_desc。

https://github.com/zhayujie/chatgpt-on-wechat/blob/master/channel/wechatmp/README.md

https://github.com/zhayujie/chatgpt-on-wechat/blob/master/channel/wechatmp/README.md

### 奇迹提问5  在发表文章 和登录时候提示验证码

what：
why：
https://github.com/weizhimeng/selenium-Sliding-verification-code
https://blog.csdn.net/lemonbit/article/details/112300728
how：


### 奇迹提问4： 解决

selenium滑动到元素可见状态

selenium——鼠标操作ActionChains：点击、滑动、拖动

https://www.cnblogs.com/mini-monkey/p/12169030.html

https://www.selenium.dev/documentation/webdriver/elements/finders/

### 奇迹提问3：

vcode连接远程服务器 编写python不能函数定义跳转问题解决





### 遇到错误


How to Install Google Chrome Web Browser on CentOS 8

查看 已经提供的 谷歌浏览器版本
dnf provides google-chrome

选择113版本：
google-chrome-beta-113.0.5672.37-1.x86_64

google-chrome --version
Google Chrome 113.0.5672.37 beta









~~~~

[root@VM-8-8-centos python]# google-chrome --version
Google Chrome 112.0.5615.121
[root@VM-8-8-centos python]# chromedriver --version
ChromeDriver 112.0.5615.49 (bd2a7bcb881c11e8cfe3078709382934e3916914-refs/branch-heads/5615@{#936})


Message: unknown error: Chrome failed to start: exited abnormally.
  (unknown error: DevToolsActivePort file doesn't exist)
  (The process started from chrome location /root/bin/chromedriver is no longer running, so ChromeDriver is assuming that Chrome has crashed.)

Traceback (most recent call last):
  File "./weibo.py", line 170, in post_sleep_weibo
    driver = init_browser(weibo_driver_path)
  File "./weibo.py", line 69, in init_browser
    driver = webdriver.Chrome(options=chrome_options)
  File "/usr/local/lib/python3.6/site-packages/selenium/webdriver/chrome/webdriver.py", line 81, in __init__
    desired_capabilities=desired_capabilities)
  File "/usr/local/lib/python3.6/site-packages/selenium/webdriver/remote/webdriver.py", line 157, in __init__
    self.start_session(capabilities, browser_profile)
  File "/usr/local/lib/python3.6/site-packages/selenium/webdriver/remote/webdriver.py", line 252, in start_session
    response = self.execute(Command.NEW_SESSION, parameters)
  File "/usr/local/lib/python3.6/site-packages/selenium/webdriver/remote/webdriver.py", line 321, in execute
    self.error_handler.check_response(response)
  File "/usr/local/lib/python3.6/site-packages/selenium/webdriver/remote/errorhandler.py", line 242, in check_response
    raise exception_class(message, screen, stacktrace)
selenium.common.exceptions.WebDriverException: Message: unknown error: Chrome failed to start: exited abnormally.
  (unknown error: DevToolsActivePort file doesn't exist)
  (The process started from chrome location /root/bin/chromedriver is no longer running, so ChromeDriver is assuming that Chrome has crashed.)


一定采用python,google-chrome,selenium 最新版本

https://www.luzhining.com/blog/article/detail?id=127

通过设置软链接简化命令输入

ln -sf /usr/local/python311/bin/python3.11 /usr/bin/python3
ln -sf /usr/local/python311/bin/pip3.11  /usr/bin/pip3

/usr/local/python311/bin;

[root@VM-8-8-centos Python-3.11.0]# pip3 show selenium
Name: selenium
Version: 4.8.3
Summary:


 python3 --version
Python 3.11.0

这个警告是由于在使用Selenium时，executable_path参数已经被弃用，建议使用Service对象代替。您可以这样修改代码：

复制
from selenium import webdriver
from selenium.webdriver.chrome.service import Service

service = Service('chromedriver.exe') # 将文件路径作为参数传入Service对象
driver = webdriver.Chrome(service=service) # 将Service对象作为参数传入Chrome浏览器对象
注意：chromedriver.exe是Chrome浏览器驱动程序的文件名，您需要将其替换为您自己的驱动程序文件名。

3 Ways to Install Google Chrome on CentOS 8

~~~~

#### 版本过低问题：

how to set path in chrome_opt  


~~~~
To set the path in Chrome options, you can use the binary_location argument. This allows you to specify the path to the Chrome binary on your system.

Here's an example of how you can set the path using Python:

from selenium import webdriver

chrome_options = webdriver.ChromeOptions()
chrome_options.binary_location = "/path/to/chrome/binary"

driver = webdriver.Chrome(options=chrome_options)
Replace "/path/to/chrome/binary" with the actual path to the Chrome binary on your system. Once you've set the path, you can create a new webdriver.Chrome instance with the chrome_options object as an argument. This will launch Chrome with the specified binary locat
~~~~
### liunx 环境安装

google-chrome --version
Google Chrome 112.0.5615.121
https://chromedriver.storage.googleapis.com/index.html?path=112.0.5615.49/



pip3 install -U  selenium

pip3 show selenium
Name: selenium
Version: 3.141.0


chrome://version/
Google Chrome	98.0.4758.102 (正式版本) （64 位） (cohort: Stable)
https://chromedriver.chromium.org/downloads


liunx谷歌浏览器
https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm
rpm  -ivh  google-chrome-stable_current_x86_64.rpm
google-chrome --version
Google Chrome 98.0.4758.102
https://chromedriver.chromium.org/downloads

依赖：
yum install -y libgtk*

yum provides */libappindicator3.so.1
yum install libappindicator-gtk3
yum install liberation-fonts

https://seeyon.ren/blog/index.php/archives/199/
whereis google-chrome
google-chrome: /usr/bin/google-chrome /usr/share/man/man1/google-chrome.1.gz
[root@VM-8-8-centos bin]#



### 安装 chrome
~~~
chrome://version/
112.0.5615.86 (正式版本) （64 位)
https://chromedriver.storage.googleapis.com/index.html?path=112.0.5615.49/
~~~
### 安装 selenium

~~~
pip install -U  selenium
pip show selenium
Name: selenium
Version: 4.8.3
d:\tools\python3\lib\site-packages

pip 安装后 pycharm 找不到selenium

~~~

### 修改pip 默认安装目录
一、使用命令查看pip默认安装目录

python -m site

USER_BASE和USER_SITE其实就是默认的启用Python通过pip自动下载的脚本和依赖安装包的基础路径

python -m site -help

D:\tools\python3\lib\site.py
export PYTHONPATH=$PYTHONPATH:/usr/local/python3.6/site-packages
使用命令查看pip默认安装目录


>一定采用python,google-chrome,selenium 最新版本

- python 4.10

- Selenium 是 ThoughtWorks 提供的一个强大的基于浏览器的开源自动化测试工具
~~~
pip3 install -U  selenium
pip3 show selenium
Name: selenium
Version: 4.1.1

~~~

- ChromeDriver 是 google 为网站开发人员提供的自动化测试接口，它是 selenium2 和 chrome浏览器 进行通信的桥梁。
~~~
chrome://version/
Google Chrome	98.0.4758.102 (正式版本) （64 位） (cohort: Stable)
https://chromedriver.chromium.org/downloads


liunx谷歌浏览器
https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm
rpm  -ivh  google-chrome-stable_current_x86_64.rpm
google-chrome --version
Google Chrome 98.0.4758.102
https://chromedriver.chromium.org/downloads

~~~
- 修改配置路径：
~~~
路径：去掉相对路径
self.cookpath="cookies.pkl"
path = r"D:\local\Python\tool\chromedriver.exe"
~~~

cookies.pkl 默认~用户下：/root

- 调度控制：启动
腾讯云的主机 根本无法启动 好事3小时无法解决。暂停

win10 系统设置定时自动执行 python 脚本
D:\local\python3\python.exe D:\code\go\src\gitee.com\wang_cyi\TryEverything\python\toutiao\toutiao1.0.py
/root/local/python/toutiao_btc.py

# to do

1. 代理模式下启动失败 情况出现。


2. 查看源代码方式 看不到web浏览器内存，必须F12查看
<button class="byte-btn byte-btn-primary byte-btn-size-default byte-btn-shape-square publish-content" type="button"><span>发布</span></button>

selenium.common.exceptions.ElementClickInterceptedException: Message: element click intercepted: Element <button class="byte-btn byte-btn-primary byte-btn-size-default byte-btn-shape-square publish-content" type="button">...</button> is not clickable at point (836, 855). Other element would receive the click: <div class="byte-drawer-scroll">...</div>
  (Session info: chrome=87.0.4280.66)


<button class="byte-btn byte-btn-primary byte-btn-size-large byte-btn-shape-square" type="button" style="margin-left: 32px;"><span>开始创作</span></button>


3. selenium实现自动发送微博
python 利用selenium模拟登录帐号验证网站并获取cookie
class="woo-button-main woo-button-flat woo-button-primary woo-button-m woo-button-round Tool_btn_2Eane"

---B---

<div placeholder="有什么想和大家分享的？" contenteditable="true" class="bili-rich-textarea__inner empty" style="font-size: 15px; line-height: 24px; min-height: 24px;" data-listener-added_b42a6bd6="true">​
</div>

<div class="bili-dyn-publishing__action launcher disabled">
        发布
      </div>
<div class="bili-dyn-publishing__action launcher">
        发布
      </div>

- 如何通过selenium获取到弹窗里的按钮并模拟点击确定或取消？
 selenium实践-如何验证警告、确认、提示信息
https://blog.csdn.net/whowhowhoisimportant/article/details/115708083
How to Handle

<button class="bili-dyn-specification-popup__btn bili-button primary bili-button--medium">
        确认并发送  
      </button>

selenium 自动化 携带cookies模拟登陆哔哩哔哩并发送弹幕和评论（解决多窗口切换、规避检测、评论无法输入等问题
https://blog.csdn.net/m0_50944918/article/details/112148216

How to switch to frames in Selenium?

https://www.cnblogs.com/hiyong/p/14163153.html



https://github.com/SK-415/HarukaBot/issues/335



用Flask 写一个python程序 处理post get请求 的例子？


~~~
好的，以下是使用 Flask 处理 POST 和 GET 请求的示例代码：

```python
from flask import Flask, request

app = Flask(__name__)

# 处理 GET 请求
@app.route('/get_example', methods=['GET'])
def get_example():
    name = request.args.get('name')
    return f'Hello, {name}! This is a GET request.'

# 处理 POST 请求
@app.route('/post_example', methods=['POST'])
def post_example():
    data = request.get_json()
    name = data['name']
    return f'Hello, {name}! This is a POST request.'

if __name__ == '__main__':
    app.run()
```

在上面的代码中，我们定义了两个路由 `/get_example` 和 `/post_example`，分别用于处理 GET 和 POST 请求。

当我们向 `/get_example` 发送 GET 请求时，可以通过 `request.args.get()` 方法获取 URL 中的参数。在上面的例子中，我们获取了名为 `name` 的参数，并将其插入到返回的字符串中。

当我们向 `/post_example` 发送 POST 请求时，可以通过 `request.get_json()` 方法获取请求体中的 JSON 数据。在上面的例子中，我们获取了名为 `name` 的属性，并将其插入到返回的字符串中。

希望这个例子可以帮助你更好地理解如何使用 Flask 处理 POST 和 GET 请求。
~~~

Java SE 8 Archive Downloads (JDK 8u211 and later)
https://www.oracle.com/java/technologies/javase/javase8u211-later-archive-downloads.html





用 go 写一个爬虫 （why 不用python呢）
https://juejin.cn/post/6949557048203804709
https://juejin.cn/post/6949557048203804709
写一个发帖程序 用 go + Selenium ?

go get github.com/tebeka/selenium
go get github.com/tebeka/selenium/chrome 

go get github.com/tebeka/selenium/chrome 

https://github.com/tebeka/selenium/blob/master/example_test.go
单元测试

testing: warning: no tests to run
go help testfunc



A test function is one named TestXxx (where Xxx does not start with a
lower case letter) and should have the signature,

        func TestXxx(t *testing.T) { ... }


export GOPROXY=https://proxy.golang.com.cn,direct


问：go 下载慢 ，必须设置代理， 云主机 必须设置代理i板瓦工必须设置代理

 配置 GOPROXY 环境变量
export GOPROXY=https://proxy.golang.com.cn,direct
一个全球代理
为 Go 模块而生


问：go 是如何管理依赖的？

https://blog.csdn.net/weixin_39003229/article/details/97638573


## 部署运维

### 环境安装 

go：

export GOROOT=/root/bin/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=/root/code

https://gitee.com/wang_cyi/goTryEverything
1.  设置开启启动

window11---任务计划程序--当计算机启动时 --D:\local\goTryEverything.exe


go mod tidy



