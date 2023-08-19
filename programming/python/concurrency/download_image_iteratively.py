import urllib.request


# API_ADDRESS = 'http://localhost:8100/image.png'
API_ADDRESS = 'https://www.google.com/'

def download_image(url, file):
    print('Downloading image from ', url)
    return urllib.request.urlretrieve(url, file)

if __name__ == '__main__':
    for i in range(8):
        image_name = f'./temp/image_{i}.jpg'
        download_image(API_ADDRESS, image_name)

