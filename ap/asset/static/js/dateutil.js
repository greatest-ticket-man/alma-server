window.Alma = window.Alma || {};
(function(_Alma) {

    class DateUtil {
        constructor() {

        }

        /**
         * Date型からTimestampの型を作成する
         * @param {Date} date 
         */
        DateToTimestamp(date) {
            return {
                seconds: date / 1000,
                nanos: (date % 1000) * 1e6,
            };
        }
    }


    _Alma.dateutil = new DateUtil();
})(window.Alma);