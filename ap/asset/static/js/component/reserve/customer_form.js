'use strict';

class ReserveCustomerForm {
    constructor() {

        // EL
        this.reserveNameEl = document.querySelector('#js-reserve-name');
        this.reserveFuriganaEl = document.querySelector('#js-reserve-furigana');
        this.reserveEmailEl = document.querySelector('#js-reserve-email');

        // bind
        this.getName = this.getName.bind(this);
        this.getFurigana = this.getFurigana.bind(this);
        this.getEmail = this.getEmail.bind(this);
    }

    getName() {
        return this.reserveNameEl.value;
    }

    getFurigana() {
        return this.reserveFuriganaEl.value;
    }

    getEmail() {
        return this.reserveEmailEl.value;
    }
}

const reserveCustomerForm = new ReserveCustomerForm();
