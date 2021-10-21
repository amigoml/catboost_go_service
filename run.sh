docker build -t catboost_v1 .
docker-compose -f docker-compose.yml up --no-start
docker-compose -f docker-compose.yml start