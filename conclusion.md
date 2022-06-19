# Go 进阶训练营总结 

> 感慨良多，为期十三周的学习和训练。进步了许多，从对Go的恐慌，到对Go 的亲和；从对工程的剪不断理还乱，到慢慢有大心脏可以应对；从知识的学习碎片和断续，到系统大量文章积累和长期持续训练。这是训练营中获得的变化，感受很好。 


对于讲师，毛大的课程听起来就像嗑瓜子一样味道，停不下来。大明老师真是"渣男"哈哈，讲授的知识非常实用。包子老师也辛苦解答我们的疑惑。感谢三位。 
特别宣班班，做事有方法条理，有亲和力，漂亮，漂亮，漂亮。有一次我问班班，每次领教直播会议链接，明明在群里已经发了，为何还单独每个人私发呢？ 
班班的回复说，因为不是每个人都会看群，我要确保每个人都能收到会议链接。顿时被班班迷住了。太棒了吧这样！！！ 

对于知识，学到了很体系的后端服务的方方面面，包括错误码规范，API设计，工程设计，
服务高可用，链路追踪，服务监控等等。 这一切让我对整体的后端体系有了系统和感性认知，以后我得反复思考学习和训练。  

每次作业都逼着自己进步。 

# 《以下是附着一些 自己沉淀的知识》

## Go env
- GOROOT  : go安装目录
- GOPATH : go项目工作区
- GO111MOUDLE: 控制是否启用 go mod 
- GOPROXY : go 访问依赖包的代理 , https://goproxy.cn/
- GOPRIVATE: 执行自己的私有仓，譬如说

## 将Go 源码编译成 汇编语言 
> go build -gcflags -S dedup.go

````
    go tool compile -S -N -l anything.go
````
-N -l 参数禁止编译器优化，否则输出的汇编结果会有较大差别 

## 逃逸分析  
````
    go run -gcflags "-m -l" escape2.go 
    go build -gcflags "-m -l" escape2.go 
````

Reference 
https://www.imooc.com/article/317776 
https://cloud.tencent.com/developer/article/1585486



## 常量定义未使用可以通过编译，变量定义未使用编译报错

## Go 的 string 类型 
> Go 中的 string 类型是只读的字节序列   

信息量1 - len(str) 计算的是字符所占的字节数，而不是字符数，go采用 Unicode 字符编码集，采用UTF-8 存储实现

````
    str := "Hello 世界"
    len(str) // 返回的是字节个数，而非字符个数
````


## 切片 
- 切片结构 
- 切片的文法
- 切片的切片操作 
- 切片的默认行为 
- 切片的零值 nil 
- 切片的越界操作
- make 创建切片


### 切片的结构，类似于下面结构体 
```
    type slice strct{
        data *[]T
        len int 
        cap int 
    }
```
**data** 是slice 第一个元素指向数组元素的地址 ，第一个元素，但地址指向底层数组  
**len** slice 拥有的元素个数 
**cap** slice 的容量，指slice 第一个元素到 底层数组末尾元素的个数  

### 切片的文法 
切片的文法类似于没有长度的数组文法   
下面是一个数组的文法
```
    [3]bool{true,false,true}
```
这是一个切片的文法,首先创建一个[3]int{1,2,3}数组，然后构建出一个切片 
```
    [3]int{1,2,3}
```


### 切片操作 
```
    a[low:high]
```
切片通过上下标，以冒号：分隔来产生 。是一个左闭又开的区间  
元素包含从 low下标开始，直到 high-1 下标，不包含high 下标 
```
    a[1:4]
```
创建了一个切片，它包含a中的元素从下标 1 到 3  
它包含a中的元素从下标 1 到 3 

### 切片默认行为 
切片也支持 上下标省略 ，可省略上标，也可省略下标，上下标可同时省略
```
    a[:]
```
它表示，包含a中的元素从下标 0 到 len(a)-1

上标省略，默认为0 ；小标省略默认为 len

### 切片的零值 nil  
切片的零值是 nil ,其len 和cap 都是0  
```
    func main(){
        var s []int 
        fmt.Println(s,len(s),cap(s))
        if s==nil{
            fmt.Println("nil")
        }
    }
```

### 切片的越界操作 
如果 切片的操作s[:cap(s)+1] 越过了cap上界，则会引发一个panic ;如果操作只是突破了 len,则切片扩展了，长度增加了 

