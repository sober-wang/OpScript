#-*- coding:utf-8 -*-
# 作用：按照运行时间，运行内存过滤出对应的 yarn 作业
# 作者：尚墨
# 使用方式： python yarn-rest-getapp.py 192.168.1.1
# 邮箱：ws1992jx@163.com
# 日期：2018年9月14日

import requests
import sys
import json

def app_select(app_msg_list):
    '''
    过滤、精简，yarn 作业信息。
    只保留 RUNNING 作业信息
    '''
    msg_list = []
    for app_msg in app_msg_list:
        if app_msg["state"] == "RUNNING":
            app_msg["elapsedTime"] = app_msg["elapsedTime"] / 1000 / 60
            msg_list.append(app_msg)
    return msg_list

def app_filter(app_list):
    '''
    通过用户输入的内存，时间大小筛选出 yarn 作业
    '''
    try:
        app_memory = float(input("Please enter app memory number lower limter (mb): "))
        app_run_time = float(input("Please enter app elapsed time (m):"))
    except:
        print("[ ERROR ] Pleace enter int number")
    for app_len in app_list:
        if app_len["allocatedMB"] > app_memory and app_len["elapsedTime"] > app_run_time:
            print("Appid:[ %s ] AppElapsedTime:[ %.2f m ] Appmemory: [ %d M ]"%(
                app_len["id"],
                app_len["elapsedTime"],
                app_len["allocatedMB"]
                )
             )

if __name__ == "__main__":
    try:
        resc_ip = sys.argv[1]
        url = "http://%s:8088/ws/v1/cluster/apps"%resc_ip
        r = requests.get(url)
    except:
        print("[ ERROR ]ResourceManager IP is error")
    yarn_app_msg = r.json()["apps"]["app"]
    run_app_list = app_select(yarn_app_msg)
    app_filter(run_app_list)
