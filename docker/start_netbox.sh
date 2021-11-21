#! /bin/sh 

set -e

if [ -z "$NETBOX_SERVER_URL" ]
  then
    echo "No netbox server url supplied"
    exit 2
fi

echo "Running tests for Netbox version: ${NETBOX_VERSION}"

SCRIPTPATH="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

DOCKER_BUILDKIT=1 docker-compose -f $SCRIPTPATH/docker-compose.yml up -d

echo "### Waiting for Netbox to become available on ${NETBOX_SERVER_URL} \n\n"

attempt_counter=0
max_attempts=48
until $(curl --connect-timeout 1 --output /dev/null --silent --head --fail ${NETBOX_SERVER_URL}); do
    if [ ${attempt_counter} -eq ${max_attempts} ];then
      echo "Max attempts reached"
      docker-compose -f $SCRIPTPATH/docker-compose.yml logs
      exit 1
    fi

    printf '.'
    attempt_counter=$(($attempt_counter+1))
    sleep 5
done

docker-compose -f $SCRIPTPATH/docker-compose.yml logs