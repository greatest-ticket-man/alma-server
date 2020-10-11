'use strict';

class TicketForm {

    constructor() {

        this.ticketIdEl = document.getElementById('js-ticket-id');
        this.ticketNameEl = document.getElementById('js-ticket-name');
        this.ticketPriceEl = document.getElementById('js-ticket-price');
        this.ticketDescEl = document.getElementById('js-ticket-desc');

        // this.ticketStockEl = document.getElementById('js-ticket-stock');
        // this.ticketEventStartTimeEl = document.getElementById('js-ticket-event-start-time');

        this.scheduleStockTableEl = document.querySelector('.js-schedule-stock-table-body');
        this.scheduleStockTableAddButtonEl = document.querySelector('.js-schedule-stock-add-button');
        this.scheduleStockRowTmpEl = document.querySelector('#js-schedule-stock-row-tmp');

        // bind
        this.addMultiStockTable = this.addMultiStockTable.bind(this);

        this.addEventListener();

    }

    addEventListener() {

        this.scheduleStockTableAddButtonEl.addEventListener('click', this.addMultiStockTable);

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

    // addMultiStockTable テーブルのrowを追加する
    addMultiStockTable() {
        const t = this.scheduleStockRowTmpEl;
        const clone = t.content.cloneNode(true);
        this.scheduleStockTableEl.appendChild(clone);
    }

    deleteMultiStockRow(elem) {
        let tr = elem.parentNode.parentNode;
        this.scheduleStockTableEl.deleteRow(tr.sectionRowIndex);
    }

    // getTicketStock() {
    //     return Number(this.ticketStockEl.value);
    // }

    // getTicketDesc() {
    //     return this.ticketDescEl.value;
    // }

    // /**
    //  * 2020-09-18T17:05
    //  */
    // getTicketEventStartTime() {
    //     return new Date(this.ticketEventStartTimeEl.value);
    // }

}

const ticketForm = new TicketForm();
