# Moon-Counter

[English]((readme.md)) | 中文

快速，简单 & 易于使用的网页浏览量计数，但并不只局限于网站。

![rule34](/assets/rule34.svg)

[> 预览 <](https://mini.moonlab.top/post/20231224-14/)

#### 🚀 Fast and Simple

#### 🎉 部署简单

只需一个二进制文件即可启动计数服务器，零依赖。没有繁琐的安装过程。

#### 🔒 安全的 [CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) 支持

让陌生人难以私自使用你的计数服务器，来为他们计数

#### 🌟 SQLite Database.

易于控制和搬迁

## 图片计数器

### 普通方法

请确保每个网页的 id 参数是独一无二的

```
# You can use this in Github Profile
![](https://yoursite.com/counter/img?id=uniqueID)

<img src="//yoursite.com/counter/img?id=uniqueID"></img>
```

### 安全 CORS

Unique id 参数会自动被处理

如果在配置文件中启用了 CORS，服务器将检查请求的来源是否合法，并返回 CORS 资源。在这种情况下，你应该仅以这种方式使用图片计数器。

```
<script src="//yoursite.com/moon-counter/js/img"></script>
<img id="moon-counter-img"></img>
```

## 文字计数器

将以下 html 代码放在你想要计数的地方

```
<script src="//yoursite.com/moon-counter/js"></script>
<span id="moon-counter"></span>
```

# 许可协议

MIT
