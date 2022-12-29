import urllib, os

testfile = urllib.URLopener()

for volumeID in range(1,18):
    if not os.path.exists(str(volumeID)):
        os.makedirs(str(volumeID))
    for problemID in range(100*volumeID, 100*volumeID + 100):
        urlString = "https://onlinejudge.org/external/" + str(volumeID) +"/" + str(problemID) + ".pdf"
        nameString = str(volumeID) + "/" + str(problemID) + ".pdf"
        print str(urlString)
        try:
            testfile.retrieve(urlString, nameString)
        except:
            print str(problemID) + " not found."
    
