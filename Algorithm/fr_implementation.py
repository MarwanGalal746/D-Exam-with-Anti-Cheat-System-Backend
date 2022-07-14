import face_recognition


def run(img1, img2):
    image1 = face_recognition.load_image_file(img1)
    image2 = face_recognition.load_image_file(img2)

    image1_encoding = face_recognition.face_encodings(image1)[0]
    image2_encoding = face_recognition.face_encodings(image2)[0]

    results = face_recognition.compare_faces([image1_encoding], image2_encoding)
    return results
