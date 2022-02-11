Последовательность команд:  
ansible-playbook -i inventory/hosts.yml site.yml  
docker run --name centos7 -d pycontribs/centos:7 sleep 6000000000  
ansible-vault encrypt group_vars/prod/custom.yml  
ansible-vault decrypt group_vars/prod/custom.yml  
ansible-vault view group_vars/prod/custom.yml  
ansible-vault edit group_vars/prod/custom.yml  
ansible-vault rekey group_vars/prod/custom.yml  
ansible-vault encrypt_string  
ansible-inventory -i inventory/hosts.yml --graph  
ansible-inventory -i inventory/hosts.yml --list  
ansible-inventory -i inventory/hosts.yml --host localhost  
