#-*- coding:utf-8 -*-
# 作用：多线程调用 hadoop fs shell 命令，清理 Spark 过期日志
# 作者：尚墨
# 日期：2018年10月29日
# 邮箱：ws1992jx@163.com

import os
import time
import re
import threading

def clean_log(path,day):
    '''
    path: 需要删除目录的地址
    '''
    smp.acquire()
    print("I will clean [ %s | %s ] "%(day,path))
    os.system("hdfs dfs -rm -r %s > /dev/null 2>&1"%path)
    smp.release()

def thread_start(path_list):
    '''
    path_list: 待删目录列表
    '''
    #t_list = []
    for d,p in path_list:
        # 创建线程
        t = threading.Thread(target=clean_log,args=(p,d,))
        # 启动线程
        t.start()

def select_dir(now_time,dir_path):
    '''
    now_time: 传入现在的时间
    return: 挑选出待删除的目录
    '''
    list_dir = os.popen('hdfs dfs -ls %s'%dir_path).read().split("\n")
    result_list = set()
    for i in list_dir[1:-2]:
        dir_msg = re.split("\s+",i)
        dir_date = time.strptime(dir_msg[5],"%Y-%m-%d")
        dir_stamp = int(time.mktime(dir_date))
        time_diff = now_time - dir_stamp
        # 定义删除时间区间
        five_day_stamp = 5 * 24 * 60 * 60
        if time_diff > five_day_stamp:
            result_list.add((dir_stamp,dir_msg[7]))
    return result_list
            
def main(conf_path):
    '''
    conf_path: 配置文件路径
    '''
    now_time = time.time()
    read_conf = open(conf_path,"r")
    for i in read_conf:
        wait_path = i.strip()
        print("Start Scan [ %s ]"%wait_path)
        wait_rm = select_dir(now_time,wait_path)
        thread_start(wait_rm)
    read_conf.close()

if __name__ == "__main__":
    BASE_PATH = os.path.dirname(os.path.abspath(__file__))
    conf_file = "clean.conf"
    conf_path = os.path.join(BASE_PATH,conf_file)
    # 定义线程信号量，用于充当线程池
    smp = threading.BoundedSemaphore(10)
    main(conf_path)

