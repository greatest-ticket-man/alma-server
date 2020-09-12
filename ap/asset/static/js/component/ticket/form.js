'use strict';

class TicketForm {

    constructor() {

        this.ticketIdEl = document.getElementById('js-ticket-id');
        this.ticketNameEl = document.getElementById('js-ticket-name');
        this.ticketPriceEl = document.getElementById('js-ticket-price');
        this.ticketDescEl = document.getElementById('js-ticket-desc');

    }

    getTicketId() {
        return this.ticketIdEl.value;
    }

    getTicketName() {
        return this.ticketNameEl.value;
    }

    getTicketPrice() {
        return this.ticketPriceEl.value;
    }

    getTicketDesc() {
        return this.ticketDescEl.value;
    }

}

const ticketForm = new TicketForm();
