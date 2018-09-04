#!/bin/bash
# 作用：备份 CDH NameNode 元数据
# 作者：尚墨
# 日期：2018年9月4日
# 邮箱：ws1992jx@163.com

hdfs dfsadmin -fetchImage /home/ops/namenode-back/namenode.backup-`date +%F`
