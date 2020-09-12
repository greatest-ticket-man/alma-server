'use strict';

class TicketCreate {
    constructor() {

        this.ticketInfoBackButtonEl = document.querySelector('.js-ticket-info-back');

        this.goTicketInfoPage = this.goTicketInfoPage.bind(this);

        this.addEventListener();
    }

    addEventListener() {
        this.ticketInfoBackButtonEl.addEventListener('click', this.goTicketInfoPage);
    }

    // goTicketInfoPage
    goTicketInfoPage() {
        window.Alma.location.href(window.Alma.location.ticket_info);
    }
}

new TicketCreate();
