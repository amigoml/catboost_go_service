FROM catboost_v1

ENV GOPATH="/app"
ENV PATH=$PATH:$GOPATH/bin

COPY ./ /app/go_app
RUN go mod init catboost_serving && go get -u github.com/gorilla/mux && go get -u github.com/rs/cors

RUN go mod download && go get github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon --build="go build ." -command="./catboost_serving"