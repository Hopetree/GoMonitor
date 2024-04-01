# GoMonitor
主机信息监控客户端，用来向服务端推送当前主机信息

## 编译
直接在 main.go 所在目录进行编译，生成二进制文件 GoMonitor

```shell
go mod tidy
bo build
```

## 调试
不给任何参数，则只输出主机信息到控制台

```shell
./GoMonitor
```

输出内容：
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
  "udp": 1
}
```

## 后台运行

```shell
./GoMonitor ${SecretKey} ${SecretValue}  ${Interval} > /dev/null &
```
- SecretKey：密钥的Key
- SecretValue：密钥的值
- Interval：上报间隔(单位：秒)
