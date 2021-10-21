FROM ubuntu:latest

RUN apt update && apt install -y git
ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt install -y golang && apt install -y python

WORKDIR /app/go_app

RUN git clone https://github.com/catboost/catboost.git
RUN cd catboost/catboost/libs/model_interface && ../../../ya make -r .

ENV CATBOOST_DIR="/app/go_app/catboost/catboost/libs/model_interface"
ENV C_INCLUDE_PATH=$CATBOOST_DIR:$C_INCLUDE_PATH
ENV LIBRARY_PATH=$CATBOOST_DIR:$LIBRARY_PATH
ENV LD_LIBRARY_PATH=$CATBOOST_DIR:$LD_LIBRARY_PATH
