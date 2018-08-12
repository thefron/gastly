from rest_framework import serializers
from migrations.models import Migration


class MigrationSerializer(serializers.ModelSerializer):
    class Meta:
        model = Migration
        fields = (
            'id', 'cluster_name', 'database', 'table', 'alter_statement',
            'status', 'copy_percentage',
            'started_at', 'completed_at', 'created_at', 'updated_at',
        )
