from django.db import models


class Migration(models.Model):
    cluster_name = models.CharField(max_length=256, null=False)
    database = models.CharField(max_length=256, null=False)
    table = models.CharField(max_length=256, null=False)
    alter_statement = models.TextField(null=False)

    status = models.IntegerField(default=0)
    copy_percentage = models.FloatField(default=0)

    started_at = models.DateTimeField(null=True)
    completed_at = models.DateTimeField(null=True)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)
