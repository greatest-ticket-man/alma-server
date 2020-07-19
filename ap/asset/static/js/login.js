const inputs = document.querySelectorAll('.input');

// focusFunc 
function focusFunc() {
    let parent = this.parentNode.parentNode;
    parent.classList.add('focus');
}

// blurFunc
function blurFunc() {
    let parent = this.parentNode.parentNode;
    if (this.value === '') {
        parent.classList.remove('focus');
    }
}

inputs.forEach(input => {
    input.addEventListener('focus', focusFunc);
    input.addEventListener('blur',blurFunc );
});
