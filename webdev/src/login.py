import requests
import urllib
import getpass

def login(username, password):
    r = requests.get('http://www.google.com') # GET // HL

    if r.url.find('google.com') == -1 :
        magic = urllib.splitquery(r.url)[1]
        values = {
                'username':username,
                'password':password,
                'magic':magic,
                }
        r2 = requests.post(r.url, data=values) # POST // HL
        print 'Logout: ', r2.content[4816:4870]
    else:
        print 'Already connected'

if __name__ == '__main__':
    password = getpass.getpass("Enter Password:")
    login('falak16018', password)