```
func TestSliceOutBound(t *testing.T) {
	var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	s := arr[:]
	//t.Log(s[:20])  //panic: runtime error: slice bounds out of range [:20] with capacity 10 

	s = s[:5]
	s1 := s[:6]
	t.Log(s1, len(s1), cap(s1))

}

```

### make 创建切片
```
    make([]T,len,cap)
```



## UTF-8 存储实现  
一个字节  0xxx-xxxx 完全兼容 Ascii 	 					如果 Unicode码值在 [0-127](d) 之间，则采用一个字节存储
两个字节  110x-xxxx 10xx-xxxx 		 					如果 Unicode码值在 [128-2047](d) 之间，则采用两个字节存储
三个字节  1110-xxxx 10xx-xxxx 10xx-xxxx  				如果 Unicode码值在 [2048-65535](d) 之间，则采用三个字节存储
四个字节  1111-0xxx 10xx-xxxx  10xx-xxxx  10xx-xxxx 	如果 Unicode码值在 [65536-0x10ffff] 之间，则采用四个字节存储



## 方法 
 * 可以为任何已命名的（除指针和接口）类型定义方法，方法接收者不一定要是结构体 



## 方法的接收者  

1、Go的方法接收者可以是 指针类型或非指针类型，不管形参是指针还是非指针类型，传入实参既可以是 指针类型也可以是非指针类型，编译器会帮我们 将实参类型与形参类型匹配对等上。

````
    type P T  
    func (p P) XX() {}

    var p = new(P)

    p.XX() // 编译器会隐私转换 (*p).XX()
````

2、可以给任意命名类型定义方法，但底层类型是 指针或接口就不行   
例如
````
    type P *int 
    func (p P) XX(){} //编译就不通过  

    type Eat interface{} 
    type E Eat 
    func (e E) EE(){} //invalid receiver E (pointer or interface type)
````
3、基础类型 和未命名类型 不可以作为 方法接受者

````
    // invalid receiver int (basic or unnamed type)compilerInvalidRecv
    func (n int) NHxx() {
        fmt.Println("ha hha")
    }
    //不能直接给 基础类型int 定义方法 
````

## new 和 make  

### new 分配内存，将内存置为类型的零值，（注意不是初始化内存）将内存地址返回   
也就是返回 指针   `*T`  

**请注意**，函数返回局部变量地址是可行的，这和C语言不同，局部变量对应的数据，在函数返回后依然有效。这一点需要借助 Go汇编来探究为何会如此 todo  

### make 只用户创建 切片、map、channel , make会返回一个 T 类型的已初始化内存的值（不是零值）。和 new 的区别就在这里，原因是 这三种类型本质上是 引用数据类型，使用他们之前，必须初始化。 
**注意啦**,make仅用于切片、map、channel 的创建且返回的不是指针。 


## 接口值 
> 接口值由两部分组成，类型和类型的值

声音一个接口变量 ，来看看接口值情况  
````
    var w io.Writer 
````
![nil_interface_value](imgs/nil_interface_value_1.png)

````
    var buf bytes.Buffer  
    w = &buf  
````

![non-nil_interface_value](imgs/non-nil_interface_value.png)

以下是陷阱 
````
    const debug = false  
    func main(){
        var w *bytes.Buffer  
        if debug{
            w = new(bytes.Buffer)
        }
        f(w)
    }

    func f(w io.Writer){
        if w !=nil{
            w.Write([]byte("hello"))  //invalid memory address or nil pointer dereference
        }
    }
````
**这段代码会报错**,此时 接口值w 就不是 nil 值了。它值的类型部分不是nil ，虽然值部分是 nil  

将上面代码改以下，就不会发生错误 
````
    var w io.Writer  
    if debug {
        w = new(bytes.Buffer)
    }
````



## Error of Go  
信息量 

- errors.New 返回的是结构体的指针，而不是直接返回结构体值；这样做的好处防止意外相同情况  
- 自定义 error 格式，好的编码应该带上位置信息 
    譬如 bufio.go 包中，自定义了一些error  
        ErrBufferFull = errors.New("bufio: buffer full")    
        这里错误消息就分成两个部分， 包名：+ 错误信息  
    在实践中确实也很恼火，别人把错误wrap ，然后返回一个自定义错误，里面没有具体报错位置，你要定位一个错误，找死了，特别恼火。 

