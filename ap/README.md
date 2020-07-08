# AP Server
基本的なhttp server  
requestの型はGRPCのものを実装してみたい  

DBはなれているMongoDBを使用する  

# 方針
構造的に、
httpとhtmlというものを作りたい

基本的には、domain層で完全にロジックをどうにかする

httpではjsonを返すようにする
htmlではviewも含めて返すようにしてしまう


## bootstrap5
ここのサンプルを見て、構築していく
https://getbootstrap.com/docs/4.5/examples/

# stripe
テストで使えるカード一覧
https://qiita.com/mimoe/items/8f5d9ce46b72b7fecff5  

クレジットカード番号	カードの種類
4111111111111111	Visa
4242424242424242	Visa
4012888888881881	Visa
5555555555554444	MasterCard
5105105105105100	MasterCard
378282246310005	American Express
371449635398431	American Express
30569309025904	Diner's Club
38520000023237	Diner's Club
3530111333300000	JCB
3566002020360505	JCB
