FROM ubuntu:16.04

MAINTAINER Mazedur Rahman <mazedur.rahman.liton@gmail.com>

RUN apt-get update \
	&& apt-get install -y ack-grep curl git iputils-ping man screen tree vim wget \
	&& apt-get clean

ARG DEF_GROUP
ARG DEF_GROUP_ID
ARG DEF_USER
ARG DEF_USER_ID

RUN groupadd --gid $DEF_GROUP_ID $DEF_GROUP \
	&& useradd --uid $DEF_USER_ID \
	--gid $DEF_GROUP_ID \
	--create-home \
	--home-dir /home/$DEF_USER \
        $DEF_USER

RUN mkdir -p /home/$DEF_USER/.vim/ \
	&& git clone https://github.com/sickill/vim-monokai.git /tmp/monokai \
	&& cp -r /tmp/monokai/* /home/$DEF_USER/.vim/ \
	&& rm -rf /tmp/monokai

USER $DEF_USER
