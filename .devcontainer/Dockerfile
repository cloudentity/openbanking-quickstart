FROM mcr.microsoft.com/devcontainers/base:ubuntu

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install ripgrep pgcli tig

RUN sudo apt update && sudo apt install -y curl ca-certificates gpg \
    && curl https://apt.fury.io/authzed/gpg.key | sudo apt-key add - \ 
    && echo "deb https://apt.fury.io/authzed/ * *" > /etc/apt/sources.list.d/fury.list \
    && sudo apt update && sudo apt install -y zed

RUN SNIPPET="export PROMPT_COMMAND='history -a' && export HISTFILE=/commandhistory/.bash_history" \
    && echo "$SNIPPET" >> "/root/.bashrc"
