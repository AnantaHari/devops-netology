all: init plan apply deploy

init:
	cd ./terraform/demo && terraform init

plan:
	cd ./terraform/demo && terraform plan

apply:
	cd ./terraform/demo &&  terraform apply -auto-approve

destroy:
	cd ./terraform/demo &&  terraform destroy -auto-approve

clean:
	cd ./terraform/demo &&  rm -f terraform.tfplan
	cd ./terraform/demo &&  rm -f .terraform.lock.hcl
	cd ./terraform/demo &&  rm -f terraform.tfstate*
	cd ./terraform/demo &&  rm -fR .terraform/

deploy:
	cd ./ansible && source .env.news-app && ansible-playbook -i inventory/demo site.yml

reconfig:
	cd ./ansible && source .env.news-app && ansible-playbook -i inventory/demo site.yml -t config