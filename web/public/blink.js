try {
const blink1 = document.getElementById('blink1');
const blink2 = document.getElementById('blink2');
const blink3 = document.getElementById('blink3');

setInterval(() => {
  blink1.style.visibility = blink1.style.visibility === 'hidden' ? 'visible' : 'hidden';
  blink2.style.visibility = blink2.style.visibility === 'hidden' ? 'visible' : 'hidden';
  blink3.style.visibility = blink3.style.visibility === 'hidden' ? 'visible' : 'hidden';
}, 2000);
} catch {}
