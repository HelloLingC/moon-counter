# Moon-Counter

A fast, simple & easy-to-use webpage visitor counter, but not only limited to websites.

With a visual admin panel, put Moon-Counter at every corner

🚀 Fast and Simple

🎉 Self-Host & Easy-To-Use
Deploy with only one file, zero dependency. No annoying complex installation

🔒 Secure
Support [CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS), make it hard for strangers to use your self-host service without permisson

🌟 SQLite Database.
Reeeallly easy to control and move

Two modes to choose, text & image

## Image Counter

```
# Markdown style
# Make sure id arg is unique for each webpage
# You can use this in Github Profile
![]()[https://yoursite.com/counter/img?id=uniqueID]

# HTML style
# Unique id arg is automatically handled
# If cors is on, you should only use image counter in this way
<script src="https://yoursite.com/moon-counter/js/img"></script>
<img id="moon-counter-img"></img>
```

## Text Counter

Add the following code to where you wanna place a text counter.

```
<script src="https://yoursite.com/moon-counter/js"></script>
<span id="moon-counter"></span>
```

# Self Host

Download the compressed file in Release Page.

1.Decompress

```bash
$ tar -xf moon-counter.tar.gz
```

2.Adjust the config.yaml

```bash
nano config.yaml
```

3.Run

```bash
$ moon-counter
```

For more details and configuration help, Please visit [my blog](https://mini.moonlab.top/)

# Credits


# Lisence

MIT






