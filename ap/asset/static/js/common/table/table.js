'use strict';

class TableInfo {
    constructor(tableClass = "js-table", tableRowClass = "js-table-row", tableHeadCheckboxClass = "js-table-head-checkbox", tableCheckboxClass = "js-table-checkbox") {

        this.tableEl = document.querySelector(`.${tableClass}`);

        // getEL
        this.tableRowElList = this.tableEl.querySelectorAll(`.${tableRowClass}`);
        this.tableHeadCheckboxEl = this.tableEl.querySelector(`.${tableHeadCheckboxClass}`);
        this.tableCheckboxEl = this.tableEl.querySelectorAll(`.${tableCheckboxClass}`);

        // bind
        this.checkRow = this.checkRow.bind(this);
        this.checkAllRow = this.checkAllRow.bind(this);
        this.checkAll = this.checkAll.bind(this);
        this.uncheckAll = this.uncheckAll.bind(this);
        this.blockClicks = this.blockClicks.bind(this);

        this.addEventListener();
    }

    addEventListener() {
        const me = this;
        this.tableRowElList.forEach(function(elem) {
            elem.addEventListener('click', me.checkRow);
        });

        // クリックイベントを無効にして2重チェックされるのを防ぐ
        this.tableCheckboxEl.forEach(function(elem) {
            elem.addEventListener('click', me.blockClicks);
        });


        this.tableHeadCheckboxEl.addEventListener('click', this.checkAllRow);
    }


    // blockClicks クリックイベントを無効にする
    blockClicks(evt) {
        evt.stopPropagation();
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