Go 的error 模式 
- error 交给 调用者处理，立即处理  
- 将致命错误，专用panic 机制处理，它会让程序 fatal eror 掉，而不是交给 调用者处理，
    这点不像java 普通业务error 和致命错误 都上抛给调用者，由调用者处理 ，那如果致命错误，Java 程序会假死在那里
- 在net/http 包中，遇到一个 处理不了的request 会panic ,但会用 recover 兜底，将该request 牺牲掉，不耽搁其它request 处理 


Go 选中 Error 模式的考量   

- 简单 
- 没有隐藏的控制流
- 完全交给你控制 error  
- plan for failure ,not success 
- Errors are values  
- 由于强调立即处理error ,就比 checked exception 更容易保持 被调函数的语义原子性  



Err Type  

Sentinel Error  
    预定义的一些error ,这样在 方法提供者和方法调用者之间 建立了错误类型的约定。 
    调用者根据 判断返回的 err 等于 什么 Sentinel Error 来决定做什么事情。  

这不是一个好的设计，灵活的错误处理模式  
- 这是一个脆弱的 约定，如果返回的error 是 wrap Sentinel Error 增加了更多错误的上下文信息,那么你的 相等判断就会失效了  
    那么通常这个时候，程序员会有一个 能立马解决，但很糟糕的编码决定，匹配错误的 字符串信息。 这是一个很脆弱的 错误类型识别方法。虽然不好，但是能最快想到而又快速的方法。 
- Sentinel errors  使得你 API 表面积增大，当你的包暴露的内容越多，其实这个抽象是脆弱的 
- Sentinel errors 使得你的包之间产生了 依赖，源代码级别的依赖 

结论是：避免使用 Sentinel  errors  

Error Type  
一个实现了 error 接口的自定义类型，比Sentinel Error好在 提供了更多上下文信息 。  
调用者，根据 类型断言来判断错误类型  
````
    type MyError Struct {
        Msg  string 
        File string 
        Line int 
    }

    func (e *MyError) Error() string{
        return fmt.Sprintf("%s:%d   %s",e.File,e.Line,e.Msg)
    }
    func NewMyError() error{
        return &MyError{Msg:"no suce file",File:"no/suce.file",Line:10}
    }
    func TestMyError(t *testing.T){
        var e error = NewMyError() 
        switch e := e.(type){
            case nil:
                //no error  
            case *MyError: 
                l.Logf("error occurred on line:%d",e.Line)
            default: 
                //unkonw 
        }
    }
````
但是呢，也避免使用  Error Type 
也是和 Sentinel Error 一样的怀味道
- 增加了包之间源代码级别耦合   
- 也增大了你API面积 ，使得API变得脆弱  


Opaque errors  

只返回错误，不假定内容 

func f(){
    v ,err := bar.Xxx() 
    if err !=nil{
        // deal err 
    }
}

这种不透明的错误，调用者只知道成功失败与否，不知道具体内容，也不依赖具体错误类型，使得调用者和提供者之间 依赖最小  

但是，仅仅这种二分法的错误处理，是不够的，需要知道具体的错误，来以此做应对策略，譬如说超时啊，那么我的调用者依据策略重试几次 

这种情况，对 错误做类型断言，判断其行为。 

Assert error for behaviour ,not type  

少暴露类型，这样被人依赖后，扩展就很难了 

````
    func f(){
        v ,err := bar.Xxx()
        if err !=nil && bar.IsTemporary(err){
            // do sth 
        }
    }

    // bar.go  
    func IsTemporary(e error){
        type temporary interface{
            IsTemporary() bool
        }
        sw ,ok := e.(temporary) 
        return ok && sw.IsTemporary()
    }
````



但是 还是没彻底解决 获取更多上下文信息，和错误类型的 问题  

### 比较糟糕处理错误的方式  
1、将错误结果滞后处理，先处理正常逻辑。 譬如 
````
 err := f()  
 if err ==nil{
     //do stuff  
 }
 //错误处理
 ````
  
 > 这种做法最恶心，代码最难读，用if 嵌套，使得阅读代码视觉特别难受，你读着读着代码，得分辨读到的 代码是处在那一个嵌套层级里面，这是最强烈批判的写法  
 **应该及时处理错误，及时返回** 

 2、第二种，写多余的代码，脱裤子放屁  
 譬如
 ```` 
 err := f1() 
 if err !=nil{
     return err 
 } 
 err = f2() 
 if err !=nil{
     return err 
 }
 return nil  
