1. Ознакомьтесь с графическим интерфейсом VirtualBox, посмотрите как выглядит виртуальная машина, которую создал для вас Vagrant, какие аппаратные ресурсы ей выделены. Какие ресурсы выделены по-умолчанию?
```
Ей выделено 1Гб памяти, 2 ядра процессора, 64Гб HDD, video memory 4Гб
```

2. Ознакомьтесь с возможностями конфигурации VirtualBox через Vagrantfile: документация. Как добавить оперативной памяти или ресурсов процессора виртуальной машине?
```
config.vm.provider "virtualbox" do |v|
  v.memory = 2048
  v.cpus = 4
end
```


3. Ознакомиться с разделами man bash, почитать о настройках самого bash:
какой переменной можно задать длину журнала history, и на какой строчке manual это описывается?
```
HISTSIZE
Строчка 862
```
что делает директива ignoreboth в bash?
```
A  colon-separated  list of values controlling how commands are saved on the history list.  If the list of values includes ignorespace, lines which begin with a space character are not saved in  the  history list.  A value of ignoredups causes lines matching the previous history entry to not be saved.  A value of ignoreboth is shorthand for ignorespace and ignoredups.
Т.е. ignoreboth объединяет действие команд ignorespace и ignoredups (в историю не запишутся строки, которые начинаются с пробела и если в истории уже есть такая запись, то такая же записана не будет.
```

4. В каких сценариях использования применимы скобки {} и на какой строчке man bash это описано?
```
{ list; }
              list is simply executed in the current shell environment.  list must be terminated with  a  newline  or
              semicolon.  This is known as a group command.  The return status is the exit status of list.  Note that
              unlike the metacharacters ( and ), { and } are reserved words and must occur where a reserved  word  is
              permitted  to be recognized.  Since they do not cause a word break, they must be separated from list by
              whitespace or another shell metacharacter.
 Строка 257
{ список; }
Список просто выполняется в среде текущего командного интерпретатора.
Список должен завершаться переводом строки или точкой с запятой.
Эту команду называют командой группировки. Статусом возврата
является статус выхода списка.
```

5. Основываясь на предыдущем вопросе, как создать однократным вызовом touch 100000 файлов? А получилось ли создать 300000?
```
Для создания файлов - touch {000001..100000}.txt

Создать 300000 файлов не получилось - 
touch {000001..300000}.txt
-bash: /usr/bin/rm: Argument list too long
```

6. В man bash поищите по /\[\[. Что делает конструкция [[ -d /tmp ]]
```
Проверяет наличие директории tmp.
Конструкция [[ ... ]] более предпочтительна, нежели [ ... ], поскольку поможет избежать некоторых логических ошибок.
```

7. Основываясь на знаниях о просмотре текущих (например, PATH) и установке новых переменных; командах, которые мы рассматривали, добейтесь в выводе type -a bash в виртуальной машине наличия первым пунктом в списке:
bash is /tmp/new_path_directory/bash
bash is /usr/local/bin/bash
bash is /bin/bash
(прочие строки могут отличаться содержимым и порядком)

```
mkdir /tmp/new_path_directory/
cp /bin/bash /tmp/new_path_directory/
PATH=/tmp/new_path_directory/:$PATH
```

Чем отличается планирование команд с помощью batch и at?
 ```
  at      executes commands at a specified time. (выполняет команды в указанное время)
 batch   executes  commands  when  system load levels permit; in other words, when the load average drops below 1.5, or the
               value specified in the invocation of atd. (выполняет команды когда определенная загрузка системы, например ниже 1.5)
```
