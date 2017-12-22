# Use a map

Use the code in the "starting code" folder.

Remove mongodb from the code.

Instead of using mongodb, store all of the data in a map.

IMPORTANT:
Make sure you update your import statements to import packages from the correct location!

curl -X POST -H "Content-Type: application/json" -d '{"name":"Waqqas","gender":"M","age":29}' http://localhost:8080/user

{"id":"312ce5bf-1bcf-499f-a918-9d9c96d485bb","name":"Waqqas","gender":"M","age":29}

curl http://localhost:8080/user/312ce5bf-1bcf-499f-a918-9d9c96d485bb

{"id":"312ce5bf-1bcf-499f-a918-9d9c96d485bb","name":"Waqqas","gender":"M","age":29}

curl -X DELETE http://localhost:8080/user/312ce5bf-1bcf-499f-a918-9d9c96d485bb

Deleted user312ce5bf-1bcf-499f-a918-9d9c96d485bb