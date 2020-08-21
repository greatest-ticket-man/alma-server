'use strict';

class EventForm {

    constructor() {

        // getEl
        this.emailTextEl = document.querySelector('.js-email-text');
        this.addEmailTableButtonEl = document.querySelector('.js-email-table-add');
        this.emailTableEl = document.querySelector('.js-email-table');

        this.addMemberToTable = this.addMemberToTable.bind(this);

        this.addEventListener();
    }

    addEventListener() {

        // this.cancelCreateEventButtonEl.addEventListener('click', this.backBeforePage);
        this.addEmailTableButtonEl.addEventListener('click', this.addMemberToTable);
        // this.createEventButtonEl.addEventListener('click', this.createEvent);

        // Enterで発火するように変更
        const me = this;
        this.emailTextEl.addEventListener('keypress', function (element) {
            if (element.keyCode === 13) {
                element.preventDefault();
                me.addEmailTableButtonEl.click();
            }
        })
    }


    // addMemberToTable tableにemailを追加する 
    addMemberToTable() {
        // 値を取得
        const emailText = this.emailTextEl.value;

        if (!emailText) {
            return;
        }

        // TODO ,でparse

        // TODO serverと通信, idとかnameを取得する, email検索, なければ新規扱い

        // tableに追加
        let row = this.emailTableEl.insertRow(-1);
        row.insertCell(0).innerHTML = `<span class="input-container__email-table__icon material-icons">perm_identity</span>`;
        row.insertCell(1).innerHTML = `id`;
        row.insertCell(2).innerHTML = `name`;
        row.insertCell(3).innerHTML = emailText;
        row.insertCell(4).innerHTML = `観覧`;
        row.insertCell(5).innerHTML = `<button class="input-container__email-table__delete-button  material-icons" onclick="formCreate.deleteMemberToTable(this);"'>delete</button>`;

        // emailTextをcliear
        this.emailTextEl.value = '';
    }

    // deleteMemberToTable memberのtableを削除する
    deleteMemberToTable(elem) {
        let tr = elem.parentNode.parentNode;
        this.emailTableEl.deleteRow(tr.sectionRowIndex);
    }

}

const formCreate = new EventForm();
