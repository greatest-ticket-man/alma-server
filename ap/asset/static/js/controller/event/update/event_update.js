'use strict';

console.log("aaa")

// use /static/js/component/event/form.js

// EventFormが存在するかを確認する
if (typeof eventForm === 'undefined') {
    alert('依存しているeventFormが見つかりません');
    console.error('eventFormが見つかりません');
    console.error('/static/js/component/event/form.jsをimportしてください');
}

class EventUpdate{

    constructor() {

        // getEl
        this.cancelUpdateEventButtonEl = document.querySelector('.js-event-update-cancel');
        this.updateEventButtonEl = document.querySelector('.js-event-update');

        this.backBeforePage = this.backBeforePage.bind(this);
        this.updateEvent = this.updateEvent.bind(this);

        this.addEventListener();
    }

    addEventListener() {

        this.cancelUpdateEventButtonEl.addEventListener('click', this.backBeforePage);
        this.updateEventButtonEl.addEventListener('click', this.updateEvent);

    }

    // backBeforePage 前のページに戻る
    backBeforePage() {
        window.history.back(-1);
        return false;
    }

    // updateEvent
    async updateEvent() {

        const eventId = window.Alma.location.getParam('event');
        const eventName = eventForm.getEventName();
        const organizationName = eventForm.getOrganizationName();
        const memberInfoList = eventForm.getMemberInfoList();

        const data = {
            event_id: eventId, 
            event_name: eventName,
            organization_name: organizationName,
            invite_member_list: memberInfoList,
        };

        let response = await window.Alma.req.post(window.Alma.req.event_update, window.Alma.req.createPostData(data), { reload: false });

        if (!response || !response.success) {
            alert('更新に失敗しました');
            return;
        }

        // TODO message toast
        alert('イベントの編集に成功しました');

        // 遷移
        window.Alma.location.href(window.Alma.location.event_info);
    }

}

new EventUpdate();
