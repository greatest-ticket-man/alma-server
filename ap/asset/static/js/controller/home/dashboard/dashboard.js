'use strict';

class Dashboard {
    constructor() {

        // getEl
        this.eventInfoPanelEl = document.querySelector('.js-event-info');
        this.memberInfoPanelEl = document.querySelector('.js-member-info');
        this.reserveInfoPanelEl = document.querySelector('.js-reserve-info');
        this.ticketInfoPanelEl = document.querySelector('.js-ticket-info');

        this.goEventInfoPage = this.goEventInfoPage.bind(this);
        this.goMemberInfoPage = this.goMemberInfoPage.bind(this);
        this.goReserveInfoPage = this.goReserveInfoPage.bind(this);
        this.goTicketInfoPage = this.goTicketInfoPage.bind(this);

        this.addEventListener();
    }

    addEventListener() {
        this.eventInfoPanelEl.addEventListener('click', this.goEventInfoPage);
        this.memberInfoPanelEl.addEventListener('click', this.goMemberInfoPage);
        this.reserveInfoPanelEl.addEventListener('click', this.goReserveInfoPage);
        this.ticketInfoPanelEl.addEventListener('click', this.goTicketInfoPage);
    }

    goEventInfoPage() {
        window.Alma.location.href(window.Alma.location.event_info);
    }
    goMemberInfoPage() {
        window.Alma.location.href(window.Alma.location.member_info);
    }
    goReserveInfoPage() {
        window.Alma.location.href(window.Alma.location.reserve_info);
    }
    goTicketInfoPage() {
        window.Alma.location.href(window.Alma.location.ticket_info);
    }

}

new Dashboard();
