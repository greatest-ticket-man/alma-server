'use strict';

class Test {
    constructor() {

        // el
        this.modalButtonEl = document.querySelector('.js-test-modal-button');

        // bind
        this.testModal = this.testModal.bind(this);

        this.addEventListener();

    }

    addEventListener() {
        this.modalButtonEl.addEventListener('click', this.testModal);
    }

    testModal() {
        window.Alma.modal.info('title', 'body', function() {
            alert('ok!');
            window.Alma.modal.delete();
        }, '確認');
    }
}

new Test();
