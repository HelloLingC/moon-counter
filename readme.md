# Moon-Counter

English | [中文](readme_cn.md)

A fast, simple & easy-to-use webpage visitor counter, but not only limited to websites.

![rule34](https://mini.moonlab.top/post/20231224-14/rule34.svg)

[> Preview <](https://mini.moonlab.top/post/20231224-14/)

#### 🚀 Fast and Simple

#### 🎉 Easy Deployment

Run counter server with only one binary file, zero dependency. No annoying complex installation

#### 🔒 Secure [CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) Support

 Make it hard for strangers to use your counter service without permisson to tally for them

#### 🌟 SQLite Database.

Reeeallly easy to control and move

# Public Counter

You can use my public MoonCounter instance. But I don't have responsibility for its stablility. In fact, no one does. But as long as my blog website continues, the counter service is always public.
But i still recommand you self-host your owned service.

**Ofc, the use of my service is prohibited for illegal websites.**

## Image Counter

**Replace the UniqueID parameter with the webpage's url. At least it must be unique to identify your page.**

```
![](https:///counter.moonlab.top/counter/img?id=UniqueID)

<img src="//counter.moonlab.top/counter/img?id=UniqueID"></img>
```

If you don't wanna handle UniqueID by yourself, you can use this. It will automatically handle the id param. You can just directly put it on your webpage to count.

```
<script src="//counter.moonlab.top/moon-counter/js/img"></script>
<img id="moon-counter-img"></img>
```

## Text Counter

And this will only output plain count text. You also don't need to handle the uniqueid parameter by yourself.

```
<script src="//counter.moonlab.top/moon-counter/js"></script>
<span id="moon-counter"></span>
```

# Self Host

Download the compressed file in [Release Page](/releases).

1.Decompress

```bash
$ tar -xf moon-counter.tar.gz
```

2.Adjust the config.yaml

```bash
$ nano config.yaml
```

3.Run

```bash
$ ./moon-counter
```

For more details and configuration help, Please visit [my blog](https://mini.moonlab.top/post/20231224-14/)

# Credits

visit https://mini.moonlab.top/post/20231224-14/

# License

MIT






