FROM nvcr.io/nvidia/pytorch:22.04-py3
# FROM python:3.8

RUN ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

USER root

RUN apt-get update
RUN apt-get -y install locales && \
    localedef -f UTF-8 -i ja_JP ja_JP.UTF-8
ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
ENV TZ JST-9
ENV TERM xterm

RUN apt-get install -y vim \
    less \
    poppler-utils \
    ffmpeg \
    libsm6 \
    libxext6

RUN pip install --upgrade pip
RUN pip install --upgrade setuptools
RUN pip install flask flask_cors beautifulsoup4 lxml layoutparser torchvision "git+https://github.com/facebookresearch/detectron2.git@v0.5#egg=detectron2" \
    "layoutparser[layoutmodels]" reportlab pypdf transformers sentencepiece opencv-python==4.5.5.64 \
    mysql-connector-python

ENV NVIDIA_VISIBLE_DEVICES all
ENV NVIDIA_DRIVER_CAPABILITIES utility,compute
