'use strict';

// use /static/js/component/event/form.js

// EventFormが存在するかを確認する
if (typeof eventForm === 'undefined') {
    alert('依存しているeventFormが見つかりませんでした');
    console.error('eventFormが見つかりません');
    console.error('/static/js/component/event/form.jsをimportしてください');
}

class EventCreate {
    constructor() {

        // getEl
        this.cancelCreateEventButtonEl = document.querySelector('.js-event-create-cancel');
        this.createEventButtonEl = document.querySelector('.js-event-create');

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

        const eventName = eventForm.getEventName();
        const organizationName = eventForm.getOrganizationName();
        const memberInfoList = eventForm.getMemberInfoList();

        const data = {
            event_name: eventName,
            organization_name: organizationName,
            invite_member_list: memberInfoList,
        };

        let response = await window.Alma.req.post(window.Alma.req.event_create, window.Alma.req.createPostData(data), { reload: false });

        if (!response || !response.success) {
            window.Alma.toast.error('イベントの作成に失敗しました');
            return;
        }

        // LocalStorageに追加
        window.Alma.localStorage.set(window.Alma.localStorage.event_id, response.result.event_id);

        // 遷移
        window.Alma.location.href(window.Alma.location.home_dashboard);
    }
}

new EventCreate();
