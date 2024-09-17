try {
	const blink1 = document.getElementById('blink1');
	const blink2 = document.getElementById('blink2');
	const blink3 = document.getElementById('blink3');
	const blink4 = document.getElementById('blink4');
	const blink5 = document.getElementById('blink5');

	setInterval(() => {
	  blink1.style.visibility = blink1.style.visibility === 'hidden' ? 'visible' : 'hidden';
	  blink2.style.visibility = blink2.style.visibility === 'hidden' ? 'visible' : 'hidden';
	  blink3.style.visibility = blink3.style.visibility === 'hidden' ? 'visible' : 'hidden';
	  blink4.style.visibility = blink4.style.visibility === 'hidden' ? 'visible' : 'hidden';
	  blink5.style.visibility = blink5.style.visibility === 'hidden' ? 'visible' : 'hidden';
	}, 2000);
} catch {}

