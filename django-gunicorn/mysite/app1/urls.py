from django.urls import path
from .views import hello, site1

urlpatterns = [
    path('', hello, name='hello'),
    path('site1/', site1, name='site1'),
]
