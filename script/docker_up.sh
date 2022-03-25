#! /bin/bash

reset=`tput sgr0`
red=`tput setaf 1`
green=`tput setaf 2`
yellow=`tput setaf 3`
blue=`tput setaf 4`

# Finding configuration file for chosen environment.
if [ $ENV ]; then
    if [[ -f "config/$ENV.yml" || -f "config/$ENV.yaml" ]]; then
        echo "${green}Application is running on ${blue}${ENV}${green} environment."
    else
        echo "${red}Failed to find configurations for ${yellow}${ENV}${red} environment.${reset}"
        exit;
    fi
else
    if [[ -f "config/local.yml" || -f "config/local.yaml" ]]; then
        export ENV="local"
        echo "${green}Application is running on ${blue}${ENV}${green} environment.${reset}"
    else
        echo "${red}Failed to find any configurations for any environment!"
        echo "Please consider making one under path ${yellow}config/${reset}"
        exit;
    fi
fi

# Finding docker-compose file related to chosen environment.
if [ -f "deploy/${ENV}/docker-compose.yaml" ]; then
    composeFile="$(pwd)/deploy/${ENV}/docker-compose.yaml"
elif [ -f "deploy/${ENV}/docker-compose.yml" ]; then
    composeFile="$(pwd)/deploy/${ENV}/docker-compose.yml"
else
    echo "${red}Failed to find any docker-compose file for your environment!"
    echo "Please consider making one under path ${yellow}deploy/${ENV}/${reset}"
    exit;
fi

DOCKER_BUILDKIT=1 docker-compose \
    --file $composeFile \
    --project-name bookstore_user \
    up -d --build
