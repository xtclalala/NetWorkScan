## NetWorkScan
> 连接linux，并执行相对应命令，利用配置文件控制程序

命令行参数
1. path
   - 执行配置文件路径，默认：../config.yml
   - 例子：xxxx --path=../config.yml

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
  // 没有识别出系统的，执行默认命令
  base:
    - cd / & ls
    - ls
  // 识别出对应系统，并执行一下命令
  openEuler:
    - cd / & ls
```