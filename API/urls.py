from django.urls import path
from . import views

urlpatterns = [
    path('verify/', views.FRView.as_view(), name='verify'),
]
