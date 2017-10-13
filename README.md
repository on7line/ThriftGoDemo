# A Thrift Demo

[https://qiknow.com/](https://qiknow.com/thrift%E4%B8%8Egolang%E5%BA%94%E7%94%A8/)

install `git.apache.org/thrift.git/lib/go/thrift`

```sh
go get git.apache.org/thrift.git/lib/go/thrift
```


├── README.md
└── thrift
    ├── handler
    │   ├── user.go
    │   └── user_test.go
    ├── idl
    │   ├── Base.thrift
    │   ├── Req.thrift
    │   ├── Resp.thrift
    │   └── UserService.thrift
    └── proto
        ├── base
        │   ├── Base-consts.go
        │   ├── Base.go
        │   └── GoUnusedProtection__.go
        ├── req
        │   ├── GoUnusedProtection__.go
        │   ├── Req-consts.go
        │   └── Req.go
        ├── resp
        │   ├── GoUnusedProtection__.go
        │   ├── Resp-consts.go
        │   └── Resp.go
        └── userservice
            ├── GoUnusedProtection__.go
            ├── UserService-consts.go
            └── UserService.go

