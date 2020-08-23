'use strict';

// use /static/js/component/event/form.js

// TODO
// EventFormが存在するかを確認する

class EventCreate {
    constructor() {

        // getEl
        this.cancelCreateEventButtonEl = document.querySelector('.js-event-create-cancel');
        this.createEventButtonEl = document.querySelector('.js-event-create');

        this.emailTextEl = document.querySelector('.js-email-text');
        this.emailTableEl = document.querySelector('.js-email-table-body');

        // value
        this.eventTitleEl = document.getElementById('js-event-title');
        this.organizationNameEl = document.getElementById('js-event-organization');


        this.backBeforePage = this.backBeforePage.bind(this);
        this.createEvent = this.createEvent.bind(this);

        this.addEventListener();
    }

    addEventListener() {

        this.cancelCreateEventButtonEl.addEventListener('click', this.backBeforePage);
        this.createEventButtonEl.addEventListener('click', this.createEvent);

    }

    // backBeforePage 前のページに戻る
    backBeforePage() {
        window.history.back(-1);
        return false;
    }

    // createEvent
    async createEvent() {

        const eventName = this.eventTitleEl.value;
        const organizationName = this.organizationNameEl.value;


        // TODO form.jsでこれを実装する
        let memberInfoList = [];
        for (let row of this.emailTableEl.rows) {

            let memberInfo = {
                email: '',
                authority: '',
            };

            for (let cell of row.cells) {
                if (cell.classList.contains('js-email-table-email')) {
                    memberInfo.email = cell.innerText;
                }
                else if (cell.classList.contains('js-email-table-auth')) {
                    const select = cell.children[0];
                    memberInfo.authority = select.value;
                }
            }

            memberInfoList.push(memberInfo);
        }

        const data = {
            event_name: eventName,
            organization_name: organizationName,
            member_info_list: memberInfoList,
        };

        console.log("data is ", data);

        let response = await window.Alma.req.post('/event/create', window.Alma.req.createPostData(data), { reload: false });

        console.log("reseponse is ", response);

        if (!response || !response.success) {
            alert('作成に失敗しました');
            return;
        }

        // LocalStorageに追加
        window.Alma.localStorage.set(window.Alma.localStorage.event_id, response.result.event_id);

        // 遷移
        window.Alma.location.href(window.Alma.location.home_dashboard);
    }
}

new EventCreate();
