'use strict';

class EventForm {

    constructor() {

        // getEl
        this.emailTextEl = document.querySelector('.js-email-text');
        this.addEmailTableButtonEl = document.querySelector('.js-email-table-add');
        this.emailTableEl = document.querySelector('.js-email-table');

        // this.emailTableRowTemplateEl = document.querySelector('.js-email-table-row-template');
        this.emailTablePulldownEl = document.querySelector('.js-email-table-pulldown');

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

        // tableに追加
        let row = this.emailTableEl.insertRow(-1);

        // pullDownを取得
        let pullDown = this.emailTablePulldownEl;
        pullDown.removeAttribute('.display-none');

        row.innerHTML = unescape(`
            <tr>
                <td>
                    <span class="input-container__email-table__icon material-icons">perm_identity</span>
                </td>
                <td>id</td>
                <td>name</td>
                <td>${emailText}</td>
                <td>
                    ${this.emailTablePulldownEl.innerHTML}
                </td>
                <td><button class="input-container__email-table__delete-button  material-icons"
                        onclick="formCreate.deleteMemberToTable(this);">delete</button></td>
            </tr>
        `);


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
