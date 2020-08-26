'use strict';
import miniToastr from './mini-toastr.js';

window.Alma = window.Alma || {};
(function(_Alma) {

    class Toast {

        // constructor
        constructor() {

            // const
            this.EMPTY_STRING = '';

        }

        init(aConfig) {
            const newConfig = {};
            Object.assign(newConfig, config);
            Object.assign(newConfig, aConfig);
            this.config = newConfig;

            const cssStr = this.makeCss(newConfig.style);

        }

        makeCss(obj) {
            const flat = this.flatten(obj);
            let str = JSON.stringify(flat, null, 2);
            str = str.replace(/"([^"]*)": {/g, '$1 {')
                    .replace(/"([^"]*)"/g, '$1')
                    .replace(/(\w*-?\w*): ([\w\d .#]*),?/g, '$1: $2;')
                    .replace(/},/g, '}\n')
                    .replace(/ &([.:])/g, '$1');
            
            return str.substr(1, str.lastIndexOf('}' -1));
        }

        flatten(obj, into, prefix) {
            into = into || {};
            prefix = prefix || EMPTY_STRING;

            for (const k in obj) {
                if (obj.hasOwnProperty(k)) {
                    const prop = obj[k];
                    if (prop && typeof prop === 'object' && !(prop instanceof Date || prop instanceof RegExp)) {
                        this.flatten(prop, into, prefix + k + ' ');
                    } else {
                        if (into[prefix] && typeof into[prefix] === 'object') {
                            into[prefix][k] = prop;
                        } else {
                            into[prefix] = {};
                            into[prefix][k] = prop;
                        }
                    }
                }
            }

            return into;
        }

        appendStyle(css) {
            let head = document.head || document.getElementsBy
        }



        

    }

    _Alma.toast = new Toast();

})(window.Alma);