import json

data = []
with open('t.txt', 'rb') as f:
    for line in f.readlines():
        line = line.decode('utf-8').strip()
        items = [item[1:-1] if item.startswith('"') else item for item in line.split('+')]
        data.append({
            'name': items[0],
            'params': items[1],
            'summary': items[2],
            'group': int(items[3]),
            'since': items[4]
        })

with open('../command_help.json', 'wb') as f:
    f.write(json.dumps(data, indent=2).encode('utf-8'))