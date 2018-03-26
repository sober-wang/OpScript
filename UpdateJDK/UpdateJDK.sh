#!/bin/bash
#作用升级更新JDK
#作者：尚墨
#邮箱：ws1992jx@163.com
#日期：2018年3月26日

echo "查看当前rpm jdk包"
rpm -qa | grep java
echo "删除openJDK"
rpm -qa | grep java | xargs rpm -e --nodeps

#解压jdk,tar包
tar xzvf $1 -C $2
#获取JAVA_HOME路径
jdkv=`ls $2 | grep jdk`
echo $jdkv
	
#添加环境变量
echo "export JAVA_HOME=$2/$jdkv" >> /etc/profile
echo "export CLASSPATH=\$JAVA_HOME/lib/dt.jar:\$JAVA_HOME/lib/tools.jar:\$JAVA_HOME/jre/lib/rt.jar" >> /etc/profile
echo "export PATH=\$PATH:\$JAVA_HOME/bin" >> /etc/profile

#使新jdk生效
source /etc/profile
java -version