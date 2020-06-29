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
