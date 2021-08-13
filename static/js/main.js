// Preloader
window.onload = setTimeout(Preloader, 1500);

function Preloader() {
    document.querySelector('#loader_bg').setAttribute('class', 'hidden');
    document.querySelector('#page-wrapper').setAttribute('class', 'visible');
    document.querySelector('#particles-js').setAttribute('class', 'visible');
};
// Menu
var toggle = document.getElementById('toggle');
toggle.onclick = function(){
    toggle.classList.toggle('active'); 
    document.getElementById('nav').classList.toggle('active');
};
