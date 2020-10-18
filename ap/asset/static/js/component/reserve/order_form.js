'use strict';

class ReserveOrderForm {

    constructor() {

        // EL
        this.reserveTicketKindEl = document.querySelector('#js-reserve-ticket-kind');
        this.reserveDateSelectEl; // 
        this.reserveDateEl = document.querySelector('.js-reserve-date');
        this.reserveNumEl = document.querySelector('.js-reserve-num');
        this.reserveNumInputEl;
        this.reserveNumDescEl;

        // bind
        this.addReserveDateSelect = this.addReserveDateSelect.bind(this);
        this.addReserveTicketNum = this.addReserveTicketNum.bind(this);

        this.addEventListener();
    }


    addEventListener() {
        this.reserveTicketKindEl.addEventListener('change', this.addReserveDateSelect);
    }

    // addReserveDateSelect チケットの種類別の予約できる日付を取得する
    addReserveDateSelect(e) {

        // 先にKindがある場合は削除する
        if (this.reserveDateSelectEl) {
            // 削除
            this.reserveDateSelectEl.remove();
            this.reserveDateSelectEl == null;
        }

        const ticketId = e.target.value;
        if (ticketId === '') {
            // 何も選択されなかった場合は、何も表示しない
            return;
        }

        const t = document.querySelector(`#js-reserve-date-select-${ticketId}-tmp`);
        const clone = t.content.cloneNode(true);

        this.reserveDateEl.appendChild(clone);
        this.reserveDateSelectEl = document.querySelector('#js-reserve-date-select');

        // add.
        this.reserveDateSelectEl.addEventListener('change', this.addReserveTicketNum);
    }

    // addReserveTicketNum チケット枚数選択フィールドを追加する
    addReserveTicketNum(e) {

        if (this.reserveNumInputEl) {
            this.reserveNumInputEl.remove();
            this.reserveNumInputEl == null;
        }
        if (this.reserveNumDescEl) {
            this.reserveNumDescEl.remove();
            this.reserveNumDescEl == null;
        }
        
        const scheduleStockId = e.target.value;
        if (scheduleStockId === '') {
            return;
        }

        const t = document.querySelector(`#js-reserve-stock-${scheduleStockId}-tmp`);
        const clone = t.content.cloneNode(true);

        this.reserveNumEl.appendChild(clone);

        this.reserveNumInputEl = document.querySelector('#js-reserve-num-input');
        this.reserveNumDescEl = document.querySelector('.js-reserve-num-desc');
    }


}

new ReserveOrderForm();