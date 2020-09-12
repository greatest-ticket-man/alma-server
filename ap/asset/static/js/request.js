window.Alma = window.Alma || {};
(function(_Alma) {

    class Req {

        constructor() {
            this.event_create = '/event/create';
            this.event_update = '/event/update';
            this.event_list = '/event/list'; // get
            this.ticket_create = '/ticket/create';
        }

        async get(url, params = {}, options = { reload: false }) {

            const data = this.createGetData();

            // パラメータを追加
            const urlObj = new URL(url, `${location.protocol}//${document.domain}:${location.port}`);
            for (const [key, value] of Object.entries(params)) {
                urlObj.searchParams.set(key, value);
            }

            return this.fetch(urlObj.toString(), data, options);
        }

        async post(url, data, options = { reload: false }) {
            return this.fetch(url, data, options);
        }

        // fetch jsonを取得する
        async fetch(url, data, options) {
            
            try {
                const response = await fetch(url, data);
                const json = await response.json();
                console.log("json is ", json);
                if (json.success) {
                    console.log(`通信に成功しました: ${JSON.stringify(json)}`);
                    if (options.reload === true) {
                        window.Alma.toast.success('成功しました');
                        setTimeout(() => location.reload(), 2000);
                    }
                    return json;
                } else {
                    window.Alma.toast.error("失敗しました");
                    console.log(json);
                    return json;
                }

            } catch (e) {
                window.Alma.toast.error('通信に失敗しました');
                console.log(e);
            }

        }

        createData(data = {}, cache = 'no-cache', method = '') {
            const d = {
                method: method, // *GET, POST, PUT, DELETE, etc.
                mode: 'cors', // no-cors, cors, *same-origin
                cache: cache, // *default, no-cache, reload, force-cache, only-if-cached
                credentials: 'same-origin', // include, same-origin, *omit
                headers: {
                    'Content-Type': 'application/json; charset=utf-8',
                    // 'Content-Type': 'application/x-www-form-urlencoded',
                },
                redirect: 'follow', // manual, *follow, error
                referrer: 'no-referrer', // no-referrer, *client
            };

            // GETはbodyは追加できない
            if (method === 'POST') {
                d.body = JSON.stringify(data); // 本文のデータ型は 'Content-Type' ヘッダーと一致する必要があります
            }
            return d;
        }

        createPostData(data = {}, cache = 'no-cache') {
            return this.createData(data, cache, 'POST');
        }

        createGetData(cache = 'no-cache') {
            return this.createData(null, cache, 'GET'); 
        }
        

    }

    _Alma.req = new Req();
    
})(window.Alma);
