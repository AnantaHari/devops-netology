#!/usr/bin/env python3
import requests
import base64
import json
import os

username = 'anantahari'
repo = 'leftjoin'
token = "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"
new_branch_name = 'automatic'

r = requests.post('https://api.github.com/user/repos',
                  data=json.dumps({'name':repo}),
                  auth=(username, token),
                  headers={"Content-Type": "application/json"})

content = "Hello, world!"

b_content = content.encode('utf-8')
base64_content = base64.b64encode(b_content)
base64_content_str = base64_content.decode('utf-8')

#Теперь создадим коммит в формате словаря: он состоит из пути (пока его оставим пустым), сообщения коммита и содержания:

f = {'path':'',
     'message': 'Automatic update',
     'content': base64_content_str}

#Отправим файл в репозиторий put-запросом, передав путь до README.md прямо в ссылке запроса:

f_resp = requests.put(f'https://api.github.com/repos/{username}/{repo}/contents/README.md',
                auth=(username, token),
                headers={ "Content-Type": "application/json" },
                data=json.dumps(f))


