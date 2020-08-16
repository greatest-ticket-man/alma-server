'use strict';

class EventMenu {
    constructor() {
        this.showButtonElList = document.querySelectorAll('.js-event-menu-show');

        this.eventMenuEl = document.querySelector('.js-event-menu');

        this.eventMenuContainerEl = document.querySelector('.js-event-menu-container');

        this.eventNameEl = document.querySelector('.js-event-name');

        this.hideButtonEl = document.querySelector('.js-event-menu-hide');

        this.goEventCreateFormButtonEl = document.querySelector('.js-go-event-create-form');

        this.showEventMenu = this.showEventMenu.bind(this);
        this.hideEventMenu = this.hideEventMenu.bind(this);
        this.blockClicks = this.blockClicks.bind(this);
        this.goEventCreateForm = this.goEventCreateForm.bind(this);

        // Eventの登録
        this.addEventListener();

        this.setEventName();
    }

    // addEventListener 各要素にイベントを追加する
    addEventListener() {

        const showEventMenuFunc = this.showEventMenu;
        this.showButtonElList.forEach(function(elem) {
            elem.addEventListener('click', showEventMenuFunc);
        })

        this.eventMenuEl.addEventListener('click', this.hideEventMenu);
        this.eventMenuContainerEl.addEventListener('click', this.blockClicks);

        this.hideButtonEl.addEventListener('click', this.hideEventMenu);

        this.goEventCreateFormButtonEl.addEventListener('click', this.goEventCreateForm);

    }

    // setEventName イベントの名前を表示させる
    setEventName() {   

        const eventName = window.Alma.localStorage.get(window.Alma.localStorage.event_name);
        
        if (eventName == null) {
            return;
        }

        this.eventNameEl.innerHTML = eventName;
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

    // goEventCreateForm 
    goEventCreateForm() {
        window.location.href = '/event/create';
    }
}

new EventMenu();
