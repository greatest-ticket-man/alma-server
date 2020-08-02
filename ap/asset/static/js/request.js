window.Alma = window.Alma || {};
(function(_Alma) {

    class Req {
        // async post(url, data, options = { reload: true }) {

        //     try {
        //         const response = await fetch(url, data);
        //         const json = await response.json();
        //         if (response.status == 200) {
        //             console.log(`通信に成功しました: ${JSON.stringify(json)}`);
        //             if (options.reload === true) {
        //                 window.Alma.toast.success('成功しました');
        //                 setTimeout(() => location.reload(), 2000);
        //             }
        //             return json;
        //         } 
        //         if (response.status === 500) {
        //             window.Alma.toast.error('失敗しました...');
        //             console.log(json);
        //         } else if (response.status !== 200) {
        //             window.Alma.toast.error('失敗しました...');
        //             console.log(json);
        //         }
        //     } catch (e) {
        //         console.log("error====");
        //         console.log(e);
        //     }
        // }

        async post(url, data, options = { reload: true }) {

            try {
                const response = await fetch(url, data);
                const json = await response.json();
                console.log("json is ", json);
                if (json.success) {
                    console.log(`通信に成功しました: ${JSON.stringify(json)}`);
                    if (options.reload === true) {
                        window.Alma.toast.success('成功しました');
                        setTimeout(() => location.reload(), 2000);
                    }
                    return json;
                }  else {
                    window.Alma.toast.error("失敗しました");
                    console.log(json);
                    return json;
                }

            } catch (e) {
                console.log(e);
            }
        }

        createPostData(data = {}, cache = 'no-cache') {
            return {
                method: 'POST', // *GET, POST, PUT, DELETE, etc.
                mode: 'cors', // no-cors, cors, *same-origin
                cache: cache, // *default, no-cache, reload, force-cache, only-if-cached
                credentials: 'same-origin', // include, same-origin, *omit
                headers: {
                    'Content-Type': 'application/json; charset=utf-8',
                    // 'Content-Type': 'application/x-www-form-urlencoded',
                },
                redirect: 'follow', // manual, *follow, error
                referrer: 'no-referrer', // no-referrer, *client
                body: JSON.stringify(data), // 本文のデータ型は 'Content-Type' ヘッダーと一致する必要があります
            };
        }
    }

    _Alma.req = new Req();
    
})(window.Alma);
