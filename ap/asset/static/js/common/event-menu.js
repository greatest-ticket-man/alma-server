'use strict';

class EventMenu {
    constructor() {
        this.showButtonElList = document.querySelectorAll('.js-event-menu-show');
        this.eventMenuEl = document.querySelector('.js-event-menu');
        this.eventMenuContainerEl = document.querySelector('.js-event-menu-container');
        this.eventMenuTableEl = document.querySelector('.js-event-menu-table-body');
        this.hideButtonEl = document.querySelector('.js-event-menu-hide');
        this.goEventCreateFormButtonEl = document.querySelector('.js-go-event-create-form');

        this.showEventMenu = this.showEventMenu.bind(this);
        this.getEventList = this.getEventList.bind(this);
        this.hideEventMenu = this.hideEventMenu.bind(this);
        this.blockClicks = this.blockClicks.bind(this);
        this.goEventCreateForm = this.goEventCreateForm.bind(this);

        // Eventの登録
        this.addEventListener();   
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

    // blockClicks クリックイベントを無効にする
    blockClicks(evt) {
        evt.stopPropagation();
    }

    // showEventMenu イベントメニューを表示する
    showEventMenu() {
        this.eventMenuEl.classList.add('event-menu--visible');

        // イベントtableを作成する
        this.createEventTable();
    }

    // createEventTable イベントテーブルを作成する
    async createEventTable() {

        // 先にTableをリセットする
        this.eventMenuTableEl.innerHTML = '';

        let eventInfoList = await this.getEventList();
        if (!eventInfoList) {
            return;
        }

        const me = this;
        eventInfoList.forEach(eventInfo => {
            let row = me.eventMenuTableEl.insertRow(-1);
            row.innerHTML = unescape(`
                <tr>
                    <td><span class="event-menu__container__table__icon material-icons">star</span></td>
                    <td>${eventInfo.event_name}</td>
                    <td>${eventInfo.event_id}</td>
                </tr>
            `);
        });
    }
    

    // getEventList 自分が参加しているイベントのリストを取得する
    async getEventList() {

        const param = {
            search_text: "todo",
        };

        let response = await window.Alma.req.get(window.Alma.req.event_list, param);

        if (!response || !response.success) {
            window.Alma.toast.error('イベントのリスト取得に失敗しました');
            return;
        }

        return response.result.event_info_list;
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
