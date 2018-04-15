#作用：生成文件列表
#研发和测试往往会给你写半成品工具，这个时候你不如自己手动写一个
#作者:尚墨
#邮箱：ws1992jx@163.com
#日期：2018年4月15日

import os
import configparser

def printFL(filepath,writfile):
    '''
    :param filepath:文件集合路径
    :param writpath: 生成文件列表路径
    :return:
    :print:干活的时候总得让人看见吧，打印你都干了啥
    '''
    filelist = os.listdir(filepath)
    wf = open(writfile,'a')
    for f in filelist:
        '''
        网上有人测试，较少的字符串拼接时，使用+号会比，join等方法都快。
        这里使用join是为了熟练一种大规模环境下字符串拼接最快的方法
        '''
        every_fp = "".join([filepath,'/',f,"\n"])
        wf.write(every_fp)
    wf.close()

if __name__ == "__main__":
    BASE_PATH = os.path.dirname(os.path.abspath(__file__))
    conf_path = "".join([BASE_PATH,'/file.conf'])
    conf = configparser.ConfigParser()
    #read()方法中加入encoding是为了识别中文
    conf.read(conf_path,encoding="utf-8-sig")

    fp = conf["filepath"]["fp"]
    wp = conf["listoutpath"]["ltp"]
    printFL(fp,wp)