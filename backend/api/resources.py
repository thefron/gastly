from tastypie.resources import ModelResource
from tastypie.authorization import Authorization
from api.models import MetaRequest


class MetaRequestResource(ModelResource):
    class Meta:
        queryset = MetaRequest.objects.all()
        resource_name = 'meta_request'
        authorization = Authorization()
