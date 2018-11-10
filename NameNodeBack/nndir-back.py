#!/usr/bin/env python
#-*- coding:utf-8 -*-
# 作用：利用 paramiko 库备份 NameNode 目录
# 作者：尚墨
# 日期：2018年11月10日
# 邮箱：ws1992jx@163.com

import paramiko
import os
import configparser
import time

class nndbk(object):
    '''
    备份 HDFS NameNode 目录
    '''
    def __init__(self,username,passwd,hostname,srcdir,destdir,port=22):
        '''
        :param username: 服务器用户名
        :param passwd: 密码
        :param hostname: 服务器 ip 地址
        :param port: ssh 端口
        :param srcdir: NameNode nn 目录路径
        :param destdir: 备份临时路径或备份目标服务器路径
        '''
        self.user = username
        self.passwd = passwd
        self.hostname = hostname
        self.port = port
        self.srcdir = srcdir
        self.destdir = destdir
        self.dt = time.strftime("%Y%m%d",time.localtime())

    def ssh_obj(self,code):
        '''
        执行命令函数
        :param code: 需要执行的命令
        :return: True or False
        '''
        ssh = paramiko.SSHClient()
        ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy)
        try:
            ssh.connect(hostname=self.hostname,
                        username=self.user,
                        password=self.passwd,
                        port=self.port
            )

            stdin,stdout,stderr = ssh.exec_command(code)
            print(stdout.read())
            ssh.close()
            return True
        except Exception as e:
            print("ERROR : %s "%e)
            ssh.close()
            return False

    def tar_nn(self):
        '''
        :return: 生成打包命令
        '''
        print(self.dt)
        # 服务器上必须存在 /home/ops 目录
        tar_name = "{}{}{}".format("/home/ops/nndbk-",self.dt,".tar.gz")
        print(tar_name)
        tar_code = "tar czvf {} {}".format(tar_name,self.srcdir)
        print(tar_code)
        return tar_name,tar_code

    def get_file(self,file_path):
        '''
        下载文件
        :param file_path: 目标服务器文件绝对路径
        :return:
        '''
        back_file_path = "%s/%s"%(self.destdir,file_path.split("/")[-1])
        print(back_file_path)
        try:
            sftp_msg = paramiko.Transport((self.hostname,self.port))
            sftp_msg.connect(username=self.user,password=self.passwd)
            sftp_cont = paramiko.SFTPClient.from_transport(sftp_msg)
            sftp_cont.get(file_path,back_file_path)
            sftp_msg.close()
            return True,back_file_path
        except Exception as e:
            print("ERROR get_file : %s"%e)
            sftp_msg.close()
            return False

    def put_file(self,file_path,decthost=None):
        '''
        :param file_path: 需要上传的文件
        :return: True or False
        '''
        if decthost == None:
            decthost = self.hostname
        else:
            pass
        try:
            sftp_msg = paramiko.Transport((decthost,self.port))
            sftp_msg.connect(username=self.user,password=self.passwd)
            sftp_cont = paramiko.SFTPClient.from_transport(sftp_msg)
            # file_path : 需要上传的文件，self.destdir: 目标服务器路径
            sftp_cont.put(file_path,self.destdir)
            sftp_msg.close()
            return True
        except Exception as e:
            print("ERROR : %s"%e)
            sftp_msg.close()
            return False

if __name__ == "__main__":
    BASE_PATH = os.path.dirname(os.path.abspath(__file__))
    config_file = "/".join([BASE_PATH,"nndbk.ini"])
    client_control = 123
    conf = configparser.ConfigParser()
    conf.read(config_file)
    HOSTNAME = conf["host"]["hostname"]
    USERNAME = conf["host"]["username"]
    PASSWORD = conf["host"]["password"]
    SRCDIR = conf["dir"]["src"]
    DESTDIR = conf["dir"]["dest"]
    active_back = nndbk(USERNAME,PASSWORD,HOSTNAME,SRCDIR,DESTDIR)
    file_path,code = active_back.tar_nn()
    active_back.ssh_obj(code)
    n,back_path = active_back.get_file(file_path)
    active_back.put_file(back_path,decthost="192.168.1.1")
