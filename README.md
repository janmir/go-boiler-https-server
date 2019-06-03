# :book: Go HTTP Server Resource
- [Intro: Go Lang Web Server](https://gowebexamples.com)
- [HTTPS in Go](https://www.kaihag.com/https-and-go/)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [Override Mime Type](https://www.lemoda.net/go/override-mime-type/index.html)
- [Static File](https://lets-go.alexedwards.net/sample/02.08-serving-static-files.html)

# TLS
- [Certificate Conversion](https://stackoverflow.com/questions/13732826/convert-pem-to-crt-and-key)
- [Trust Self Signed Cert](https://www.nullalo.com/en/chrome-how-to-install-self-signed-ssl-certificates)
```
    1. Visit your website. You will get the above warning, click on “ADVANCED” and then on “Proceed to <domain name> (unsafe)”.
    2. Open Chrome Developer tools by hitting F12 key and go to “Security” tab.
    3. Click on “View certificate“, go to “Details” tab and click on “Copy to File…”.
    4. Hit “Next” and select “Cryptographic Message Syntax Standard – PKCS #7 Certificates (.P7B)“, then “Next” again.
    5. Click “Browse…” and save certificate to a location on your PC (ie: Desktop) by clicking “Next”. Then Finish.
    6. Open Chrome Settings, scroll down and click on “Show advanced settings…“, then under “HTTPS/SSL” click on “Manage certificates…“.
    7. Click “Import…” and then “Next”.
    8. Click “Browse…” and, since it’s not default, in the open dialog make sure you have “PKCS #7 Certificates (*.spc,*.p7b)” selected as filter, otherwise you won’t see your previously saved certificate.
    9. Click “Next” and then “Browse…” and make sure you select “Trusted Root Certification Authorities” as destination for your certificate, then hit “OK”, then “Next” and finally “Finish”.
    10. Now simply confirm certificate installation by clicking “Yes”.
    11. Now close all open dialogs, close Chrome (important), restart it and try to visit your website… you will see it’s now “Secure” (green lock) and without warnings!
```
- [More..](https://stackoverflow.com/questions/7580508/getting-chrome-to-accept-self-signed-localhost-certificate)

# Create Self Signed Cert
```
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem -subj /CN=Hostname -reqexts SAN -extensions SAN -config <(cat /etc/ssl/openssl.cnf <(printf '[SAN]\nsubjectAltName=DNS:hostname,IP:192.168.0.1')) -sha256
```