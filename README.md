# lambda-http

![technology Go](https://img.shields.io/badge/technology-go-blue.svg)
![1.19](https://img.shields.io/badge/go-1.19--mini-green.svg)
![AWS](https://img.shields.io/badge/infra-aws-orange)
![Version 0.0.1](https://img.shields.io/badge/version-0.0.1-green)

## Configure LocalStack

In `LocalStack` CLI (docker) run the following scripts:

- Private url: 
    ```bash
    aws configure set cli_follow_urlparam false
    awslocal ssm put-parameter --name "PRIVATE_API_URL" --value "CHANGE_IT" --type "SecureString"
    ```