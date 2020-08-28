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

        this.eventTitleEl = document.getElementById('js-event-title');
        this.organizationNameEl = document.getElementById('js-event-organization');

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
                window.Alma.toast.warn(`${emailText}はメールアドレスではありません`);
                return;
            }

            this.pushTable(emailText);
        })

        // emailTextをclier
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

    // getEventName eventNameのValueを取得する
    getEventName() {
        return this.eventTitleEl.value;
    }

    // getOrganizationName
    getOrganizationName() {
        return this.organizationNameEl.value;
    }

    // getMemberInfoList tableのデータを取得する
    getMemberInfoList() {

        let memberInfoList = [];

        for (let row of this.emailTableEl.rows) {

            let memberInfo = {
                email: '',
                authority: '',
            };

            for (let cell of row.cells) {

                if (cell.classList.contains('js-email-table-email')) {
                    memberInfo.email = cell.innerText;
                } else if (cell.classList.contains('js-email-table-auth')) {
                    const select = cell.children[0];
                    memberInfo.authority = select.value;
                }

            }
            memberInfoList.push(memberInfo);
        }

        return memberInfoList;
    }

    

}

const eventForm = new EventForm();
