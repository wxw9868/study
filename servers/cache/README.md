- [手把手教你用 Go 语言实现缓存系统](#手把手教你用go语言实现缓存系统)
  - [面试题内容](#面试题内容)
  - [题目分析](#题目分析)
  - [动手写代码](#动手写代码)
    - [构建大体框架](#构建大体框架)
    - [逐个实现方法](#逐个实现方法)
      - [SetMaxMemory](#setmaxmemory)
      - [Set](#set)
      - [Get](#get)
      - [Del](#del)
      - [Exists](#exists)
      - [Flush](#flush)
      - [Keys](#keys)
    - [实现代理层](#实现代理层)
    - [加分项](#加分项)
      - [轮询检查删除过期键](#轮询检查删除过期键)
      - [单元测试](#单元测试)
  - [小结](#小结)

# 手把手教你用 Go 语言实现缓存系统

[原文链接](https://studygolang.com/topics/16993)

​ 今天我们围绕一个面试题来实现一个内存缓存系统，大家也可以在完成后，自己增加一些额外的功能。

## 面试题内容

1. 支持设置过期时间，精度到秒
2. 支持设置最大内存，当内存超出时做出合理的处理
3. 支持并发安全
4. 按照以下接口要求实现

```go
type Cache interface {
    // SetMaxMemory size : 1KB 100KB 1MB 2MB 1GB
    SetMaxMemory(size string) bool
    // Set 将 value 写入缓存
    Set(key string, val interface{}, expire time.Duration) bool
    // Get 根据 key 值获取 value
    Get(key string) (interface{}, bool)
    // Del 删除 key 值
    Del(key string) bool
    // Exists 判断 key 是否存在
    Exists(key string) bool
    // Flush 清空所有 key
    Flush() bool
    // Keys 获取缓存中所有 key 的数量
    Keys() int64
}
```

使用示例

```go
func main() {
    cache := NewMemCache()
    cache.SetMaxMemory("100MB")
    cache.Set("int", 1)
    cache.Set("bool", false)
    cache.Set("data", map[string]interface{}{"a": 1})
    cache.Get("int")
    cache.Del("int")
    cache.Flush()
    cache.Keys()
}
```

## 题目分析

题目乍一看没有什么难点，就依据题目实现对应的接口以及对应的方法就行了。但其实有一个坑，那就是接口中的 Set 方法参数和使用示例中的对不上，使用示例中没有传入过期时间。难道是题目出错了？

​ 显然不是的，这里是需要我们去做一个代理层，去实现一个可选参数的 Set 方法。我们可以在实现了带过期时间参数的方法后，再去封装一层，然后设置成可选参数即可。

这样子，题目的要求，就差不多没什么问题了。但这是面试题，我们想要面试官对我们的评价更好，就要做到更多的内容，寻找一些加分项，比如我们可以增加一个功能，定期删除过期缓存键，又或者我们可以写一些单元测试，让面试官知道我们有写单元测试的好习惯，这些都是一些加分项，能让我们更加地突出。

## 动手写代码

下面就带着大家一步一步来完成这个缓存系统，当然只是一个具有基础的功能缓存系统，大家在后续也可以自行在其中丰富更多的功能。

### 构建大体框架

​ 首先，我们可以先在项目根目录下创建一个 cache 包，然后在 cache 包里创建一个 cache.go 文件，然后将题目中要求实现的接口放在里面：

```go
// cache/cache.go
type Cache interface {
    // SetMaxMemory size : 1KB 100KB 1MB 2MB 1GB
    SetMaxMemory(size string) bool
    // Set 将 value 写入缓存
    Set(key string, val interface{}, expire time.Duration) bool
    // Get 根据 key 值获取 value
    Get(key string) (interface{}, bool)
    // Del 删除 key 值
    Del(key string) bool
    // Exists 判断 key 是否存在
    Exists(key string) bool
    // Flush 清空所有 key
    Flush() bool
    // Keys 获取缓存中所有 key 的数量
    Keys() int64
}
```

​ 接着我们再在 cache 包下创建一个 memCache.go 文件，并在该文件中创建一个 memCache 结构体，去实现题目中要求的 Cache 接口：

```go
type memCache struct {
}

func (mc *memCache) SetMaxMemory(size string) bool {
    return false
}

// Set 将 value 写入缓存
func (mc *memCache) Set(key string, val interface{}, expire time.Duration) bool {
    return false
}

// Get 根据 key 值获取 value
func (mc *memCache) Get(key string) (interface{}, bool) {
    return nil, false
}

// Del 删除 key 值
func (mc *memCache) Del(key string) bool {
    return true
}

// Exists 判断 key 是否存在
func (mc *memCache) Exists(key string) bool {
    return true
}

// Flush 清空所有 key
func (mc *memCache) Flush() bool {
    return true
}

// Keys 获取缓存中所有 key 的数量
func (mc *memCache) Keys() int64 {
    return 0
}
```

​ 可以看到使用样例中有一个 NewMemCache 函数，于是我们还需要在 memCache.go 文件中添加一个 NewMemCache() 函数，返回一个实例，供 main 函数调用：

```go
// cache/memCache.go
func NewMemCache() Cache {
    return &memCache{}
}
```

​ 接着，我们就可以先去主函数 main 中用使用示例跑一下，看看有没有什么问题。在项目根目录下创建一个 main 函数，如下：

```go
// main.go
package main

import cache2 "main/cache"

func main() {
    cache := cache2.NewMemCache()
    cache.SetMaxMemory("100MB")
    cache.Set("int", 1)
    cache.Set("bool", false)
    cache.Set("data", map[string]interface{}{"a": 1})
    cache.Get("int")
    cache.Del("int")
    cache.Flush()
    cache.Keys()
}
```

​ 然后你就会发现报错了，这个问题就是我们上面说的那个坑，这里需要做一个代理层，但为了方便我们可以先修改使用样例，使他能够先跑通，最后再来做一个代理就可以了。

```go
// main.go
cache.Set("int", 1, 3)
cache.Set("bool", false, 1)
cache.Set("data", map[string]interface{}{"a": 1}, 2)
```

为了看出效果，我们可以在所有的方法中都打印一个信息，比如 Set 方法中打印 `fmt.Println("我是 Set 方法")。

​ 最后如果所有的方法都打印出了对应的信息，就说明这个大体框架我们已经搭建好了，下面再去慢慢实现各个方法就可以了。

### 逐个实现方法

​ 下面就带着大家逐个实现每个具体的方法：

#### SetMaxMemory

这个方法是用于设置我们缓存系统的最大缓存大小的，因此我们的 memCache 结构体中，就至少应该需要两个字段：最大内存大小 和当前内存大小，因为我们肯定需要去判断当前内存是否超过了最大内存大小，为了方便，我们再增加一个 最大内存大小字段的字符串表示，如下：

```go
// cache/memCache.go
type memCache struct {
    // 最大内存大小
    maxMemorySize int64
    // 最大内存字符串表示
    maxMemorySizeStr string
    // 当前内存大小
    currMemorySize int64
}
```

然后再来看我们的题目要求，SetMaxMemory size : 1KB 100KB 1MB 2MB 1GB 要求支持多种单位的表示，所以这里我们肯定需要对输入的内存大小做一个转换，因此我们需要去实现一个 parseSize 函数去解析用户的输入。

​ 我们在 cache 包下，创建一个 util.go 文件，用于存放一些通用功能和工具函数。

​parseSize 函数的实现思路是：将用户输入的字符串中的数字部分和单位部分分别提取出来，再进行校验和单位转换等的处理。

​ 利用正则表达式先将用户的输入中的数字部分提取出来，然后再将用户输入的字符串中的数字部分用空格替换，这样剩下的部分就是单位了。

​ 同样的，将用户输入字符串中的单位部分用空格替换，得到的就是数字部分了。

​ 接下来，就是对用户的单位做一个转换，这里我们利用 Go 语言中的预定义标识符，采用小技巧来做一个单位的转换，如下：

```go
// cache/util.go
const (
    B = 1 << (iota * 10) // 1
    KB                      // 2024
    MB                   // 1048576
    GB                   // 1073741824
    TB                      // ...
    PB                      // ...
)
​ 有了不同的单位，我们就可以对解析出来的单位进行处理了，我们这里统一将所有的单位转换成字节，也就是 B：

// cache/util.go
var byteNum int64 = 0
// 1KB 100KB 1MB 2MB 1GB，单位统一为 byte
switch unit {
case "B":
    byteNum = num
case "KB":
    byteNum = num * KB
case "MB":
    byteNum = num * MB
case "GB":
    byteNum = num * GB
case "TB":
    byteNum = num * TB
case "PB":
    byteNum = num * PB
default: // 设置不合法，设置为 0，后续设置为 默认值
    num = 0
}
```

​ 如果用户输入的单位不合法，就是通过后续处理设置为默认值 100MB：

```go
// cache/util.go
// 用户使用不合法，打印日志并设置默认值
if num == 0 {
    log.Println("ParseSize 仅支持 B、KB、MB、GB、TB、PB")
    num = 100
    byteNum = num * MB
    unit = "MB"
}
```

​ 最后由于我们需要返回的有两种形式，即字符串形式和数字形式，所以这里还需要拼接一下字符串形式。这里没有直接使用用户传入的值，是因为可能用户的输入有问题，然后我们采用的是默认值，故这里直接统一全部重新拼接：

```go
// cache/util.go
sizeStr := strconv.FormatInt(num, 10) + unit
```

​ 至此，ParseSize 函数，我们就实现完毕了，完整代码如下：

```go
func ParseSize(size string) (int64, string) {
    // 默认大小为 100

    // 利用正则表达式匹配一个或者多个数字
    re, _ := regexp.Compile("[0-9]+")
    // 获取单位：使用编译好的正则表达式 re，将 size 字符串中匹配的数字字符替换为空字符串
    unit := string(re.ReplaceAll([]byte(size), []byte("")))

    // 获取数字：将 size 字符串中的单位部分 unit 用空字符串替换，即可获取数字部分。最后再将数字转换为 int64 类型
    num, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)

    // 单位转换为大写
    unit = strings.ToUpper(unit)

    var byteNum int64 = 0
    // 1KB 100KB 1MB 2MB 1GB，单位统一为 byte
    switch unit {
    case "B":
        byteNum = num
    case "KB":
        byteNum = num * KB
    case "MB":
        byteNum = num * MB
    case "GB":
        byteNum = num * GB
    case "TB":
        byteNum = num * TB
    case "PB":
        byteNum = num * PB
    default: // 设置不合法，设置为 0，后续设置为 默认值
        num = 0
    }

    // 用户使用不合法，打印日志并设置默认值
    if num == 0 {
        log.Println("ParseSize 仅支持 B、KB、MB、GB、TB、PB")
        num = 100
        byteNum = num * MB
        unit = "MB"
    }

    sizeStr := strconv.FormatInt(num, 10) + unit
    return byteNum, sizeStr
}
```

​ 然后我们的 SetMaxMemory 函数就简单了，如下：

```go
// cache/memCache.go
func (mc *memCache) SetMaxMemory(size string) bool {
    mc.maxMemorySize, mc.maxMemorySizeStr = ParseSize(size)
    fmt.Println(mc.maxMemorySize, mc.maxMemorySizeStr)
    return true
}
```

然后我们可以运行 main.go 函数，打印一下检查是否有问题：

```sh
$ go run main.go
104857600 100MB
```

​ 可以用计算器算一下，没有问题，我们的 SetMaxMemory 函数就完成了。

#### Set

​ 然后是我们的 Set 方法，这个方法是用来将键值对存入我们的缓存系统中的。

​ 首先，我们需要考虑用什么类型来存储键值对，很自然就可以想到用 Go 语言内置的字典，即 map 来实现。那新的问题又来了，那 map 的 key-value 分别用什么类型呢？key 的类型，毫无疑问肯定是 string 类型；value 的类型的话，这里如果也用一个单独的 interface{} 类型的话，可能也会存在一些问题，因为我们的 value 需要携带很多附加信息，比如 value 的值、过期时间、value 大小等，故这里的 value 需要用一个结构体去表示，故我们要先创建一个 memCacheValue 结构体：

```go
// cahce.memCache.go
type memCacheValue struct {
    // value 值
    val interface{}
    // 过期时间
    expireTime time.Time
    // 有效时长
    expire time.Duration
    // value 大小。用于计算当前内存大小
    size int64
}
```

​ 有了 memCacheValue，就可以在 memCache 中新增一个字段了：

```go
// cahce.memCache.go
type memCache struct {
    // 最大内存大小
    maxMemorySize int64
    // 最大内存字符串表示
    maxMemorySizeStr string
    // 当前内存大小
    currMemorySize int64
    // 缓存键值对
    values map[string]*memCacheValue
}
```

​ 由于这里使用了 map，故在初始化 memCache 实例的时候，需要进行内存的分配，所以我们要修改 NewMemCache 函数：

```go
// cahce.memCache.go
func NewMemCache() Cache {
    mc := &memCache{
        values: make(map[string]*memCacheValue),
    }
    return mc
}
```

​ 言归正传，继续分析我们的 Set 方法的实现，由于 Set 方法是写操作，Map 并非线程安全的，所以我们在进行写操作的时候需要进行加锁保护，故这里 memCache 结构中还需要加一个锁：

```go
type memCache struct {
    ...
    // 读写锁
    locker sync.RWMutex
}
```

​ 这里我们采用读写锁，这样可以利用 map 的读写机制：读操作兼容、写操作互斥，最大化提升读写 map 的性能。

​ 因为我们的键可能存在过期时间，如果是重复 Set 一个已存在的值，就需要去重新计算更新对应的时间，会需要分情况讨论，比较复杂。所以，这里我们统一使用，先删除对应键值，再添加对应键值，写起来会比较方便，下面我们实现三个方法，以便我们调用：

```go
// 判断是否存在对应的 value
func (mc *memCache) get(key string) (*memCacheValue, bool) {
    val, ok := mc.values[key]
    return val, ok
}

// 删除：当前内存大小更新、删除当前 key 值
func (mc *memCache) del(key string) {
    tmp, ok := mc.get(key)
    if ok && tmp != nil {
        mc.currMemorySize -= tmp.size
        delete(mc.values, key)
    }
}

// 添加：当前内存大小更新、删除当前 key 值
func (mc *memCache) add(key string, val *memCacheValue) {
    mc.values[key] = val
    mc.currMemorySize += val.size
}
```

​ 上述三个方法比较简单，就不做过多赘述了。有了这三个函数，我们后面其他的方法实现起来，都会很简单。

​ 然后我们的 Set 方法就可以写了：

```go
func (mc *memCache) Set(key string, val interface{}, expire time.Duration) bool {
    // map 非线程安全需要加锁访问
    mc.locker.Lock()
    defer mc.locker.Unlock()
    // 确定一个 value 值
    v := &memCacheValue{
        val:        val,
        expireTime: time.Now().Add(expire),
        expire:     expire,
        size:       GetValSize(val),
    }
    // 为了简化代码复杂度，这里用 “删除再添加” 来代替 “更新” 操作
    if _, ok := mc.get(key); ok { // 存在则删除
        mc.del(key)
    }
    mc.add(key, v)

    // 新增后缓存是否超过最大内存：超过则直接删除刚刚添加的这个 key，并报 panic
    if mc.currMemorySize > mc.maxMemorySize {
        mc.del(key)
        // 这里可以自己完善一下，通过一些内存淘汰策略来选择删除一些 key，来判断是否还会超过最大内存
        log.Println(fmt.Sprintf("max memory size %s", mc.maxMemorySizeStr))
    }
    return false
}
```

在开始操作之前，加写锁保护
根据用户输入，构建对应的 value 值
如果存在对应键值对，就先删除，然后再添加对应键值
新增后判断是否超过内存，超过了的话，就直接删除刚刚添加的这个键
​
上述 Set 方法中还用到一个函数 GetValSize ，我们可以先不去实现这个函数具体逻辑，后面再回过头来看：

```go
// cache/util.go
// GetValSize 计算 value 值大小
func GetValSize(val interface{}) int64 {
    return 0
}
```

#### Get

​Get 方法，是根据用户输入的键，来获取对应的 value 值的。实现很简单，如下：

```go
func (mc *memCache) Get(key string) (interface{}, bool) {
    mc.locker.RLock()
    defer mc.locker.RUnlock()

    // 拿到对应的值
    mcv, ok := mc.get(key)
    // 判断是否过期
    if ok {
        if mcv.expire != 0 && mcv.expireTime.Before(time.Now()) { // 过期时间早于当前时间，删除
            mc.del(key)
            return nil, false
        }
        return mcv.val, ok
    }
    return nil, false
}
```

加读锁保护
先通过 get 方法拿到对应的值
如果存在该键，且该值没有过期或者没有过期时间，则返回该值
否则返回 nil，并删除该过期键

#### Del

​Del 方法，是用于删除对应键值对的。直接加写锁操作，并调用先前实现的 del 函数即可：

```go
func (mc *memCache) Del(key string) bool {
    mc.locker.Lock()
    defer mc.locker.Unlock()
    mc.del(key)
    return true
}
```

加写锁保护
直接调用 del 函数进行删除对应键值对即可

#### Exists

​Exists 方法，用于判断某个键是否存在于我们的缓存系统。实现也非常简单：

```go
func (mc *memCache) Exists(key string) bool {
    mc.locker.RLock()
    defer mc.locker.RUnlock()
    _, ok := mc.values[key]
    return ok
}
```

加读锁保护
直接获取对应键值对，以此判断是否存在该键

#### Flush

​Flush 方法，是在整个缓存系统的缓存数据不需要使用之后，用来清空所有的缓存时使用的。这里我们利用 Go 语言的垃圾回收机制，直接将整个 map 置空，Go 语言的垃圾回收机制会直接将没有使用的内存进行回收：

```go
func (mc *memCache) Flush() bool {
    mc.locker.Lock()
    defer mc.locker.Unlock()
    // 直接将整个 map 置空，go 的垃圾回收机制会自行将没有使用的内存进行回收
    mc.values = make(map[string]*memCacheValue, 0)
    mc.currMemorySize = 0

    return true
}
```

加写锁保护
将整个 map 置空，并将当前使用内存大小清空

#### Keys

​Keys 方法，用于获取缓存中 key 的数量。直接用 len() 函数获取即可：

```go
func (mc *memCache) Keys() int64 {
    mc.locker.RLock()
    defer mc.locker.RUnlock()
    return int64(len(mc.values))
}
```

加读锁保护
利用 len() 函数直接获取
​
现在我们再来看看这个 GetValSize 函数，这里有两种思路实现：

利用反射包 unsafe.Sizeof(val) 来获取对应的值的大小
野路子：利用 json 包，将 val 序列化为字节数组，然后求字节数组的长度，就知道该值占用了多少字节了
​
通过测试，可以发现第一种方法是不可靠的，因为 unsafe.Sizeof() 方法只是算出对应类型的字节大小，而不是你所存储的内容的具体大小。于是我们这里采用第二种方法来实现：

```go
// cache/util.go
// GetValSize 计算 value 值大小
func GetValSize(val interface{}) int64 {
    // 野路子：利用 json 包，将 val 序列化为字节数组，然后求字节数组的长度，就知道占用了多少字节了
    bytes, _ := json.Marshal(val)
    size := int64(len(bytes))
    return size
}
```

​ 至此，我们的基础功能，就差不多实现了。大家可以通过在 main 函数打印对应的信息来检查。在这里，我就不再带着大家检查了。

### 实现代理层

​ 首先我们得明白，为什么要再加一层代理层？在这里，题目给的接口是包含过期时间的，但是我们的使用示例却没有使用过期时间，这就是说明需要加一层代理层来进行封装。

​ 添加代理层还有一些好处，比如：

- 安全性：它可以过滤和阻止对系统的未经授权的访问，通过身份验证和授权机制确保只有合法的用户或服务可以访问底层的资源。这有助于防范潜在的安全威胁。
- 性能优化：代理层可以缓存某些请求的结果，以避免重复计算或数据库查询。此外，代理层还可以对请求进行负载均衡，确保各个后端服务得到合理的分配，以提高整体性能。
- 抽象底层实现： 代理层可以用于隐藏底层实现的复杂性，提供简化的接口给上层系统。这有助于实现系统的模块化和降低耦合度，使得系统更容易维护和扩展。
  ​

下面我们来看看怎么实现代理层：

​ 首先，在项目根目录创建一个文件夹 cache-server ，并在该目录下创建一个 cache.go 文件，完整代码如下：

```go
package cache_server

import (
    "main/cache"
    "time"
)

// 代理层/适配层
type cacheServer struct {
    memCache cache.Cache
}

func NewMemCache() *cacheServer {
    return &cacheServer{
        memCache: cache.NewMemCache(),
    }
}

// SetMaxMemory size : 1KB 100KB 1MB 2MB 1GB
func (cs *cacheServer) SetMaxMemory(size string) bool {
    return cs.memCache.SetMaxMemory(size)
}

// Set 将 value 写入缓存
// 代理层：将 有效时长参数设置为可有可无
func (cs *cacheServer) Set(key string, val interface{}, expire ...time.Duration) bool {
    // 默认值为 0 秒
    expireTs := time.Second * 0
    if len(expire) > 0 {
        expireTs = expire[0]
    }
    return cs.memCache.Set(key, val, expireTs)
}

// Get 根据 key 值获取 value
func (cs *cacheServer) Get(key string) (interface{}, bool) {
    return cs.memCache.Get(key)
}

// Del 删除 key 值
func (cs *cacheServer) Del(key string) bool {
    return cs.memCache.Del(key)
}

// Exists 判断 key 是否存在
func (cs *cacheServer) Exists(key string) bool {
    return cs.memCache.Exists(key)
}

// Flush 清空所有 key
func (cs *cacheServer) Flush() bool {
    return cs.memCache.Flush()
}

// Keys 获取缓存中所有 key 的数量
func (cs *cacheServer) Keys() int64 {
    return cs.memCache.Keys()
}
```

1. 首先我们仍然需要先将 cache 接口封装起来，并写一个构造函数返回一个实例
2. 除了 Set 方法外，其他方法直接调用实现即可
3. 在 Set 方法中，将过期时间参数设置为可选参数，即 expire ...time.Duration，然后通过判断是否传入该参数来构建新的参数 expireTs 去调用实现好的方法
   ​
   实现完成后，我们就可以将 main 函数的代码修改一下，调用代理层提供的方法来进行使用：

```go
package main

import (
    cache_server "main/cache-server"
)

func main() {
    cache := cache_server.NewMemCache()
    cache.SetMaxMemory("100MB")
    cache.Set("int", 1)
    cache.Set("bool", false)
    cache.Set("data", map[string]interface{}{"a": 1})
    cache.Get("int")
    cache.Del("int")
    cache.Flush()
    cache.Keys()
}
```

​ 这样，即使我们不使用带过期时间的 Set 方法，也不会报错了。

### 加分项

​ 最后我们再来看看我们的加分项：

#### 轮询检查删除过期键

​ 我们可以在创建缓存系统实例的时候，同时开启我们的 ” 轮询检查删除过期键 “ 功能。

```go
func NewMemCache() Cache {
    mc := &memCache{
        values:                       make(map[string]*memCacheValue),
        cleanExpiredItemTimeInterval: time.Second, // 定期清理缓存
    }
    // 轮询检查删除过期键
    go mc.cleanExpiredItem()
    return mc
}
```

​ 这里需要新添加一个字段 cleanExpiredItemTimeInterval 表示清理周期，还需要写一个轮询的函数，如下：

```go
type memCache struct {
    ...
    // 清楚过期缓存时间间隔
    cleanExpiredItemTimeInterval time.Duration
}
```

​ 下面是轮询的函数：

```go
// 轮询清空过期 key
func (mc *memCache) cleanExpiredItem() {
    // 设置一个定时触发器：定时向 Ticker.C 中发送一个消息，即触发了一次
    timeTicker := time.NewTicker(mc.cleanExpiredItemTimeInterval)
    defer timeTicker.Stop()
    for {
        select {
        case <-timeTicker.C: // 每个周期做一个缓存清理
            // 遍历所有缓存的键值对，将有过期时间且过期的键删除掉
            for key, item := range mc.values {
                if item.expire != 0 && time.Now().After(item.expireTime) {
                    mc.locker.Lock()
                    mc.del(key)
                    mc.locker.Unlock()
                }
            }
        }
    }
}
```

1. 采用 time.NewTicker，定义一个制定周期的定时器
2. 由于需要不断轮询，故需要放在 for 循环中
3. 配合 select 实现一个阻塞式的轮询检查并删除过期键的操作操作

#### 单元测试

​ 单元测试是一种用于验证程序各个独立单元是否能按照预期工作的测试方法。Go 语言的测试工具内置于语言本身，通过 testing 包提供了一套简单而有效的测试框架。平时不论是学习、还是工作，都应该养成写单元测试的习惯。

​ 我们在 cache 包下，创建一个 memCache_test.go 文件，并在里面写我们测试内容：

```go
package cache

import (
    "testing"
    "time"
)

func TestCacheOP(t *testing.T) {
    testData := []struct {
        key    string
        val    interface{}
        expire time.Duration
    }{
        {"baer", 678, time.Second * 10},
        {"hrws", false, time.Second * 11},
        {"gddfas", true, time.Second * 12},
        {"rwe", map[string]interface{}{"a": 3, "b": false}, time.Second * 13},
        {"rqew", "fsdfas", time.Second * 14},
        {"fsdew", "这里是字符串这里是字符串这里是字符串", time.Second * 15},
    }

    c := NewMemCache()
    c.SetMaxMemory("10MB")
    // 测试 set 和 get
    for _, item := range testData {
        c.Set(item.key, item.val, item.expire)
        val, ok := c.Get(item.key)
        if !ok {
            t.Error("缓存取值失败")
        }
        if item.key != "rwe" && val != item.val {
            t.Error("缓存取值数据与预期不一致")
        }
        _, ok1 := val.(map[string]interface{})
        if item.key == "rwe" && !ok1 {
            t.Error("缓存取值数据与预期不一致")
        }
    }
    // 测试 Keys()
    if int64(len(testData)) != c.Keys() {
        t.Error("缓存数量不一致")
    }
    // 测试 Del()
    c.Del(testData[0].key)
    c.Del(testData[1].key)

    if int64(len(testData)) != c.Keys()+2 {
        t.Error("缓存数量不一致")
    }

    // 测试过期时间
    time.Sleep(time.Second * 16)

    if c.Keys() != 0 {
        t.Error("缓存清空失败")
    }
}
```

先用匿名结构体，构建需要用到的各类测试数据

然后对 Set、Get、Del 等方法进行调用，然后对比结果

## 小结

这篇文章，通过一个面试题，从题目到各种坑点的分析，带大家了解并实现了一个简易版的 内存缓存系统 。相信大家在看完后肯定会收货满满。
