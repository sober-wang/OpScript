#!/bin/bash
#作用：升级glibc
#作者：尚墨
#邮箱：ws1992jx@163.com
#日期：2018年3月29日

gcbag=$1
gcpath=$2

tar xzvf $gcbag -C $gcpath
gcdir=$gcpath`ls $gcpath| grep glibc`

cd $gcdir
mkdir $gcdir/build
cd $gcdir/build
../configure --prefix=/usr --disable-profile --enable-add-nos --with-headers=/usr/include --with-binutils=/usr/bin/ --disable-sanity-checks
make && make install
