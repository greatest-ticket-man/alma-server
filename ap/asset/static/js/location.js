window.Alma = window.Alma || {};
(function(_Alma) {
    class Location {

        constructor() {
            this.event_info = '/event';
            this.event_create = '/event/create';
            this.event_update = '/event/update';
            this.home_dashboard = '/home/dashboard';

            this.baseURL = `${location.protocol}//${document.domain}:${location.port}`;
        }


        /**
         * 遷移する
         * ex) window.Alma.location.href(window.Alma.location.event_info);
         */
        href(path, options = { eventPath: true }) {

            const url = new URL(path, this.baseURL);
            const params = new URLSearchParams(url.search.slice(1));

            // eventのparamがない場合は、追加する
            if (options.eventPath == true && !params.has('event')) {
                // セッションストレージから取得
                const eventId = window.Alma.localStorage.get(window.Alma.localStorage.event_id);
                if (eventId) {
                    params.append('event', eventId);
                } 
            }

            // 遷移
            window.location.href = `${path}?${params.toString()}`;
        }

        /**
         * URLのKeyを取得する
         * @param {string} key param key
         */
        getParam(key = '', path) {

            if (!path) {
                path = location.href;
            }

            const url = new URL(path, this.baseURL);
            const parmas = new URLSearchParams(url.search.slice(1));
            return parmas.get(key);

        }

    }

    _Alma.location = new Location();
})(window.Alma);
