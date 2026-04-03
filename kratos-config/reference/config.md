---
id: config
title: Configuration
description: Kratos configuration supports multiple sources, and the config will be merged into map[string]interface{}, then you can get the value content through Scan or Value.
keywords:
  - Go 
  - Kratos
  - Toolkit
  - Framework
  - Microservices
  - Protobuf
  - gRPC
  - HTTP
---
The best practice for configuring microservices or cloud-native applications is to separate configuration files from application code - do not put configuration files in code repositories or package them into container images, but mount the configuration files or load them directly from the configuration center at runtime. Kratos' config component is used to help applications load configurations from various sources.

## Design Philosophy
### 1. Support for Multiple Configuration Sources
Kratos defines standardized [Source and Watcher interfaces](https://github.com/go-kratos/kratos/blob/main/config/source.go) to adapt to various configuration sources.

The framework comes with built-in implementations of [local file (file)](https://github.com/go-kratos/kratos/tree/main/config/file) and [environment variable (env)](https://github.com/go-kratos/kratos/tree/main/config/env).

In addition, in [contrib/config](https://github.com/go-kratos/kratos/tree/main/contrib/config), we also provide adapters for the following configuration centers:

* [Apollo](https://github.com/go-kratos/kratos/tree/main/contrib/config/apollo)
* [Consul](https://github.com/go-kratos/kratos/tree/main/contrib/config/consul)
* [Etcd](https://github.com/go-kratos/kratos/tree/main/contrib/config/etcd)
* [Kubernetes](https://github.com/go-kratos/kratos/tree/main/contrib/config/kubernetes)
* [Nacos](https://github.com/go-kratos/kratos/tree/main/contrib/config/nacos)
* [Polaris](https://github.com/go-kratos/kratos/tree/main/contrib/config/polaris)

If the above configuration loading methods do not cover your environment, you can also implement the interface to adapt your own configuration loading method.

### 2. Support for Multiple Configuration Formats
The config component reuses the deserialization logic in `encoding` as the configuration parsing. It supports the following formats by default:

* JSON
* Protobuf
* XML
* YAML

The framework will parse the configuration file based on its type by matching the corresponding codec. You can also implement [Codec](https://github.com/go-kratos/kratos/blob/main/encoding/encoding.go#L10) and register it with the `encoding.RegisterCodec` method to parse other formats of configuration files.

The extraction of configuration file types varies slightly depending on the specific implementation of the configuration source. The built-in file source uses the file extension as the file type. Please refer to the documentation of the other configuration source plugins for their specific logic.

### 3. Hot Reloading
Kratos' config component supports hot reloading of configurations. You can use the configuration center to update the configuration of a service online without re-deploying, stopping, or restarting the service, and modify some behaviors of the service.

### 4. Configuration Merge
In the config component, the configurations (files) from all configuration sources will be read one by one, parsed into maps, and merged into one map. Therefore, after loading, you don't need to consider the file names or search for configurations by file names. Instead, you can use the structure of the contents to index the values of the configurations. When designing and writing configuration files, please note that **the root-level keys in different configuration files should not be duplicated, otherwise they may be overwritten**.

When using the configuration, you can use `.Value("foo.bar")` to directly get the value of a specific field, or use the `.Scan` method to read the entire map into a specific structure. Please refer to the following sections for specific usage.

## Usage
### 1. Initialize Configuration Sources
Use file, which loads from a local file:
Here, the `path` is the path to the configuration file. You can also specify a directory name, and all files in the directory will be parsed and loaded into the same map.
```go
import (
    "github.com/go-kratos/kratos/v2/config"
    "github.com/go-kratos/kratos/v2/config/file"
)

path := "configs/config.yaml"
c := config.New(
    config.WithSource(
        file.NewSource(path),
    ),
)

```

### 2. Read Configuration
First, define a structure to parse the fields of the configuration file.
```go
var v struct {
  Service struct {
    Name    string `json:"name"`
    Version string `json:"version"`
  } `json:"service"`
}
```

Using the initialized config instance, call the `.Scan` method to read the configuration file into the structure.
```go
// Unmarshal the config to struct
if err := c.Scan(&v); err != nil {
  panic(err)
}
```

You can use the `.Value` method of the config instance to get the content of a specific field.
```go
name, err := c.Value("service.name").String()
if err != nil {
  panic(err)
}
```

### 3. Watch Configuration Changes
You can use the `.Watch` method to listen for changes to a specific field in the configuration.
```go
if err := c.Watch("service.name", func(key string, value config.Value) {
  fmt.Printf("config changed: %s = %v\n", key, value)
}); err != nil {
  log.Error(err)
}
```
