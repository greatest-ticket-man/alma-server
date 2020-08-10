'use strict';

class AccountMenu {
    constructor() {
    
        this.showHideButtonEl = document.querySelector('.js-account-menu-show-hide');

        this.accountMenuEl = document.querySelector('.js-account-menu');
        this.accountMenuContainerEl = document.querySelector('.js-account-menu-container');

        this.accountLogoutButtonEl = document.querySelector('.js-account-logout');


        this.detabinator = new Detabinator(this.accountMenuContainerEl);
        this.detabinator.inert = true;

        // thisの登録
        this.showAccountMenu = this.showAccountMenu.bind(this);
        this.hideAccountMenu = this.hideAccountMenu.bind(this);
        this.showHideAccountMenu = this.showHideAccountMenu.bind(this);
        this.logout = this.logout.bind(this);

        // field
        this.supportsPassive = undefined;

        // Eventの登録
        this.addEventListener();
    }

    // applyPassive passive event listeningがサポートされている場合は有効にする
    applyPassive() {
        if (this.supportsPassive !== undefined) {
            return this.supportsPassive ? { passive: true } : false;
        }

        // 機能があるか確認する
        let isSupported = false;
        try {
            document.addEventListener('test', null, {
                get passive() {
                    isSupported = true;
                }
            });
        } catch (e) { }

        this.supportsPassive = isSupported;
        return this.applyPassive();
    }

    // addEventListener 各要素にイベントを追加する
    addEventListener() {
        this.showHideButtonEl.addEventListener('click', this.showHideAccountMenu);

        this.accountMenuEl.addEventListener('click', this.hideAccountMenu);
        this.accountMenuContainerEl.addEventListener('click', this.blockClicks);

        this.accountLogoutButtonEl.addEventListener('click', this.logout);
    }

    // blockClicks クリックイベントを無効にする
    blockClicks(evt) {
        evt.stopPropagation();
    }

    // showHideAccountMenu .
    showHideAccountMenu() {
        if (this.accountMenuEl.classList.contains('account-menu--visible')) {
            this.hideAccountMenu();
            return;
        }
        this.showAccountMenu();
    }

    // showAccountMenu アカウントメニューを表示する
    showAccountMenu() {
        this.accountMenuEl.classList.add('account-menu--visible');
        this.detabinator.inert = false;
    }

    // hideAccountMenu アカウントメニューを非表示にする
    hideAccountMenu() {
        this.accountMenuEl.classList.remove('account-menu--visible');
        this.detabinator.inert = true;
    }

    // logout ログアウトする
    logout() {

        // ログアウト画面に遷移
        window.location.href = '/logout';
    }

}

new AccountMenu();
