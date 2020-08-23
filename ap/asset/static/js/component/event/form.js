'use strict';

// use /static/js/util/validation/validation.js validationUtil

// 依存Classの存在を確認
if (typeof ValidationUtil === 'undefined') {
    alert('依存しているValidationUtilが見つかりませんでした');
    console.error('ValidationUtilが見つかりません');
    console.error('/static/js/util/validation/validation.jsをimportしてください');
}


class EventForm {

    constructor() {



        // getEl
        this.emailTextEl = document.querySelector('.js-email-text');
        this.addEmailTableButtonEl = document.querySelector('.js-email-table-add');
        this.emailTableEl = document.querySelector('.js-email-table-body');

        this.emailTablePulldownEl = document.querySelector('.js-email-table-pulldown');

        this.addMemberToTable = this.addMemberToTable.bind(this);
        this.pushTable = this.pushTable.bind(this);

        this.addEventListener();
    }

    addEventListener() {

        this.addEmailTableButtonEl.addEventListener('click', this.addMemberToTable);

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

        // split
        let emailTextList = emailText.split(',');

        emailTextList.forEach(emailText => {

            if (!ValidationUtil.email(emailText)) {
                // TODO toast
                alert(`${emailText}はメールアドレスではありません`);
                return;
            }

            this.pushTable(emailText);
        })

        // emailTextをcliear
        this.emailTextEl.value = '';
    }

    
    pushTable(email) {
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
                <td class="js-email-table-email">${email}</td>
                <td class="js-email-table-auth input-container__email-table__td">
                    ${this.emailTablePulldownEl.innerHTML}
                </td>
                <td><button class="input-container__email-table__delete-button  material-icons"
                        onclick="eventForm.deleteMemberToTable(this);">delete</button></td>
            </tr>
        `);
    }

    // deleteMemberToTable memberのtableを削除する
    deleteMemberToTable(elem) {
        let tr = elem.parentNode.parentNode;
        this.emailTableEl.deleteRow(tr.sectionRowIndex);
    }

}

const eventForm = new EventForm();
