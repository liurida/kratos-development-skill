.
├── go.mod
├── go.sum
├── LICENSE
├── README.md
├── api
│   └── helloworld
│       ├── errors
│       │   ├── helloworld.pb.go
│       │   ├── helloworld.proto
│       │   └── helloworld_errors.pb.go
│       └── v1
│           ├── greeter.pb.go
│           ├── greeter.proto
│           ├── greeter_grpc.pb.go
│           └── greeter_http.pb.go
├── cmd
│   └── server
│       ├── main.go
│       ├── wire.go
│       └── wire_gen.go
├── configs
│   └── config.yaml
└── internal
    ├── conf
    │   ├── conf.pb.go
    │   └── conf.proto
    ├── data
    │   ├── README.md
    │   ├── data.go
    │   └── greeter.go
    ├── biz
    │   ├── README.md
    │   ├── biz.go
    │   └── greeter.go
    ├── service
    │   ├── README.md
    │   ├── greeter.go
    │   └── service.go
    └── server
        ├── grpc.go
        ├── http.go
        └── server.go
