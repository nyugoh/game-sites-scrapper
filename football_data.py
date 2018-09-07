import bs4 as bs
from urllib import request
import os

url = 'http://football-data.co.uk/data.php'

# Return a bs object when given a url
def parse_page(url):
    sauce = request.urlopen(url).read()
    parsed = bs.BeautifulSoup(sauce, 'lxml')
    return parsed

soup = parse_page(url)
tables = soup.find_all('table')

# Extract all the league urls
def process_leagues(table):
    print('Parsing league page')
    rows = table.find_all('tr')
    urls = [td.find(class_='menus').find('a') for td in rows]
    urls = [url.get('href') for url in urls[1:]]
    return urls

main_leagues = process_leagues(tables[17])
extra_leagues = process_leagues(tables[18])

# Open each link and download the csv
def make_file(file):
    base = '/media/joe/Jarvis/Ripos/Go/src/github.com/nyugoh/game-sites-scrapper'
    segments = file.split('/')
    print("Making file: ", end=" ")
    print(segments)
    if os.getcwd() != base:
        os.chdir(base)
    if os.path.exists(base+'/'+file):
        return
    if os.path.exists(base+'/'+segments[0]):
        pass
    else:
        os.mkdir(base+'/'+segments[0])

    if segments[1].find('.csv') == -1:
        if os.path.exists(base + '/' + segments[0] + '/' + segments[1]):
            pass
        else:
            os.mkdir(base + '/' + segments[0] + '/' + segments[1])

    # return to base
    os.chdir('/media/joe/Jarvis/Ripos/Go/src/github.com/nyugoh/game-sites-scrapper')


def download_main_leagues():
    print('Getting csv')
    for lg in main_leagues:
        if lg == 'http://www.football-data.co.uk/matches.php':
            continue
        parsed_page = parse_page(lg)
        links = [link.get('href') for link in parsed_page.find_all('a')]
        csv_links = []
        for link in links:
            if link.find('.csv') > -1:
                loc = os.getcwd() + '/' + link
                if os.path.exists(loc):
                    print(loc+" Already exits")
                    continue
                csv_links.append(link)
                make_file(link)
                print("Retriving file :"+ link)
                request.urlretrieve('http://football-data.co.uk/'+link, loc)
        print(csv_links)
        print('Compeleted getting main league files')

def download_extra_leagues():
    print('Getting extra leagues csv')
    for lg in extra_leagues:
        if lg == 'http://www.football-data.co.uk/matches_new_leagues.php':
            continue
        parsed_page = parse_page(lg)
        links = [link.get('href') for link in parsed_page.find_all('a')]
        csv_links = []
        for link in links:
            if link.find('.csv') > -1:
                loc = os.getcwd() + '/' + link
                if os.path.exists(loc):
                    print(loc+" Already exits")
                    continue
                csv_links.append(link)
                make_file(link)
                print("Retriving file :"+ link)
                request.urlretrieve('http://football-data.co.uk/'+link, loc)
        print(csv_links)
        print('Completed getting extra leagues files.')


download_main_leagues()
download_extra_leagues()