#! /bin/sh 

if [ -z "$1" ]
  then
    echo "No netbox server url supplied"
fi

SERVER_URL=$1
SCRIPTPATH="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

docker-compose -f $SCRIPTPATH/docker-compose.yml up -d

echo "### Waiting for Netbox to become available on ${SERVER_URL} \n\n"

attempt_counter=0
max_attempts=24
until $(curl --connect-timeout 1 --output /dev/null --silent --head --fail ${SERVER_URL}); do
    if [ ${attempt_counter} -eq ${max_attempts} ];then
      echo "Max attempts reached"
      # Debug logs
      docker-compose -f $SCRIPTPATH/docker-compose.yml logs
      exit 1
    fi

    printf '.'
    attempt_counter=$(($attempt_counter+1))
    sleep 5
done

docker-compose -f $SCRIPTPATH/docker-compose.yml logs
