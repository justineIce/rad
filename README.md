# rad 自动生成echo+model+api+mysql的web框架
> Automatic generation of echo + model + API + MySQL Web Framework.
### 快速使用
 * 默认参数

|参数|值|备注|
|:----:|:----:|:----:|
|sqltype|mysql|数据库类型|
|dbHost|127.0.0.1|数据库主机地址|
|dbPort|3306|数据库默认|
|dbName|test|数据库名称|
|dbUser|root|数据库用户名|
|dbPassword|123456|数据库密码|
|dbParameters|charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true|数据库连接字符串|
|table|*|表名|
|package|当前执行路径/example|项目包路径|
|json|--no-json|添加json标记|
|gorm|gorm|添加gorm标记|
|guregu|guregu|支持可为空值的字段类型|
 ````
 > go run main.go --package "github.com/devopsmi/rad/example" --json --gorm --guregu
 ````
* 自定义方式
 ````
 > go run main.go --package "github.com/devopsmi/rad/example" --dbHost 192.168.2.140 --dbName daas --dbUser cox --dbPassword "tqojg&WIGggN89bm"  --json --gorm --guregu
 ````
