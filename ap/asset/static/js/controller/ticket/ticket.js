'use strict';

// TableInfoが存在するかを確認する
if (typeof TableInfo === 'undefined') {
    alert('依存しているTableInfoが見つかりませんでした');
    console.error('TableInfoがみつかりません');
    console.error('/static/js/common/table/table.jsをimportしてください');
}

class TicketInfo {
    constructor() {

        this.tableInfo = new TableInfo('js-ticket-table');

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

        // checkされているチケットを取得
        let ticketIDList = [];
        this.tableInfo.getCheckRowList().forEach(function(row) {
            ticketIDList.push(row.ticketId);
        });

        if (ticketIDList.length === 0) {
            window.Alma.toast.warn('削除対象のチケットが選択されていません', 'Greatest Ticket Man');
            return;
        }

        const data = {
            event_id: window.Alma.location.getParam('event'),
            ticket_id_list: ticketIDList,
        };

        const response = await window.Alma.req.post(window.Alma.req.ticket_delete, window.Alma.req.createPostData(data));
        if (!response || !response.success) {
            window.Alma.toast.error('チケットの削除に失敗しました');
            return;
        }

        window.Alma.toast.success('チケットの削除に成功しました', 'Greatest Ticket Man', 1500, function () {
            // reload
            window.Alma.location.href(window.Alma.location.ticket_info);

        });
    }

    // reloadTicketPage
    reloadTicketPage() {
        window.Alma.location.href(window.Alma.location.ticket_info);
    }

}

new TicketInfo();
