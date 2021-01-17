set GOOS=linux
set GOARCH=amd64
go build -o gprc_example-linux64
# docker build --no-cache -t gprc_example-linux64:dev .
# docker build -t gprc_example-linux64:dev .
docker run -p 5300:5300 --name gprc_example-linux64 gprc_example-linux64:dev