window.Alma = window.Alma || {};
(function(_Alma) {
    class Location {

        constructor() {
            // TODO href path

            this.event_create = '/event/create';
            this.home_dashboard = '/home/dashboard';

        }


        /**
         * 遷移する
         */
        href(path) {

            const url = new URL(path, `${location.protocol}//${document.domain}:${location.port}`);
            const params = new URLSearchParams(url.search.slice(1));

            // eventのparamがない場合は、追加する
            if (!params.has('event')) {
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