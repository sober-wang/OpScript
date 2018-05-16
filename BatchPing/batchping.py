#批量测试服务器联通性
#使用方式:
#python batchping.py 12 21 192.168.1
#
#作者：尚墨
#日期：2018年5月16日
#邮箱：ws1992jx@163.com

import os
import sys

start_v = int(sys.argv[1])
print("起始点：%d"%start_v)
end_v = int(sys.argv[2])
print("结束点：%d"%end_v)
ip_d = sys.argv[3]
print("IP地址端：%s"%ip_d)

for i in range(start_v,end_v):
	order = "ping -c 3 %s.%s"%(ip_d,str(i))
	os.system(order)