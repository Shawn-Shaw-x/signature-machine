# 📦 signature-machine web3通用型签名机
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/Shawn-Shaw-x/signature-machine)
![Go Version](https://img.shields.io/badge/go-1.20+-brightgreen)
![Build](https://img.shields.io/github/actions/workflow/status/Shawn-Shaw-x/signature-machine/go.yml?label=build)
![Tests](https://img.shields.io/github/actions/workflow/status/Shawn-Shaw-x/signature-machine/test.yml?label=tests)
![License](https://img.shields.io/github/license/Shawn-Shaw-x/signature-machine)
![Stars](https://img.shields.io/github/stars/Shawn-Shaw-x/signature-machine?style=social)

> 本项目是一个基于 Go 实现的通用性签名机服务。提供 ECDSA、EdDSA 地址批量生成能力，并支持消息签名。

> ps：项目大部分代码基于交易所签名机抽象、简化而成。
---

## ✨ 功能特性

- 🚀 支持批量地址生成
- 🔐 多种签名方式支持（ECDSA、EdDSA）
- ⭐️ 程序内安全保障、私钥永不泄漏
- 🛠 高性能、轻量级设计
- 📦 提供 Cli 工具 / Grpc 服务

---

## 📥 安装与使用

### 环境要求

- Go 1.20+
- Git
- Make

### 获取源码

```bash
git clone https://github.com/Shawn-Shaw-x/signature-machine
cd signature-machine
```
### 构建项目
```bash
make ./signature
```
### 运行项目
```bash
./signature rpc
```
```bash
grpcui -plaintext 127.0.0.1:8983
```
### 运行测试
```bash
make test
```
## 🗂 项目结构
```bash
signature-machine/
├── bin/            # protobuf 生成管理
├── cmd/            # 程序入口
├── common/         # 公共库
├── config/         # 默认配置文件
├── data/         # leveldb 生成文件
├── flags/           # 环境变量定义
├── leveldb/        # leveldb 代码
├── protobuf/        # protobuf 代码
├── services/        # Grpc 代码及接口实现
├── ssm/        # ecdsa、eddsa 库封装
├── Makefile    # shell 命令管理
├── go.mod
├── README.md
```

## 📚 架构图
- 交易所签名机架构
  ![image.png](https://img.learnblockchain.cn/attachments/2025/05/oTgu7VHp6817db0fcfbed.png)

- 签名机内部架构
  ![image.png](https://img.learnblockchain.cn/attachments/2025/05/GNWDLsOa6817db7430004.png)

- TEE 地址生成流程
  ![image.png](https://img.learnblockchain.cn/attachments/2025/05/YqvaQliK68165d1fe7087.png)

- TEE 签名流程
  ![image.png](https://img.learnblockchain.cn/attachments/2025/05/OibBv1ea6816608a90d25.png)


