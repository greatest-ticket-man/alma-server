'use strict';

// use /static/js/component/ticket/form.js

// TicketFormが存在するかを確認する
if (typeof ticketForm === 'undefined') {
    alert('依存しているticketFormが見つかりませんでした');
    console.error('ticketFormがみつかりません');
    console.error('/static/js/component/ticket/form.jsをimportしてください');
}

class TicketCreate {
    constructor() {

        this.ticketInfoBackButtonEl = document.querySelector('.js-ticket-info-back');
        this.ticketCreateCancelButtonEl = document.querySelector('.js-ticket-create-cancel');
        this.ticketCreateButtonEl = document.querySelector('.js-ticket-create');

        this.goTicketInfoPage = this.goTicketInfoPage.bind(this);
        this.createTicket = this.createTicket.bind(this);

        this.addEventListener();
    }

    addEventListener() {
        this.ticketInfoBackButtonEl.addEventListener('click', this.goTicketInfoPage);
        this.ticketCreateCancelButtonEl.addEventListener('click', this.goTicketInfoPage);
        this.ticketCreateButtonEl.addEventListener('click', this.createTicket);
    }

    // createTicket 
    async createTicket() {

        const data = {
            ticket_info: {
                ticket_id: ticketForm.getTicketId(),
                ticket_name: ticketForm.getTicketName(),
                ticket_price: ticketForm.getTicketPrice(),
                ticket_desc: ticketForm.getTicketDesc(),
                event_id: window.Alma.location.getParam('event'),
                ticket_stock: ticketForm.getTicketStock(),
                ticket_event_start_time: ticketForm.getTicketEventStartTime(),
            },
        };

        const response = await window.Alma.req.post(window.Alma.req.ticket_create, window.Alma.req.createPostData(data));
        if (!response || !response.success) {
            window.Alma.toast.error('チケットの作成に失敗しました');
            return;
        }

        window.Alma.toast.success('チケットの作成に成功しました', 'Greatest Ticket Man', 1500, function() {
            // 遷移
            window.Alma.location.href(window.Alma.location.ticket_info);
        });
    }


    // goTicketInfoPage
    goTicketInfoPage() {
        window.Alma.location.href(window.Alma.location.ticket_info);
    }
}

new TicketCreate();
