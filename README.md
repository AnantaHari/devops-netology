1. Найдите полный хеш и комментарий коммита, хеш которого начинается на aefea.

$ git show aefea
commit aefead2207ef7e2aa5dc81a34aedf0cad4c32545
Комментарий: Update CHANGELOG.md


2. Какому тегу соответствует коммит 85024d3?

$ git log 85024d3
(tag: v0.12.23)


3. Сколько родителей у коммита b8d720? Напишите их хеши.

$ git show b8d720
commit b8d720f8340221f2146e4e4870bf2ee0bc48f2d5
Merge: 56cd7859e 9ea88f22f

Два родителя
$ git show b8d720^
commit 56cd7859e05c36c06b56d013b55a252d0bb7e158
$ git show b8d720^2
commit 9ea88f22fc6269854151c571162c5bcf958bee2b


4. Перечислите хеши и комментарии всех коммитов которые были сделаны между тегами v0.12.23 и v0.12.24.

$ git log v0.12.23..v0.12.24 --oneline
33ff1c03b (tag: v0.12.24) v0.12.24
b14b74c49 [Website] vmc provider links
3f235065b Update CHANGELOG.md
6ae64e247 registry: Fix panic when server is unreachable
5c619ca1b website: Remove links to the getting started guide's old location
06275647e Update CHANGELOG.md
d5f9411f5 command: Fix bug when using terraform login on Windows
4b6d06cc5 Update CHANGELOG.md
dd01a3507 Update CHANGELOG.md
225466bc3 Cleanup after v0.12.23 release


5. Найдите коммит в котором была создана функция func providerSource, ее определение в коде выглядит так func providerSource(...) (вместо троеточего перечислены аргументы).

$ git log -S'func providerSource(' --oneline
8c928e835 main: Consult local directories as potential mirrors of providers


6. Найдите все коммиты в которых была изменена функция globalPluginDirs.

$ git log -SglobalPluginDirs --oneline
35a058fb3 main: configure credentials from the CLI config file
c0b176109 prevent log output during init
8364383c3 Push plugin discovery down into command package
Функция globalPluginDirs упоминается в перечисленных выше коммитах.
Изменена в этом коммите:
$ git log -L :globalPluginDirs:plugins.go
commit 78b12205587fe839f10d946ea3fdc06719decb05


7. Кто автор функции synchronizedWriters?

$ git show 5ac311e2a
commit 5ac311e2a91e381e2f52234668b49ba670aa0fe5
Author: Martin Atkins <mart@degeneration.co.uk>
Date:   Wed May 3 16:25:41 2017 -0700

Автор: Martin Atkins <mart@degeneration.co.uk>
