'use strict';

class SideNav {
    constructor() {
        // 表示も非表示も同じボタンにする
        this.showHideButtonEl = document.querySelector('.js-menu-show-hide');
        
        // SideNavのElement
        this.sideNavEl = document.querySelector('.js-side-nav');

        // SideNavのContainerElement
        this.sideNavContainerEl = document.querySelector('.js-side-nav-container');


        // Containerの子にフォーカスできるかどうかを制御
        // 初期状態が画面外のため、フォーカスを無効にする
        this.detabinator = new Detabinator(this.sideNavContainerEl);
        this.detabinator.insert = true;

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

    // blkockClicks クリックイベントを無効にする
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
        this.detabinator.insert = false;

        this.transitionEndProperty = 'transform';

        // 遷移の時間(遷移を区別するために一意にしている)
        this.transitionEndTime = 0.33;

        this.sideNavEl.addEventListener('transitionend', this.onTransitionEnd);
    }

    // hideSideNav SideNavを隠す
    hideSideNav() {
        this.sideNavEl.classList.add('side-nav--animatable');
        this.sideNavEl.classList.remove('side-nav--visible');
        this.detabinator.insert = true;

        this.transitionEndProperty = 'transform';
        this.transitionEndTime = 0.13;

        this.sideNavEl.addEventListener('transitionend', this.onTransitionEnd);
    }
}

new SideNav();
