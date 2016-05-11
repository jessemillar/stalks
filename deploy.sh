#!/usr/bin/env bash

echo "Stopping running application"
ssh $DEPLOY_USER@$DEPLOY_HOST 'docker stop stalks'
ssh $DEPLOY_USER@$DEPLOY_HOST 'docker rm stalks'

echo "Pulling latest version"
ssh $DEPLOY_USER@$DEPLOY_HOST 'docker pull jessemillar/stalks:latest'

echo "Starting the new version"
ssh $DEPLOY_USER@$DEPLOY_HOST 'docker run -d -e DATABASE_USERNAME="'$DATABASE_USERNAME'" -e DATABASE_PASSWORD="'$DATABASE_PASSWORD'" -e DATABASE_HOST="'$DATABASE_HOST'" -e DATABASE_PORT="'$DATABASE_PORT'" -e DATABASE_NAME="'$DATABASE_NAME'" --restart=always --name stalks -p 15000:8000 jessemillar/stalks:latest'

echo "Success!"

exit 0
