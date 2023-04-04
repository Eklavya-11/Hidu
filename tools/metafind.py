import argparse
import os
from urllib.parse import urlparse

import requests
from bs4 import BeautifulSoup
from googlesearch import search
from termcolor import colored

class MetaFinder:
def init(self, domain: str, output: str) -> None:
self.domain = domain
self.output = output
self.banner()

    if output:
        self.download_files()
    else:
        self.search_files()

def banner(self) -> None:
    print(colored("""
    ╔══════════════════════════════════╗
    ║  Tool Created by Eklavya11       ║
    ║  Version 1.1                     ║
    ╚══════════════════════════════════╝
""", "cyan"))

def search_files(self) -> None:
    print(colored(f"[*] Starting the search for {self.domain}", "blue"))
    urls = self.get_urls()

    for url in urls:
        print(colored(f"[+] {url}", "green"))

def get_urls(self) -> list[str]:
    urls = []
    query = f"site:*.{self.domain} ext:pdf | ext:doc | ext:xls | ext:ppt | ext:odp | ext:ods | ext:docx | ext:xlsx | ext:pptx"
    for url in search(query, pause=3, lang="en", start=0, stop=40):
        urls.append(url)
    return urls

def download_files(self) -> None:
    print(colored(f"[*] Starting the search and download for {self.domain}", "blue"))
    urls = self.get_urls()
    os.makedirs(self.output, exist_ok=True)

    for url in urls:
        print(colored(f"[+] {url}", "green"))

        try:
            r = requests.get(url, allow_redirects=True, timeout=3)
            if r.status_code == 200:
                soup = BeautifulSoup(r.content, 'html.parser')
                title = soup.title.string.strip()
                file_name = f"{title}.pdf"
                file_name = file_name.replace("/", "-")
                parsed_url = urlparse(url)
                file_path = os.path.join(self.output, file_name)
                open(file_path, 'wb').write(r.content)
                print(colored(f"[+] Downloaded file: {file_path}", "green"))

        except requests.exceptions.RequestException as e:
            print(colored(f"[!] Error downloading file {url} - {e}", "red"))

if name == "main":
parser = argparse.ArgumentParser()
parser.add_argument('-d', '--domain', help='target domain', required=True)
parser.add_argument('-o', '--output', help='output folder', required=False)
args = parser.parse_args()
meta_finder = MetaFinder(args.domain, args.output)
