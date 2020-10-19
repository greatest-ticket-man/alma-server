'use strict';

class ReservePayForm {

    constructor() {

        // EL
        this.reservePayEl = document.querySelector('#js-reserve-pay');

        // bind
        this.getPay = this.getPay.bind(this);
    }

    getPay() {
        return this.reservePayEl.value;
    }

}

const reservePayForm = new ReservePayForm();
