#-*- coding:utf-8 -*-
# 作用: 计算需要参数下每个目录中 文件的平均大小，输出平均值小于 4M 的目录
# 使用方式： python new-filterHDFS.py 或者 python new-filterHDFS.py /user/hive/warehous
# 作者：尚墨
# 日期：2018年11月2日
# 邮箱：ws1992jx@163.com

import configparser
import os
import requests
import sys
import time


class get_fs_msg(object):
    def __init__(self, user_name,start_path,filter_size, ip, port=14000):
        self.user_name = user_name
        self.ip = ip
        self.port = port
        self.start_path = start_path
        self.iport = ":".join([self.ip, self.port])
        self.filter_size = filter_size

    def get_path_msg(self, abp=""):

        '''
        获取 HDFS 目录下的目录
        :param abp: 用户传入的路径
        :return: 路径列表
        '''

        if abp == "":
            abp = self.start_path

        dir_name_list = set()
        url = "http://%s/webhdfs/v1%s?op=LISTSTATUS&user.name=%s" % (
            self.iport,
            abp,
            self.user_name
        )
        r = requests.get(url)
        dir_msg = r.json()["FileStatuses"]["FileStatus"]
        for dir_json in dir_msg:
            dir_name_list.add(
                "/".join(
                    [
                        self.start_path,
                        dir_json["pathSuffix"]
                    ]
                )
            )
        return dir_name_list

    def get_size(self, abp=None):

        '''
        获取目录大小
        :param abp: 路径
        :return:
        '''

        size = {}

        if abp == None:
            abs_path = self.start_path
        else:
            abs_path = abp

        url = "http://%s/webhdfs/v1%s?op=GETCONTENTSUMMARY&user.name=%s" % (
            self.iport,
            abs_path,
            self.user_name
        )
        r = requests.get(url)
        size["path"] = abs_path
        size["fileSize"] = r.json()["ContentSummary"]["length"]
        size["spaceConsumed"] = r.json()["ContentSummary"]["spaceConsumed"]
        size["fileCount"] = r.json()["ContentSummary"]["fileCount"]
        if size["fileCount"] == 0:
            size["averageSize"] = 0
        else:
            size["averageSize"] = round(
                size["fileSize"] / size["fileCount"] /
                1024 / 1024
                ,2)
        return size

    def filter_avg(self,path_list):
        for p in path_list:
            table_msg = self.get_size(p)
            if table_msg["averageSize"] < self.filter_size:
                print(table_msg)
#                print("这条信息不合规")
#            else:
#                print(table_msg)

def main(hdfs_path=None):
	
	'''
	hdfs_path: 执行脚本时传入的参数
	'''

    BASE_PATH = os.path.dirname(os.path.abspath(__file__))
    conf_file = "hive-bigtable-site.ini"
    conf_path = os.path.join(BASE_PATH, conf_file)

    conf = configparser.ConfigParser()
    conf.read(conf_path)

    if hdfs_path== None:
        first_path = conf["WebHDFS"]["hivepath"]
    else:
        first_path = hdfs_path

    user_name = conf["WebHDFS"]["username"]
    wb_hdfs_ip = conf["WebHDFS"]["ip"]
    wb_hdfs_port = conf["WebHDFS"]["port"]
    filter_size = int(conf["WebHDFS"]["filtersize"])

    get_fs = get_fs_msg(
			user_name,
			first_path,
			filter_size,
			wb_hdfs_ip, 
			port=wb_hdfs_port
			)
    path_list = get_fs.get_path_msg()
    get_fs.filter_avg(path_list)


if __name__ == "__main__":
    try:
        abs_path = sys.argv[1]
        main(abs_path)
    except :
        print("《《《《《《《《《《《《《《      未传入路径参数,将使用配置文件中的路径       》》》》》》》》》》》》》》")
        time.sleep(1)
        main()


