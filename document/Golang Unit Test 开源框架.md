Golang Unit Test 开源框架

## 背景
随着组内越来越多的新服务使用Golang开发，Gopher需要一套规范来保证Golang服务质量。包括单元测试，CR规范等。本篇主要用于调研开源社区内对Golang Unit Test（UT）有良好支持的开源框架&工具。

首先需要对UT进行初次筛选，筛选条件为：

1. 满足基本的UT需求，包括但不限于Assert，Mock，UT gen等。
2. 社区拥有足够的关注度（Star数量）和活跃度（Contributor数量&30天内Commit数量）

## 项目目标
1. 帮助rd快速的生成UT，减少繁杂的重复劳动.
2. 提供丰富的自动生成工具，帮助rd模拟各种情况下的输入。
3. 规范单元测试用例的输出，便于统计和展示。

## 统计结果
||Github Star数量|支持Assert|支持Mock|支持Gen|支持结果Web化展示|是否支持web环境构造|支持自动运行UT|支持随机Input|最近30天内是否有Commit提交|学习成本|链接|项目特点|
|-|-|-|-|-|-|-|-|-|-|-|-|-|
|testify|6.8k|✅|✅|❌|❌|❌|❌|❌|✅|中|https://github.com/stretchr/testify|Assert包使用简便，可以快速上手。Mock对代码有入侵，有一定的使用成本。包之间相互独立，可独立对外提供服务。额外提供suite包。|
|goconvey|4.3k|✅|❌|❌|✅|❌|❌|❌|✅|低|https://github.com/smartystreets/goconveyhttp://goconvey.co/|支持超丰富的断言语句。学习成本低，可以快速上手|
|httpexpect|1k|✅|❌|❌|❌|✅|❌|❌|❌|中|https://github.com/gavv/httpexpecthttps://godoc.org/github.com/gavv/httpexpect|如果Handler可以被引用，可以不启动web server对handler进行测试支持echo，iris，gae等框架的web模拟。|
|go-fuzz|2.6k|❌|❌|❌|❌|❌|❌|✅|✅|中|https://github.com/dvyukov/go-fuzz|对代码有少量入侵性，需要编译运行。可以生成大量随机数据，挖掘潜在的bug。命令行操作，提供简单的report。更适合Library的UT。非传统的UT，与UT互补。|
|gotests|1.8k|❌|❌|✅|❌|❌|❌|❌|✅|低|https://github.com/cweill/gotests|自动生成测试用例，可以极快速开始测试用例。|
|monkey|1k|❌|✅|❌|❌|❌|❌|❌|❌|中|https://github.com/bouk/monkey|上手较快，Mock部分函数有一些难度对于一些带有私有声明的方法无法mock对代码没有入侵性可以Mock方法，过程，|
|gostub|59|❌|✅|❌|❌|❌|❌|❌|❌|低|https://github.com/prashantv/gostub|可以mock全局变量，和闭包函数。大部分情况下函数调用并不使用闭包函数，所以在mock函数的时候，有一定的侵入性。|
|net/http/httptest|
|❌|❌|❌|❌|✅|❌|❌|❌|低|https://golang.org/pkg/net/http/httptest/|官方包|

## 分工
|分类|相关的开源框架&工具|代码库&相关教程地址|负责人|
|-|-|-|-|
|Assert断言|||唐国琦|
|Report生成&展示|||唐国琦|
|Mock变量&方法&函数|||俞洋|
|Input数据生成构造|||仲维祎|
|UT自动生成|||仲维祎|
|工具顶层入口|||仲维祎|




## 相关代码库
[http://icode.baidu.com/repos/baidu/gdp/got/tree/master](http://icode.baidu.com/repos/baidu/gdp/got/tree/master)

## 相关分享
[GOT分享](./Golang%20Unit%20Test%20开源框架.md)
