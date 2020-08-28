'use strict';

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
            window.Alma.toast.error('イベントの更新に失敗しました');
            return;
        }


        window.Alma.toast.success('イベントの更新に成功しました', 'イベント', 2000, function() {
            // 遷移
            window.Alma.location.href(window.Alma.location.event_info);
        });
    }

}

new EventUpdate();
