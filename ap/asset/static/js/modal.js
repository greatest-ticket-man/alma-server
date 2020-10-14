'use strict';
/**
 * 
 * modal.jsは、/static/css/modal.cssに依存しています
 * 
 * 今のところ、modalは1画面に1つのみサポートしています
 * 2つ以上表示する必要が出てきた場合は、prefixを変更して別々に動作するように切り替えます
 * 
 * 使用例
 *  window.Alma.modal.info('title', 'body', function() {
 *           alert('ok!');
 *           window.Alma.modal.delete();
 *       }, '確認');
 */

window.Alma = window.Alma || {};
(function(_Alma) {

    class Modal {

        constructor() {

            // 動的に作成されるため、constructorでaddEventListenerは呼ばない
            this.modalCancelButtonEl;
            this.modalContainerEl;
            this.modalOkEl;
            this.modalEl;

            // bind
            this.show = this.show.bind(this);
            this.info = this.info.bind(this);
            this.delete = this.delete.bind(this);
            this.blockClicks = this.blockClicks.bind(this);
        }

        /**
         * 
         * @param {function} okFunc okボタンが押されたときの、func
         */
        addEventListener(okFunc) {
            this.modalContainerEl.addEventListener('click', this.delete);
            this.modalEl.addEventListener('click', this.blockClicks);
            this.modalCancelButtonEl.addEventListener('click', this.delete);
            this.modalOkEl.addEventListener('click', okFunc);
        }

        /**
         * info
         * @param {string} title title
         * @param {string} body modalの内容
         * @param {function} func okボタンを押したときの動作
         * @param {string} funcMessage okボタンのメッセージ
         */
        info(title, body, func, funcMessage) {
            this.show('info', title, body, func, funcMessage);
        }

        /**
         * show modalを表示する
         * @param {string} icon titleの前のicon, material-iconsをサポート
         * @param {string} title title
         * @param {string} body modalの　boday
         * @param {function} func okボタンを押したときの動作
         * @param {string} funcMessage okボタンのメッセージ
         */
        show(icon, title, body, func, funcMessage) {
            let modalContainerEl = document.createElement('div');
            modalContainerEl.classList.add('js-modal-container', 'modal__container');
            modalContainerEl.innerHTML = `
                <div class="js-modal modal">
                    <div class="modal__title">
                        <i class="material-icons">${icon}</i>
                        <div>${title}</div>
                    </div>
                    <div class="modal__body">${body}</div>
                    <div class="modal__button-group">
                        <button class="js-modal-cancel modal__button-cancel">キャンセル</button>
                        <button class="js-modal-ok modal__button-primary">${funcMessage}</button>
                    </div>
                </div>
            `;
            document.body.appendChild(modalContainerEl);

            // el
            this.modalCancelButtonEl = document.querySelector('.js-modal-cancel');
            this.modalContainerEl = document.querySelector('.js-modal-container');
            this.modalOkEl = document.querySelector('.js-modal-ok');
            this.modalEl = document.querySelector('.js-modal');

            // add event litener
            this.addEventListener(func);
        }

        /**
         * modalを削除する
         */
        delete() {
            if (this.modalContainerEl) {
                this.modalContainerEl.remove();
            }
        }

        blockClicks(evt) {
            evt.stopPropagation();
        }
    }

    _Alma.modal = new Modal();

})(window.Alma);
