BASEDIR=$(PWD)
APPNAME=middlewarekam
PROJECTID=kamestery

gen-rsa:
	# Generating params file
	openssl ecparam -name brainpoolP512t1 -out $(BASEDIR)/resources/ecparams.pem
	# Generate a private key from params file
	openssl ecparam -in $(BASEDIR)/resources/ecparams.pem -genkey -noout -out $(BASEDIR)/resources/ecprivkey.pem
	# Generate a public key from private key
	openssl ec -in $(BASEDIR)/resources/ecprivkey.pem -pubout -out $(BASEDIR)/resources/ecpubkey.pem

container:
	gcloud builds submit --tag gcr.io/$(PROJECTID)/$(APPNAME)

deploy: container
	gcloud beta run deploy --image gcr.io/$(PROJECTID)/$(APPNAME) --platform managed

