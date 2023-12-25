// Text
fetch('//%s/counter/text?id=' + window.location.href)
    .then(r => {
		return r.text();
    })
	.then(d => {
		document.getElementById("moon-counter").innerText = d;
	});

// Img
fetch('//%s/counter/img?id=' + window.location.href)
    .then(r => {
		return r.text();
    })
	.then(d => {
		document.getElementById("moon-counter-img").src='data:image/svg+xml,' + d;
	});