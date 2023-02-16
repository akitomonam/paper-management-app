from flask import Flask, request
from flask_cors import CORS
import mysql.connector
import json
import sys
import os
import json
import re
from bs4 import BeautifulSoup
sys.path.append('../grobid_client_python')
from grobid_client.grobid_client_one_file import GrobidClient as GCOne
from pdfTranslater import PdfTranslater
import socket

app = Flask(__name__)
CORS(app)

# mysqlに接続する関数
def connect_mysql():
    # jsonファイル読み込み
    with open('../config.json') as f:
        config = json.load(f)

    # mysqlに接続
    connect = mysql.connector.connect(host=config["mysql"]["host"],
                                                  user=config["mysql"]["user"],
                                                  password=config["mysql"]["password"],
                                                  database=config["mysql"]["dbname"])
    return connect

connection = connect_mysql()

@app.route("/api/python/get_paper_meta_data", methods=["POST"])
def get_paper_meta_data():
    print(request.json)
    data = request.json
    metadata = grobid_client(data["file_path"])
    print("metaData:", metadata)
    return json.dumps(metadata, ensure_ascii=False) # 日本語に対応

@app.route("/api/python/pdf_translate", methods=["POST"])
def pdf_translate():
    print(request.json)
    data = request.json
    sessionToken = data["sessionToken"]
    paper_id = data["paper_id"]
    file_path, file_name = pdf_translater.make_translate_pdf(data["file_path"])
    # dbからsessionTokenに紐づくuser_idを取得
    cursor = connection.cursor()
    cursor.execute("SELECT user_id FROM sessions WHERE session_token = %s", (sessionToken,))
    user_id = cursor.fetchone()[0]
    # dbに登録
    cursor.execute("INSERT INTO support_files (file_path, file_name, user_id, paper_id) VALUES (%s, %s, %s, %s)", (file_path, file_name, user_id, paper_id))
    connection.commit()
    cursor.close()
    return json.dumps({"file_path": file_path, "file_name": file_name}, ensure_ascii=False)
    # return "ok"

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
    # info_dict["abstract"] = soup.find("abstract").find("p").text
    abstract = soup.find("abstract").find("p")
    if abstract:
        info_dict["abstract"] = abstract.text
    else:
        info_dict["abstract"] = ""
    year = re.search(r'\d{4}$', soup.find("date").text)
    if year:
        info_dict["year"] = int(year.group())
    else:
        info_dict["year"] = 1
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
    pdf_translater = PdfTranslater()
    app.run(host=socket.gethostbyname(socket.gethostname()), port=12340)
