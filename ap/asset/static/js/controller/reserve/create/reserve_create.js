'use strict';

class ReserveCreate {

    constructor() {

        // EL
        this.backButtonEl = document.querySelector('.js-reserve-info-back');
        this.cancelButtonEl = document.querySelector('.js-reserve-create-cancel');

        // Bind
        this.goReserveInfoPage = this.goReserveInfoPage.bind(this);

        this.addEventListener();

    }

    addEventListener() {

        this.backButtonEl.addEventListener('click', this.goReserveInfoPage);
        this.cancelButtonEl.addEventListener('click', this.goReserveInfoPage);




    }


    // goReserveInfoPage 
    goReserveInfoPage() {
        window.Alma.location.href(window.Alma.location.reserve_info);
    }


}

new ReserveCreate();