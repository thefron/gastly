from django.urls import path
from .views import ListMigrationsView

urlpatterns = [
    path('migrations/', ListMigrationsView.as_view(), name='migrations-all'),
]
