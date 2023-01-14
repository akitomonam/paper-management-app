from flask import Flask, request, jsonify
from flask_cors import CORS
import json
import sys
import os
from bs4 import BeautifulSoup
sys.path.append('../grobid_client_python')
from grobid_client.grobid_client_one_file import GrobidClient as GCOne
import socket

app = Flask(__name__)
CORS(app)

@app.route("/api/python/get_paper_meta_data", methods=["POST"])
def get_paper_meta_data():
    print(request.json)
    data = request.json
    metadata = grobid_client(data["file_path"])
    print("metaData:", metadata)
    return json.dumps(metadata, ensure_ascii=False) # 日本語に対応

def grobid_client(file_path):
    client_one = GCOne(config_path="../config.json")
    client_one.process("processHeaderDocument", file_path,
                       output="../out", tei_coordinates=True, force=True)
    text = extract_information_from_xml_file("../out/" + os.path.splitext(os.path.basename(file_path))[0] + ".tei.xml")
    return text

def extract_information_from_xml_file(file_path):
    info_dict = {}
    with open(file_path) as f:
        soup = BeautifulSoup(f, "lxml-xml")

    info_dict["title"] = soup.find("title").text
    authors = soup.find_all("author")
    author_list = []
    for author in authors:
        try:
            persName = author.find("persName")
            firstname = persName.find("forename", {"type":"first"})
            surname = persName.find("surname")
        except AttributeError:
            continue
        if firstname:
            author_list.append(firstname.text + " " + surname.text)
        # else:
        #     author_list.append(surname.text)
    author_text = ", ".join(author_list)
    info_dict["author"] = author_text
    return info_dict

if __name__ == '__main__':
    app.run(host=socket.gethostbyname(socket.gethostname()), port=12340)
