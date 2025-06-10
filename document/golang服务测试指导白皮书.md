golang服务测试指导白皮书

# 背景
百度APP服务端计划2020年完成50%机器的PHP服务迁移golang，这就意味着，从2020年开始，将有越来越多的go服务提测，由于RD和QA都没有丰富的golang开发和测试经验，所以，2020年golang服务的线上质量就面临了严峻挑战

[百度APP服务端go服务迁移规划&进展](https://wiki2ku.baidu-int.com/pubapi/urlmap?id=1045917785)

# 目标
基于上面提到的背景和挑战，需要通过调研、分析和整理厂内&厂外的golang服务测试经验、工具和方法，并结合百度APP服务端业务现状，产出并持续维护一份golang服务测试指导白皮书

# 相关调研
## golang与PHP语言层面的区别

## [golang与PHP的利弊](./php%20vs%20go.md)
## golang与PHP服务测试层面的区别


## golang相关测试工具链
    [Go服务测试工具链路](https://wiki2ku.baidu-int.com/pubapi/urlmap?id=996562610)
    [golang相关工具链](https://wiki2ku.baidu-int.com/pubapi/urlmap?id=1037843373)

## 厂内其他部门测试经验和方法【持续补充】
    

### BFE：使用的工具主要来自go tool，常规需求以功能测试为主
    * 单测：go test，RD负责
    * 功能测试：常规方法，语言无关
    * 性能测试：常规方法，语言无关
    * CI：基于agile+脚本等实现
    * 自动化：基于go test自研的框架

### 贴吧：常规需求以功能测试为主，迁移类需求，会增加线上流量diff和性能测试
    * 单测：没做
    * 功能测试：常规方法，语言无关
    * 性能测试：常规方法，语言无关
    * CI：agile+脚本实现
    * 自动化：基于robotframework

### 网盘：常规需求以功能测试为主，部分模块会采用线上引流进行性能测试
    * 单测：go test，RD负责，QA在调研自动生成单测case的方法，期望接入RD开发工具链中
    * 功能测试：常规方法，语言无关
    * 性能测试：部分模块采用线上引流（压力云的引流方案）进行
    * CI：agile接入bugbye，bugbye定制化了一些扫描规则
    * 其他：自研的CR辅助工具

## 调研结论

### PHP与GO的差异
    * Go是静态编译型语言，PHP动态解释型语言，go每次修改后都需要重新编译，PHP不需要，go需要在初始化变量和对象时提前想清楚。在PHP中，你永远不会初始化变量，需要时当场使用就可以了。
    * Go 性能更优，绝大多数目前的服务可以节约 60% 以上的资源
    * Go的开发速度很快，测试运行速度更快、内存使用效率更高、CPU使用率更低。
    * 由于Golang内置的错误检查机制，由于开发人员疏忽而出现漏洞的可能性非常低。PHP经常会出现因为低级的语法错误或是不规范导致的定位问题耗时
    * Go 语言生态要比 PHP 丰富，而且很多是官方提供，还有Google的强大支持，规范且有保障，PHP的生态比较散
    * Go的代码更规范和统一，因为有官方的格式化工具
    * Go需要程序员考虑的更多，比如json解析，PHP只需decode后，就可以随便访问了，go需要先定一个结构体，需要明确json中的字段及类型，但是结构明确，后期维护效率高，也更加安全
    * Go原生支持并发，开发并发代码方便快捷，PHP需要通过扩展实现，比如phaster
    * Go语言的语法简单，语法糖并不多，容易上手，且与PHP的语法和其他主流面向对象编程语言有一些差异，比如go中没有类、继承等概念，取而代之的是接口类型，所以在写和读代码时需要转换一下思维
    * GO和PHP的单元测试框架都很出色，Go拥有嵌入式测试包，而PHP有 PHPUnit，对于性能测试，Go的测试包中拥有很多性能测试的功能。pprof 等许多库都可以使用这些功能来创建华丽的数据报告，虽然PHP也有一套可用于性能测试的库和技术，但Go的更加易于使用

### PHP迁移GO在测试层面的差异及影响
    * 调试
        * PHP：以在代码中打断点的方式调试，需要不断的修改代码
        * GO：也可以通过在代码中打断点调试，但是，由于go是编译型语言，每次打完断点都需要重新编译，所以，还是建议使用dlv这样的调试工具

    * 单测
        * PHP：PHP 的单元测试工具不是太好用，使用也比较麻烦，对于一些情况无法 mock，无法写单元测试
        * GO：Go 自带单元测试工具，测试规范也比较明确，基本不存在无法写单元测试的情况，但对于一些并发测试没有更好的方法。

    * 稳定性测试：
        * PHP：PHP 运行模式是单进程模式，单个进程挂掉不影响其它的进程；PHP 每个请求运行完会主动释放内存，不存在内存泄露问题，基本不需要稳定性测试
        * GO：Go 的goroutine 都是平级的，不存在父子goroutine的概念，如果使用不当会导致整个程序崩溃；Go 程序由于是单进程一直运行，可能会存在内存泄露的情况

    * 测试环境：
        * PHP：一套ODP环境即可，测试开发环境配置方便快速，有成熟的环境自动申请&部署方案（otp）
        * GO：自动申请&部署测试环境能力建设中，目前使用自己的测试开发环境部署，同时需要部署go的基础环境与相关调试工具

    * 竞争检测：
        * PHP：PHP和hhvm目前都无法实现多个请求间的内存共享
        * GO：通过启动goroutine处理不同请求，不同goroutine可以共享内存，如果不加锁，或是加锁不合理，就会很容易出现竞争问题，最终导致执行结果不可预期

# 测试指导
## golang代码能力提高 

### 学习资料
[Golang 新手可能会踩的 50 个坑](https://colobu.com/2015/09/07/gotchas-and-common-mistakes-in-go-golang/)；[Go开发中的十大常见陷阱](https://golangtc.com/t/5d57c0ffb17a820430406e04)；[《go语言圣经》](https://book.itsfun.top/gopl-zh/)；[《go语言实战》](https://baike.baidu.com/item/Go%E8%AF%AD%E8%A8%80%E5%AE%9E%E6%88%98/22384518?fr=aladdin)；[《go web编程》](https://www.kancloud.cn/kancloud/web-application-with-golang/44105)；[《go语言编程》](https://book.douban.com/subject/11577300/)；[《go语言学习笔记》](https://book.douban.com/subject/26832468/)；[《Go并发编程实践》提取码: jcgj](https://pan.baidu.com/s/1DeOSTivQb_V5FcaEY9RYXQ)；[《Go语言标准库》](https://books.studygolang.com/The-Golang-Standard-Library-by-Example/)；[go语言标准库文档](https://studygolang.com/pkgdoc)；[《Go by Example 中文》](https://books.studygolang.com/gobyexample/)


### 相关认证
【推荐】go goodcoder

## 测试工具链
    
### 单测：
【外部开源】gotest，testify，gomock
【内部开源】[GOT](./Golang%20Unit%20Test%20开源框架.md)

**CR:**

[Golang 新手可能会踩的 50 个坑](https://yinzige.com/2018/03/07/50-shades-of-golang-traps-gotchas-mistakes/)

[Go开发中的十大常见陷阱](https://golangtc.com/t/5d57c0ffb17a820430406e04)

[CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)

### 静态代码扫描：
【推荐】[bugbye](https://wiki2ku.baidu-int.com/pubapi/urlmap?id=324946650)（可以并行使用govet、golint和golangci）【包含主流检查工具，且与agile流水新集成，满足工程能力地图要求】
【官方】go lint，go vet
【外部开源】revive，golangci-lint，staticcheck
【内部开源】[goc](http://icode.baidu.com/repos/baidu/gdp/goc/tree/master)【已接入icode】

### P0级case自动化：
【推荐】agile+itp【公司推荐的方案，且与agile流水线集成，满足工程能力地图要求】

### 安全测试：
【推荐】猫头鹰+啄木鸟【公司推荐的方案，且与agile流水线集成，满足工程能力地图要求】

### 自动化回归测试：
【推荐】itp【robot+后续已不再升级，TC建议统一使用itp】

### 测试环境：
【推荐】基于sofacloud搭建的测试环境

### 性能测试：
【推荐】locust（master）+boomer（slave）【操作简单，灵活，开源，boomer为go版本的locust slave，基于go出色的并发能力，相比于其他压测工具，在同等发压资源下，可以发出更大的压力】

【可选】压力云（[流量与压力云](https://wiki2ku.baidu-int.com/pubapi/urlmap?id=489556202)），[糯米压测平台](http://perf.baidu.com/simulation/index)

【外部开源】[pprof](https://github.com/google/pprof)，[火焰图](https://github.com/uber-archive/go-torch)【方便性能分析】

【文章】[Debugging performance issues in Go programs](https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs)( [中文翻译版本](http://www.oschina.net/translate/debugging-performance-issues-in-go-programs))

   

### DIFF测试：
【推荐】xstp



### 异常测试：
【推荐】xstp

    

### 稳定性测试：
【推荐】locust+boomer

【公司推荐】xstp


### 覆盖率统计
【推荐】[系统/接口测试覆盖率方案](https://wiki2ku.baidu-int.com/pubapi/urlmap?id=993191241)

### 其他测试方法
【推荐】go race检测：[https://www.cnblogs.com/yjf512/p/5144211.html](https://www.cnblogs.com/yjf512/p/5144211.html)

### 调试
【推荐】【外部开源】[delve](https://github.com/go-delve/delve)

## 项目测试建议

### PHP迁移golang项目测试指导：
必选：
    * 功能测试：确保迁移功能正常【方法&工具与PHP服务一样】
    * diff测试：线上流量回放diff【迁移后兼容老接口协议，且通过线上切流放量的必选】
    * 性能测试：确定迁移后服务的性能指标
    * 稳定性测试：单实例峰值流量下的服务稳定性，重点关注内存使用情况
    * CR：重点关注一下[Golang 新手可能会踩的 50 个坑](https://yinzige.com/2018/03/07/50-shades-of-golang-traps-gotchas-mistakes/)

建议：
    * 异常测试：下游服务异常时的容错能力，请求异常时的容错能力


### golang常规项目测试指导：
必选：
    * 功能测试：方法&工具与PHP服务一样
    * CR：重点关注一下[Golang 新手可能会踩的 50 个坑](https://yinzige.com/2018/03/07/50-shades-of-golang-traps-gotchas-mistakes/)**（****值得推荐学习！！！****）**

### **go学习书籍和课程推荐**
|序号|书名|作者|推荐理由|难易程度|电子书/介绍链接|推荐人/登记人|
|-|-|-|-|-|-|-|
|1|《GO语言实战》![](https://rte.weiyun.baidu.com/wiki/attach/image/api/imageDownloadAddress?attachId=eae4dc0ca585432b9034ffca4edd1a51&docGuid=WTJh1ZdsY0SP-w "")|作者: 威廉·肯尼迪 (William Kennedy) / 布赖恩·克特森 (Brian Ketelsen) / 埃里克·圣马丁 (Erik St.Martin) 出版社: 人民邮电出版社 出品方: 异步图书 译者: 李兆海|比较经典，系统，结构完整|一般|https://awesome-programming-books.github.io/golang/Go%E8%AF%AD%E8%A8%80%E5%AE%9E%E6%88%98.pdf|程艳青|
|2|《GO语言学习笔记》![](https://rte.weiyun.baidu.com/wiki/attach/image/api/imageDownloadAddress?attachId=79750ffc95e84e8ea165edd2b3bcc5cf&docGuid=WTJh1ZdsY0SP-w "")|雨痕|GO常见基本语法入门，看完后对GO的各种基本数据结构有一些整体的了解适合小白入门学习，但是前半部分都是语法，可能会有些枯燥，边看边动手实践更好|入门级||程艳青|
|3|极客时间付费课程![](https://rte.weiyun.baidu.com/wiki/attach/image/api/imageDownloadAddress?attachId=788193b78f524e7b8d9e266b1237ec7b&docGuid=WTJh1ZdsY0SP-w "")|晁岳攀（鸟窝）--目前就职于百度|比较系统地讲解了并发编程的特点以及一些实践开源项目中，go并发编程的坑|中等（需要深入了解并发，部分需要了解源码）|https://time.geekbang.org/column/intro/100061801|程艳青|
|4|Go语言并发之道|作者: Katherine Cox-Buday 出版社: 中国电力出版社 译者: 于畅 / 马鑫 / 赵晨光|更系统地了解并发系统以及实现并发的各种组件和常见模型，了解不同规模并发系统所需要用到的技巧和工具。|中等（需要对GO有一定了解后学习）|https://book.douban.com/subject/30424330/|程艳青|
|5|《高并发高可用的Go服务开发常用模式与 GAP 工作原理》|陈鑫(go tc主席)|待学习||http://learn.baidu.com/review.html?reviewId=7840||



# 链接：
* [Go语言服务开发生态平台](https://wiki2ku.baidu-int.com/pubapi/urlmap?id=883677829)
* [百度Golang规范委员会](https://wiki2ku.baidu-int.com/pubapi/urlmap?id=106372269)
* [go测试工具小组](https://wiki2ku.baidu-int.com/pubapi/urlmap?id=1037841605)
* [论go语言中goroutine的使用](https://www.cnblogs.com/yjf512/archive/2012/06/30/2571247.html)
* [[Go语言]无辜的goroutine](http://blog.sina.com.cn/s/blog_9be3b8f10101dsr6.html)
* [Introducing the Go Race Detector](https://blog.golang.org/race-detector)
* [Go语言资料收集](https://github.com/wonderfo/wonderfogo/wiki)
* [go社区维护的学习资料](https://github.com/golang/go/wiki/Learn)