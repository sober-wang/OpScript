#-*- coding:utf-8 -*-
#作用：分发部署包
#   尚未实现多进程防阻塞方式分发部署包。
#   需要分发服务器的ssh公钥部署在被分发部署包的服务器上
#   使用格式：
#           脚本名称    被分发包的路径     目的路径
#   python DispBag.py   /sourcePath   /targetPath
#作者：尚墨
#邮箱：ws1992jx@163.com
#时间：2018年3月26日

import configparser
import sys
import os

def ergIp(ipGroup,sourceDir,tagDir):
    for ipValue in ipGroup:
        os.system("scp -r " + sourceDir + " root@" + ipValue + ":" + tagDir)

if __name__ == "__main__":
    filePath = os.path.dirname(__file__)
    conFile = '%s/ip.ini'%filePath
    conf = configparser.ConfigParser()
    conf.read(conFile)
    ipList = conf["ipvalue"]["ip"].split(",")
    sourceDir = sys.argv[1]
    tagDir = sys.argv[2]
    ergIp(ipList,sourceDir,tagDir)

