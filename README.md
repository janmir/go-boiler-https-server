# :book: Go HTTP Server Resource
- [Intro: Go Lang Web Server](https://gowebexamples.com)
- [HTTPS in Go](https://www.kaihag.com/https-and-go/)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [Override Mime Type](https://www.lemoda.net/go/override-mime-type/index.html)
- [Static File](https://lets-go.alexedwards.net/sample/02.08-serving-static-files.html)

# TLS
- [Certificate Conversion](https://stackoverflow.com/questions/13732826/convert-pem-to-crt-and-key)
- [Trust Self Signed Cert](https://www.nullalo.com/en/chrome-how-to-install-self-signed-ssl-certificates)
- [More..](https://stackoverflow.com/questions/7580508/getting-chrome-to-accept-self-signed-localhost-certificate)

# Create Self Signed Cert
```
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem -subj /CN=Hostname -reqexts SAN -extensions SAN -config <(cat /etc/ssl/openssl.cnf <(printf '[SAN]\nsubjectAltName=DNS:hostname,IP:192.168.0.1')) -sha256
```