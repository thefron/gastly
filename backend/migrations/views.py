from rest_framework import generics
from migrations.models import Migration
from migrations.serializers import MigrationSerializer


class ListMigrationsView(generics.ListAPIView):
    """
    List migrations.
    """
    queryset = Migration.objects.all()
    serializer_class = MigrationSerializer
