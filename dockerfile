# Need to decouple build from runtimea
FROM hashicorp/terraform:full
COPY . /go/src/github.com/flynnhandley/terraform-provider-azdevops/
RUN cd /go/src/github.com/flynnhandley/terraform-provider-azdevops \
&& apk update && apk add --no-cache --virtual .build-deps gcc docker \
&& go get ./... \
&& GOOS=linux go build -o ./terraform-provider-azdevops \
&& mkdir --parents ~/.terraform.d/plugins \
&& mv ./terraform-provider-azdevops ~/.terraform.d/plugins/
WORKDIR /go/src/github.com/flynnhandley/terraform-provider-azdevops/examples
ENTRYPOINT ["terraform"]