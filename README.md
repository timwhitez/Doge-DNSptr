# Doge-DNSptr

🐶 通过PTR记录使用IP反查内网域名

🐶 Reverse lookup of internal domain names using IP through PTR records

## 项目简介 (Project Introduction)

Doge-DNSptr 是一个用Go语言编写的工具，用于通过PTR记录对IP地址进行反向DNS查询，以发现内网域名。这个工具特别适用于网络探测和信息收集。

Doge-DNSptr is a tool written in Go for performing reverse DNS lookups on IP addresses using PTR records to discover internal domain names. This tool is particularly useful for network reconnaissance and information gathering.

## 特性 (Features)

- 并发查询以提高效率
- 可自定义DNS服务器
- 支持从文件读取IP地址列表
- 结果输出到文件

- Concurrent queries for improved efficiency
- Custom DNS server option
- Supports reading IP address list from a file
- Results output to a file

## 安装 (Installation)

确保您的系统上安装了Go。然后，您可以通过以下命令克隆并构建项目：

Ensure you have Go installed on your system. Then, you can clone and build the project with the following commands:

```bash
git clone https://github.com/timwhitez/Doge-DNSptr.git
cd Doge-DNSptr
go build
```

## 使用方法 (Usage)

运行编译后的二进制文件，使用以下命令行参数：

Run the compiled binary with the following command-line arguments:

```bash
./Doge-DNSptr -f <ip_file> -o <output_file> -d <dns_server> -t <threads>
```

参数说明 (Arguments):
- `-f`: 包含IP地址列表的文件路径（必需）
- `-o`: 输出结果的文件路径（必需）
- `-d`: 自定义DNS服务器地址（可选，默认使用系统DNS）
- `-t`: 并发线程数（可选，默认为10）

- `-f`: Path to the file containing the list of IP addresses (required)
- `-o`: Path to the output file for results (required)
- `-d`: Custom DNS server address (optional, uses system DNS by default)
- `-t`: Number of concurrent threads (optional, default is 10)

示例 (Example):
```bash
./Doge-DNSptr -f ips.txt -o results.txt -d 8.8.8.8 -t 20
```

## 输入文件格式 (Input File Format)

输入文件应包含每行一个IP地址，支持IPv4和IPv6：

The input file should contain one IP address per line, supporting both IPv4 and IPv6:

```
192.168.1.1
10.0.0.1
2001:db8::1
```

## 输出格式 (Output Format)

输出文件将包含成功的反向查询结果，格式为：

The output file will contain successful reverse lookup results in the format:

```
IP地址 -> 域名
IP Address -> Domain Name
```

## 注意事项 (Notes)

- 该工具主要用于内网环境的域名发现。
- 请确保您有权限对目标IP地址执行DNS查询。
- 大量并发查询可能会对网络和DNS服务器造成压力，请谨慎使用。

- This tool is primarily intended for domain name discovery in internal network environments.
- Ensure you have permission to perform DNS queries on the target IP addresses.
- Large numbers of concurrent queries may stress the network and DNS servers, use with caution.

## 贡献 (Contributing)

欢迎提交问题报告、功能请求和Pull请求。对于重大更改，请先开启一个问题进行讨论。

Feel free to submit issue reports, feature requests, and Pull Requests. For major changes, please open an issue first to discuss what you would like to change.

## 免责声明 (Disclaimer)

本工具仅用于教育和研究目的。使用者应负责任地使用本工具，并遵守所有适用的法律和法规。作者不对因使用本工具而导致的任何滥用或损坏负责。

This tool is for educational and research purposes only. Users are responsible for using this tool responsibly and in compliance with all applicable laws and regulations. The author is not responsible for any misuse or damage caused by this tool.
