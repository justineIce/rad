`# http配置
[http]
# http监听地址
host = "0.0.0.0"
# http监听端口
port = 8081
# http优雅关闭等待超时时长(单位秒)
shutdown_timeout = 30

# gorm配置
[gorm]
# 是否开启调试模式
debug = true
# 数据库类型(目前支持的数据库类型：mysql)
db_type = "mysql"
# 设置连接可以重用的最长时间(单位：秒)
max_lifetime = 7200
# 设置数据库的最大打开连接数
max_open_conns = 150
# 设置空闲连接池中的最大连接数
max_idle_conns = 50

# mysql数据库配置
[mysql]
# 连接地址
host = "{{.host}}"
# 连接端口
port= {{.port}}
# 用户名
user = "{{.user}}"
# 密码
password = "{{.password}}"
# 数据库
db_name = "{{.name}}"
# 连接参数
parameters = "{{.parameters}}"

# redis配置
[redis]
# 地址
addr = "{{.rdHost}}:{{.rdPort}}"
# 密码
password = "{{.rdPassword}}"
# 数据库
database = {{.rdDatabase}}
# 最大连接数
maxOpenConns = 10
# 最大空闲连接数
maxIdleConns = 10

# 文件上传
[fileUpload]
# 前缀地址
basePath = ""
# 存放位置
path = "files"
# 域名
doMain = ""
# 文件最大大小
maxFileSize = 52428800
# 文件格式限制
extFilter = "255216,7173,13780,6677,51200,208207,3780,8075,119101,97115"

# 日志配置
[log]
path="{{.name}}.log"
level=7
maxdays=10
separate="error"`
