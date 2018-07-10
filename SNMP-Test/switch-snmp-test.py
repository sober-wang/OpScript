#!/usr/bin/env python
#-*- coding:utf-8 -*-
# 作用snmp v2连通性检测
# 作者：尚墨
# 邮箱：ws1992jx@163.com

from subprocess import *

public_orde = "snmpwalk -v 2c -c public "
orde = "%s 192.168.1."%public_orde

orde_set = []

def read_execut(orde_exe):
    p = Popen(orde_exe,shell=True,stdout=PIPE,stderr=STDOUT)
    r_text = p.stdout.readline()
    if "Timeout" in r_text:
        print("[ ERROR ] %s"%r_text)
    else:
        print("[ INFO ] connet sucessful")

for i in range(1,33,2):
    orde_exe = "%s%s"%(orde,i)
    orde_set.append(orde_exe)

orde_set.append("%s 192.168.1.254"%public_orde)
orde_set.append("%s 192.168.1.69"%public_orde)
orde_set.append("%s 192.168.1.96"%public_orde)

for o in orde_set:
    print(o)
    read_execut(o)
