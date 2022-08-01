## NetWorkScan
> 利用配置文件控制程序，配置文件名：config.yml
```yaml
// 文件相关配置
file:
  // 带有需要扫描ip的文件名
  inFileName: in.xlsx
  // 输出扫描内容的文件名
  outFileName: out.xlsx
  // 扫描的工作表名和输出内容的工作表名
  sheet: Sheet1
  // 带扫描的内容是哪一列
  ip: 0
  port: 1
  user: 2
  password: 3
// 程序在扫描时的并发相关配置
burst:
  // 并发量
  burstNum: 2
// 远程连接的相关配置
connect:
  // 远程连接的等待时长
  timeout: 5
// 系统相关配置
os:
  // 没有在程序中添加的系统执行默认命令
  base:
    - cd / & ls
    - ls
  // 对应系统需要执行命令
  openEuler:
    - cd / & ls
```