'use strict';

class TicketCreate {
    constructor() {

        this.ticketInfoBackButtonEl = document.querySelector('.js-ticket-info-back');
        this.ticketCreateCancelButtonEl = document.querySelector('.js-ticket-create-cancel');
        this.ticketCreateButtonEl = document.querySelector('.js-ticket-create');

        this.goTicketInfoPage = this.goTicketInfoPage.bind(this);
        this.createTicket = this.createTicket.bind(this);

        this.addEventListener();
    }

    addEventListener() {
        this.ticketInfoBackButtonEl.addEventListener('click', this.goTicketInfoPage);
        this.ticketCreateCancelButtonEl.addEventListener('click', this.goTicketInfoPage);
        this.ticketCreateButtonEl.addEventListener('click', this.createTicket);
    }

    // createTicket 
    createTicket() {
        // TODO チケットの作成
        console.log("create ticket !! 未実装");
    }


    // goTicketInfoPage
    goTicketInfoPage() {
        window.Alma.location.href(window.Alma.location.ticket_info);
    }
}

new TicketCreate();
