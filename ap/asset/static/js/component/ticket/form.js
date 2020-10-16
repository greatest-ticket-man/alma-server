'use strict';

class TicketForm {

    constructor() {

        this.ticketIdEl = document.getElementById('js-ticket-id');
        this.ticketNameEl = document.getElementById('js-ticket-name');
        this.ticketPriceEl = document.getElementById('js-ticket-price');
        this.ticketDescEl = document.getElementById('js-ticket-desc');

        this.scheduleStockTableEl = document.querySelector('.js-schedule-stock-table-body');
        this.scheduleStockTableAddButtonEl = document.querySelector('.js-schedule-stock-add-button');
        this.scheduleStockRowTmpEl = document.querySelector('#js-schedule-stock-row-tmp');

        // bind
        this.addMultiStockTable = this.addMultiStockTable.bind(this);
        this.deleteMultiStockRow = this.deleteMultiStockRow.bind(this);
        this.getScheduleStockInfoList = this.getScheduleStockInfoList.bind(this);
        this.setReadOnlyTicketId = this.setReadOnlyTicketId.bind(this);

        this.addEventListener();
    }

    addEventListener() {
        this.scheduleStockTableAddButtonEl.addEventListener('click', this.addMultiStockTable);
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

    getScheduleStockInfoList() {
        let scheduleStockInfoList = [];

        for (let row of this.scheduleStockTableEl.rows) {

            const scheduleStockId = row.querySelector('.js-schedule-stock-id').value;
            const stock = Number(row.querySelector('.js-stock').value);
            const eventStartTime = new Date(row.querySelector('.js-schedule').value);
            let scueduleStockInfo = {
                schedule_stock_id: scheduleStockId,
                event_start_time: window.Alma.dateutil.DateToTimestamp(eventStartTime),
                stock: stock,
            };

            scheduleStockInfoList.push(scueduleStockInfo);
        }

        return scheduleStockInfoList
    }

}

const ticketForm = new TicketForm();
