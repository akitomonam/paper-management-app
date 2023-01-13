from flask import Flask, request, jsonify
import json
import sys
import os
# import xml.etree.ElementTree as ET
from bs4 import BeautifulSoup
sys.path.append('../grobid_client_python')
from grobid_client.grobid_client_one_file import GrobidClient as GCOne


app = Flask(__name__)

@app.route("/example", methods=["POST"])
def example():
    print(request.json)
    data = request.json
    data["echo"] = "echo"
    data["key"] += ":fuga:hoge"
    text = grobid_client(data["file_path"])
    data["text"] = text
    return json.dumps(data, ensure_ascii=False) # 日本語に対応

def grobid_client(file_path):
    client_one = GCOne(config_path="../config.json")
    client_one.process("processHeaderDocument", file_path,
                       output="../out", tei_coordinates=True, force=True)
    print("filepash:", os.path.splitext(os.path.basename(file_path))[0])
    text = extract_information_from_xml_file("../out/" + os.path.splitext(os.path.basename(file_path))[0] + ".tei.xml")
    return text

def extract_information_from_xml_file(file_path):
    with open(file_path) as f:
        soup = BeautifulSoup(f, "lxml-xml")

    title = soup.find("title")
    # title = soup.find("title", {"level": "a", "type": "main"})
    print(title.text)
    title_text = title.text

    print("title_text:", title_text)
    return title_text

if __name__ == '__main__':
    app.run(host='192.168.192.2', port=12340)
