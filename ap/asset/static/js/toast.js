window.Alma = window.Alma || {};
(function(_Alma) {

    class Toast {

        // constructor 先にtoastを入れるdivを生成する
        constructor() {
            let div = document.createElement('div');
            div.style = 'position: absolute; top: 5vh; left: 50vw';

            this.div = div;
            document.body.appendChild(div);
        }

        // info 
        info(body) {
            this.toast('bg-info', 'fa-info-circle', 'info', body, 2000);
        }

        // warning
        warning(body) {
            this.toast('bg-warning', 'fa-exclamation-triangle', 'warning', body, 3000);
        }

        // success
        success(body) {
            this.toast('bg-success', 'fa-check', 'success', body, 2000);
        }

        // error
        error(body) {
            this.toast('bg-danger', 'fa-exclamation-circle','error', body, 999999);
        }

        // toastを作成する
        toast(level, levelIcon, title, body, delay) {
            
            let toastDiv = document.createElement('div');
            toastDiv.classList.add('toast', level, 'text-white');
            toastDiv.setAttribute('role', 'alert');
            toastDiv.setAttribute('aria-live', 'assertive');
            toastDiv.setAttribute('aria-atomic', 'true');
            toastDiv.style = 'width: 80vw; transform: translate(-50%);'; // 中央に表示するためのやつ

            let toastHeaderDiv = document.createElement('div');
            toastHeaderDiv.classList.add('toast-header', level, 'text-white');
            
            let toastHeaderIcon = document.createElement('i');
            toastHeaderIcon.classList.add('fas', levelIcon, 'fa-lg', 'mr-2');

            let toastHeaderTitle = document.createElement('strong');
            toastHeaderTitle.className = 'mr-auto';
            toastHeaderTitle.innerText = title;

            let toastHeaderDesc = document.createElement('small');
            toastHeaderDesc.innerText = '';

            let toastButton = document.createElement('button');
            toastButton.classList.add('ml-2', 'my-1', 'close', 'text-white');
            toastButton.setAttribute('data-dismiss', 'toast');
            toastButton.setAttribute('aria-label', 'Close');

            let toastButtonSpan = document.createElement('span');
            toastButtonSpan.setAttribute('aria-hidden', 'true');
            toastButtonSpan.innerHTML = '×';


            let toastBody = document.createElement('div');
            toastBody.className = 'toast-body';
            toastBody.innerText = body;


            // 要素組み立て
            
            // button
            toastButton.appendChild(toastButtonSpan);

            // header
            toastHeaderDiv.appendChild(toastHeaderIcon);
            toastHeaderDiv.appendChild(toastHeaderTitle);
            toastHeaderDiv.appendChild(toastHeaderDesc);
            toastHeaderDiv.appendChild(toastButton);

            // div
            toastDiv.appendChild(toastHeaderDiv);
            toastDiv.appendChild(toastBody);


            // 初期化
            let option = {
                animation: true,
                autohide: true,
                delay: delay,
            };

            let toast = new bootstrap.Toast(toastDiv, option);

            // toast fieldに追加
            this.div.appendChild(toastDiv);

            // show
            toast.show();
        }
    }


    _Alma.toast = new Toast();


})(window.Alma);
