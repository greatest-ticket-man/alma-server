'use strict';

class Base {
    constructor() {
        this.eventNameEl = document.querySelector('.js-event-name');

        // this.setEventName();
        this.setEvent();
    }

    // setEvent eventidのURLが変わっている場合は、そのイベントに切り替える
    // TODO 名前は適宜変わる可能性があるため、毎度取得する
    // TODO https://www.it-swarm.dev/ja/http/url%E3%82%AF%E3%82%A8%E3%83%AA%E3%83%91%E3%83%A9%E3%83%A1%E3%83%BC%E3%82%BF%E3%82%92http-get%E3%83%AA%E3%82%AF%E3%82%A8%E3%82%B9%E3%83%88%E3%81%AB%E8%BF%BD%E5%8A%A0%E3%81%99%E3%82%8B%E6%96%B9%E6%B3%95%E3%81%AF%EF%BC%9F/834905187/
    // ↑サーバーがわでGetRequest変更できないかな
    async setEvent() {

        const eventIdParam = this.getParam('event'); // 優先度1
        const eventIdLocal = window.Alma.localStorage.get(window.Alma.localStorage.event_id); // 優先度2
    

        if (!eventIdParam && !eventIdLocal) {
            // 何もしない
            return;
        }

        let eventId = eventIdLocal;
        if (eventIdParam) {
            eventId = eventIdParam;
        } 

        const data = {
            event: eventId,
        };

        let response = await window.Alma.req.post('/event/get', window.Alma.req.createPostData(data), { reload: false });

        let eventName = response.result.event_name;

        if (eventIdLocal !== eventId) {
            window.Alma.localStorage.set(window.Alma.localStorage.event_id, eventId);
            window.Alma.localStorage.set(window.Alma.localStorage.event_name, response.result.event_name);
        }

        if (!eventIdParam) {
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
