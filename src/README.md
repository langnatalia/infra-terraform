#infra-terraform

This is a Terraform project for infrastructure as code.

## Requirements

* Terraform >= 0.14.0
* AWS CLI >= 1.16.0

## Usage

1. Clone the repository: `git clone https://github.com/your-username/infra-terraform.git`
2. Initialize the Terraform working directory: `terraform init`
3. Review and accept the plan before applying it: `terraform plan`
4. Apply the configuration: `terraform apply`
5. Destroy the infrastructure: `terraform destroy`

## Modules

This project uses the following Terraform modules:

* `aws-ec2-instance`: Creates a single EC2 instance
* `aws-security-group`: Creates a single security group
* `aws-vpc`: Creates a single VPC

## Variables

This project uses the following input variables:

* `aws_region`: The AWS region to deploy to
* `vpc_cidr`: The CIDR block for the VPC
* `ec2_instance_type`: The type of the EC2 instance
* `security_group_name`: The name of the security group

## Files

This project consists of the following files:

* `main.tf`: The main Terraform configuration file
* `variables.tf`: The Terraform input variables file
* `outputs.tf`: The Terraform output values file
* `modules/*`: The Terraform modules directory