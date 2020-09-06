'use strict';

class SideNav {
    constructor() {
        // 表示も非表示も同じボタンにする
        this.showHideButtonEl = document.querySelector('.js-menu-show-hide');
        
        // SideNavのElement
        this.sideNavEl = document.querySelector('.js-side-nav');

        // SideNavのContainerElement
        this.sideNavContainerEl = document.querySelector('.js-side-nav-container');
        this.sideNavSubContainerElList = document.querySelectorAll('.side-nav__sub-menu__container');

        this.sideNavMenuRowElList = document.querySelectorAll('.js-side-nav-row');
        this.sideNavSubMenuRowElList = document.querySelectorAll('.js-side-nav-sub-row');

        this.sideNavHeader = document.querySelector('.js-side-nav-header');


        // Containerの子にフォーカスできるかどうかを制御
        // 初期状態が画面外のため、フォーカスを無効にする
        this.detabinator = new Detabinator(this.sideNavContainerEl);
        this.detabinator.inert = true;

        // 関数を登録している
        this.showSideNav = this.showSideNav.bind(this);
        this.hideSideNav = this.hideSideNav.bind(this);
        this.showHideSideNav = this.showHideSideNav.bind(this);
        this.blockClicks = this.blockClicks.bind(this);
        this.onTouchStart = this.onTouchStart.bind(this);
        this.onTouchMove = this.onTouchMove.bind(this);
        this.onTouchEnd = this.onTouchEnd.bind(this);
        this.onTransitionEnd = this.onTransitionEnd.bind(this);
        this.update = this.update.bind(this);
        this.goPath = this.goPath.bind(this);
        this.goSubPath = this.goSubPath.bind(this);
        this.goDashboard = this.goDashboard.bind(this);
        this.showSubMenu = this.showSubMenu.bind(this);
        this.hideAllSubMenu = this.hideAllSubMenu.bind(this);
        

        // fieldの初期化
        this.startX = 0;
        this.currentX = 0;
        this.touchingSideNav = false;

        this.transitionEndProperty = null;
        this.transitionEndTime = 0;

        this.supportsPassive = undefined;

        // Eventの登録
        this.addEventListener();

        
    }

    // applyPassive passive event listeningがサポートされている場合は有効にする
    applyPassive() {
        if (this.supportsPassive !== undefined) {
            return this.supportsPassive ? {passive: true} : false;
        }

        // 機能があるか確認する
        let isSupported = false;
        try {
            document.addEventListener('test', null, {get passive() {
                isSupported = true;
            }});
        } catch(e){}

        this.supportsPassive = isSupported;
        return this.applyPassive();
    }
 
    // addEventListener 各要素にイベントを追加する
    addEventListener() {

        this.showHideButtonEl.addEventListener('click', this.showHideSideNav);

        this.sideNavEl.addEventListener('click', this.hideSideNav);
        this.sideNavContainerEl.addEventListener('click', this.blockClicks);

        this.sideNavEl.addEventListener('touchstart', this.onTouchStart, this.applyPassive());
        this.sideNavEl.addEventListener('touchmove', this.onTouchMove, this.applyPassive());
        this.sideNavEl.addEventListener('touchend', this.onTouchEnd);

        const me = this;
        this.sideNavMenuRowElList.forEach(function(elem) {
            elem.addEventListener('click', me.goPath);
            elem.addEventListener('mouseenter', me.showSubMenu);
        });


        this.sideNavSubContainerElList.forEach(function (elem) {
            // menuのclick eventが優先されてしまうため、先にブロックする
            elem.addEventListener('click', me.blockClicks);
        });

        this.sideNavSubMenuRowElList.forEach(function (elem) {
            elem.addEventListener('click', me.goSubPath);
        })

        this.sideNavHeader.addEventListener('click', this.goDashboard);
    }

    // onTouchStart .
    onTouchStart(evt) {
        // 表宇治状態なら、何もしない
        if (!this.sideNavEl.classList.contains('side-nav--visible')) {
            return;
        }

        this.startX = evt.touches[0].pageX;
        this.currentX = this.startX;

        this.touchingSideNav = true;
        window.requestAnimationFrame(this.update);
    }

