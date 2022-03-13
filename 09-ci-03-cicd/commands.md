ansible-playbook -i inventory/cicd site.yml
ansible-doc -t connection paramiko - проверка установки paramiko
ansible mitogen - это искать в инете чтобы изучить как ускорить работу
скачали сонар сканер
pushd ~/Downloads
cd sonar-scanner
export PATH=$PATH:$(pwd)/bin
popd - вернуться в сохраненную директорию

sonar-scanner \
  -Dsonar.projectKey=netology \
  -Dsonar.sources=. \
  -Dsonar.host.url=http://51.250.13.160:9000 \
  -Dsonar.login=4ed19a3dc154d039ca7eba1f478ec6564f39e936

cd apache-maven/bin
export PATH=$PATH:$(pwd)
mvn --version
popd
