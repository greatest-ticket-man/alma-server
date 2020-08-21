'use strict';

class Base {
    constructor() {
        this.baseHeaderTitleEl = document.querySelector('.js-base-header-title');

        this.addEventListener();
    }

    addEventListener() {
        this.baseHeaderTitleEl.addEventListener('click', this.goDashboardPage);
    }

    goDashboardPage() {

        console.log("hoge");
        // 遷移
        window.Alma.location.href(window.Alma.location.home_dashboard);
    }
}

new Base();
