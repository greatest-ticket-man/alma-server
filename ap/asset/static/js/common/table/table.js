'use strict';

class TableInfo {
    constructor(tableClass = "js-table") {

        // 複数のTableを表示する場合に、NameSpaceで分ける必要があるためこのように実装
        this.tableEl = document.querySelector(`.${tableClass}`);

        this.tableRowElList = this.tableEl.querySelectorAll('.js-row');
        this.tableHeadCheckboxEl = this.tableEl.querySelector('.js-head-checkbox');
        this.tableCheckboxEl = this.tableEl.querySelectorAll('.js-checkbox');

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

    // getChckRowList チェックされているRowデータを取得する
    getCheckRowList() {
        // TODO

        let list = [];
        this.tableRowElList.forEach(function(row) {

            if (row.querySelector('.js-checkbox').checked) {
                console.log(row);

                let obj = {};

                // TODO data-keyとdata-valueを合わせてobjectを作成する
                // http://sarchitect.net/10929


                // arrayに変換する https://qiita.com/fivestar/items/074671e137497a8347ee
                console.log(row.children);
                // row.children.forEach(function(cell) {

                //     console.log("cell is ", cell);

                // });


            }

        });


    }


    // checkRow そのテーブルのチェックする
    checkRow(elem) {
        let checkBoxEl = elem.currentTarget.children[0].children[0];
        checkBoxEl.checked = !checkBoxEl.checked;

        if (checkBoxEl.checked) {
            elem.currentTarget.classList.add('table__row--selected')
        } else {
            elem.currentTarget.classList.remove('table__row--selected')
        }
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
            elem.classList.add('table__row--selected');
        });
    }

    // uncheckAll すべてのRowのcheckを外す
    uncheckAll() {
        this.tableRowElList.forEach(function(elem) {
            elem.children[0].children[0].checked = false;
            elem.classList.remove('table__row--selected')
        });
    }
}