````

你会看到 最后判断一下 err 不为nil 是多余的，直接return err 即可 。 
虽然抽身出来，你会很容易发现 这是脱裤子放屁的写法，可我刚刚就脱裤子放了一次屁  
如下： 
````

    type Header struct {
        Key, Value string
    }

    type Status struct {
        Code   int
        Reason string
    }

    func writeResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
        _, err := fmt.Fprintf(w, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
        if err != nil {
            return err
        }
        for _, h := range headers {
            _, err := fmt.Fprintf(w, "%s: %s", h.Key, h.Value)
            if err != nil {
                return err
            }
        }
        if _, err := fmt.Fprint(w, "\r\n"); err != nil {
            return err
        }

        // 我在下面代码段就 脱裤子放屁了。反思挖掘为何会这么做，究其原因还是陷进去了，已经进入了 if err != nil 惯性，已经没去思考了，靠着惯性去写代码了。
        _, err = io.Copy(w, body)
        if err != nil {
            return err
        }
        return nil
    }

````

**你会发现，上面代码有重复的  if err !=nil 的处理** 代码还能优化，精简 

那我们有一个技巧，把写入的错误暂存起来，等到返回时，一次性处理。优先后代码如下  
````
    type errWriter struct{
        io.Writer 
        err error 
    }

    func (e *errWriter) Write(buf []byte) (int,error){
        if e.err !=nil{
            return 0,e.err 
        }

        n:=0
        n,e.err = e.Writer.Write(buf)
        return n,e.err
    }

    func writeResponse(w io.Writer,st Status,headers []Header ,body io.Reader) error{
        sw := &errWriter{Writer:w}

        fmt.Fprintf(sw,"HTTP/1.1 %d %s\r\n",st.Code,st.Reason)

        for _,h := range headers{
            fmt.Fprintf(sw,"%s: %s\r\n",h.Key,h.Value)
        }
        fmt.Fprint(sw,"\r\n")
        
        io.Copy(sw,body)

        return sw.err 
    }
````







我是假设，譬如在 io.Copy 调用我们的 errWriter.Write ,
那么Copy 函数内做如下事情 
 
 err := errWriter.Write(xx) 
 if err !=nil{  //由于这里永远是nil 报错了不会return 代码接着往下执行 
     return err 
 }
 
 依赖写入的 buf 做一些什么事情，因为写入失败，导致后续操作就 crash 拉，会不会有这种风险。 


Wrap error   

github.com/pkg/errors  第三方库提供的  错误处理包，解决了 错误的更多上下文信息和错误堆栈的 问题   

它的用法是这样的 

````
    import "github.com/pkg/errors" 

    //sth code 

    func xx() error{
            
        err := f() 
        if err!=nil{
            return errors.Wrap(err,"wrap error")
        }
    }

    func main(){
        err := xx()

        if err !=nil{
            fmt.Printf("stack trace : %+v",err)
            err.Cause() == Sentinel error  
            //err.Cause() 返回原始错误 
        }
    }
````

Wrap Error 总结 
* 重用性高的包，只返回根错误，不要wrap 错误返回。 
    怎么理解呢？ 基础库、第三方库不要wrap error ，而应该 直接返回原始错误。应该是我们业务程序 才适合使用 wrap error ,包含更多上下文信息 返回到调用顶端处理 
* 如果函数不处理错误，就携带上足够多上下文（输入参数，返回结果等）包装错误返回。 
* 如果函数选择处理错误，那么就不要在把这个错误 向上返回了  


## close channel  
* 向 closed channel 发生数据会 引发 panic；而接收却不会，只是会拿到零值 
* 针对待缓冲的 channel ,如果 接收方尚未读取完数据，就已经被close 了，放心！队列中还剩多少数据没读完一条都不会少。即使关闭 已经在channel中的数据都是会被读到的，不会立马 读到零值 

### 单方向的channel 类型 
    chan<- T 只用于发送操作 
    <-chan T 只能用于接收操作 

其实我也分辨不出那个是发送哪个是接收，但只要记得，一旦代码中出现此类写法，就是单方向channel ，别不认识就行了 ，也不会有什么问题，编译器会纠正我们的写法错误 

