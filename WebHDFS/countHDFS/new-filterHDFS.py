# 作用：Hive 库表统计，并用钉钉机器人发送消息，发送邮件附件格式 csv
# 作者：尚墨
# 日期：2018年12月4日
# 邮箱：ws1992jx@163.com

import configparser
import os
import requests
import json
import csv
import math
import re

import smtplib
from email.mime.text import MIMEText
from email.header import Header
from email.mime.multipart import MIMEMultipart



class get_fs_msg(object):
    def __init__(self,user_name,ip,port,avg_size,
                 robot_link,tlph_list,mail_host,
                 mail_user,mail_pass,receives,atta,mail_port):
        '''
        获取 HDFS hive 表目录下文件信息
        :param user_name: HDFS 用户名，建议传入 hdfs 作为用户
        :param ip: HttpFS ip地址
        :param port: HttpFS 端口
        :param avg_size: 文件平均大小
        :param robot_link: 钉钉机器人连接
        :param tlph_list: 钉钉联系人，传入 str,格式 12345678912,12345678912 如果只有一个电话号码以","结尾
        :param mail_host: smtp 服务器地址
        :param mail_user: 邮件用户
        :param mail_pass: 邮件密码
        :param receives: 接收方列表,传入 str,格式 12345678912,12345678912 如果只有一个电话号码以","结尾
        :param atta: 附件文件目录
        :param mail_port: smtp 端口号
        注：smtp 只使用 服务器端口号(加密)
        '''
        self.user_name = user_name
        self.ip = ip
        self.port = port
        self.iport = ":".join([self.ip, self.port])
        self.avg_size = int(avg_size)
        self.robot_lik = robot_link
        self.tlph_list = tlph_list.split(",")
        self.mail_host = mail_host
        self.mail_user = mail_user
        self.mail_pass = mail_pass
        self.receives = receives.split(",")
        self.atta = atta
        self.mail_port = mail_port

    def get_path_msg(self,abp):

        '''
        获取 HDFS 目录下的目录
        :param abp: 用户传入的路径
        :return: 路径列表 dir_list
        '''

        dir_list = []

        # 这段代码为了应对哪些狗屎表
        if re.search("^.+[\$\%\(\)].+$",abp) != None:
            print("[ ERROR ] 这条路径有问题 => %s"%abp)
            return dir_list

        url = "http://%s/webhdfs/v1%s?op=LISTSTATUS&user.name=%s" % (
            self.iport,
            abp,
            self.user_name
        )
        try:
            r = requests.get(url)
            dir_msg = json.loads(r.text,encoding="utf-8")["FileStatuses"]["FileStatus"]

            for dir_json in dir_msg:
                dir_name_mdftime = {}
                dir_name_mdftime["path"] =(
                    "/".join(
                        [
                            abp,
                            dir_json["pathSuffix"]
                        ]
                    )
                )
                dir_name_mdftime["accessTime"] = dir_json["accessTime"]
                dir_name_mdftime["modificationTime"] = dir_json["modificationTime"]
                dir_list.append(dir_name_mdftime)
                #print(json.dumps(dir_name_mdftime))

            return dir_list
        except:
            return dir_list

    def get_size(self, abp):

        '''
        获取目录大小，并计算目录下文件平均大小
        :param abp: 路径
        :return: 统计结果 字典 size
        '''

        size = {}

        # 这段代码为了应对哪些狗屎表
        if re.search("^.+[\$\%\(\)].+$",abp) != None:
            print("[ ERROR ] 这条路径有问题 => %s"%abp)
            return False

        url = "http://%s/webhdfs/v1%s?op=GETCONTENTSUMMARY&user.name=%s" % (
            self.iport,
            abp,
            self.user_name
        )
        r = requests.get(url)
        size["path"] = abp
        size["fileSize"] = r.json()["ContentSummary"]["length"]
        size["spaceConsumed"] = r.json()["ContentSummary"]["spaceConsumed"]
        size["fileCount"] = r.json()["ContentSummary"]["fileCount"]
        if size["fileCount"] == 0:
            size["averageSize"] = 0
        else:
            size["averageSize"] = round(
                size["fileSize"] / size["fileCount"] /
                1024 / 102,2)

        if size["averageSize"] < self.avg_size :
            return  size
        else:
            return False

    def send_DingTalk(self,msg):
        '''
        作用： 调用 钉钉的机器人发送信息
        :param msg: 钉钉 消息内容
        :return:
        '''
        headers = {'Content-Type': 'application/json;charset=utf-8'}
        json_text = {
            "msgtype": "text",
            "at": {
                "atMobiles": self.tlph_list,
                "isAtAll": False
            },
            "text": {
                "content": msg
            }
        }

        requests.post(self.robot_lik, json.dumps(json_text), headers=headers).content

    def send_mail(self,msg):
        '''
        作用：发送邮件 ，附件格式:csv
        :param msg: 统计表目录文件结果信息
        :return:
        '''
        mini_file = os.path.join(self.atta,"Hive-MoreMiniFile.csv")
        null_table = os.path.join(self.atta,"Null-table.csv")

        # 制作 小文件 csv 文件
        csv_headers = ["path", "fileSize", "spaceConsumed", "fileCount", "averageSize"]
        with open(mini_file,"w",newline="",encoding="utf-8") as f:
            write_head = csv.DictWriter(f,csv_headers)
            write_head.writeheader()
            for m in msg:
                if m == False or m["fileCount"] == 0:
                    continue
                else:
                    m["fileSize"] = self.unit_count(m["fileSize"])
                    m["spaceConsumed"] = self.unit_count(m["spaceConsumed"])
                    write_head.writerow(m)

        # 制作 空表统计 csv 文件
        with open(null_table,"w",newline="",encoding="utf-8") as null_f:
            write_head = csv.DictWriter(null_f,csv_headers)
            write_head.writeheader()
            for null_m in msg:
                if null_m == False:
                    continue
                elif null_m["fileCount"] == 0:
                    write_head.writerow(null_m)

        message = MIMEMultipart()
        # 发件方地址
        message["From"] = "{}".format(self.mail_user)
        # 收件方地址
        message["To"] = ",".join(self.receives)
        # 邮件的标题
        message["Subject"] = Header(
            """
            Hive 库表文件统计
            """,
             "utf-8")
        # 邮件中正文内容
        message.attach(MIMEText(
            """
            Hive-MoreMiniFile.csv 统计 Hive 库表中文件小于4MB，文件中统计的库表请尽快合并\n
            Null-table.csv 统计 Hive 中的空表，如果不需要请删除\n
            表头说明：\n
            path: HDFS路径\n
            fileSize: 目录大小\n
            spaceConsunmed: 所占空间大小\n
            fileCount: 目录中文件个数\n
            averageSize: 目录中文件平均大小
            ""","plain","utf-8"))

        # 创建 小文件统计 附件
        # comma-separated-values：csv 格式编码，参阅：https://www.cnblogs.com/doNetTom/p/4277182.html
        att = MIMEText(open(mini_file,"rb").read(),"comma-separated-values","utf-8")
        att["Content-Type"] = "application/comma-separated-values"
        att["Content-Disposition"] = "attachment;filename=Hive-MoreMiniFile.csv"
        message.attach(att)

        # 创建 空表统计 csv 附件
        att1 = MIMEText(open(null_table,"rb").read(),"comma-separated-values","utf-8")
        att1["Content-Type"] = "application/comma-separated-values"
        att1["Content-Disposition"] = "attachment;filename=Null-table.csv"
        message.attach(att1)

        try:
            smtp_obj = smtplib.SMTP_SSL(self.mail_host,self.mail_port)
            smtp_obj.login(self.mail_user, self.mail_pass)
            smtp_obj.sendmail(self.mail_user, self.receives, message.as_string())
            print("[ INFO ]  Hive库表 文件统计邮件发送成功")
        except smtplib.SMTPException as e:
            print("[ ERROR ] Hive库表 文件统计邮件发送失败 %s => 类 get_fs_msg() 函数 send_mail()" % e)

    def unit_count(self,number):
        '''
        作用：单位转换，将字节信息转换成容易识别的信息
        number: 字节数
        :return: 返回 易于识别的大小;INFO 数据类型：str;ERROR 数据类型：int
        '''
        lst = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', "ZB"]
        try:
            # 舍弃小数点，取小 求对数(对数：若 a**b = N 则 b 叫做以 a 为底 N 的对数)
            i = int(math.floor(math.log(number, 1024)))
            if i >= len(lst):
                i = len(lst) - 1
            result = ('%d' + " " + lst[i]) % (number / math.pow(1024, i))
            return result
        except Exception as e:
            print("[ ERROR ] 类 get_fs_msg() 函数 unit_count() %s"%e)
            return 0

