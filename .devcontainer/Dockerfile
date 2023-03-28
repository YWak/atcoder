FROM ubuntu:22.04

# japanese environment
ENV LANG="ja_JP.UTF-8" \
    LANGUAGE="ja_JP:ja" \
    LC_ALL="ja_JP.UTF-8" \
    TZ="Asia/Tokyo"

ENV NODE_VERSION=node_18.x

RUN \
  : setup japanese environment \
  && apt-get update \
  && DEBIAN_FRONTEND=noninteractive \
    apt-get -y install \
      apt-transport-https \
      ca-certificates \
      curl \
      gnupg \
      language-pack-ja \
      lsb-release \
      tzdata \
  && update-locale LANG=ja_JP.UTF-8 \
  \
  && : setup node installation \
  && curl -fsSL https://deb.nodesource.com/gpgkey/nodesource.gpg.key | gpg --dearmor -o "/usr/share/keyrings/nodesource.gpg" \
  && gpg --no-default-keyring --keyring "/usr/share/keyrings/nodesource.gpg" --list-keys \
  && chmod a+r /usr/share/keyrings/nodesource.gpg \
  && echo \
      "deb [signed-by=/usr/share/keyrings/nodesource.gpg] https://deb.nodesource.com/$NODE_VERSION $(lsb_release -cs) main" \
      | tee /etc/apt/sources.list.d/nodesource.list \
  && echo \
      "deb-src [signed-by=/usr/share/keyrings/nodesource.gpg] https://deb.nodesource.com/$NODE_VERSION $(lsb_release -cs) main" \
      | tee -a /etc/apt/sources.list.d/nodesource.list \
  \
  && : setup docker installation \
  && curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o "/usr/share/keyrings/docker.gpg" \
  && echo \
      "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" \
      | tee /etc/apt/sources.list.d/docker.list > /dev/null \
  \
  && : install github cli \
  && curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | gpg --dearmor -o "/usr/share/keyrings/githubcli-archive-keyring.gpg" \
  && chmod go+r /usr/share/keyrings/githubcli-archive-keyring.gpg \
  && echo \
      "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" \
      | tee /etc/apt/sources.list.d/github-cli.list > /dev/null \
  \
  && : install from repo \
  && apt-get update \
  && apt-get -y install --no-install-recommends \
      direnv \
      dnsutils \
      docker-ce-cli \
      docker-compose-plugin \
      fonts-noto-cjk \
      g++ \
      gcc \
      gh \
      git \
      golang-go \
      gosu \
      iputils-ping \
      less \
      libglib2.0-0 \
      libnss3 \
      libnspr4 \
      libatk1.0-0 \
      libatk-bridge2.0-0 \
      libcups2 \
      libdrm2 \
      libdbus-1-3 \
      libatspi2.0-0 \
      libxcomposite1 \
      libxdamage1 \
      libxfixes3 \
      libxrandr2 \
      libgbm1 \
      libxkbcommon0 \
      libpango-1.0-0 \
      libcairo2 \
      libasound2 \
      make \
      nano \
      net-tools \
      nodejs \
      openssh-client \
      python3 \
      python3-pip \
      sudo \
      unzip \
      vim \
  && : install aws cli \
  && curl -sSL "https://awscli.amazonaws.com/awscli-exe-linux-$(uname -m).zip" -o "awscliv2.zip" \
  && unzip -q awscliv2.zip \
  && ./aws/install \
  && rm -rf awscliv2.zip ./aws \
  \
  && : enable git-remote-codecommit \
  &&  pip install git-remote-codecommit \
  \
  && : install nodejs modules \
  && npm install -g \
      aws-cdk \
      pnpm@7.13.0 \
      yarn \
  \
  && : add user \
  && useradd -m user -s /bin/bash \
  && mkdir -p /home/user/.ssh /home/user/.aws /home/user/work \
  && chmod 700 /home/user /home/user/.ssh \
  && chown -R user:user /home/user \
  && echo "user:user" | chpasswd \
  && echo "%user ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers.d/user \
  && chmod 0440 /etc/sudoers.d/user

# configure bash
WORKDIR /home/user
RUN \
  : install nvm && \
      gosu user sh -c \
        "export HOME=/home/user; curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash" && \
  : add git completion && \
    mkdir -p .local/bin && \
    mkdir -p .local/lib && \
    curl -sSL -o git-completion.bash \
      https://raw.githubusercontent.com/git/git/master/contrib/completion/git-completion.bash && \
    curl -sSL -o git-prompt.sh \
      https://raw.githubusercontent.com/git/git/master/contrib/completion/git-prompt.sh && \
    mv git-completion.bash git-prompt.sh .local/lib/ && \
  : enable git completion && \
    echo '# enable git completion'               >> .bashrc && \
    echo '. ~/.local/lib/git-completion.bash'    >> .bashrc && \
    echo '. ~/.local/lib/git-prompt.sh'          >> .bashrc && \
    echo 'GIT_PS1_SHOWDIRTYSTATE=1'              >> .bashrc && \
    echo 'GIT_PS1_SHOWUPSTREAM=1'                >> .bashrc && \
    echo 'GIT_PS1_SHOWUNTRACKEDFILES='           >> .bashrc && \
    echo 'GIT_PS1_SHOWSTASHSTATE=1'              >> .bashrc && \
    echo "export PS1='\[\033[1;32m\]\u\[\033[00m\]@gonow:\[\033[1;34m\]\w\[\033[1;31m\]\$(__git_ps1)\[\033[00m\] \n\$ '" >> .bashrc && \
    echo ''                                      >> .bashrc && \
  : enable direnv && \
    echo '# enable direnv'                       >> .bashrc && \
    echo 'eval "$(direnv hook bash)"'            >> .bashrc

USER user
WORKDIR /home/user/work
COPY ./entrypoint.sh /

ENTRYPOINT [ "/bin/bash", "/entrypoint.sh" ]
CMD [ "sleep", "infinity" ]
