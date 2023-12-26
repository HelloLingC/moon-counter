fetch('//%s/counter/text?id=' + window.location.href)
    .then(r => {
		return r.text();
    })
	.then(d => {
		document.getElementById("moon-counter").innerText = d;
	});