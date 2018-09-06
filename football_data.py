import bs4 as bs
from urllib import request

url = 'http://football-data.co.uk/data.php'

sauce = request.urlopen(url).read()

soup = bs.BeautifulSoup(sauce, 'lxml')

tables = soup.find_all('table')

def process_leagues(table):
    rows = table.find_all('tr')
    urls = [td.find(class_='menus').find('a') for td in rows]
    urls = [url.get('href') for url in urls[1:]]
    return urls

leagues = process_leagues(tables[16])
leagues.append(process_leagues(tables[18]))

print(leagues)

# Open each link and download the csv
