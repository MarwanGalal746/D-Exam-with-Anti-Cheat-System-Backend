from rest_framework import status
from rest_framework.parsers import MultiPartParser, FormParser
from rest_framework.response import Response
from rest_framework.views import APIView
from API.serializer import DataSerializer
from Algorithm import fr_implementation


class FRView(APIView):
    parser_classes = (MultiPartParser, FormParser)

    def post(self, request, *args, **kwargs):
        serializer = DataSerializer(data=request.data)
        if serializer.is_valid():
            result = fr_implementation.run(serializer.validated_data['profilePicture'],
                                           serializer.validated_data['toComparePicture'])
            return Response({"result": result}, status=status.HTTP_200_OK)
        else:
            return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)
