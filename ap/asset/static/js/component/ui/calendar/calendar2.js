'use strict';

class Calendar {
    constructor() {

        // valiable
        this.date = new Date();
        this.today = new Date();
        this.selectDate;
        this.months = ["1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "９月", "10月", "11月", "12月"];

        // EL
        this.dateTitleEl = document.querySelector('.js-calendar-date-title');
        this.dateDescEl = document.querySelector('.js-calendar-date-desc');
        this.daysEl = document.querySelector('.js-calendar-days');
        this.prevEl = document.querySelector('.js-calendar-prev');
        this.nextEl = document.querySelector('.js-calendar-next');
        this.dayListEl; // renderCalendar();ないで呼び出し

        // bind
        this.renderCalendar = this.renderCalendar.bind(this);
        this.prevMonth = this.prevMonth.bind(this);
        this.nextMonth = this.nextMonth.bind(this);
        this.selectDay = this.selectDay.bind(this);
        this.resetSelectDay = this.resetSelectDay.bind(this);

        // render 
        this.renderCalendar();

        this.addEventListener();
    }

    addEventListener() {
        const me = this;
        this.dayListEl.forEach(function(elem) {
            elem.addEventListener('click', me.selectDay);
        });

        this.prevEl.addEventListener('click', this.prevMonth);
        this.nextEl.addEventListener('click', this.nextMonth);
    }

    // selectDay 日付を選択する
    selectDay(elem) {

        let dayEl = elem.currentTarget;
        if (!dayEl.classList.contains('calendar__day--selected')) {
            this.resetSelectDay();
            
            dayEl.classList.add('calendar__day--selected');

            this.selectDate = new Date(
                this.date.getFullYear(),
                this.date.getMonth(),
                Number(dayEl.innerText),
            );
            this.dateDescEl.innerHTML = `${this.date.getFullYear()}/${this.date.getMonth()+1}/${this.selectDate.getDate()}`;
        } else {
            this.dateDescEl.innerHTML = `----/--/--`;
            this.selectDate = null;
            this.resetSelectDay();
        }
    }

    // resetSelectDay 選択状態を解除する
    resetSelectDay() {
        this.dayListEl.forEach(function(elem) {
            elem.classList.remove('calendar__day--selected');
        });
    }

    // renderCalendar カレンダーをレンダリングする
    renderCalendar() {

        // 日数を1日にする
        this.date.setDate(1);

        // 最終日
        const lastDate = this.getLastDate();
        const prevLastDate = this.getPrevLastDate();

        const firstDayIndex = this.getFirstDayIndex();
        const lastDayIndex = this.getLastDayIndex();

        const nextDate = 7 - lastDayIndex - 1;

        // title
        this.dateTitleEl.innerHTML = `${this.date.getFullYear()}年 ${this.date.getMonth()+1}月`;
        this.dateDescEl.innerHTML = `----/--/--`;

        // 前の月の日付
        let days = '';
        for (let i = firstDayIndex; i > 0; i--) {
            days += `<div class="calendar__day calendar__prev-date">${prevLastDate - i + 1}</div>`;
        }

        // 今月の日付
        for (let i = 1; i <= lastDate; i++) {
            if (i === this.today.getDate() 
                && this.date.getMonth() === this.today.getMonth()) {
                days += `<div class="js-calendar-day calendar__day calendar__day--today">${i}</div>`;
            } else {
                days += `<div class="js-calendar-day calendar__day">${i}</div>`;
            }
        }

        // 来月の日付
        for (let i = 1; i <= nextDate; i++) {
            days += `<div class="calendar__day calendar__next-date">${i}</div>`;
        }

        // 表示を更新
        this.daysEl.innerHTML = days;

        // eventをもう一度作成
        this.dayListEl = document.querySelectorAll('.js-calendar-day');
    }

    prevMonth() {
        this.date.setMonth(this.date.getMonth() - 1);
        this.renderCalendar();
        this.addEventListener();
    }

    nextMonth() {
        this.date.setMonth(this.date.getMonth() + 1);
        this.renderCalendar();
        this.addEventListener();
    }

    getLastDate() {
        let date = this.date;
        return new Date(
            date.getFullYear(),
            date.getMonth() + 1,
            0,
        ).getDate();
    }

    getPrevLastDate() {
        let date = this.date;
        return new Date(
            date.getFullYear(),
            date.getMonth(),
            0,
        ).getDate();
    }

    // 最初の曜日のindexを取得する
    getFirstDayIndex() {
        return this.date.getDay();
    }

    // 最後の日の曜日indexを取得する
    getLastDayIndex() {
        let date = this.date;
        return new Date(
            date.getFullYear(),
            date.getMonth() + 1,
            0,
        ).getDay();
    }

    // 次の曜日のindexを取得する
    prevNextDayIndex(dayIndex) {
        return 7 - dayIndex - 1;
    }

}

new Calendar();
