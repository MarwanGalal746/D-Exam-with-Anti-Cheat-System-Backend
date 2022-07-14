import cv2
import face_recognition


def run(img1, img2):
    image = face_recognition.load_image_file(img1)
    rgb = cv2.cvtColor(image, cv2.COLOR_BGR2RGB)

    boxes = face_recognition.face_locations(rgb)[0]
    encoding1 = face_recognition.face_encodings(rgb)[0]

    # read 2nd image and store encodings
    image = face_recognition.load_image_file(img2)
    rgb = cv2.cvtColor(image, cv2.COLOR_BGR2RGB)

    boxes = face_recognition.face_locations(rgb)[0]
    encoding2 = face_recognition.face_encodings(rgb)[0]

    # now you can compare two encodings
    # optionally you can pass threshold, by default it is 0.6
    matches = face_recognition.compare_faces([encoding1], encoding2)
    return matches
