import logging
import weibo
import os
import toutiao
import maimai

# easy sleep
def show_sleep():
    os.system("ps -ef | grep google-chrome  | grep -v grep | awk '{print $2}' | xargs kill")
    logging.info("22点打卡，关机睡觉，手机一定要放客厅")
    weibo.post_sleep_weibo()
    weibo.post_sleep_weibo()

    toutiao.post_sleep_toutiao()  # easy sleep

    maimai.post_sleep_maimai()
