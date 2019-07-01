module github.com/devopsmi/rad

go 1.12

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/droundy/goopt v0.0.0-20170604162106-0b8effe182da
	github.com/go-playground/locales v0.12.1 // indirect
	github.com/go-playground/universal-translator v0.16.0 // indirect
	github.com/guregu/null v3.4.0+incompatible
	github.com/jimsmart/schema v0.0.0-20181113191328-8d0563922e25
	github.com/jinzhu/gorm v1.9.8
	github.com/jinzhu/inflection v0.0.0-20190603042836-f5c5f50e6090
	github.com/kr/pretty v0.1.0 // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.2.8
	github.com/leodido/go-urn v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/noaway/dateparse v0.0.0-20171117034806-ad2b19d7b298
	github.com/serenize/snaker v0.0.0-20171204205717-a683aaf2d516
	github.com/stretchr/testify v1.3.0
	github.com/valyala/fasttemplate v1.0.1 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.29.0
)

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
