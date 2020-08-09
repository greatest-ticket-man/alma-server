'use strict';

class Signup {
    constructor() {

        this.signupEmailInputEl = document.getElementById('js-email');
        this.signupPassInputEl = document.getElementById('js-password');

        this.signupButtonEl = document.querySelector('.js-signup');

        this.signup = this.signup.bind(this);

        this.addEventListener();
    }

    addEventListener() {

        console.log("b el is ", this.signupButtonEl);

        this.signupButtonEl.addEventListener('click', this.signup);
    }

    // signup
    async signup() {

        const email = this.signupEmailInputEl.value;
        const pass = CybozuLabs.MD5.calc(this.signupPassInputEl.value);

        const data = {
            email: email,
            password: pass,
        };

        let response = await window.Alma.req.post('/signup', window.Alma.req.createPostData(data), {reload: false});

        console.log('response is ', response);

        if (response.success) {
            alert('サインアップが成功しました、ログインを試してください');

            window.location.href = '/login';
        }
    }


}

new Signup();