
from PIL import Image
import os
import PIL
import glob

base_width = 712
dir = './assets/'

def resize_and_save(image_file):
    full_image_path = dir + image_file
    image = Image.open(full_image_path)
    width, height = image.size
    p = base_width / float(width)
    height_resize = int((float(height) * float(p)))
    image = image.resize((base_width, height_resize), PIL.Image.NEAREST)
    image.save(full_image_path, optimize=True, quality=40)
    
image_files = [file for file in os.listdir(dir) if file.endswith(('png'))]
for image_file in image_files:
    resize_and_save(image_file)