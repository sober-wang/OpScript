#-*- coding:utf-8 -*-
# 作用：筛选出 异常的 yarn 作业。详情请看 app_kill 函数。
# 运行方式：python yarn-kill-getapp.py 192.168.1.1
# 作者：尚墨
# 日期：2018年9月17日
# 邮箱：ws1992jx@163.com
import requests
import sys
import os


def app_select(app_msg_list):
    '''
    :param app_msg_list: 过滤出正在运行的 app ，并且修改 使用内存信息，运行总时长，
    :return:
    '''
    msg_list = []
    for app_msg in app_msg_list:
        if app_msg["state"] == "RUNNING":
            # 这里 app 使用的 总内存 单位是:G
            app_msg["allocatedMB"] = int(app_msg["allocatedMB"])/1024
            app_msg["elapsedTime"] = app_msg["elapsedTime"] / 1000 / 60 / 60
            msg_list.append(app_msg)
    return msg_list

def app_kill(app_list):
    '''
    :param app_list: 挑选出 需要处理的 app 进程
    :return:
    '''
    app_memory = 1
    app_run_time = 3
    for app_len in app_list:
        if app_len["allocatedMB"] > app_memory and app_len["startedTime"] < app_run_time:
            print(u"3 hores start app : [ %s ] Memory used : [ %s ]"%(app_len["id"],app_len["allocatedMB"]))
            os.system("yarn application -kill %s"%app_len["id"])

if __name__ == "__main__":
    try:
        resc_ip = sys.argv[1]
        url = "http://%s:8088/ws/v1/cluster/apps"%resc_ip
        print(url)
        r = requests.get(url)
    except Exception as e:
        print(e)
        print("[ ERROR ] Resource Manager IP is error ip")
    yarn_app_msg = r.json()["apps"]["app"]
    run_app_list = app_select(yarn_app_msg)
    app_kill(run_app_list)
