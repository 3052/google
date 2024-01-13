// If the page was opened with an anchor (e.g. #foo),
// and the destination is a <details> element, open it.
function openDetailsAnchor() {
	let hash = window.location.hash
	if (!hash) {
		return
	}
	let el = document.getElementById(hash.slice(1)) // remove leading '#'
	if (!el) {
		return
	}

	let details = el.closest("details")
	while (details) {
		details.open = true
		details = details.parentElement.closest("details")
	}

	// New elements may have appeared.
	// Set hash again to scroll to the right place.
	window.location.hash = hash;
	return false;
}

window.addEventListener('hashchange', openDetailsAnchor)

window.addEventListener('load', () => {
	document.querySelectorAll("h2, h3, h4, h5, h6").forEach((el) => {
		if (!el.id) {
			return
		}
		el.innerHTML += ' <a class="permalink" href="#'+el.id+'">&para;</a>'
	})

	document.querySelectorAll("details.example > summary").forEach((el) => {
		let id = el.parentElement.id;
		if (!id) {
			return
		}
		el.innerHTML += ' <a class="permalink" href="#'+id+'">&para;</a>'
	})

	openDetailsAnchor()
})
