'use strict';

class Dashboard {
    constructor() {

        // getEl
        this.eventInfoPanelEl = document.querySelector('.js-event-info');
        this.memberInfoPanelEl = document.querySelector('.js-member-info');
        this.reserveInfopanelEl = document.querySelector('.js-reserve-info');

        this.goEventInfoPage = this.goEventInfoPage.bind(this);
        this.goMemberInfoPage = this.goMemberInfoPage.bind(this);

        this.addEventListener();
    }

    addEventListener() {
        this.eventInfoPanelEl.addEventListener('click', this.goEventInfoPage);
        this.memberInfoPanelEl.addEventListener('click', this.goMemberInfoPage);
        this.reserveInfopanelEl.addEventListener('click', this.goReserveInfopage);
    }

    goEventInfoPage() {
        window.Alma.location.href(window.Alma.location.event_info);
    }
    goMemberInfoPage() {
        window.Alma.location.href(window.Alma.location.member_info);
    }

    goReserveInfopage() {
        window.Alma.location.href(window.Alma.location.reserve_info);
    }
}

new Dashboard();
