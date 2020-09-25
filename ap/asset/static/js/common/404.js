'use strict';

let container = document.querySelector('.js-container');

window.onmousemove = function(e) {
    const x = -e.clientX / 5;
    const y = -e.clientY / 5;
    container.style.backgroundPositionX = `${x}px`;
    container.style.backgroundPositionY = `${y}px`;
};

class PageNotFound {
    constructor() {
        
        this.goBackDashboardButtonEl = document.querySelector('.js-back-dashboard');
        this.goDashboardPage = this.goDashboardPage.bind(this);

        this.addEventListener();
    }

    addEventListener() {
        
        this.goBackDashboardButtonEl.addEventListener('click', this.goDashboardPage);

    }

    goDashboardPage() {
        window.Alma.location.href(window.Alma.location.home_dashboard);
    }
}

new PageNotFound();
