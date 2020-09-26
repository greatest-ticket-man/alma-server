'use strict';

class TicketForm {

    constructor() {

        this.ticketIdEl = document.getElementById('js-ticket-id');
        this.ticketNameEl = document.getElementById('js-ticket-name');
        this.ticketPriceEl = document.getElementById('js-ticket-price');
        this.ticketDescEl = document.getElementById('js-ticket-desc');
        this.ticketStockEl = document.getElementById('js-ticket-stock');
        this.ticketEventStartTimeEl = document.getElementById('js-ticket-event-start-time');

    }

    getTicketId() {
        return this.ticketIdEl.value;
    }

    getTicketName() {
        return this.ticketNameEl.value;
    }

    getTicketPrice() {
        return Number(this.ticketPriceEl.value);
    }

    getTicketStock() {
        return Number(this.ticketStockEl.value);
    }

    getTicketDesc() {
        return this.ticketDescEl.value;
    }

    /**
     * 2020-09-18T17:05
     */
    getTicketEventStartTime() {

        console.log("t ", new Date(this.ticketEventStartTimeEl.value).toISOString());
        return new Date(this.ticketEventStartTimeEl.value);
    }

}

const ticketForm = new TicketForm();
