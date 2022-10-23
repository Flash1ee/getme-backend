OUT = ./certs
.PHONY: gen-cert
gen-cert:
	openssl req -x509 -out $(OUT)/getme.crt -keyout $(OUT)/getme.key -newkey rsa:2048 -nodes -sha256 -subj '/CN=localhost' -extensions EXT -config $(OUT)/cfg/cert.cnf