def main():
    BASE_PATH = os.path.dirname(os.path.abspath(__file__))
    conf_file = "hive-bigtable-site.ini"
    conf_path = os.path.join(BASE_PATH, conf_file)

    conf = configparser.ConfigParser()
    conf.read(conf_path,encoding="utf-8")

    # webhdfs 配置
    user = conf["WebHDFS"]["username"]
    ip = conf["WebHDFS"]["ip"]
    port = conf["WebHDFS"]["port"]
    avg_size = conf["WebHDFS"]["avgSize"]

    # 钉钉机器人配置
    robot_link = conf["DingTalk"]["RobotLink"]
    tlph_list = conf["DingTalk"]["TelephoneList"]

    # 邮箱配置
    smtp_server = conf["E-mail"]["host"]
    mail_user = conf["E-mail"]["user"]
    mail_pass = conf["E-mail"]["pass"]
    receives = conf["E-mail"]["receivers"]
    atta = conf["E-mail"]["atta"]
    mail_port = conf["E-mail"]["port"]

    db_list = conf["WebHDFS"]["HiveDBList"].split(",")
    get_fs = get_fs_msg(user,ip,port,avg_size,
                        robot_link,tlph_list,smtp_server,
                        mail_user,mail_pass,receives,atta,mail_port)
    # 数据库表目录列表
    table_path = []
    for db in  db_list:
        tmp = "/user/hive/warehouse/{}.db".format(db)
        table_path = get_fs.get_path_msg(tmp) + table_path

    # 数据库分区(pt)列表获取
    pt_path = []
    for table in  table_path:
        if len(get_fs.get_path_msg(table["path"])) == 0:
            pt_path.append(table)
            print("[ INFO ] 这是没有分区的表 => %s"%json.dumps(table))
        else:
            have_pt = get_fs.get_path_msg(table["path"])
            pt_path = have_pt + pt_path
            print("[ INFO ] 这是有pt的表 => %s"%json.dumps(table))


    # 数据库表统计列表
    count_list = []
    for pt in pt_path:
        count_table = get_fs.get_size(pt["path"])
        count_list.append(count_table)

    get_fs.send_mail(count_list)

    DingTalk_msg = " Hive 文件统计已发邮件，请查收！"
    get_fs.send_DingTalk(DingTalk_msg)



if __name__ == "__main__":
    main()



