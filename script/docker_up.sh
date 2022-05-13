#! /bin/bash

reset=`tput sgr0`
red=`tput setaf 1`
green=`tput setaf 2`
yellow=`tput setaf 3`
blue=`tput setaf 4`

# Finding configuration file for chosen environment.
if [ $APP_MODE ]; then
    if [[ -f "config/$APP_MODE.yml" || -f "config/$APP_MODE.yaml" ]]; then
        echo "${green}Application is running on ${blue}${APP_MODE}${green} environment."
    else
        echo "${red}Failed to find configurations for ${yellow}${APP_MODE}${red} environment.${reset}"
        exit;
    fi
else
    if [[ -f "config/local.yml" || -f "config/local.yaml" ]]; then
        export APP_MODE="local"
        echo "${green}Application is running on ${blue}${APP_MODE}${green} environment.${reset}"
    else
        echo "${red}Failed to find any configurations for any environment!"
        echo "Please consider making one under path ${yellow}config/${reset}"
        exit;
    fi
fi

# Finding docker-compose file related to chosen environment.
if [ -f "deploy/${APP_MODE}/docker-compose.yaml" ]; then
    composeFile="$(pwd)/deploy/${APP_MODE}/docker-compose.yaml"
elif [ -f "deploy/${APP_MODE}/docker-compose.yml" ]; then
    composeFile="$(pwd)/deploy/${APP_MODE}/docker-compose.yml"
else
    echo "${red}Failed to find any docker-compose file for your environment!"
    echo "Please consider making one under path ${yellow}deploy/${APP_MODE}/${reset}"
    exit;
fi

DOCKER_BUILDKIT=1 docker-compose \
    --file $composeFile \
    --project-name restaurant_user \
    up -d --build
