# AWS CDK Workshop Setup

This guide will walk you through the prerequisites and setup steps for the AWS CDK (Cloud Development Kit) workshop.

## Prerequisites

Make sure to install the following prerequisites:

- [Node.js](https://nodejs.org/) (npm version 10.2.3, Node.js version 20.10.0)
- [Python](https://www.python.org/) (version 3.8)
- [AWS CLI](https://aws.amazon.com/cli/)
- AWS user account

Check the documentation for AWS CDK Python package: [aws_cdk.aws_s3](https://docs.aws.amazon.com/cdk/api/v2/python/aws_cdk.aws_s3/README.html)

## Workshop Setup

1. Create a new directory for the CDK project:

    ```bash
    mkdir cdk
    cd cdk
    ```

2. Initialize a sample CDK app in Python:

    ```bash
    cdk init sample-app --language=python
    ```

    This command will create all the configuration files.

3. If the virtual environment (venv) is not activated automatically, activate it manually:

    For Venv:

    ```bash
    sudo apt install python3-virtualenv
    ```

    Activate venv:

    ```bash
    source .venv/bin/activate
    ```

    Deactivate venv:

    ```bash
    deactivate
    ```

4. Install dependencies:

    ```bash
    pip install -r requirements.txt
    ```

5. Bootstrap the CDK:

    ```bash
    cdk bootstrap
    ```

6. Check for any errors in basic compilation:

    ```bash
    cdk ls
    ```

7. Create the stack:

    ```bash
    cdk synth
    ```

8. Deploy the stack (similar to Terraform apply):

    ```bash
    cdk deploy
    ```

That's it! You have successfully set up your AWS CDK project.

For more details, refer to the [CDK documentation](https://docs.aws.amazon.com/cdk/latest/guide/home.html).
