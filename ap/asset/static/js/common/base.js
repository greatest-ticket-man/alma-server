'use strict';

class Base {
    constructor() {
        this.eventNameEl = document.querySelector('.js-event-name');
        this.setEventName();
    }

    // setEventName イベントの名前を表示させる
    setEventName() {   

        const eventName = window.Alma.localStorage.get(window.Alma.localStorage.event_name);
        
        // なければそのまま返す
        if (eventName == null) {
            return;
        }

        this.eventNameEl.innerHTML = eventName;
    }
}

new Base();
