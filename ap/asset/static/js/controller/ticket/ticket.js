'use strict';

// TableInfoが存在するかを確認する
if (typeof TableInfo === 'undefined') {
    alert('依存しているTableInfoが見つかりませんでした');
    console.error('TableInfoがみつかりません');
    console.error('/static/js/common/table/table.jsをimportしてください');
}

class TicketInfo {
    constructor() {

        this.tableInfo = new TableInfo('js-ticket-table', 'js-ticket-table-row', 'js-ticket-table-head-checkbox', 'js-table-checkbox');

        this.createTicketButtonEl = document.querySelector('.js-ticket-create');

        this.goCreateTicketPage = this.goCreateTicketPage.bind(this);
        
        this.addEventListener();
    }

    addEventListener() {
        this.createTicketButtonEl.addEventListener('click', this.goCreateTicketPage);
    }

    // goCreateTicketPage
    goCreateTicketPage() {
        window.Alma.location.href(window.Alma.location.ticket_create);
    }

}

new TicketInfo();
