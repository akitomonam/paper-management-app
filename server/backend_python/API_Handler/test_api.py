from flask import Flask, request, jsonify
import json

app = Flask(__name__)

@app.route("/example", methods=["POST"])
def example():
    print(request.json)
    data = request.json
    data["echo"] = "echo"
    data["key"] += ":fuga:hoge"
    return json.dumps(data, ensure_ascii=False) # 日本語に対応

if __name__ == '__main__':
    app.run(host='192.168.192.2', port=12340)

# curl -X POST -H 'Content-Type: application/json' http://172.21.65.107:12340/example -d '{"key": "value"}'
