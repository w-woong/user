# Curl

```bash
curl -X GET http://localhost:8080/v1/user/85bf6aeb-459c-445a-be1e-0b67b8c100ef


curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X POST \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
-d '{"status":200,"document":{"login_id":"wonkwonkwonk","login_type":"id","password":{"value":"asdfasdfasdf"},"personal":{"first_name":"wonk","last_name":"sun","birth_year":2002,"birth_month":1,"birth_day":2,"gender":"M","nationality":"KOR"},"emails":[{"email":"wonk@wonk.orgg","priority":0}]}}' \
http://localhost:8080/v1/user


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