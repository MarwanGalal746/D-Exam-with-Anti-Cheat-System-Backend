docker pull eyadzz/dexam_exam_service:latest &
docker pull eyadzz/dexam_user_service:latest &
docker pull eyadzz/dexam_course_service:latest &
docker pull eyadzz/dexam_api_gateway:latest &
docker pull eyadzz/dexam_overseer_service:latest &
docker-compose -f backend.yaml up --build