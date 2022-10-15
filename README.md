## Curl
```bash
curl -X GET http://localhost:8080/v1/user/85bf6aeb-459c-445a-be1e-0b67b8c100ef


curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X POST \
-d '{"document":{"login_id":"wonksing","first_name":"wonk","last_name":"sun","birth_date":"2022-01-01T00:00:00+09:00","gender":"M","nationality":"KOR"}}' \
http://localhost:8080/v1/user/register


curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X PUT \
-d '{"document":{"login_id":"wonksing","first_name":"wonk","last_name":"sun","birth_date":"20010101","gender":"F","nationality":"KOR"}}' \
http://localhost:8080/v1/user/17345530-53ea-4c27-8f0c-dcd96f8a4a9d

curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X DELETE \
http://localhost:8080/v1/user/17345530-53ea-4c27-8f0c-dcd96f8a4a9d

curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X POST \
-d '{"document":{"login_id":"wonksing","first_name":"wonk","last_name":"sun","birth_date":"20000101","gender":"M","nationality":"KOR","UserEmails":[{"email":"wonk@wonk.orgg"}],"UserSecrets":null}}' \
http://localhost:8080/v1/user/register


```