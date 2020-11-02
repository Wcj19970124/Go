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
    Ⅴ、strings:主要提供一些字符串的基本操作方法，例如:
                
