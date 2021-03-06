# 基于 kratos 微服务框架，而做的 album 商城项目 

## 《项目要点》

	- 微服务架构（BFF、Service、Admin、Job、Task 分模块）
	- API 设计（包括 API 定义、错误码规范、Error 的使用）   
	- gRPC 的使用		  
	- Go 项目工程化（项目结构、DI、代码分层、ORM 框架） 
	- 并发的使用（errgroup 的并行链路请求）
	- 微服务中间件的使用（ELK、Opentracing、Prometheus、Kafka）  
	- 缓存的使用优化（一致性处理、Pipeline 优化） 

## 《项目结构(album/service) 》
````
app\album\service
					\cmd
						\server   
							main.go
							wire.go 
							wire_gen.go 
					\configs
						config.yaml
						registry.yaml 
					\internal
						\biz 
						\conf 
						\data
						\server 
						\service  
app\album\admin
app\album\interface
app\album\job
app\album\task
````


## 《新建工程步骤 (album/service) 》
~工程，真是一个考验组织心力的活啊。~
````
1. 新建了 configs/configs.yaml 
2. 新建了 conf/conf.proto 并 protoc 生成了 conf.pb.go 
3. 新建了 api/album.proto 并 protoc album_grpc.pb.go 和 album.pb.go 
4. 新建了 internal/biz/album.go 主要定义了domain object struct album{},还有 AlbumRepo 接口，还有 AlbumUseCase 已经方法
5. 新建了 internal/service/album.go 主要是定义了 AlbumService struct{} ,它的作用是 实现 album.proto 定义的服务，当然他依赖 biz 
6. 新建了 internal/server/grpc.go 主要提供了NewGRPCServer 将 AlbumService 注入到 grpc 服务中  
7. 新建了 internal/data/album.go 主要实现了 biz.AlbumRepo 接口 ，定义PO album ,与数据库交互  
8. 新建了 cmd/main.go 胶水代码，主要是加载config,初始化日志,初始化 conf proto模型, 
	提供 ProviderSet 需要的参数，注入，并最终 new出一个 app  
		依赖关系 
			conf *conf.Data -> NewDB() -> *gorm.DB 
			*gorm.DB -> NewData() -> *data.Data{包裹了 db链接} 
			*data.Data -> NewAlbumRepo() -> data.albumRepo 实现了 interface  biz.AlbumRepo 
			biz.AlbumRepo -> NewAlbumUseCase() -> AlbumUseCase{包裹了一个实现了biz.AlbumRepo 的对象 }
			AlbumUseCase-> NewAlbumService() -> AlbumService{包裹了一个AlbumUseCase对象} 
			AlbumService & conf *conf.Server ->  NewGRPCServer()  -> *grpc.Server  
			*grpc.Server  -> newApp() -> *kratos.App
			data *conf.Data,*conf.Server,logger log.Logger 
9. 编译 生成二进制文件 
	cd D:\github.com\RoseWsp\graduate\app\album\service\cmd\server 
	go build -ldflags "-X main.Version=1.0.0"
````


protoc --go_out=. --go_opt=paths=source_relative ^
    --go-grpc_out=. --go-grpc_opt=paths=source_relative ^
    album.proto

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    conf.proto

INSERT INTO albums
  (title, artist, price,create_at)
VALUES
  ('Blue Train', 'John Coltrane', 56.99,now()),
  ('Giant Steps', 'John Coltrane', 63.99,now()),
  ('Jeru', 'Gerry Mulligan', 17.99,now()),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98,now());
  
  

## 《表设计》
	# albums 专辑表 

	# orders 订单表 - 记录用户下单 ,订单表超简单，一个订单就买一张专辑。

	# integrating  积分表- 用户下单赚积分  

	# integrating_count 积分汇总，统计每个用户的积分总数 

	# operation  操作记录表，主要用在 admin 后台修改数据的流水信息，用于业务数据"对账"
	
## 《grpc服务》
````
	album/service    addr: 0.0.0.0:8001	-> 主体的微服务 ,查专辑，查订单，下订单 等  
	album/admin      addr: 0.0.0.0:8002	-> 运营管理后端服务 修改订单和专辑信息，并保存修改记录  
	album/job 		 addr: 0.0.0.0:8003	-> 从kafka 中捞取订单，生成积分记录     
	album/task 		 addr: 0.0.0.0:8004  -> 汇总计算 积分 
	album/interface  addr: 0.0.0.0:8005 -> BFF 分别从 service 和task 查 用户的订单和积分数据，组合起来返回给调用者 
````
