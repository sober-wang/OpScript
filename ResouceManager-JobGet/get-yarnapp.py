#!/usr/bin/env python
#-*- coding:utf-8 -*-
# 作用：从 ResourceManager web 页面获取 Yarn Job 信息
# 作者：尚墨
# 日期：2018年8月29日
# 邮箱：ws1992jx@163.com

import requests
import time
import sys
import re
import os
from bs4 import BeautifulSoup

def js_app_msg(url):
    '''
    爬取 yarn ResouceManager Web 页面
    '''
    res = requests.get(url)
    print('web code varlue ------> %d'%res.status_code)
    if res.status_code != 200:
        print("页面请求失败，请检查 url 变量值")
        sys.exit()
    yarn_run = res.text
    print('***************************************************************')
    bs_obj = BeautifulSoup(yarn_run,'html.parser')
    for link in bs_obj.findAll('script'):
        if 'var' in link.string:
            #ls_tmp = link.string.decode('utf-8')
            #print(ls_tmp)
            app_list_tmp = link.string.split('\n')
            return app_list_tmp

def start_time(time_st):
    '''
    处理 web 页面中的时间戳
    '''
    point_before = time_st[0:-3]
    point_later = time_st[-3:-1]
    st_t = ".".join([point_before,point_later])
    return st_t
    

def slt_app(msg):
    ''' 
    洗出 yarn ResourceManager Web 页面中的任务
    ''' 
#    print(msg)
    work_list = []
    del(msg[0],msg[0],msg[-2],msg[-1],)
    for i in msg:
        work_dis = {}
        slt_tmp = i.split(',')
        slt_tmp_1 = slt_tmp[0].split('>')
        slt_tmp_2 = slt_tmp_1[1].split('<')
        work_dis["id"] = slt_tmp_2[0]
        work_dis["user"] = slt_tmp[1][1:-1]
        work_dis["name"] = slt_tmp[2][1:-1]
        work_dis["start_time"] = start_time(slt_tmp[5][1:-1])
        work_dis["run_contain"] = slt_tmp[9][1:-1]
        work_dis["cpu_num"] = slt_tmp[10][1:-1]
        work_dis["memory"] = slt_tmp[11][1:-1]
        work_dis["progress"] = slt_tmp[14].split("\'")[1]
        work_dis["keep_time"] = round(time.time()-float(work_dis["start_time"]),2)
        work_list.append(work_dis)
    return work_list

def slt_crim(web_msg,local_job):
    '''
    挑选待杀进程
    '''
    for m in web_msg:
        if int(m["memory"]) > 100 and re.match("^[0-9]+.*[0-9]$",m["user"])\
        and "Spark" in m["name"]  and m["keep_time"] > 7200:
            if len(local_job) == 0:
                print "本机无待处理进程"
            else:
                print "待杀任务 ----->  %s"%(m)
                print(m)

def yarn_job():
    '''
    获取本地的 yarn 作业进程，和 PID
    '''
    app_pid = {}
    pro_text = os.popen('jcmd | grep yarn | grep jar$').read().split("\n")
    del pro_text[-1]
    for jvm_msg in pro_text:
        app_name = jvm_msg.split(" ")[-1].split("/")
        app_name_tmp = app_name[-3]
        app_pid[app_name_tmp] = jvm_msg.split(" ")[0]
    print("本机 yarn 作业进程 -----> %s"%app_pid)
    return app_pid

if __name__ == '__main__':
    url = 'http://192.168.1.1:8088/cluster/apps/RUNNING'
    app_msg = js_app_msg(url)
    web_msg = slt_app(app_msg)
    local_yarn_job = yarn_job()
    slt_crim(web_msg,local_yarn_job)
    
