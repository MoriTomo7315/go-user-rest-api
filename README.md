# GO-USER-REST-API


## How to run

### prepare .env.local

Here is env file.
please set firebase_credential.json on the root directory of this project(the same layer with main.go)
```
GOOGLE_APPLICATION_CREDENTIALS=./firebase_credential.json
```

### firebase_credential.json
Please create this json file (it is ok by just copying from your firebase project console) to access firestore database.
```
{
    "type": "service_account",
    "project_id": "xxxxxxxxxxxx",
    "private_key_id": "xxxxxxxxxxx",
    "private_key": "xxxxxxxxxxxxx",
    "client_email": "xxxxxxxxx",
    "client_id": "xxxxxx",
    "auth_uri": "xxxxxxxx",
    "token_uri": "xxxxxxxx",
    "auth_provider_x509_cert_url": "xxxxxxxx",
    "client_x509_cert_url": "xxxxxxxxxx"
}
```

### Run Command
```
# for console

GO_ENV=local go run main.go


# for docker

docker build --build-arg _STAGE=local -t go-user-api-image .

docker run -d --rm -p 50001:50001 --name go-user-api-container go-user-api-image
```