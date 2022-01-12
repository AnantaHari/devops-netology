SHELL := /bin/bash
all: init plan apply

init:
	terraform init

plan:
	terraform plan

apply:
	terraform apply -auto-approve

destroy:
	terraform destroy -auto-approve

clean:
	rm -f terraform.tfplan
	rm -f .terraform.lock.hcl
	rm -f terraform.tfstate*
	rm -fR .terraform/
