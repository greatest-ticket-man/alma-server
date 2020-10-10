'use strict';

class Calendar {
    constructor() {

        console.log("calendar...");

        // valiable
        this.date = new Date();
        this.today = new Date(); // TODO select day
        this.months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];

        // EL
        this.dateTitleEl = document.querySelector('.js-calendar-date-title');
        this.dateDescEl = document.querySelector('.js-calendar-date-desc');
        this.daysEl = document.querySelector('.js-calendar-days');
        this.prevEl = document.querySelector('.js-calendar-prev');
        this.nextEl = document.querySelector('.js-calendar-next');

        // bind
        this.renderCalendar = this.renderCalendar.bind(this);
        this.prevMonth = this.prevMonth.bind(this);
        this.nextMonth = this.nextMonth.bind(this);

        // render ...
        this.renderCalendar();

        this.addEventListener();
    }

    addEventListener() {
        this.prevEl.addEventListener('click', this.prevMonth);
        this.nextEl.addEventListener('click', this.nextMonth);
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

        const nextDate = 7 - lastDayIndex;

        // title
        this.dateTitleEl.innerHTML = this.months[this.date.getMonth()];
        this.dateDescEl.innerHTML = this.date.toDateString();

        // 前の月の日付
        let days = '';
        for (let i = firstDayIndex; i > 0; i--) {
            days += `<div class="calendar__day calendar__prev-date">${prevLastDate - i + 1}</div>`;
        }

        // 今月の日付
        for (let i = 1; i <= lastDate; i++) {
            if (i === this.today.getDate() 
                && this.date.getMonth() === this.today.getMonth()) {
                days += `<div class="calendar__day calendar__day--today">${i}</div>`;
            } else {
                days += `<div class="calendar__day">${i}</div>`;
            }
        }

        // 来月の日付
        for (let i = 1; i <= nextDate; i++) {
            days += `<div class="calendar__day calendar__next-date">${i}</div>`;
        }

        this.daysEl.innerHTML = days;
    }

    prevMonth() {
        this.date.setMonth(this.date.getMonth() - 1);
        this.renderCalendar();
    }

    nextMonth() {
        this.date.setMonth(this.date.getMonth() + 1);
        this.renderCalendar();
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
