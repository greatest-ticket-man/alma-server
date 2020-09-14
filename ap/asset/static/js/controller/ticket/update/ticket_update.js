'use strict';

// TicketFormが存在するかを確認する
if (typeof ticketForm === 'undefined') {
    alert('依存しているticketFormが見つかりませんでした');
    console.error('ticketFormが見つかりません');
    console.error('/static/js/component/ticket/form.jsをimportしてください');
}

class TicketUpdate {
    constructor() {

        this.ticketInfoBackButtonEl = document.querySelector('.js-ticket-info-back');
        this.ticketUpdateCancelButtonEl = document.querySelector('.js-ticket-update-cancel');
        this.ticketUpdateButtonEl = document.querySelector('.js-ticket-update');

        this.updateTicket = this.updateTicket.bind(this);
        this.goTicketInfoPage = this.goTicketInfoPage.bind(this);

        this.addEventListener();
    }

    addEventListener() {

        this.ticketInfoBackButtonEl.addEventListener('click', this.goTicketInfoPage);
        this.ticketUpdateCancelButtonEl.addEventListener('click', this.goTicketInfoPage);
        this.ticketUpdateButtonEl.addEventListener('click', this.updateTicket);

    }

    // updateTicket
    async updateTicket() {
        alert('未実装');
    }

    // goTicketInfoPage
    goTicketInfoPage() {
        window.Alma.location.href(window.Alma.location.ticket_info);
    }
}

new TicketUpdate();
