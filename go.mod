module jinbao-cs

go 1.16

replace (
	cloud.google.com/go => github.com/GoogleCloudPlatform/google-cloud-go v0.33.1
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20181112202954-3d3f9f413869
	golang.org/x/lint => github.com/golang/lint v0.0.0-20181026193005-c67002cb31c3
	golang.org/x/net => github.com/golang/net v0.0.0-20181114220301-adae6a3d119a
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20181128211412-28207608b838
	golang.org/x/sync => github.com/golang/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys => github.com/golang/sys v0.0.0-20181116161606-93218def8b18
	golang.org/x/text => github.com/golang/text v0.3.1-0.20181030141323-6f44c5a2ea40
	golang.org/x/time => github.com/golang/time v0.0.0-20181108054448-85acf8d2951c
	golang.org/x/tools => github.com/golang/tools v0.0.0-20181117154741-2ddaf7f79a09
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20191204190536-9bdfabe68543
	google.golang.org/appengine => github.com/golang/appengine v1.3.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20181109154231-b5d43981345b
	google.golang.org/grpc => github.com/grpc/grpc-go v1.2.1-0.20181115212939-ef2b8e2f53fc
	honnef.co/go/tools => github.com/dominikh/go-tools v0.0.0-20180920025451-e3ad64cb4ed3
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/snappy v0.0.1 // indirect
	github.com/lib/pq v1.7.0
	github.com/mojocn/base64Captcha v1.3.5
	github.com/spf13/viper v1.10.1
	github.com/ugorji/go v1.1.7 // indirect
	go.uber.org/zap v1.17.0
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
	golang.org/x/net v0.0.0-20211029224645-99673261e6eb
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	xorm.io/xorm v1.0.3
)
