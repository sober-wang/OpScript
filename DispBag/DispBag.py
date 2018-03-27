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
from multiprocessing import Pool


def ergScp(ipv,sDir,tDir):
    '''
    job函数，执行拷贝命令
    '''
    orderShell = "scp -r %s root@%s:%s"%(sDir,ipv,tDir)
    print(orderShell)
    #os.system(orderShell)

if __name__ == "__main__":
    filePath = os.path.dirname(__file__)
    conFile = '%s/ip.ini' % filePath
    conf = configparser.ConfigParser()
    conf.read(conFile)
    ipList = conf["ipvalue"]["ip"].split(",")
    sourceDir = sys.argv[1]
    tagDir = sys.argv[2]
    p = Pool(processes = 5)
    for ipValue in ipList:
        p.apply_async(ergScp,(ipValue,sourceDir,tagDir))
    p.close()
    p.join()


