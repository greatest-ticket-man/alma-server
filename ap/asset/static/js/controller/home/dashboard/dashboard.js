'use strict';

class Dashboard {
    constructor() {

        // getEl
        this.eventInfoPanelEl = document.querySelector('.js-event-info');
        this.memberInfoPanelEl = document.querySelector('.js-member-info');

        this.goEventInfoPage = this.goEventInfoPage.bind(this);
        this.goMemberInfoPage = this.goMemberInfoPage.bind(this);

        this.addEventListener();
    }

    addEventListener() {
        this.eventInfoPanelEl.addEventListener('click', this.goEventInfoPage);
        this.memberInfoPanelEl.addEventListener('click', this.goMemberInfoPage);
    }

    goEventInfoPage() {
        window.Alma.location.href(window.Alma.location.event_info);
    }
    goMemberInfoPage() {
        window.Alma.location.href(window.Alma.location.member_info);
    }
}

new Dashboard();
