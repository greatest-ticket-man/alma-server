'use strict';

// TableInfoが存在するかを確認する
if (typeof TableInfo === 'undefined') {
    alert('依存しているTableInfoが見つかりませんでした');
    console.error('TableInfoが見つかりません');
    console.error('/static/js/common/table/table.jsをimportしてください');
}

class ReserveInfo {
    constructor() {
        this.tableInfo = new TableInfo('js-reserve-table');

        // EL
        this.createButtonEl = document.querySelector('.js-reserve-create');
        this.deleteButtonEl = document.querySelector('.js-reserve-delete');
        this.reloadButtonEl = document.querySelector('.js-reserve-reload');

        // bind
        this.goCreateReservePage = this.goCreateReservePage.bind(this);
        this.deleteReserev = this.deleteReserev.bind(this);
        this.reloadReservePage = this.reloadReservePage.bind(this);


        this.addEventListener();
    }

    addEventListener() {
        this.createButtonEl.addEventListener('click', this.goCreateReservePage);
        this.deleteButtonEl.addEventListener('click', this.deleteReserev);
        this.reloadButtonEl.addEventListener('click', this.reloadReservePage);

    }

    // goCreateReservePage 
    goCreateReservePage() {
        window.Alma.location.href(window.Alma.location.reserve_create);
    }

    // deleteReserve
    async deleteReserev() {
        alert("未実装");
    }

    // reloadReservePage
    reloadReservePage() {
        window.Alma.location.href(window.Alma.location.reserve_info);
    }
}

new ReserveInfo();