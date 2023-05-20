# Full Cycle 3.0 - EDA - Event Driven Architecture

This repository contains the microservices related to the Event Driven Architecture Module of Full Cycle 3.0 course.

# How it works

- Run docker compose

```bash
docker compose up -d --build
```
- The following containers will be started
```bash
[+] Running 9/9
 ✔ Network fc-eda_default            Created                                                             0.0s 
 ✔ Container zookeeper               Started                                                             1.9s 
 ✔ Container fc-eda-mysql-wallet-1   Started                                                             2.2s 
 ✔ Container fc-eda-mysql-balance-1  Started                                                             2.2s 
 ✔ Container fc-eda-goapp-balance-1  Started                                                             2.8s 
 ✔ Container fc-eda-goapp-wallet-1   Started                                                             2.8s 
 ✔ Container kafka                   Started                                                             2.4s 
 ✔ Container fc-eda-init-kafka-1     Started                                                             2.3s 
 ✔ Container control-center          Started     
```

- Ensure that Go Apps are running
```bash
docker compose logs goapp-wallet
fc-eda-goapp-wallet-1  | Server is running

docker compose logs goapp-balance
fc-eda-goapp-balance-1  | *** GoAppClient Started ***
fc-eda-goapp-balance-1  | Server is running
```

- Call endpoints

By using the api/client.http in the root directory, execute the transfer on the "wallet" microservice and check the result on the "balance" microservice.

- Step 1 expected result:
```bash
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 20 May 2023 00:38:52 GMT
Content-Length: 169
Connection: close

{
  "id": "aecba931-3667-44be-9699-7919d51a4374",
  "account_id_from": "4926e6cf-9d18-49d6-ae4e-32c9ddcb9b81",
  "account_id_to": "bf3a2451-f5cd-463a-845f-10cb9ee46d4f",
  "amount": 1
}

```

- Step 2 expected result:
```bash
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 20 May 2023 00:40:06 GMT
Content-Length: 76
Connection: close

{
  "account_id": "4926e6cf-9d18-49d6-ae4e-32c9ddcb9b81",
  "account_balance": 998
}
```

- Step 3 expected result:
```bash
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 20 May 2023 00:40:27 GMT
Content-Length: 74
Connection: close

{
  "account_id": "bf3a2451-f5cd-463a-845f-10cb9ee46d4f",
  "account_balance": 2
}
```