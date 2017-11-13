var AppCanvas = document.querySelector('#AppCanvas'),
    AppInputArea = document.querySelector('#AppInputArea');

//event listeners
window.addEventListener('load', function() {
    Particles.init({
        selector: '#AppCanvas'
    });
});