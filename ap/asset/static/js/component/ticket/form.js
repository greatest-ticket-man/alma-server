'use strict';

class TicketForm {

    constructor() {

        this.ticketIdEl = document.getElementById('js-ticket-id');
        this.ticketNameEl = document.getElementById('js-ticket-name');
        this.ticketPriceEl = document.getElementById('js-ticket-price');
        this.ticketDescEl = document.getElementById('js-ticket-desc');
        this.ticketStockEl = document.getElementById('js-ticket-stock');
        this.ticketStartTimeEl = document.getElementById('js-ticket-start-time');
        this.ticketEndTimeEl = document.getElementById('js-ticket-end-time');

        // bind
        this.setReadOnlyTicketId = this.setReadOnlyTicketId.bind(this);

        this.addEventListener();
    }

    addEventListener() {
        // this.scheduleStockTableAddButtonEl.addEventListener('click', this.addMultiStockTable);
    }

    // setReadOnlyTicketId ticketIDを編集できないようにする
    setReadOnlyTicketId() {
        this.ticketIdEl.readOnly = true;
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

    getTicketDesc() {
        return this.ticketDescEl.value;
    }

    getTicketStock() {
        return Number(this.ticketStockEl.value);
    }

    getTicketStartTime() {
        const startTime = new Date(this.ticketStartTimeEl.value);
        return window.Alma.dateutil.DateToTimestamp(startTime);
    }

    getTicketEndTime() {
        const endTime = new Date(this.ticketEndTimeEl.value);
        return window.Alma.dateutil.DateToTimestamp(endTime);
    }

}

const ticketForm = new TicketForm();
