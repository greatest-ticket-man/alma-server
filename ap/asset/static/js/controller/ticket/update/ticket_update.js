'use strict';

// TicketFormが存在するかを確認する
if (typeof ticketForm === 'undefined') {
    alert('依存しているticketFormが見つかりませんでした');
    console.error('ticketFormが見つかりません');
    console.error('/static/js/component/ticket/form.jsをimportしてください');
}

class TicketUpdate {
    constructor() {

        this.ticketInfoBackButtonEl = document.querySelector('.js-ticket-info-back');
        this.ticketUpdateCancelButtonEl = document.querySelector('.js-ticket-update-cancel');
        this.ticketUpdateButtonEl = document.querySelector('.js-ticket-update');

        this.updateTicket = this.updateTicket.bind(this);
        this.goTicketInfoPage = this.goTicketInfoPage.bind(this);

        // TicketIdを編集できないようにする
        ticketForm.setReadOnlyTicketId();

        this.addEventListener();
    }

    addEventListener() {

        this.ticketInfoBackButtonEl.addEventListener('click', this.goTicketInfoPage);
        this.ticketUpdateCancelButtonEl.addEventListener('click', this.goTicketInfoPage);
        this.ticketUpdateButtonEl.addEventListener('click', this.updateTicket);

    }


    async updateTicket() {

        // request
        const data = {
            ticket_info: {
                ticket_id: ticketForm.getTicketId(),
                name: ticketForm.getTicketName(),
                price: ticketForm.getTicketPrice(),
                desc: ticketForm.getTicketDesc(),
                event_id: window.Alma.location.getParam('event'),
                stock: ticketForm.getTicketStock(),
                start_time: ticketForm.getTicketStartTime(),
                end_time: ticketForm.getTicketEndTime(),
            },
            event_id: window.Alma.location.getParam('event'),
        };

        const response = await window.Alma.req.post(window.Alma.req.ticket_update, window.Alma.req.createPostData(data));
        if (!response || !response.success) {
            window.Alma.toast.error('チケットの編集に失敗しました');
            return;
        }
        window.Alma.toast.success('チケットの編集に成功しました', 'Greatest Ticket Man', 1500, function () {
            // 遷移
            window.Alma.location.href(window.Alma.location.ticket_info);
        });

    }

    // goTicketInfoPage
    goTicketInfoPage() {
        window.Alma.location.href(window.Alma.location.ticket_info);
    }
}

new TicketUpdate();
