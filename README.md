# Go
   日常学习，提交Go相关的学习代码，包括以下内容：<br/>
    Ⅰ、Go 基础 <br/>
    Ⅱ、Go Web知识  <br/>
    Ⅲ、Go 相关框架  <br/>
    当前可能只有代码，没有具体的文档细则，不过代码中也有大量注释，可以作为相关学习的资料，之后有时间，会详细的整理一份每一部分的细则文档 <br/>
### Go web
   概述：主要介绍Go web相关的知识，含有大量示例，基本上全部使用的是标准库，主要依据《Go Web编程》一书
    前六章分别介绍：<br/>
                Ⅰ、Go web的基本流程 <br/>
                Ⅱ、net/http包的使用：处理器和处理器函数,接收请求(Request的使用),响应请求(Response的使用，严格说这里时ResponseWriter),Cookie的获取和设置 <br/>
                Ⅲ、模板引擎：主要是标准库中的text/template,html/template <br/>
                Ⅳ、数据存储：主要分为内存存储、文件存储(csv,gob),数据库存储(net/sql)<br>
                Ⅴ、Go web的基本工具：xml和json的编码和解码(encoding/xml,encoding/json),这里写了一个小的Go web服务,学生信息管理的Demo，涉及到前面所学习的知识的80%<br>
                Ⅵ、Go的测试相关：单元测试，基准测试，其它 <br/>
                Ⅶ、并发和并行方面的知识,包括goroutine,channel  <br/>
### Go Source Code
   概述：主要是阅读Golang的一些常用标准库，对常用标准库内的导出函数的使用进行剖析 <br/>
    Ⅰ、fmt：主要介绍fmt包下的常用的输入输出，包含(Scan()类型,Print()类型,Errorf()类型) <br/>
    Ⅱ、time：主要是时间、日期相关的操作，下面主要有几个重要的go文件，time文件,format文件,tick文件.主要需要掌握日期的格式化,Tick类似于定时器的操作 <br/>
    Ⅲ、log：Go的日志包短小而精悍，并且使用互斥锁的机制保证了其并发安全，而且可以利用log.Logger自定义日志级别，满足自己的需求 <br/>
    Ⅳ、strconv：主要的作用就是用于将各种类型和字符串类型进行相互的转换  <br/>
    Ⅴ、strings：主要提供一些字符串的基本操作方法，例如:大小写的转换，去空格，子串的匹配，字符串字符的替换，字符串转字符切片，字符切片转字符串，字符串转为*io.Reader等等 <br/>
    Ⅵ、ioutil：主要是提供了对io操作的封装,例如：文件的读取、写入、创建等等 <br/>
    Ⅶ、os:主要提供对于底层os操作的封装，内部包含对于os的基本操作(环境变量的获取设置等等),FileMode对象(表示文件的模式：文件/目录,权限),FileInfo对象(文件对象，可以进行基本的权限操作，但是无法进行读写),File对象(表示已打开的文件对象，可以用于io操作，例如:Create(),Open(),Close(),Write()等),另外还包含对于进程的操作，进程的获取启动等  <br/>
    Ⅷ、bufio:提供缓冲式的io,主要针对io.Reader和io.Writer进行包装,获取带有缓冲的io.Reader和io.Writer，然后可以利用io.Reader将缓冲中的数据读取出来，也可以利用io.Writer将数据写入到缓冲中 <br/> 
    Ⅸ、net/http,net/url:前者主要是用于建立服务器和客户端，最主要的是注意里面的两个struct：request,response/responseWriter,前者用于对请求消息进行处理,后者用于对响应消息津行处理,除此之外,还有一些服务器的设置,监听,多路复用器分发请求,处理器,处理器函数等,见Go Web中的第二节;net/url包主要是针对url进行处理,包括对参数的转码便于安全的用于url中,解析url为一个结构体便于获取url中的详细信息,提供了url.Values{}对url中的queryString进行处理 <br/>
    Ⅹ、math:提供一些基本的数学操作,例如两者的最大值,最小值,取余,绝对值,二次方根，三次方根等等,还有一些对复数，大数的操作，这一部分可以自行看一下标准库的中文文档。另外math包中还提供的有一个随机数的包rand <br/>

   不常用的包: <br/>
    Ⅰ、unsafe:主要提供一些跳过go安全类型检查的方法,主要有三个,SizeOf()获取类型所占内存的大小,Offset()获取字段的偏移量,AlignOf()获取类型在内存中的对齐方式 <br/>
    Ⅱ、archive/tar：主要提供对于tar文件的存取操作,基本步骤都是先获取tar的读取器或者写入器,然后利用读取器或者写入器进行操作即可
        archive/zip:主要提供对于zip文件的存取操作，基本操作和上述的对于tar文件的操作类似 <br/>
    Ⅲ、builtin：主要是golang中的预定义标识符，这里并没有提供具体的实现，只是用来方便生成godoc文档，不过还是要提及一下里面涉及的函数，比较需要区分的就是make()和new()函数，前者返回的是一个参数类型的具体类型值，而后者返回的是一个指向类型的指针，另外还提供了一些常用的函数，例如cap(),len(),append(),delete(),copy(),close()等 <br/>
    Ⅳ、bytes：主要提供对于[]byte的操作,函数和方法和strings基本一致 <br/>
    Ⅴ,  compress/gzip:主要用于对gzip格式的压缩文件进行读写操作,实际上,基本所有的对于压缩文件的处理都是非常类似的: <br/>
        几个比较重要的文件操作:(1)创建文件,os.Create();(2)打开文件,os.Open();(3)获取文件头信息,file.Stat()/os.Stat(file) 
                             (4)读取文件,file.Read();(5)写入文件,file.Write();其中(4)(5)可以合并使用io.Copy()直接进行复制
                              另外需要注意,文件头信息中的Name是文件名而不是文件的路径+文件名,因此有时候需要对其进行处理
        其它的几个格式的压缩文件的操作基本都一致,了解即可<br/>
### WebSocket
   概述：主要是介绍一下golang的WebSocket实现,因为业务中经常会遇到消息推送类似的需求,而实现消息推送,使用WebSocket是比较常见的,不过golang提供的WebSocket好像不是很好用,所以这里选用github.com/gorilla/websocket这个实现<br/>