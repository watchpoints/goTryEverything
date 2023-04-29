import logging
from apscheduler.schedulers.blocking import BlockingScheduler
from apscheduler.triggers.cron import CronTrigger
from apscheduler.schedulers.background import BackgroundScheduler
import sleep
import http.server
from flask import Flask, request, jsonify

LOG_FORMAT = "[%(asctime)s][%(levelname)s][%(filename)s:%(funcName)s:%(lineno)d] %(message)s"
DATE_FORMAT = '%Y-%m-%d %H:%M:%S'
#  https://github.com/zhayujie/bot-on-anything
# https://github.com/zhayujie/chatgpt-on-wechat

app = Flask(__name__)
@app.route('/')
def hello_world():
    return 'Hello, World!'

@app.route('/api/json', methods=['POST'])
def example():
    if request.method == 'POST':
        # 从 POST 请求的 JSON 数据中获取参数
        data = request.json
        name = data.get('name')
        age = data.get('age')

        # 在服务端进行一些处理
        # ...

        # 返回 JSON 格式的响应
        response = {'message': 'Hello, {}, you are {} years old.'.format(name, age)}
        return jsonify(response)
    
if __name__ == "__main__":
    # not  execute logging  fucntion before  here
    logging.basicConfig(level=logging.DEBUG,
                        format=LOG_FORMAT,
                        datefmt=DATE_FORMAT,
                        filename="./log"
                        )
  
    logging.info("""
        ┌──────────────────────────────────────────────────────────────────────┐
        │                                                                      │    
        │                      •  Start Beemove  •                             │
        │                                                                      │
        └──────────────────────────────────────────────────────────────────────┘
    """)

    job_defaults = {
                'coalesce': False,
                'max_instances': 1
            }
    backsched = BackgroundScheduler(job_defaults=job_defaults, timezone='Asia/Shanghai')
    backsched.add_job(sleep.show_sleep, CronTrigger.from_crontab("0 22 * * *"), id="do_show_sleep_job")
    # backsched.add_job(sleep.show_sleep, CronTrigger.from_crontab("10 15 * * *"), id="do_show_sleep_job")
    backsched.start()
    
    
    # server_address = ("", 8089)
    # httpd = http.server.HTTPServer(server_address, http.server.SimpleHTTPRequestHandler)
    # httpd.serve_forever()
    # host: 绑定的ip(域名)
    # port: 监听的端口号
    # debug: 是否开启调试模式
    app.run(host="0.0.0.0", port=8000, debug=True)
    # export FLASK_APP=xx.py  # 指定flask应用所在的文件路径
    # export FLASK_ENV=development  # 设置项目的环境, 默认是生产环境
    # flask run -h 0.0.0.0 -p 8000  # 启动测试服务器并接受请求
    # http://127.0.0.1:8000/


        
