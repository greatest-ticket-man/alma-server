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
        this.deleteTicketButtonEl = document.querySelector('.js-ticket-delete');
        this.reloadTicketButtonEl = document.querySelector('.js-ticket-reload');

        this.goCreateTicketPage = this.goCreateTicketPage.bind(this);
        this.deleteTicket = this.deleteTicket.bind(this);
        this.reloadTicketPage = this.reloadTicketPage.bind(this);
        
        this.addEventListener();
    }

    addEventListener() {
        this.createTicketButtonEl.addEventListener('click', this.goCreateTicketPage);
        this.deleteTicketButtonEl.addEventListener('click', this.deleteTicket);
        this.reloadTicketButtonEl.addEventListener('click', this.reloadTicketPage);
    }

    // goCreateTicketPage
    goCreateTicketPage() {
        window.Alma.location.href(window.Alma.location.ticket_create);
    }

    // deleteTicket
    async deleteTicket() {
        // alert('未実装');

        this.tableInfo.getCheckRowList();

        // TODO チェックしたTicketIDを取得


        // TODO Modalを表示して、本当に削除していいかを確認する

        // TODO 通信して、削除する
    }

    // reloadTicketPage
    reloadTicketPage() {
        window.Alma.location.href(window.Alma.location.ticket_info);
    }

    // TODO 削除, 確認Modalを作成する必要がある、めんどくさい

}

new TicketInfo();
