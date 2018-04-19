# -*- coding: utf-8 -*-

import requests
import re
import binascii
import datetime

#パラメータ調整用、気象庁で表示している県コード
kenCd = '315'
nameFmt = u'payaneco-tenki-'

def getWeather(day, group):
    index = day[0]
    return getMark(group[index])

def getMark(w):
    if u'雪' in w:
        return u':snowman:'
    elif u'雨' == w:
        return u':umbrella:'
    elif u'雨' in w:
        return u':umbrella:'
    elif u'晴れ' == w:
        return u':sunny:'
    elif u'晴' in w:
        return u':partly_sunny:'
    elif u'曇' in w:
        return u':cloud:'
    #判別不可
    return u':bicyclist:'

#名前出力
def getName(today, tomorrow, satday, sunday):
    wd = datetime.datetime.now().weekday() #月曜が0, 金,土,日が4,5,6
    td = today[2].encode('utf-8')
    tm = tomorrow[2].encode('utf-8')
    st = satday[2].encode('utf-8')
    sn = sunday[2].encode('utf-8')
    if wd < 4:
        s = '{0}{1}→{2}{3}'.format(td, tm, st, sn)
    elif wd == 4:
        s = '{0}:soon:{1}{2}'.format(td, st, sn)
    elif wd == 5:
        s = '{0}{1}:white_flower:'.format(td, tm)
    elif wd == 6:
        s = '{0}{1}:zzz:{2}'.format(td, tm, st)
    return nameFmt.replace('-tenki-', s)

#気象庁の当日天気予報を取得
src = requests.get('http://www.jma.go.jp/jp/yoho/{0}.html'.format(kenCd)).text
th = re.findall(u'<th class="weather">.*?今[日夜](\d+)日.+? title="(.+?)"', src, flags=(re.MULTILINE | re.DOTALL))
#今日のインデックスと日付([index, 日付, 天気])
day = th[0][0]
today = [0, day, getMark(th[0][1])]

#気象庁の週間天気予報を取得
src = requests.get('http://www.jma.go.jp/jp/week/{0}.html'.format(kenCd)).text
tbl = re.findall('<table id="infotablefont".*?>(.+)</table>', src, flags=(re.MULTILINE | re.DOTALL))
div = tbl[0]

#スクレイピング
#日付取得
days = re.findall('<th class="(\w+)">(\d+)', div)
#明日のインデックスと日付([index, 日付, 天気])
if days[0][1] == day:
    tomorrow = [1, days[1][1], '']
else:
    #午後の週間予報は今日の予報を省略する
    tomorrow = [1, days[0][1], '']

#週末のインデックスと日付
for i in range(7):
    if days[i][0] == 'satday':
        satday = [i, days[i][1], '']
    elif days[i][0] == 'sunday':
        sunday = [i, days[i][1], '']

#天気取得
ws = re.findall('<td class="for".+? title="(.+?)"', div)
today[2] = getWeather(today, ws)
tomorrow[2] = getWeather(tomorrow, ws)
satday[2] = getWeather(satday, ws)
sunday[2] = getWeather(sunday, ws)

#名前出力
name = getName(today, tomorrow, satday, sunday)
#毎日5時・11時・17時に更新されるので、30分後に取得すること
print(name)
