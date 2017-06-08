var xhr = new XMLHttpRequest();
xhr.open("GET", "/getmails/100", false);
xhr.send();
console.log(xhr.status);
console.log(xhr.statusText);