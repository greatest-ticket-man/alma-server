'use strict';

class MemberInfo {
    constructor() {
        // getEL
        this.memberTableRowElList = document.querySelectorAll('.js-member-table-row');
        this.memberTableHeadCheckboxEl = document.querySelector('.js-member-table-head-checkbox');
        

        // bind
        this.checkRow = this.checkRow.bind(this);
        this.checkAllRow = this.checkAllRow.bind(this);
        this.checkAll = this.checkAll.bind(this);
        this.uncheckAll = this.uncheckAll.bind(this);


        this.addEventListener();
    }

    addEventListener() {
        const me = this;
        this.memberTableRowElList.forEach(function(elem) {
            elem.addEventListener('click', me.checkRow);
        });

        this.memberTableHeadCheckboxEl.addEventListener('click', this.checkAllRow);
    }

    // checkRow そのテーブルにRowを追加する
    checkRow(elem) {
        let checkBoxEl = elem.currentTarget.children[0].children[0];
        checkBoxEl.checked = !checkBoxEl.checked;
    }

    // checkAllRow 
    checkAllRow() {
    
        if (this.memberTableHeadCheckboxEl.checked) {
            this.checkAll();
        } else {
            this.uncheckAll();
        }
    }

    // checkAll すべてのRowをcheckする
    checkAll() {
        this.memberTableRowElList.forEach(function(elem) {
            elem.children[0].children[0].checked = true;
        })
    }

    // uncheckAll すべてのRowのcheckを外す
    uncheckAll() {
        this.memberTableRowElList.forEach(function(elem) {
            elem.children[0].children[0].checked = false;
        })
    }

    
}

new MemberInfo();
