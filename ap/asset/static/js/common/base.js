'use strict';

class Base {
    constructor() {
        this.eventNameEl = document.querySelector('.js-event-name');

        // this.setEventName();
        this.setEvent();
    }

    // setEvent eventidのURLが変わっている場合は、そのイベントに切り替える
    async setEvent() {

        const eventIdParam = this.getParam('event'); // 優先度1
        const eventIdLocal = window.Alma.localStorage.get(window.Alma.localStorage.event_id); // 優先度2
        let eventName = window.Alma.localStorage.get(window.Alma.localStorage.event_name);

        if (eventIdParam !== null) {
            if (eventIdParam !== eventIdLocal) {

                const data = {
                    event: eventIdParam,
                };

                let response = await window.Alma.req.post('/event/get', window.Alma.req.createPostData(data), { reload: false});
                console.log("event set response is ", response);

                window.Alma.localStorage.set(window.Alma.localStorage.event_id, response.result.event_id);
                window.Alma.localStorage.set(window.Alma.localStorage.event_name, response.result.event_name);
                eventName = response.result.event_name;
            }
        }

        // paramがなく、Localにデータがある場合は取得する
        if (eventIdParam === null && eventIdLocal !== null) {
            window.Alma.location.href(window.location.href);
        }

        this.eventNameEl.innerHTML = eventName;
    }

    /**
     * Get the URL parameter value
     *
     * @param  name {string} パラメータのキー文字列
     * @return  url {url} 対象のURL文字列（任意）
     */
    getParam(name, url) {
        if (!url) url = window.location.href;
        name = name.replace(/[\[\]]/g, "\\$&");
        var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
            results = regex.exec(url);
        if (!results) return null;
        if (!results[2]) return '';
        return decodeURIComponent(results[2].replace(/\+/g, " "));
    }

}

new Base();
