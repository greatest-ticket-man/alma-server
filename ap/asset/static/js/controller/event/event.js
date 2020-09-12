'use strict';

class EventInfo {
    constructor() {
        // getEl
        this.updateEventButtonEl = document.querySelector('.js-event-update');

        this.goUpdateEventPage = this.goUpdateEventPage.bind(this);

        this.addEventListener();
    }

    addEventListener() {
        this.updateEventButtonEl.addEventListener('click', this.goUpdateEventPage);
    }

    // goUpdateEventPage 編集ページに遷移
    goUpdateEventPage() {
        window.Alma.location.href(window.Alma.location.event_update);
    }
}

new EventInfo();
