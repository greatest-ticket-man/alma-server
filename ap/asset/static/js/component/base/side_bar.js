'use strict';

console.log("hoge");

class SideBar {
    constructor() {

        // getEl 
        this.sideBarHeader = document.querySelector('.js-sidebar-header');
        this.sideBarHeaderPath = document.querySelector('.js-sidebar-header-path');


        this.goHeaderpath = this.goHeaderpath.bind(this);

        this.addEventListener();
    }

    addEventListener() {

        this.sideBarHeader.addEventListener('click', this.goHeaderpath);
    }

    goHeaderpath() {
        let path =  this.sideBarHeaderPath.innerText;
        window.Alma.location.href(path);
    }


}

new SideBar();
