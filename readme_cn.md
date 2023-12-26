# Moon-Counter

[English]((readme.md)) | 中文

快速，简单 & 易于使用的网页浏览量计数，但并不只局限于网站。

![rule34](https://mini.moonlab.top/post/20231224-14/rule34.svg)

[> 预览 <](https://mini.moonlab.top/post/20231224-14/)


#### 🚀 Fast and Simple

#### 🎉 部署简单

只需一个二进制文件即可启动计数服务器，零依赖。没有繁琐的安装过程。

#### 🔒 安全的 [CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) 支持

让陌生人难以私自使用你的计数服务器，来为他们计数

#### 🌟 SQLite Database.

易于控制和搬迁


# 公用计数器

你可以使用我的 MoonCounter 实例。但是我对其稳定性不负责任。实际上，没有人需要对此负责。但只要我的博客网站持续存在，计数器服务将始终是公开的。
但我仍然建议你自己托管您拥有的服务。

**当然，禁止非法网站使用我的服务**

## 图片计数器

**用网页的 URL 替换 UniqueID 参数。至少它必须是唯一的，以标识你的页面。**

```
![](https:///counter.moonlab.top/counter/img?id=UniqueID)

<img src="//counter.moonlab.top/counter/img?id=UniqueID"></img>
```

如果你不想自己处理 UniqueID，你可以使用这个。它将自动处理 id 参数。你只需直接将其放在你的网页上进行计数。

```
<script src="//counter.moonlab.top/moon-counter/js/img"></script>
<img id="moon-counter-img"></img>
```

## 文本计数器

而这将仅输出纯文本计数。你也不需要自己处理 uniqueid 参数。

```
<script src="//counter.moonlab.top/moon-counter/js"></script>
<span id="moon-counter"></span>
```

# 自己部署

下载最新版本：[Release Page](/releases).

1.解压

```bash
$ tar -xf moon-counter.tar.gz
```

2.修改 config.yaml

```bash
$ nano config.yaml
```

3.运行

```bash
$ ./moon-counter
```

更多细节和配置教程，请查看 [我的博客](https://mini.moonlab.top/post/20231224-14/)

# 许可协议

MIT
