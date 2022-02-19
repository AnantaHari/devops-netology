ansible-galaxy role init nginx_role - создать новую роль
ansible-galaxy install -r requirements.yml -p roles
ansible-playbook -i inventory/hosts.yml site.yml
ansible-playbook -i inventory/test site.yml

Когда готова роль в ее каталоге выполняем:
git init
git remote add origin путь до репозитория

https://docs.ansible.com/ansible/latest/reference_appendices/special_variables.html