curl -X POST http://localhost:8081/notes/create \
-H "Content-Type: application/json" \
-d '{
  "title": "My First Note",
  "content": "This is the content of my first note."
}'



curl -X GET http://localhost:8081/notes
