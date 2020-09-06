'use strict';

class SideBar {
    constructor() {

        // getEl 
        this.sideBarHeaderEl = document.querySelector('.js-sidebar-header');
        this.sideBarHeaderPathEl = document.querySelector('.js-sidebar-header-path');
        this.sideBarMenuRowElList = document.querySelectorAll('.js-sidebar-menu-row');

        // bind this
        this.goHeaderPath = this.goHeaderPath.bind(this);
        this.goPath = this.goPath.bind(this);

        this.addEventListener();
    }

    addEventListener() {

        this.sideBarHeaderEl.addEventListener('click', this.goHeaderPath);

        const me = this;
        this.sideBarMenuRowElList.forEach(function(elem) {
            elem.addEventListener('click', me.goPath);
        });

    }

    goHeaderPath() {
        let path =  this.sideBarHeaderPathEl.innerText;
        window.Alma.location.href(path);
    }

    goPath(elem) {
        let sideBarMenuRow = elem.currentTarget;
        let path = sideBarMenuRow.querySelector('.js-sidebar-path').innerText;
        window.Alma.location.href(path);
    }
}

new SideBar();
