#### 安装依赖中间件

安装 redis

  docker run -d --name redis -p 6379:6379 redis:6.2.14

安装 mysql

  docker run -d --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:8.0.38

安装 etcd

  docker run -d --name etcd -p 2379:2379 -p 2380:2380 -e ALLOW_NONE_AUTHENTICATION=yes -e ETCD_ADVERTISE_CLIENT_URLS=http://etcd:2379 bitnami/etcd

安装 jaeger

  docker run -d --name jaeger -p 16686:16686 -p 14268:14268 bitnami/jaeger

#### 构建项目

  go mod tidy
  
  go mod download
  
> go mod vendor  如果需要把依赖安装到当前目录
  
