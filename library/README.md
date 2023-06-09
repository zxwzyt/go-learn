### Golang Package 介绍

* archive/tar 和 /zip-compress: 压缩(解压缩)文件功能
* fmt-io-bufio-path/filepath-flag:
  * fmt: 提供了格式化输入输出功能
  * io: 提供了基本输入输出功能，大多数是围绕系统功能的封装
  * bufio: 缓冲输入输出功能的封装
  * path/filepath: 用来操作在当前系统中的目标文件名路径
  * flag:对命令行参数的操作

* strings-strconv-unicode-regexp-bytes:
  * strings: 提供对字符串的操作
  * strconv: 提供将字符串转换为基础类型的功能
  * unicode: 为unicode型的字符串提供特殊的功能
  * regexp: 正则表达式功能
  * bytes: 体工队字符型切片的操作
  * index/suffixarray: 字符串快速查询

* math-match/cmatch-math/big-math/rand-sort:
  * math: 基本的数字函数
  * math/cmath: 对复数的操作
  * math/rand: 伪随机数生成
  * sort: 为数组排序和自定义集合
  * math/big: 大叔的实现和计算
  
* container-/list-ring-heap: 实现对集合的操作
  * list: 双链表
  * ring: 环形链表

* builtin: 描述了允许godoc提供语言标识符的文档
* compress/bzip2-flate-gzip-lzw-zlib
  * compress/bzip2: 实现了bzip2的解压缩
  * flate: 实现了DEFLATE的数据压缩格式，入RFC 1951中所述
  * gzip: 实现了gzip格式压缩文件的读写
  * lzw: 实现了Lempel-Ziv-Welch压缩数据格式
  * zlib: 实现了zlib数据压缩格式的读写

* context: 定义了Context类型，它跨API边界和进程之间携带截止日期，取消信号和其他请求范围的值
* crypto-crypto/md
  * crypto: 创建的加密常量的集合
  * crypto/md5: md5加密
  * crypto/sha1: sha1加密

* errors: 实现了操作错误的方法
* expvar: 为通用变量提供标准化接口
* hash: 所有散列函数实现的通用接口
* html: HTML文本进行转码和转义的功能
* image-color-draw-gif-jpeg-png
  * imgae:实现基本的二维图形库
  * image/color: 颜色基本库
  * image/draw: 提供图像合成功能
  * image/gif: gif图像解码器和编码器
  * image/jpeg:
  * image/png

* sort: 提供了用于排序切片和用户定义集合的原函数
* unsafe: 包含了一些打破go语言"类型安全"的命令，一般的程序中不会被使用，可用在C/C++程序的调用中
* syscall-os-os/exec:
  * os: 提供给我们一个平台无惯性的操作系统功能接口，采用类unix设计， 隐藏了不同操作系统间差异，让不同的文件系统和操作系统对象表现一致
  * os/exec: 提供我们运行外部操作系统命令和程序的方式
  * syscall: 底层的外部包，提供了操作系统底层调用的基本接口

* time-log:
  * time: 日期和时间的基本操作
  * log: 记录程序运行时产生的日志

* encoding/json-encoding/xml-text/template:
  * encoding/json: 读取并解码和写入并解码JSON数据
  * encoding/xml:简单的xml1.0解析器
  * text/template: 生成像HTML一样的数据与文本混合的数据驱动

* net-net/http:
  * net: 网络数据的基本操作
  * http: 提供了一个可扩展的HTTP服务器和基础客户端，解析HTTP请求和回复

* runtime: GO程序运行时的交互操作, 例如垃圾回收和写成操作
* reflect: 实现通过程序运行时反射，让程序操作任意类型的变量
* sql: 为数据库提供围绕SQL或类SQL的通用接口