window.Alma = window.Alma || {};
(function(_Alma) {
    class Location {

        constructor() {
            this.event_info = '/event';
            this.event_create = '/event/create';
            this.home_dashboard = '/home/dashboard';
        }


        /**
         * 遷移する
         * ex) window.Alma.location.href(window.Alma.location.event_info);
         */
        href(path, options = { eventPath: true }) {

            const url = new URL(path, `${location.protocol}//${document.domain}:${location.port}`);
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

    }

    _Alma.location = new Location();
})(window.Alma);
