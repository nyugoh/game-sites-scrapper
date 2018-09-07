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

leagues = process_leagues(tables[16])
leagues.append(process_leagues(tables[18]))

print(leagues)
# Open each link and download the csv
def make_file(file):
    base = '/media/joe/Jarvis/Ripos/Go/src/github.com/nyugoh/game-sites-scrapper/data'
    segments = file.split('/')
    print(segments)
    if os.getcwd() != base:
        os.chdir(base)
    if os.path.exists(base+'/'+file):
        return
    if os.path.exists(base+'/'+segments[0]):
        pass
    else:
        os.mkdir(base+'/'+segments[0])

    if os.path.exists(base+'/'+segments[0]+'/'+segments[1]):
        pass
    else:
        os.mkdir(base+'/'+segments[0]+'/'+segments[1])

    # return to base
    os.chdir('/media/joe/Jarvis/Ripos/Go/src/github.com/nyugoh/game-sites-scrapper')


def download_csv():
    print('Getting csv')
    for lg in leagues:
        if lg == 'http://www.football-data.co.uk/matches.php':
            continue
        parsed_page = parse_page(lg)
        links = [link.get('href') for link in parsed_page.find_all('a')]
        csv_links = []
        for link in links:
            if link.find('.csv') > -1:
                csv_links.append(link)
                make_file(link)
                loc = os.getcwd()+'/data/'+link
                request.urlretrieve('http://football-data.co.uk/'+link, loc)
        print(csv_links)

download_csv()