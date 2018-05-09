### api/v0/gen

This directory contains the Optikon API swagger spec.

To generate the Github-flavored markdown documentation from an updated spec:

```
docker run -v $(pwd):/opt swagger2markup/swagger2markup convert -i /opt/swagger.yaml -f /opt/swagger -c /opt/config.properties
```
