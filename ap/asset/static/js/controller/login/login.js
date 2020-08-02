'use strict';

class Login {

    constructor() {

        // getEl
        this.loginButtonEl = document.querySelector('.js-login');
        this.loginNameInputEl = document.getElementById('js-username');
        this.loginPassInputEl = document.getElementById('js-password');

        // 関数登録
        this.login = this.login.bind(this);

        this.addEventListener();
    }

    // addEventListener イベントを追加する
    addEventListener() {
        this.loginButtonEl.addEventListener('click', this.login);
    }

    // loginする
    async login() {

        let name = this.loginNameInputEl.value;
        let pass = this.loginPassInputEl.value;

        // passwordをmd5ハッシュに変更
        let passMd5 = CybozuLabs.MD5.calc(pass);

        const data = {
            name: name,
            pass: passMd5,
        };

        let response = await window.Alma.req.post('/login', window.Alma.req.createPostData(data), {reload: false});

        if (response.result) {

            // localStorageにデータを追加する

            // 遷移
            window.location.href = '/test';
        }
    }
}

new Login();

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

