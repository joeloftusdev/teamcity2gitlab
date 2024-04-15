#!/bin/bash

cd gitlab

docker-compose up -d

echo "Waiting for GitLab to start..."
until [ "$(curl -o /dev/null -s -w "%{http_code}\n" http://localhost:8080)" == "302" ]; do
  printf '.'
  sleep 5
done

docker exec gitlab-ce cat /etc/gitlab/initial_root_password | grep -oP 'Password: \K.*' > gitlab-root-password.txt
echo ""

GITLAB_ROOT_PASSWORD=$(cat gitlab-root-password.txt)

echo "Please Login to Gitlab at http://localhost:8080/. Username is root and your password is $GITLAB_ROOT_PASSWORD"
echo "Register your runner and enter the GitLab registration token:"
read GITLAB_REGISTRATION_TOKEN

echo "Registering a runner..."
docker exec gitlab-lab-runner gitlab-runner register --non-interactive --url "http://gitlab-ce" --registration-token "$GITLAB_REGISTRATION_TOKEN" --executor "docker" --docker-image "alpine:latest" --description "docker-runner"

echo "Updating config.toml..."
docker exec gitlab-lab-runner bash -c "echo '    network_mode = \"gitlab-network\"' >> /etc/gitlab-runner/config.toml"

echo "Gitlab build completed."

cd ../teamcity

docker-compose up -d

echo "Waiting for TeamCity to start..."
until [ "$(curl -o /dev/null -s -w "%{http_code}\n" http://localhost:8111/mnt)" == "200" ]; do
  printf '.'
  sleep 5
done

echo "Please follow the TeamCity start up wizard at http://localhost:8111/"

echo "Accept the license agreement then press any key to continue"

read -n 1 -s -r

docker exec teamcity-server cat  /opt/teamcity/logs/teamcity-server.log | grep -oP 'Super user authentication token: \K[0-9]+' | head -n 1 > teamcity-auth-token.txt

TEAMCITY_AUTH_TOKEN=$(cat teamcity-auth-token.txt)

echo "TeamCity build completed. Please Login to TeamCity at http://localhost:8111/. Username is blank and your password is $TEAMCITY_AUTH_TOKEN"

