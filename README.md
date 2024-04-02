# GoMonitor
主机信息监控客户端，用来向服务端推送当前主机信息

## 编译
直接在 main.go 所在目录进行编译，生成二进制文件 GoMonitor

```shell
go mod tidy
go build
```

## 命令使用

### 命令帮助

```shell
./GoMonitor --help
```

输出帮助指导：

```shell
NAME:
   GoMonitor - 一个简单的agent客户端，用于采集主机信息并上报到服务端

USAGE:
   GoMonitor [global options] command [command options] 

VERSION:
   0.0.1

COMMANDS:
   collect  采集主机信息
   decrypt  解密密钥，并显示原文
   report   采集主机信息，并解析密钥，将信息上报到服务端，非调试模式不输出任何信息
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

```

### 采集命令
只采集主机信息并输出，不上报

```shell
./GoMonitor collect -i 5
```

```json
{
  "interval": 5,
  "uptime": 178837,
  "system": "linux-3.10.0-1160.99.1.el7.x86_64-x86_64-centos-7.9.2009",
  "cpu_cores": 2,
  "cpu_model": "AMD Ryzen 7 5700U with Radeon Graphics",
  "cpu": 0.499999999996362,
  "load_1": "0.00",
  "load_5": "0.01",
  "load_15": "0.05",
  "memory_total": 1927139328,
  "memory_used": 473124864,
  "swap_total": 2147479552,
  "swap_used": 1048576,
  "hdd_total": 51496095744,
  "hdd_used": 20355805184,
  "network_in": "0B",
  "network_out": "0B",
  "process": 171,
  "thread": 249,
  "tcp": 7,
  "udp": 1,
  "version": ""
}
```

### 上报命令

```shell
./GoMonitor report --help
```

```shell
NAME:
   GoMonitor report - 采集主机信息，并解析密钥，将信息上报到服务端，非调试模式不输出任何信息

USAGE:
   GoMonitor report [command options] [arguments...]

OPTIONS:
   --key value, -k value       解密Key
   --secret value, -s value    待解密的密文
   --interval value, -i value  采集间隔时间（默认值：6） (default: 6)
   --cmd value                 用来采集服务版本信息的命令，可以使用带管道符的shell命令 (default: "echo ''")
   --debug, -d                 调试模式，开启则会输出采集信息，并上报一次信息到服务端 (default: false)
   --help, -h                  show help

```
