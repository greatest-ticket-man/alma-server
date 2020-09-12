'use strict';

class TicketInfo {
    constructor() {
        
        this.createTicketButtonEl = document.querySelector('.js-ticket-create');

        this.goCreateTicketPage = this.goCreateTicketPage.bind(this);
        
        this.addEventListener();
    }

    addEventListener() {
        this.createTicketButtonEl.addEventListener('click', this.goCreateTicketPage);
    }

    // goCreateTicketPage
    goCreateTicketPage() {
        window.Alma.location.href(window.Alma.location.ticket_create);
    }
}

new TicketInfo();
