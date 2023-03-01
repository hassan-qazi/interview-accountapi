## Instructions

Open terminal 

```
git clone https://github.com/hassan-qazi/interview-accountapi.git
cd interview-accountapi
docker compose up
```

### View Container Logs

Open a new terminal

```
docker logs interview-accountapi-form3test-1 
```

Please ensure container name is correct or use container id

### Sample Logs

```
=== RUN   TestCreateAccount
    integration_test.go:33: Created account with ID: 81810562-458c-484b-bf7f-a4bcd8a77d1a, successfully
--- PASS: TestCreateAccount (0.00s)
=== RUN   TestCreateAccountNilData
    integration_test.go:45: Error: Response StatusCode: 400, ErrorMessage: invalid account data
--- PASS: TestCreateAccountNilData (0.00s)
=== RUN   TestCreateAccountMissingData
    integration_test.go:61: Error: Response StatusCode: 400, ErrorMessage: validation failure list:
        validation failure list:
        id in body is required
--- PASS: TestCreateAccountMissingData (0.00s)
=== RUN   TestCreateAccountDuplicateData
    integration_test.go:80: Error: Response StatusCode: 409, ErrorMessage: Account cannot be created as it violates a duplicate constraint
--- PASS: TestCreateAccountDuplicateData (0.00s)
=== RUN   TestFetchAccount
    integration_test.go:93: Created account with ID: 88037e45-f86d-487d-9859-a1393592daf4, successfully
    integration_test.go:105: Fetched account with ID: 88037e45-f86d-487d-9859-a1393592daf4, successfully
--- PASS: TestFetchAccount (0.00s)
=== RUN   TestFetchAccountEmptyID
    integration_test.go:119: Error: id is empty
--- PASS: TestFetchAccountEmptyID (0.00s)
=== RUN   TestFetchAccountUnknownID
    integration_test.go:133: Created account with ID: dccf4435-69d4-4a66-9383-f35ca287fd38, successfully
    integration_test.go:141: Error: Response StatusCode: 404, ErrorMessage: record 0999a97e-0f7a-463e-adfd-3e1bd2380e06 does not exist
--- PASS: TestFetchAccountUnknownID (0.00s)
=== RUN   TestDeleteAccount
    integration_test.go:155: Created account with ID: 2ec4b98d-3a74-461f-8d9b-c382f04d69f2, successfully
    integration_test.go:170: Deleted account with ID: 2ec4b98d-3a74-461f-8d9b-c382f04d69f2, successfully
--- PASS: TestDeleteAccount (0.00s)
=== RUN   TestDeleteAccountEmptyId
    integration_test.go:186: Error: id is empty
--- PASS: TestDeleteAccountEmptyId (0.00s)
=== RUN   TestDeleteAccountUnknownID
    integration_test.go:202: Error: Response StatusCode: 404, ErrorMessage: Not Found
--- PASS: TestDeleteAccountUnknownID (0.00s)
=== RUN   TestDeleteAccountUnknownVersion
    integration_test.go:216: Created account with ID: c355565c-aed6-4887-b685-d43451d9a682, successfully
    integration_test.go:226: Error: Response StatusCode: 409, ErrorMessage: invalid version
--- PASS: TestDeleteAccountUnknownVersion (0.00s)
PASS
```

#### Unit Tests

Open terminal

```
cd ./interview-accountapi/form3
go test -v
```
#### Integration Tests (outside Docker)

Ensure fake api, postgres and vault containers are running

```
cd ./interview-accountapi/form3test
go test -v
```


## Notes

Used the <a href="https://github.com/google/go-github" target="_blank">Github Client Library</a> common service approach. Also using their approach to build the request buffer for POST
<br/>
<br/>
form3 is a standalone module with basic unit tests
<br/>
<br/>
form3test uses form3 and performs integration tests against the actual form3 (fake) API

## About Me

<a href="https://www.linkedin.com/in/qazihassan/" target="_blank">Hassan Qazi (New to Go)</a>
