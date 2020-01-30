seccomp 是 secure computing 的缩写，其是 Linux kernel 从2.6.23版本引入的一种简洁的 sandboxing 机制。在 Linux 系统里，大量的系统调用（system call）直接暴露给用户态程序。但是，并不是所有的系统调用都被需要，而且不安全的代码滥用系统调用会对系统造成安全威胁。seccomp安全机制能使一个进程进入到一种“安全”运行模式，该模式下的进程只能调用4种系统调用（system call），即 read(), write(), exit() 和 sigreturn()，否则进程便会被终止。  
参考文件：https://www.meiwen.com.cn/subject/vzbphftx.html  




LOOP：  
loop设备是一种伪设备，是使用文件来模拟块设备的一种技术，文件模拟成块设备后, 就像一个磁盘或光盘一样使用。在使用之前，一个 loop 设备必须要和一个文件进行连接。这种结合方式给用户提供了一个替代块特殊文件的接口。因此，如果这个文件包含有一个完整的文件系统，那么这个文件就可以像一个磁盘设备一样被 mount 起来。之所以叫loop设备（回环），其实是从文件系统这一层来考虑的，因为这种被 mount 起来的镜像文件它本身也包含有文件系统，通过loop设备把它mount起来，它就像是文件系统之上再绕了一圈的文件系统，所以称为 loop。  
原文链接：https://blog.csdn.net/lengye7/article/details/80247437