---
title: OpenSSL / TLS
icon: fa-lock
primary: "#721412"
lang: bash
---

## fa-certificate Certificate Generation (Self-Signed)

```bash
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
openssl req -x509 -newkey ec -pkeyopt ec_paramgen_curve:prime256v1 \
  -keyout key.pem -out cert.pem -days 365 -nodes \
  -subj "/CN=localhost"
openssl req -x509 -newkey ed25519 \
  -keyout key.pem -out cert.pem -days 365 -nodes \
  -subj "/CN=localhost"
openssl req -newkey rsa:2048 -keyout key.pem -out cert.pem \
  -days 365 -nodes -x509 \
  -subj "/C=US/ST=State/L=City/O=Org/CN=example.com" \
  -addext "subjectAltName=DNS:example.com,DNS:www.example.com"
```

## fa-file-signature CSR & CA Signing

```bash
openssl req -new -newkey rsa:4096 -keyout client.key -out client.csr \
  -subj "/CN=client.example.com"
openssl req -new -key client.key -out client.csr

openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key \
  -CAcreateserial -out client.crt -days 365 \
  -extfile <(printf "subjectAltName=DNS:client.example.com")

openssl ca -in client.csr -cert ca.crt -keyfile ca.key \
  -out client.crt -extensions v3_req
```

## fa-eye View Certificate Info

```bash
openssl x509 -in cert.pem -text -noout
openssl x509 -in cert.pem -subject -issuer -noout
openssl x509 -in cert.pem -dates -noout
openssl x509 -in cert.pem -ext subjectAltName -noout
openssl x509 -in cert.pem -fingerprint -sha256 -noout
openssl x509 -in cert.pem -purpose -noout
openssl req -in client.csr -text -noout
openssl rsa -in key.pem -check
openssl ec -in key.pem -text -noout
```

## fa-arrows-left-right TLS Protocol Versions

```bash
openssl s_client -connect example.com:443 -tls1_2
openssl s_client -connect example.com:443 -tls1_3
openssl s_client -connect example.com:443 -no_ssl3 -no_tls1 -no_tls1_1

openssl ciphers -v | grep -i tls
openssl ciphers -v TLSv1.3
openssl s_client -connect example.com:443 2>&1 | grep "Protocol"
```

## fa-shield Cipher Suites

```bash
openssl ciphers -v 'HIGH:!aNULL:!MD5'
openssl ciphers -v 'ECDHE+AESGCM'
openssl ciphers -v 'TLS_AES_256_GCM_SHA384:TLS_CHACHA20_POLY1305_SHA256'

openssl s_client -connect example.com:443 -cipher 'ECDHE-RSA-AES256-GCM-SHA384'
openssl s_client -connect example.com:443 -ciphersuites 'TLS_AES_256_GCM_SHA384'
openssl s_client -connect example.com:443 2>&1 | grep "Cipher"
```

## fa-key Key Types (RSA/ECDSA/Ed25519)

```bash
openssl genrsa -out rsa.key 4096
openssl rsa -in rsa.key -pubout -out rsa.pub

openssl ecparam -genkey -name prime256v1 -noout -out ec.key
openssl ecparam -genkey -name secp384r1 -noout -out ec384.key
openssl ec -in ec.key -pubout -out ec.pub

openssl genpkey -algorithm ED25519 -out ed.key
openssl pkey -in ed.key -pubout -out ed.pub

openssl genpkey -algorithm X25519 -out x25519.key
openssl pkey -in x25519.key -pubout -out x25519.pub
```

## fa-link Certificate Chain

```bash
cat intermediate.crt root.crt > chain.pem
cat cert.pem chain.pem > fullchain.pem

openssl verify -CAfile chain.pem cert.pem
openssl verify -CAfile root.crt -untrusted intermediate.crt cert.pem

openssl s_client -connect example.com:443 -showcerts
awk '/BEGIN/{f="cert"(++i)".pem"} /BEGIN/,/END/{print > f}' <<< "$(openssl s_client -showcerts -connect example.com:443 2>/dev/null)"
```

## fa-leaf Let's Encrypt (certbot)

```bash
certbot certonly --standalone -d example.com
certbot certonly --webroot -w /var/www/html -d example.com
certbot certonly --dns-cloudflare -d example.com
certbot certonly --manual -d '*.example.com' --preferred-challenges dns

certbot renew
certbot renew --dry-run
certbot certificates
certbot delete --cert-name example.com
ls /etc/letsencrypt/live/example.com/
```

## fa-handshake Mutual TLS

```bash
openssl s_server -accept 4433 -cert server.crt -key server.key -CAfile ca.crt -Verify 1
openssl s_client -connect localhost:4433 \
  -cert client.crt -key client.key -CAfile ca.crt

curl --cert client.crt --key client.key --cacert ca.crt \
  https://localhost:4433/api
```

## fa-ban CRL & OCSP

```bash
openssl ca -gencrl -out crl.pem -config openssl.cnf
openssl crl -in crl.pem -text -noout

openssl ocsp -issuer ca.crt -cert cert.pem \
  -url http://ocsp.example.com -resp_text

openssl x509 -in cert.pem -ocsp_uri -noout
openssl s_client -connect example.com:443 -status 2>&1 | grep "OCSP"
```

## fa-thumbtack Certificate Pinning

```bash
openssl x509 -in cert.pem -pubkey -noout | openssl pkey -pubin -outform der | openssl dgst -sha256 -binary | openssl enc -base64
echo "pin-sha256=\"$(openssl s_client -connect example.com:443 2>/dev/null | openssl x509 -pubkey -noout | openssl pkey -pubin -outform der | openssl dgst -sha256 -binary | openssl enc -base64)\";"
curl --pinnedpubkey sha256//BASE64 https://example.com
```

## fa-flask-vial Testing TLS (s_client)

```bash
openssl s_client -connect example.com:443
openssl s_client -connect example.com:443 -servername example.com
openssl s_client -connect example.com:443 -alpn h2,http/1.1
openssl s_client -connect example.com:443 -starttls smtp
openssl s_client -connect example.com:443 -msg -debug
openssl s_client -connect example.com:443 -tlsextdebug
echo -e "GET / HTTP/1.1\r\nHost: example.com\r\n\r\n" | openssl s_client -connect example.com:443 -quiet
```

## fa-shield-heart Hardening Config

```nginx
ssl_protocol TLSv1.2 TLSv1.3;
ssl_ciphers ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305;
ssl_prefer_server_ciphers on;
ssl_session_cache shared:SSL:10m;
ssl_session_timeout 1d;
ssl_session_tickets off;
ssl_stapling on;
ssl_stapling_verify on;
add_header Strict-Transport-Security "max-age=63072000" always;
```

## fa-bug Troubleshooting

```bash
openssl s_client -connect example.com:443 2>&1 | grep -E "verify|error|alert"
openssl s_client -connect example.com:443 -verify_return_error
openssl x509 -in cert.pem -checkend 86400 -noout
openssl verify -CAfile ca.crt cert.pem
openssl asn1parse -in cert.pem -i
openssl x509 -in cert.pem -text -noout | grep -A1 "Serial"
openssl s_client -connect example.com:443 -servername example.com 2>&1 | openssl x509 -noout -dates
```
