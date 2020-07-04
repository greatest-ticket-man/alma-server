window.Alma = window.Alma || {};
(function(_Alma) {

    class Toast {

        constructor() {
            let div = document.createElement('div');
            div.style = 'position: absolute; top: 5vh; left: 50vw';

            this.div = div;
            document.body.appendChild(div);
        }


        // toastを作成する
        toast() {
            
            // 要素作成
            // let div = document.createElement('div');
            // div.style = 'position: absolute; top: 5vh; right: 50vw;';

            let toastDiv = document.createElement('div');
            toastDiv.classList.add('toast', 'bg-info', 'text-white');
            toastDiv.setAttribute('role', 'alert');
            toastDiv.setAttribute('aria-live', 'assertive');
            toastDiv.setAttribute('aria-atomic', 'true');

            let toastHeaderDiv = document.createElement('div');
            toastHeaderDiv.classList.add('toast-header', 'bg-info', 'text-white');
            
            let toastHeaderIcon = document.createElement('i');
            toastHeaderIcon.classList.add('fas', 'fa-info-circle', 'fa-lg', 'mr-2');

            let toastHeaderTitle = document.createElement('strong');
            toastHeaderTitle.className = 'mr-auto';
            toastHeaderTitle.innerText = 'Info';

            let toastHeaderDesc = document.createElement('small');
            toastHeaderDesc.innerText = '何分前';

            let toastButton = document.createElement('button');
            toastButton.classList.add('ml-2', 'my-1', 'close', 'text-white');
            toastButton.setAttribute('data-dismiss', 'toast');
            toastButton.setAttribute('aria-label', 'Close');

            let toastButtonSpan = document.createElement('span');
            toastButtonSpan.setAttribute('aria-hidden', 'true');
            toastButtonSpan.innerHTML = '×';


            let toastBody = document.createElement('div');
            toastBody.className = 'toast-body';
            toastBody.innerText = 'あれあれがそれそれ';


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
                delay: 2000,
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
