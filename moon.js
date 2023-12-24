// Text
fetch('//%s/counter/text')
    .then(r => {
		return r.text();
    })
	.then(d => {
		document.getElementById("moon-counter").innerText = d;
	})
    .catch(e => {
        console.error(e);
    });

// Img
fetch('//%s/counter/img')
    .then(r => {
		return r.text();
    })
	.then(d => {
		document.getElementById("moon-counter-img").src='data:image/svg+xml,' + d;
	})
    .catch(e => {
        console.error(e);
    });