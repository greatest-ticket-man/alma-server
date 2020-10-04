'use strict';

class ReserveCreate {

    constructor() {

        // EL
        this.backButtonEl = document.querySelector('.js-reserve-info-back');
        this.cancelButtonEl = document.querySelector('.js-reserve-create-cancel');

        this.stepOrderEl = document.querySelector('.js-step-order');
        this.stepCustomorEl = document.querySelector('.js-step-customor');
        this.stepPayEl = document.querySelector('.js-step-pay');
        this.stepConfirmEl = document.querySelector('.js-step-confirm');

        this.formInitEl = document.querySelector('.js-form-init');
        this.formOrderEl = document.querySelector('.js-form-order');
        this.formCustomorEl = document.querySelector('.js-form-customor');
        this.formConfirmEl = document.querySelector('.js-form-confirm');
        this.formPayEl = document.querySelector('.js-form-pay');


        // Bind
        this.goReserveInfoPage = this.goReserveInfoPage.bind(this);
        this.stepOrder = this.stepOrder.bind(this);
        this.stepCustomor = this.stepCustomor.bind(this);
        this.stepPay = this.stepPay.bind(this);
        this.stepConfirm = this.stepConfirm.bind(this);
        this.stepInit = this.stepInit.bind(this);

        this.formContentAllHide = this.formContentAllHide.bind(this);
        this.stepContentAllDeselect = this.stepContentAllDeselect.bind(this);
        this.switchForm = this.switchForm.bind(this);


        this.addEventListener();

    }

    addEventListener() {
        this.backButtonEl.addEventListener('click', this.goReserveInfoPage);
        this.cancelButtonEl.addEventListener('click', this.goReserveInfoPage);
        this.stepOrderEl.addEventListener('click', this.stepOrder);
        this.stepCustomorEl.addEventListener('click', this.stepCustomor);
        this.stepPayEl.addEventListener('click', this.stepPay);
        this.stepConfirmEl.addEventListener('click', this.stepConfirm);
    }

    // init
    stepInit() {


        // all hide
        this.formContentAllHide();

        this.formInitEl.classList.add('form-content--visible');

    }

    // 購入情報のステップ
    stepOrder() {
        this.switchForm(this.stepOrderEl, this.formOrderEl);
    }

    // お客様情報のステップ
    stepCustomor() {
        this.switchForm(this.stepCustomorEl, this.formCustomorEl);
    }

    // お支払い情報のステップ
    stepPay() {
        this.switchForm(this.stepPayEl, this.formPayEl);
    }

    // 確認画面のステップ
    stepConfirm() {
        this.switchForm(this.stepConfirmEl, this.formConfirmEl);
    }

    // switchForm
    switchForm(stepEl, formEl) {
        if (formEl.classList.contains('form-content--visible')) {
            this.stepInit();
            return;
        }

        this.formContentAllHide();
        this.stepContentAllDeselect();

        formEl.classList.add('form-content--visible');
        stepEl.classList.add('step-content--selected');
    }

    // 全てのコンテンツを非表示にする
    formContentAllHide() {
        this.formInitEl.classList.remove('form-content--visible');
        this.formOrderEl.classList.remove('form-content--visible');
        this.formCustomorEl.classList.remove('form-content--visible');
        this.formPayEl.classList.remove('form-content--visible');
        this.formConfirmEl.classList.remove('form-content--visible');
    }

    // 全てのステップを選択解除する
    stepContentAllDeselect() {
        this.stepOrderEl.classList.remove('step-content--selected');
        this.stepCustomorEl.classList.remove('step-content--selected');
        this.stepPayEl.classList.remove('step-content--selected');
        this.stepConfirmEl.classList.remove('step-content--selected');
    }

    // goReserveInfoPage 
    goReserveInfoPage() {
        window.Alma.location.href(window.Alma.location.reserve_info);
    }


}

new ReserveCreate();
