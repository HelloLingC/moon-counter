# Moon-Counter

English | [中文](readme_cn.md)

A fast, simple & easy-to-use webpage visitor counter, but not only limited to websites.

![rule34](/assets/rule34.svg)

[> Preview <](https://mini.moonlab.top/post/20231224-14/)

With a visual admin panel, put Moon-Counter at every corner

#### 🚀 Fast and Simple

#### 🎉 Easy Deployment

Run counter server with only one binary file, zero dependency. No annoying complex installation

#### 🔒 Secure [CORS]() Support

 Make it hard for strangers to use your counter service without permisson to tally for them

#### 🌟 SQLite Database.

Reeeallly easy to control and move

## Image Counter

### Common Method
Make sure id argument is unique for every webpage

```
# You can use this in Github Profile
![](https://yoursite.com/counter/img?id=uniqueID)

<img src="//yoursite.com/counter/img?id=uniqueID"></img>
```

### Secure CORS

Unique id arg is automatically handled

If cors is on in the config file, server will check whether the request origin is vaild, and return cors resources.
In this case, you should only use image counter in this way

```
<script src="//yoursite.com/moon-counter/js/img"></script>
<img id="moon-counter-img"></img>
```

## Text Counter

Add the following code to where you wanna place a text counter.

```
<script src="//yoursite.com/moon-counter/js"></script>
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


# License

MIT






