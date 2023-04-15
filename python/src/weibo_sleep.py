from selenium import webdriver
import time
import requests
import json
import os

# 初始化浏览器 打开微博登录页面
def init_browser():
    path = r'D:/MyEnv/chromedriver.exe'  # 指定驱动存放目录
    chrome_options = webdriver.ChromeOptions()
    # 把允许提示这个弹窗关闭
    prefs = {"profile.default_content_setting_values.notifications": 2}
    chrome_options.add_experimental_option("prefs", prefs)
    driver = webdriver.Chrome(executable_path=path, options=chrome_options)
    driver.maximize_window()
    driver.get('https://weibo.com/login.php')
    return driver

# 读取cookies 登录微博
def login_weibo(driver):
    cookies = read_cookies()
    for cookie in cookies:
        driver.add_cookie(cookie)
    time.sleep(3)
    driver.refresh()  # 刷新网页

# 获取cookie并存入文件
def get_cookies(driver):
    driver.get('https://weibo.com/login.php')
    time.sleep(20)  # 留时间进行扫码
    Cookies = driver.get_cookies()  # 获取list的cookies
    jsCookies = json.dumps(Cookies)  # 转换成字符串保存
    with open('cookies.txt', 'w') as f:
        f.write(jsCookies)
    print('cookies已重新写入！')

# 读取本地的cookies
def read_cookies():
    # 检查文件是否存在
    files = os.path.exists('D:/Python/douyin/cookies.txt')
    if files:
        with open('cookies.txt', 'r', encoding='utf8') as f:
            Cookies = json.loads(f.read())
        cookies = []
        for cookie in Cookies:
            cookie_dict = {
                'domain': '.weibo.com',
                'name': cookie.get('name'),
                'value': cookie.get('value'),
                'expires': '',
                'path': '/',
                'httpOnly': False,
                'HostOnly': False,
                'Secure': False
            }
            cookies.append(cookie_dict)
        return cookies
    else:
        return False

# 检测cookies的有效性
def check_cookies():
    # 读取本地cookies
    cookies = read_cookies()
    s = requests.Session()
    for cookie in cookies:
        s.cookies.set(cookie['name'], cookie['value'])
    response = s.get("https://weibo.com")
    response.encoding = response.apparent_encoding
    html_t = response.text
    # 检测页面是否包含微博用户名
    if '用户7720733258' in html_t:
        return True
    else:
        return False

# 发布微博
def post_weibo(content, driver):
    time.sleep(5)
    weibo_content = driver.find_element_by_xpath('//*[ @id ="homeWrap"]/div[1]/div/div[1]/div/textarea')
    weibo_content.send_keys(content)
    bt_push = driver.find_element_by_xpath('//*[@id="homeWrap"]/div[1]/div/div[4]/div/button')
    bt_push.click()  # 点击发布
    time.sleep(5)
    driver.close()  # 关闭浏览器


if __name__ == '__main__':
    # 查看本地是否有cookie
    re = read_cookies()
    # 存在则检测
    if re:
        # 检测cookies的有效性
        res = check_cookies()
        # 无效
        if res==False:
            # 扫码登录微博
            driver = init_browser()
            # 获取cookie
            get_cookies(driver)
         else:
            driver = init_browser()
    else:
        # 扫码登录微博
        driver = init_browser()
        # 获取cookie
        get_cookies(driver)

    login_weibo(driver)
    # 自动发微博
    content = '今天的天气湿冷~~~，心情不好'
    post_weibo(content, driver)