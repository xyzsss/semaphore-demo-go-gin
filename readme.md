```
go mod init github.com/YOUR_USERNAME/semaphore-demo-go-gin
go get -u github.com/gin-gonic/gin
go build -o app
./app

# get json data
curl -X GET -H "Accept: application/json" http://localhost:8080/
# get xml data
curl -X GET -H "Accept: application/xml" http://localhost:8080/article/view/1

# test application
go test -v


# push to github
$ git init
$ git remote add YOUR_REPOSITORY_URL
$ git add -A
$ git commit -m "initial commit"
$ git push origin main# semaphore-demo-go-gin
```
