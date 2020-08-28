'use strict';

window.Alma = window.Alma || {};
(function(_Alma) {

    class Toast {

        // constructor
        constructor() {

            // bind
            this.fadeOut = this.fadeOut.bind(this);

            // const
            const ERROR = 'error';
            const WARN = 'warn';
            const SUCCESS = 'success';
            const INFO = 'info';


            this.EMPTY_STRING = '';
            this.LIB_NAME = 'mini-toastr';
            this.ERROR = ERROR;
            this.WARN = WARN;
            this.SUCCESS = SUCCESS;
            this.INFO = INFO;
            this.CONTAINER_CLASS = this.LIB_NAME;
            this.NOTIFICATION_CLASS = `${this.LIB_NAME}__notification`;
            this.TITLE_CLASS = `${this.LIB_NAME}-notification__title`;
            this.ICON_CLASS = `${this.LIB_NAME}-notification__icon`;
            this.MESSAGE_CLASS = `${this.LIB_NAME}-notification__message`;
            this.ERROR_CLASS = `-${this.ERROR}`;
            this.WARN_CLASS = `-${this.WARN}`;
            this.SUCCESS_CLASS = `-${this.SUCCESS}`;
            this.INFO_CLASS = `-${this.INFO}`;
            this.DEFAULT_TIMEOUT = 3000;

            this.isInitialised = false;

            // config
            this.config = {
                types: { ERROR, WARN, SUCCESS, INFO },
                animation: this.fadeOut,
                timeout: this.DEFAULT_TIMEOUT,
                icons: {},
                appendTarget: document.body,
                node: this.makeNode(),
                allowHtml: false,
                style: {
                    [`.${this.CONTAINER_CLASS}`]: {
                        position: 'fixed',
                        'z-index': 99999,
                        right: '12px',
                        top: '12px'
                    },
                    [`.${this.NOTIFICATION_CLASS}`]: {
                        cursor: 'pointer',
                        padding: '12px 18px',
                        margin: '0 0 6px 0',
                        'background-color': '#000',
                        opacity: 0.8,
                        color: '#fff',
                        'border-radius': '3px',
                        'box-shadow': '#3c3b3b 0 0 12px',
                        width: '300px',
                        [`&.${this.ERROR_CLASS}`]: {
                            'background-color': '#D5122B'
                        },
                        [`&.${this.WARN_CLASS}`]: {
                            'background-color': '#F5AA1E'
                        },
                        [`&.${this.SUCCESS_CLASS}`]: {
                            'background-color': '#7AC13E'
                        },
                        [`&.${this.INFO_CLASS}`]: {
                            'background-color': '#4196E1'
                        },
                        '&:hover': {
                            opacity: 1,
                            'box-shadow': '#000 0 0 12px'
                        }
                    },
                    [`.${this.TITLE_CLASS}`]: {
                        'font-weight': '500'
                    },
                    [`.${this.MESSAGE_CLASS}`]: {
                        display: 'inline-block',
                        'vertical-align': 'middle',
                        width: '240px',
                        padding: '0 12px'
                    }
                }
            };

            // init
            this.init();

        }



        init(aConfig) {
            const newConfig = {};
            Object.assign(newConfig, this.config);
            Object.assign(newConfig, aConfig);
            this.config = newConfig;

            const cssStr = this.makeCss(newConfig.style);
            this.appendStyle(cssStr);

            newConfig.node.id = this.CONTAINER_CLASS;
            newConfig.node.className = this.CONTAINER_CLASS;
            newConfig.appendTarget.appendChild(newConfig.node);

            Object.keys(newConfig.types).forEach(v => {
                this[newConfig.types[v]] = function(message, title, timeout, cb, config) {
                    this.showMessage(message, title, newConfig.types[v], timeout, cb, config);
                    return this;
                }.bind(this);
            });

            this.isInitialised = true;

            return this;
        }

        setIcon(type, nodeType = 'i', attrs = []) {
            attrs.class = attrs.class ? attrs.class + ' ' + this.ICON_CLASS : this.ICON_CLASS;
            this.config.icons[type] = {nodeType, attrs};
        }

        showMessage(message, title, type, timeout, cb, overrideConf) {

            const config = {};
            Object.assign(config, this.config);
            Object.assign(config, overrideConf);

            const notificationElem = this.makeNode();
            notificationElem.className = `${this.NOTIFICATION_CLASS} ${this.getTypeClass(type)}`;

            notificationElem.onclick = function() {
                config.animation(notificationElem, null);
            };

            if (title) {
                this.addElem(notificationElem, title, this.TITLE_CLASS, config);
            }
            if (config.icons[type]) {
                this.createIcon(notificationElem, type, config);
            }
            if (message) {
                this.addElem(notificationElem, message, this.MESSAGE_CLASS, config);
            }

            config.node.insertBefore(notificationElem, config.node.firstChild);
            setTimeout(() => config.animation(notificationElem, cb), timeout || config.timeout);

            if (cb) {
                cb();
            }

            return this;
        }

        makeCss(obj) {
            const flat = this.flatten(obj);
            let str = JSON.stringify(flat, null, 2);
            str = str.replace(/"([^"]*)": {/g, '$1 {')
                .replace(/"([^"]*)"/g, '$1')
                .replace(/(\w*-?\w*): ([\w\d .#]*),?/g, '$1: $2;')
                .replace(/},/g, '}\n')
                .replace(/ &([.:])/g, '$1');

            str = str.substr(1, str.lastIndexOf('}') - 1);

            return str;
        }

        flatten(obj, into, prefix) {
            into = into || {};
            prefix = prefix || this.EMPTY_STRING;

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

        makeNode(type = 'div') {
            return document.createElement(type);
        }

        createIcon(node, type, config) {
            const iconNode = this.makeNode(config.icons[type].nodeType);
            const attrs = config.icons[type].attrs;

            for (const k in attrs) {
                if (attrs.hasOwnProperty(k)) {
                    iconNode.setAttribute(k, attrs[k]);
                }
            }

            node.appendChild(iconNode);
        }

        addElem(node, text, className, config) {
            const elem = this.makeNode();
            elem.className = className;
            if (config.allowHtml) {
                elem.innerHTML = text;
            } else {
                elem.appendChild(document.createTextNode(text));
            }
            node.appendChild(elem);
        }

        getTypeClass(type) {
            if (type === this.SUCCESS) return this.SUCCESS_CLASS;
            if (type === this.WARN) return this.WARN_CLASS;
            if (type === this.ERROR) return this.ERROR_CLASS;
            if (type === this.INFO) return this.INFO_CLASS;

            return this.EMPTY_STRING;
        }

        appendStyle(css) {
            let head = document.head || document.getElementsByTagName('head')[0];
            let styleElem = this.makeNode('style');
            styleElem.id = `${this.LIB_NAME}-style`;
            styleElem.type = 'text/css';

            if (styleElem.styleSheet) {
                styleElem.styleSheet.cssText = css;
            } else {
                styleElem.appendChild(document.createTextNode(css));
            }

            head.appendChild(styleElem);
        }

        fadeOut(element, cb) {
            if (element.style.opacity && element.style.opacity > 0.05) {
                element.style.opacity = element.style.opacity - 0.05;
            } else if (element.style.opacity && element.style.opacity <= 0.1) {
                if (element.parenetNode) {
                    element.parentNode.removeChild(element);
                    if (cb) {
                        cb();
                    }
                }
            } else {
                element.style.opacity = 0.9;
            }

            setTimeout(() => this.fadeOut.apply(this, [element, cb]), 1000 / 30);
        }
    }

    _Alma.toast = new Toast();

})(window.Alma);
