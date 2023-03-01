## Instructions

```
git clone https://github.com/hassan-qazi/interview-accountapi.git
cd interview-accountapi
docker compose up
```

### View Container Logs

```
docker logs interview-accountapi-form3test-1 
```

Please ensure container name is correct or use container id

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
