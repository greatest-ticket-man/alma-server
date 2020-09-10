window.Alma = window.Alma || {};
(function(_Alma) {
    class Location {

        constructor() {
            this.event_info = '/event';
            this.event_create = '/event/create';
            this.event_update = '/event/update';
            this.member_info = '/member';
            this.home_dashboard = '/home/dashboard';
            this.home_dashboard_empty = '/home/dashboard/empty';
            this.reserve_info = '/reserve';
            this.ticket_info = '/ticket';

            this.baseURL = `${location.protocol}//${document.domain}:${location.port}`;
        }


        /**
         * 遷移する
         * ex) window.Alma.location.href(window.Alma.location.event_info);
         */
        href(path = "", options = { eventPath: true, ordinaryMode: false }) {
            // 通常遷移モード
            if (options.ordinaryMode) {
                window.location.href = path;
                return;
            } 

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

            // eventPath trueでイベントパス持っているときにLocalStorageにあたりを塗り替える
            if (options.eventPath == true && params.has('event')) {
                window.Alma.localStorage.set(window.Alma.localStorage.event_id, params.get('event'));
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
