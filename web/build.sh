if ! command -v nvm &> /dev/null
then
    echo "installing nvm"
    curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
    export NVM_DIR="$HOME/.nvm"
    [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"
    [ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"
else
    echo "nvm already installed"
fi
nvm install 20.11.0 && nvm use 20.11.0

npm install && npm run build
docker build -t cesslab/watchdog-web:"$1" .
docker tag cesslab/watchdog-web:"$1" cesslab/watchdog-web:latest