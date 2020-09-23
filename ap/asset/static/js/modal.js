// bootstrap依存だから削除しなければ
window.Alma = window.Alma || {};
(function(_Alma) {

    class Modal {

        /**
         * 確認Modal
         * @param {string} message メッセージ
         * @param {function} func okボタンを押したときの動作
         * @param {string} funcMessage okボタンのメッセージ
         */
        confirm(message, func, funcMessage) {
            this.show('確認', message, func, funcMessage);
        }

        /**
         * 
         * @param {string} title modalのtitle
         * @param {string} body modalのbody
         * @param {function} func okボタンを押した時の動作
         * @param {string} funcMessage okボタンのメッセージ
         */
        show(title, body, func, funcMessage) {

            let modalDiv = document.createElement('div');
            modalDiv.classList.add('modal', 'fade');
            modalDiv.setAttribute('data-backdrop', 'static');
            modalDiv.setAttribute('data-keyboard', 'false');
            modalDiv.setAttribute('tabindex', '-1');
            modalDiv.setAttribute('aria-hidden', 'true');

            let modalDialogDiv = document.createElement('div');
            modalDialogDiv.className = 'modal-dialog';

            let modalContentDiv = document.createElement('div');
            modalContentDiv.className = 'modal-content';

            let modalHeaderDiv = document.createElement('div');
            modalHeaderDiv.className = 'modal-header';

            let modalTitleH5 = document.createElement('div');
            modalTitleH5.className = 'modal-title';
            modalTitleH5.innerText = title;

            let modalCloseButton = document.createElement('button');
            modalCloseButton.type = 'button';
            modalCloseButton.className = 'close';
            modalCloseButton.setAttribute('data-dismiss', 'modal');
            modalCloseButton.setAttribute('aria-label', 'Close');

            let modalCloseButtonSpan = document.createElement('span');
            modalCloseButtonSpan.setAttribute('aria-hidden', 'true');
            modalCloseButtonSpan.innerHTML = '&times;';

            let modalBody = document.createElement('div');
            modalBody.className = 'modal-body';
            modalBody.innerHTML = body;

            let modalFooter = document.createElement('div');
            modalFooter.className = 'modal-footer';

            let modalFooterCloseButton = document.createElement('button');
            modalFooterCloseButton.type = 'button';
            modalFooterCloseButton.classList.add('btn', 'btn-secondary');
            modalFooterCloseButton.setAttribute('data-dismiss', 'modal');
            modalFooterCloseButton.innerText = 'Close';
            
            let modalFooterOkButton = document.createElement('button');
            modalFooterOkButton.type = 'button';
            modalFooterOkButton.classList.add('btn', 'btn-primary');
            modalFooterOkButton.innerText = funcMessage;

            // 組み立て

            // button
            modalCloseButton.appendChild(modalCloseButtonSpan);

            // header
            modalHeaderDiv.appendChild(modalTitleH5);
            modalHeaderDiv.appendChild(modalCloseButton);

            // footer
            modalFooter.appendChild(modalFooterCloseButton);
            modalFooter.appendChild(modalFooterOkButton);


            modalContentDiv.appendChild(modalHeaderDiv);
            modalContentDiv.appendChild(modalBody);
            modalContentDiv.appendChild(modalFooter);

            modalDialogDiv.appendChild(modalContentDiv);

            modalDiv.appendChild(modalDialogDiv);


            let modal = new bootstrap.Modal(modalDiv, {});
            modal.show();

            // okボタンが押された時の挙動
            modalFooterOkButton.onclick = function() {
                // 引数のfunctionを実行
                func();

                // modalを閉じる
                modal.hide();
            };
        }
    }


    _Alma.modal = new Modal();
})(window.Alma)
