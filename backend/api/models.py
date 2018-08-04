from django.db import models


class MetaRequest(models.Model):
    alter_statement = models.TextField()
    table_name = models.CharField(max_length=256)
    database_name = models.CharField(max_length=256)
    pr_url = models.CharField(max_length=256)

    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)
