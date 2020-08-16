window.Alma = window.Alma || {};
(function(_Alma) {

    class LocalStorage {

        constructor() {
            this.event_id = 'event_id'; // イベントのID
            this.event_name = 'event_name'; // イベント名
        }

        /**
         * セット
         * @param {string} key 
         * @param {string} value 
         */
        set(key, value) {
            localStorage.setItem(key, value);
        }

        /**
         * 取得
         * @param {string} key
         * @returns {string} value 
         */
        get(key) {
            return localStorage.getItem(key);
        }

        /**
         * ローカルストレージを削除する
         */
        clear() {
            localStorage.clear();
        }

        /**
         * 削除
         * @param {string} key 
         */
        remove(key) {
            localStorage.removeItem(key);
        }

    }   

    _Alma.localStorage = new LocalStorage();
})(window.Alma);
