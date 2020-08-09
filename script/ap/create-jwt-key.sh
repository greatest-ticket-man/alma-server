#!/bin/sh

# https://qiita.com/wh1teB0x/items/e2133eeb94f57629b5e7

# https://stackoverflow.com/questions/55470311/encode-private-key-getting-error-asn1-structure-error-tags-dont-match

ssh-keygen -t rsa -f ./jwt.rsa -m pem
ssh-keygen -f jwt.rsa.pub -e -m pkcs8 > jwt.rsa.pub.pkcs8
