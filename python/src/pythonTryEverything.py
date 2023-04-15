import logging
from apscheduler.schedulers.blocking import BlockingScheduler
from apscheduler.triggers.cron import CronTrigger
from apscheduler.schedulers.background import BackgroundScheduler
import sleep
import http.server

LOG_FORMAT = "[%(asctime)s][%(levelname)s][%(filename)s:%(funcName)s:%(lineno)d] %(message)s"
DATE_FORMAT = '%Y-%m-%d %H:%M:%S'

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
    
    server_address = ("", 8089)
    httpd = http.server.HTTPServer(server_address, http.server.SimpleHTTPRequestHandler)
    httpd.serve_forever()

        
