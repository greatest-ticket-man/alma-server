'use strict';

class SideNav {
    constructor() {
        this.showButtonEl = document.querySelector('.js-menu-show');
        this.hideButtonEl = document.querySelector('.js-menu-hide');
        this.sideNavEl = document.querySelector('.js-side-nav');
        this.sideNavContainerEl = document.querySelector('.js-side-nav-container');

        // Containerの子にフォーカスできるかどうかを制御します
        // ドロワーがスクリーンのため、初期状態を無効にします
        // 初期状態が画面外のため、フォーカスを無効にします
        console.log("sideNavContainerEl" + this.sideNavContainerEl);
        this.detabinator = new Detabinator(this.sideNavContainerEl);
        // this.detabinator.inert = true;
        this.detabinator.inert = true;

        // 何してんだろうこれ
        this.showSideNav = this.showSideNav.bind(this);
        this.hideSideNav = this.hideSideNav.bind(this);
        this.blockClicks = this.blockClicks.bind(this);
        this.onTouchStart = this.onTouchStart.bind(this);
        this.onTouchMove = this.onTouchMove.bind(this);
        this.onTouchEnd = this.onTouchEnd.bind(this);
        this.onTransitionEnd = this.onTransitionEnd.bind(this);
        this.update = this.update.bind(this);
        
        this.startX = 0;
        this.currentX = 0;
        this.touchingSideNav = false;

        this.transitionEndProperty = null;
        this.transitionEndTime = 0;

        this.supportsPassive = undefined;
        this.addEventListener();
    }


    // applyPassive passive event listeningがサポートされている場合は有効にする
    // これは、タッチ系のイベントのときに、ScrollJankなどの問題が発生しないようにする対応
    applyPassive() {
        // オプションが指定されている場合は、返す
        if (this.supportsPassive !== undefined) {
            return this.supportsPassive ? {passive: true} : false;
        }

        // ない場合は、機能があるか検出する
        let isSupported = false;
        try {
            document.addEventListener('test', null, {get passive(){
                isSupported = true;
            }});
        } catch(e) {}

        this.supportsPassive = isSupported;
        return this.applyPassive(); // これ頭いい...
    }

    // addEventListeners 各要素にイベントを追加する
    addEventListener() {
        this.showButtonEl.addEventListener('click', this.showSideNav);
        this.hideButtonEl.addEventListener('click', this.hideSideNav);
        this.sideNavEl.addEventListener('click', this.hideSideNav);
        this.sideNavContainerEl.addEventListener('click', this.blockClicks);

        this.sideNavEl.addEventListener('touchstart', this.onTouchStart, this.applyPassive());
        this.sideNavEl.addEventListener('touchmove', this.onTouchMove, this.applyPassive());
        this.sideNavEl.addEventListener('touchend', this.onTouchEnd);
    }

    onTouchStart(evt) {
        // 表示状態で無いなら、何もしない
        if (!this.sideNavEl.classList.contains('side-nav--visible')) {
            return;
        }

        this.startX = evt.touches[0].pageX;
        this.currentX = this.startX;

        this.touchingSideNav = true;
        window.requestAnimationFrame(this.update);
    }

    onTouchMove(evt) {
        if (!this.touchingSideNav) {
            return;
        }

        this.currentX = evt.touches[0].pageX;
    }

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

    update() {
        if (!this.touchingSideNav) {
            return;
        }

        window.requestAnimationFrame(this.update);

        const translateX = Math.min(0, this.currentX - this.startX);
        this.sideNavContainerEl.getElementsByClassName.transform = `translateX(${translateX}px)`;
    }

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

        this.sideNavEl.classList.remove('side-nav-animatable');
        this.sideNavEl.removeEventListener('transitionend', this.onTransitionEnd);
    }

    // showSideNav SideNavを表示する
    showSideNav() {
        this.sideNavEl.classList.add('side-nav--animatable');
        this.sideNavEl.classList.add('side-nav--visible');
        this.detabinator.inert = false;

        this.transitionEndProperty = 'transform';

        // 遷移の時間(遷移を区別するために一意にする)
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
    }

}

new SideNav();
