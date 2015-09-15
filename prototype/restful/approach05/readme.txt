1. 加入内存数据库
2. 支持POST JSON

访问：http://localhost:8080/todos

使用curl　POST后，再访问刚才的URL，会发现内容越来越多
curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos




还是来自这里：
http://thenewstack.io/make-a-restful-json-api-go/
