# Go Calculator

> 关于我，欢迎关注：
>
> 个人主页：[mrcai.space](https://mrcai.space)
>
> Github 主页：[MrCaiDev](https://github.com/MrCaiDev)
>
> 个人邮箱：1014305148@qq.com
>
> 工作邮箱（不常用）：yuwangcai@std.uestc.edu.cn

## 项目介绍

本项目为电子科技大学《Go语言与区块链技术》半期课设。

实现了一个简易的网页计算器，前端使用`HTML + CSS + Javascript`，后端使用`Golang`，前后端间通过`websocket`通信。

## 使用方法
```bash
go run ip:port
```
#### 编译运行
使用
```bash
go build -ldflags "-s -w"
```
编译出独立的可执行文件，双击运行。同时在浏览器中打开 http://localhost:1234 查看效果。
#### 一键启动
运行`run.bat`，自动打开服务器和网页。
#### 分步启动
  - 在根下执行`go run main.go`
  - 在浏览器中打开 http://localhost:1234

## 注意事项

如果遇到防火墙拦截，请选择“允许访问”，否则服务器无法启动。

## 免责声明

本项目开源仅供交流学习之用，不保证准确性与泛用性，作者也一贯反对抄袭挪用的行为。

若有因借鉴本项目源码而导致意外结果者，如代码执行出错、被发现抄袭等，项目作者不承担相应责任。
