'use strict';

// modal.jsは、/static/css/modal.cssに依存しています

window.Alma = window.Alma || {};
(function(_Alma) {

    class Modal {

        constructor() {

            // 動的に作成されるため、constructorでaddEventListenerは呼ばない
            this.modalCancelButtonEl;
            this.modalContainerEl;
            this.modalEl;

            // bind
            this.show = this.show.bind(this);
            this.delete = this.delete.bind(this);
            this.blockClicks = this.blockClicks.bind(this);


        }

        addEventListener() {
            this.modalContainerEl.addEventListener('click', this.delete);
            this.modalEl.addEventListener('click', this.blockClicks);
            this.modalCancelButtonEl.addEventListener('click', this.delete);
        }

        show() {
            let modalContainerEl = document.createElement('div');
            modalContainerEl.classList.add('js-modal-container', 'modal__container');
            modalContainerEl.innerHTML = `
                <div class="js-modal modal">
                    <div class="modal__title">
                        <i class="material-icons">create</i>
                        <div>TITLETITLETITLETITLETITLETITLETITLETITLETITLETITLETITLETITLETITLETITLETITLE</div>
                    </div>
                    <div class="modal__body">BODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODYBODY</div>
                    <div class="modal__button-group">
                        <button class="js-modal-cancel modal__button-cancel">キャンセル</button>
                        <button class="modal__button-primary">OK</button>
                    </div>
                </div>
            `;
            document.body.appendChild(modalContainerEl);

            // el
            this.modalCancelButtonEl = document.querySelector('.js-modal-cancel');
            this.modalContainerEl = document.querySelector('.js-modal-container');
            this.modalEl = document.querySelector('.js-modal');

            this.addEventListener();
        }

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
