'use strict';

class EventMenu {
    constructor() {

        this.showButtonEl = document.querySelector('.js-event-menu-show');

        this.eventMenuEl = document.querySelector('.js-event-menu');

        this.eventMenuContainerEl = document.querySelector('.js-event-menu-container');

        this.hideButtonEl = document.querySelector('.js-event-menu-hide');

        this.showEventMenu = this.showEventMenu.bind(this);
        this.hideEventMenu = this.hideEventMenu.bind(this);
        this.blockClicks = this.blockClicks.bind(this);

        // Eventの登録
        this.addEventListener();
    }

    // addEventListener 各要素にイベントを追加する
    addEventListener() {

        this.showButtonEl.addEventListener('click', this.showEventMenu);

        this.eventMenuEl.addEventListener('click', this.hideEventMenu);
        this.eventMenuContainerEl.addEventListener('click', this.blockClicks);

        this.hideButtonEl.addEventListener('click', this.hideEventMenu);

    }

    // blockClicks クリックイベントを無効にする
    blockClicks(evt) {
        evt.stopPropagation();
    }

    // showEventMenu イベントメニューを表示する
    showEventMenu() {
        this.eventMenuEl.classList.add('event-menu--visible');
    }

    // hideEventMenu イベントメニューを非表示
    hideEventMenu() {
        this.eventMenuEl.classList.remove('event-menu--visible');
    }
}

new EventMenu();
