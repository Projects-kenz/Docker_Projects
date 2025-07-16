from django.shortcuts import render
from django.http import JsonResponse

def hello(request):
    return JsonResponse({'message': 'Hello from Django + Gunicorn'})

def site1(request):
    return render(request, 'app1/index.html')

