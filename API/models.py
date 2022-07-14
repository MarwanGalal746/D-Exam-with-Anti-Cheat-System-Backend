from django.db import models


# Create your models here.

class Data(models.Model):
    profilePicture = models.ImageField(upload_to='face_images')
    toComparePicture = models.ImageField(upload_to='face_images')

    def __str__(self):
        return self.title
