#!/bin/
# -*- coding:utf-8 -*-
# requires python2.6+
 
import os
import sys
#import mcpack
#import socket
#import common_pb2
#import fwd_pb2
#import pbjson
import json
import base64
import time
import urllib2
import urllib
import struct
#import xmltodict
from xml.dom.minidom import Document
import urlparse
#import redis
import sys
 
reload(sys)
sys.setdefaultencoding('utf-8')
 
fp = open('./error_tag','w')
out = open('./out_query.txt','w')
 
def func_get_page_data(link, query):
    try:
        agents = {'User-Agent' : 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36'}
        req = urllib2.Request(link, headers = agents)
        content = urllib2.urlopen(req).read()
        # print content
        return content
    except Exception, e:
        print e
    return ""
 
def unicode_convert(input):
    if isinstance(input, dict):
        return {unicode_convert(key): unicode_convert(value) for key, value in input.iteritems()}
    elif isinstance(input, list):
        return [unicode_convert(element) for element in input]
    elif isinstance(input, unicode):
        return input.encode('utf-8')
    else:
        return input
 
def get_first100_pic_title_len(line):
    datas = line.split('\t')
    word_tru = str(datas[0])
    word = word_tru.replace(' ', '%20')
    try:
        url = "https://opendata.baidu.com/api.php?tn=wisexmlnew&dsp=iphone&alr=1&resource_id=5391&query=" + word
        content = func_get_page_data(url, word)
        dict_json = unicode_convert(json.loads(content))
        
        bookname = dict_json['Result'][0]['DisplayData']['resultData']['tplData']['content']['bookinfo'][0]['bookname']
        penname = dict_json['Result'][0]['DisplayData']['resultData']['tplData']['content']['bookinfo'][0]['penname']
        category_raw = dict_json['Result'][0]['DisplayData']['resultData']['tplData']['content']['bookinfo'][0]['category_raw']
        bd_tag = dict_json['Result'][0]['DisplayData']['resultData']['tplData']['content']['bookinfo'][0]['bd_tag']
        arr1 = category_raw.split("_")
        arr2 = bd_tag.split(",")
        category = arr1[1] + ' ' + arr2[0]
        if len(arr2) > 1: 
            category += ' ' + arr2[1]

        # print bookname+'\t'+penname+'\t'+category
        out.write(bookname+'\t'+penname+'\t'+category+'\n')
    except Exception, e:
        fp.write(line+'\n')
 
    return 0
 
if __name__ == '__main__':
    fp_col=open('query.txt','r')
    for line in fp_col:
        line=line.strip()
        # print line
        get_first100_pic_title_len(line)
        
        #time.sleep(1)
    fp.close()