## 小心 goroutine 泄露 
````
    func mirroredQuery()strig{
        resp := make(chan string)
        go func() {resp <- request("asia.gopl.io")}
        go func() {resp <- request("europe.gopl.io")}
        go func() {resp <- request("americas.gopl.io")}
        return <-resp
    }

````
上面的代码存在 goroutine 泄露问题 
假设 镜像 asia.gopl.io 最快响应，那其它两个慢请求goroutine 
将会因为 channel 无人接收 而被阻塞，永远不会退出，这就是 goroutine 泄露 
goroutine 不像 垃圾变量一样会被回收 

## select 多路复用
* 一个没有任务 case 的select 写作 select {} ,它会永远的等待下去 
* 对零值的 channel 发送和接收操作，永远会被阻塞；在select中对nil channel 操作，永远不会被select到 
例如下面 程序就会永远阻塞 
````
func TestNilChannel(t *testing.T) {
	ch := make(chan struct{})
	ch = nil
	go func() {
		fmt.Println("begin send")
		ch <- struct{}{}
		fmt.Println("after send")
	}()
	fmt.Println("begin receive")
	<-ch
	fmt.Println("after receive")
}
````
* 如果多个 case 准备就绪，那么select 会随机选择一个执行 
````
func TestRandomSelect(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 0
	for i := 1; i <= 20; i++ {
		select {
		case x := <-ch:
			fmt.Printf("receive %d\r\n", x)
		case ch <- i:
			fmt.Printf("send %d\r\n", i)
		}
	}
}
````
你会发现，他下一次执行的是 接收消息，还是发送消息，都是不确定的 。 




## Leave concurrency to the caller 
###  **把并行、后台执行、异步执行的 决定权交给 调用者决定,通常容易得多**

### never start a goroutine without konwing when it will stop 



## 除了 main 和 init 函数，其它任何地方都不要使用 os.Exit 和 log.Fatal  
> log.Fatal 内部调用了 os.Exit 。os.Exit 是无条件退出，不会执行defer 函数 
> 所以其它任务地方，都不允许调用 os.Exit 退出进程  

一些第三方库内，使用了 os.Exit 会让你很难受，因为你要对第三方库做降级，而第三方库的 os.Exit 直接将你进程退出了
你什么都控制不了，很尴尬



errgroup + functional options

一定要记住，channel 的发送方，才拥有 close channel 的资格，
因为 如果close 掉channel ,在往里写数据，会应发 panic  


TODO  以下主题要去 练习 
应用生命周期管理  errgroup + functional options 




## 并发观
* 不要去假设指令的执行顺序，你预测不到顺序。因为CPU 存在指令重排的优化。 
* 多核多缓存结构，导致程序执行和你预期结果不一致   MESI 协议 


TODO 以下主要要去扩展阅读  

Golang 内存管理 
        https://golang.org/ref/mem  
            翻译 https://segmentfault.com/a/1190000008230146    
                 https://blog.hugozhu.site/2013/04/20/31-golang-memory-model.html
        https://zhuanlan.zhihu.com/p/27807169 
COW: Copy On Write  
MESI: 缓存一致性协议   https://blog.csdn.net/muxiqingyang/article/details/6615199

context 应用场景 https://www.jianshu.com/p/6def5063c1eb 



## http 设置超时 
* http client 设置超时，在规定时间内还未响应，则中断接收响应 ，转去做别的事情 
* http server 设置超时，在规定时间内还未处理完请求，则放弃处理，直接回复client 说超时了 
````
    timeoutHandler := http.TimeoutHandler(http.DefaultServeMux, 5*time.Second, "timeout")
    err = http.ListenAndServe("localhost:8090", timeoutHandler)
````


## 分布式事务 

### 为何把付钱和发货拆解成两个动作？ 
答，本质上是因为高并发，发货是慢的。把慢的动作单独拆解出来，然后付完钱给出一个ticket ,让顾客等待 提高接待顾客的数量。 

### 暴露出口的回调接口，一定要保证幂等性，业务幂等性。 

### Best Effort 努力送达的模型，一定要预扣（预扣资源）。这样做的好处就可以回滚。 
譬如说你不预扣，先发资源，那么你回滚时发现用户已经消费了资源，根本回滚不了啦。 
先扣账，在发钱或发货或发道具 
