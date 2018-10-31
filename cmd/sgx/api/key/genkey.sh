private_key_path=RAS-rsa-key.pem
public_key_path=RAS-rsa-key-pub.pem

if [ ! -e ${private_key_path} ]; then
    openssl genrsa > RAS-rsa-key.pem
fi

if [ ! -e ${public_key_path} ]; then
    openssl rsa -in RAS-rsa-key.pem -RSAPublicKey_out > RAS-rsa-key-pub.pem
fi