    // onTouchMove .
    onTouchMove(evt) {
        if (!this.touchingSideNav) {
            return;
        }

        this.currentX = evt.touches[0].pageX;
    }

    // onTouchEnd .
    onTouchEnd(evt) {
        if (!this.touchingSideNav) {
            return;
        }

        this.touchingSideNav = false;

        const translateX = Math.min(0, this.currentX - this.startX);
        this.sideNavContainerEl.getElementsByClassName.transform = '';

        if (translateX < 0) {
            this.hideSideNav();
        }
    }

    // update .
    update() {
        if (!this.touchingSideNav) {
            return;
        }

        window.requestAnimationFrame(this.update);

        const translateX = Math.min(0, this.currentX - this.startX);
        this.sideNavContainerEl.getElementsByClassName.transform = `translateX(${translateX})`;
    }

    // blockClicks クリックイベントを無効にする
    blockClicks(evt) {
        evt.stopPropagation();
    }

    // onTransitionEnd 移動が終わったときのやつ
    onTransitionEnd(evt) {
        if (evt.propertyName != this.transitionEndProperty && evt.elapsedTime != this.transitionEndTime) {
            return;
        }

        this.transitionEndProperty = null;
        this.transitionEndTime = 0;

        this.sideNavEl.classList.remove('side-nav--animatable');
        this.sideNavEl.removeEventListener('transitionend', this.onTransitionEnd);
    }

    // showHideSideNav SideNavが非表示なら表示、表示なら非表示にする
    showHideSideNav() {
        if (this.sideNavEl.classList.contains('side-nav--visible')) {
            this.hideSideNav();
            return;
        }

        this.showSideNav();
    }

    // showSideNav SideNavを表示する
    showSideNav() {
        this.sideNavEl.classList.add('side-nav--animatable');
        this.sideNavEl.classList.add('side-nav--visible');
        this.detabinator.inert = false;

        this.transitionEndProperty = 'transform';

        // 遷移の時間(遷移を区別するために一意にしている)
        this.transitionEndTime = 0.33;

        this.sideNavEl.addEventListener('transitionend', this.onTransitionEnd);
    }

    // hideSideNav SideNavを隠す
    hideSideNav() {
        this.sideNavEl.classList.add('side-nav--animatable');
        this.sideNavEl.classList.remove('side-nav--visible');
        this.detabinator.inert = true;

        this.transitionEndProperty = 'transform';
        this.transitionEndTime = 0.13;

        this.sideNavEl.addEventListener('transitionend', this.onTransitionEnd);

        this.hideAllSubMenu();
    }

    // showSubMenu SubMenuを表示する
    showSubMenu(elem) {

        this.hideAllSubMenu();

        let sideNavMenuRowEl = elem.currentTarget;
        let sideNavSubMenuContainerEl = sideNavMenuRowEl.querySelector('.side-nav__sub-menu__container');
        sideNavSubMenuContainerEl.classList.add('side-nav__sub-menu__container-visible');

    }

    // hideAllSubMenu すべてのSubMenuを非表示にする
    hideAllSubMenu() {
        this.sideNavMenuRowElList.forEach(function(elem) {
            let sideNavSubMenuContainerEl = elem.querySelector('.side-nav__sub-menu__container');
            sideNavSubMenuContainerEl.classList.remove('side-nav__sub-menu__container-visible');
        });
    }

    goDashboard() {
        window.Alma.location.href(window.Alma.location.home_dashboard);
    }

    // menuの遷移
    goPath(elem) {
        let sideNavMenuRow = elem.currentTarget;
        let path = sideNavMenuRow.querySelector('.js-side-nav-path').innerText;

        window.Alma.location.href(path);
    }

    // subMenuの遷移
    goSubPath(elem) {
        let sideNavSubMenuRow = elem.currentTarget;
        let path = sideNavSubMenuRow.querySelector('.js-side-nav-sub-path').innerText;

        window.Alma.location.href(path);
    }
    
}

new SideNav();
