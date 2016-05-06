#!/usr/bin/env bash

echo "Stopping running application"
ssh $DEPLOY_USER@$DEPLOY_HOST 'docker stop stalks'
ssh $DEPLOY_USER@$DEPLOY_HOST 'docker rm stalks'

echo "Pulling latest version"
ssh $DEPLOY_USER@$DEPLOY_HOST 'docker pull jessemillar/stalks:latest'

echo "Starting the new version"
ssh $DEPLOY_USER@$DEPLOY_HOST 'docker run -d --restart=always --name stalks -p 8000:8000 jessemillar/stalks:latest'

echo "Success!"

exit 0
