module github.com/devopsmi/rad

go 1.12

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.40.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190605123033-f99c8df09eb5
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190510132918-efd6b22b2522
	golang.org/x/image => github.com/golang/image v0.0.0-20190523035834-f03afa92d3ff
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190607214518-6fa95d984e88
	golang.org/x/net => github.com/golang/net v0.0.0-20190607181551-461777fb6f67
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190609082536-301114b31cce
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190608022120-eacb66d2a7c3
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.6.0
	google.golang.org/appengine => github.com/golang/appengine v1.6.1
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190605220351-eb0b1bdb6ae6
	google.golang.org/grpc => github.com/grpc/grpc-go v1.21.1
)

require (
	github.com/denisenkom/go-mssqldb v0.0.0-20190515213511-eb9f6a1743f3
	github.com/droundy/goopt v0.0.0-20170604162106-0b8effe182da
	github.com/go-sql-driver/mysql v1.4.1
	github.com/guregu/null v3.4.0+incompatible
	github.com/jimsmart/schema v0.0.0-20181113191328-8d0563922e25
	github.com/jinzhu/gorm v1.9.8
	github.com/jinzhu/inflection v0.0.0-20190603042836-f5c5f50e6090
	github.com/julienschmidt/httprouter v1.2.0
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lib/pq v1.1.1
	github.com/mattn/go-sqlite3 v1.10.0
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/serenize/snaker v0.0.0-20171204205717-a683aaf2d516
)
