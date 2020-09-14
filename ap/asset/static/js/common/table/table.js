'use strict';

class TableInfo {
    constructor(tableClass = "js-table", tableRowClass = "js-table-row", tableHeadCheckboxClass = "js-head-checkbox") {

        this.tableEl = document.querySelector(`.${tableClass}`);

        // getEL
        this.tableRowElList = this.tableEl.querySelectorAll(`.${tableRowClass}`);
        this.tableHeadCheckboxEl = this.tableEl.querySelector(`.${tableHeadCheckboxClass}`);

        // bind
        this.checkRow = this.checkRow.bind(this);
        this.checkAllRow = this.checkAllRow.bind(this);
        this.checkAll = this.checkAll.bind(this);
        this.uncheckAll = this.uncheckAll.bind(this);

        this.addEventListener();
    }

    addEventListener() {
        const me = this;
        this.tableRowElList.forEach(function(elem) {
            elem.addEventListener('click', me.checkRow);
        });

        this.tableHeadCheckboxEl.addEventListener('click', this.checkAllRow);
    }

    // checkRow そのテーブルのチェックする
    checkRow(elem) {
        let checkBoxEl = elem.currentTarget.children[0].children[0];
        checkBoxEl.checked = !checkBoxEl.checked;
    }

    // checkAllRow .
    checkAllRow() {
        if (this.tableHeadCheckboxEl.checked) {
            this.checkAll();
        } else {
            this.uncheckAll();
        }
    }

    // checkAll すべてのRowをcheckする
    checkAll() {
        this.tableRowElList.forEach(function(elem) {
            elem.children[0].children[0].checked = true;
        });
    }

    // uncheckAll すべてのRowのcheckを外す
    uncheckAll() {
        this.tableRowElList.forEach(function(elem) {
            elem.children[0].children[0].checked = false;
        });
    }
}
