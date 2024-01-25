import os
import mmap
import multiprocessing

def child_process():
    # 打开共享内存
    shared_memory_name = "shared_memory"
    shared_memory = mmap.mmap(0, 1024, tagname=shared_memory_name)

    # 从共享内存中读取数据
    data = shared_memory.read().decode()
    print("Received data from Go parent process:", data)

    # 向共享内存中写入数据
    shared_memory.seek(0)
    shared_memory.write(b"Python child process")

    # 关闭共享内存
    shared_memory.close()

if __name__ == "__main__":
    child_process()