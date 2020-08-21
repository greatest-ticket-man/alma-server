'use strict';

class Dashboard {
    constructor() {

        // getEl
        this.eventInfoPanelEl = document.querySelector('.js-event-info');

        this.goEventInfoPage = this.goEventInfoPage.bind(this);

        this.addEventListener();
    }

    addEventListener() {

        this.eventInfoPanelEl.addEventListener('click', this.goEventInfoPage);

    }

    goEventInfoPage() {
        window.Alma.location.href(window.Alma.location.event_info);
    }
}

new Dashboard();
