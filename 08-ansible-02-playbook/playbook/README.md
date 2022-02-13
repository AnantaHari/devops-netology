--  
Команды с урока  
ansible-playbook pup1.yml --skip-tags always  
ansible-playbook pup1.yml --skip-tags always --tags never  
ansible-playbook pup1.yml --skip-tags always --tags "never, tagged"
ansible-playbook pup1.yml --tags untagged
ansible-playbook pup1.yml --tags pre
ansible-playbook pup1.yml --skip-tags pre2
ansible-playbook pup1.yml --tags another
ansible-playbook pup1.yml --list-tags
ansible-playbook pup1.yml --list-tags --tags "all, never"
ansible-playbook pup1.yml --list-task
ansible-playbook pup1.yml --list-task --skip-tags pre2
ansible-playbook pup1.yml --step
Когда включено:   debugger: always
    пишем [localhost] TASK: Gathering Facts (debug)> p task.args для вывода значений переменных
    [localhost] TASK: Gathering Facts (debug)> с - продолжить
    [localhost] TASK: Gathering Facts (debug)> task.args['msg'] = 'One' - переопределение переменной
    [localhost] TASK: Gathering Facts (debug)> r - перезапустить таск с новым значением переменной
    [localhost] TASK: Gathering Facts (debug)> q - закончить режим debug

ansible-playbook -i inventory/prod.yml site.yml --check
ansible-playbook -i inventory/prod.yml site.yml --diff