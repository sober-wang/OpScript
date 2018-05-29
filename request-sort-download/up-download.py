# 多线程并发请求数据，通过线程Queue的方式实现请求与写文件分离
# 你可以针对提取url的函数做修改，也可以以导入方式导入请求和下载函数。
# 作者：尚墨
# 日期：2018年5月29日
# 邮箱：ws1992jx@163.com

from urllib import request
import os
import sys
import json
import threading
import queue
import time

def write_file():
	'''
	写文件的函数
	通过smp2 信号量作为线程锁
	'''
	smp2.acquire()
	if thread_q.empty():
		print("[ ERROR ] 队列空了")
	else:
		file_name,data = thread_q.get()
		print(file_name)
		with open(file_name,"wb") as f:
			f.write(data)
			print("[ INFO ] 正在写入文件 [ %s ]"%file_name)
	smp2.release()

def run_write(numb):
	run_list = []

	#print("This is sum Threading %d"%len(numb))
	
	for i in numb:
		run_write_t = threading.Thread(target=write_file,args=())
		run_write_t.start()
		run_list.append(run_write_t)
	return run_list

def get_file(url,file_name):
	'''
	下载请求函数函数
	通过smp 信号量标示作为线程锁
	'''
	smp.acquire()
	try:
		with request.urlopen(url) as f:
			data = f.read()
			data_tmp = (file_name,data)
			ctlng_file_siz = float(f.getheaders()[3][1])/1024/1024
			#for k,v in f.getheaders():
			# 	print("%s:%s"%(k,v))
			if ctlng_file_siz > 3.0:
				#将请求到的数据放入队列
				thread_q.put(data_tmp)
			else:
				print("[ ERROR ] 文件大小不足3M，忽略下载")
	except Exception as e:
		print("[ ERROR ] -------------------------文件下载失败：%s----------------------"%e)
	smp.release()


def download_data(url_list,dl_path):
	'''
	请求并发函数，生成多线程并发
	'''
	run_list = []

	if os.path.exists(dl_path) is False:
		os.mkdir(dl_path)

	for url in url_list:
		
		file_name = os.path.join(dl_path,url.split("/")[-1])
		
		if os.path.exists(file_name) is True:
			print("[ ERROR ] ----------------------文件已存在---------------------")
			#time.sleep(10)
			continue
		else:
			#创建请求进程
			run_t = threading.Thread(target=get_file,args=(url,file_name,))
			run_t.start()
			run_list.append(run_t)
	print("总线程数 ： %d"%len(run_list))
	return run_list

def scan_urllist_file(file_path):
	file_url_list = []

	#从文件中提取url链接
	with open(file_path,"r") as file_object:
		for json_line in file_object:
			json_line = json_line.strip()
			try:
				file_url_b = json.loads(json_line)["file_url_b"]
				file_url_a = json.loads(json_line)["file_url_a"]
				file_url_list.append(file_url_b)
				file_url_list.append(file_url_a)
			except Exception as e:
				print("[ ERROR ]------------------ 这个 [%s] 信息为空----------------"%e)
	return file_url_list

def wait_thread_stop(download_thread,write_thread):
	for dl_thread in download_thread:
		dl_thread.join()
	for rt_thread in write_thread:
		rt_thread.join()

if __name__ == '__main__':
	old_time = time.time()
	
	file_path = sys.argv[1]
	download_dir = sys.argv[2]
	
	smp = threading.BoundedSemaphore(20)
	smp2 = threading.BoundedSemaphore(10)
	thread_q = queue.Queue(maxsize=1000)
	
	url_list = scan_urllist_file(file_path)
	
	print("start run download_yy")
	run_download_list = download_data(url_list,download_dir)
	print("Start run write file")
	run_write_sort = run_write(run_download_list)
	wait_thread_stop(run_download_list,run_write_sort)

	now_time = time.time()-old_time
	print("总共耗时：%.2f 秒"%now_time)