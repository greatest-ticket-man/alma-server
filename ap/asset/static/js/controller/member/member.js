'use strict';

class MemberInfo {
    constructor() {
        // getEL
        this.memberTableRowElList = document.querySelectorAll('.js-member-table-row');

        // bind
        this.checkRow = this.checkRow.bind(this);

        this.addEventListener();
    }

    addEventListener() {

        const me = this;
        this.memberTableRowElList.forEach(function(elem) {
            elem.addEventListener('click', me.checkRow);
        })

    }

    // checkRow そのテーブルにRowを追加する
    checkRow(elem) {
        let checkBoxEl = elem.currentTarget.children[0].children[0];
        checkBoxEl.checked = !checkBoxEl.checked;
    }
}

new MemberInfo();
