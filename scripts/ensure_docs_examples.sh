#!/bin/bash

exitcode=0

resources="$(grep "resource.*()," netbox/provider.go | sed 's/.*"\(.*\)":.*/\1/')"

for resource in ${resources}; do
    ls -1 examples/resources/"$resource"/*.tf >/dev/null 2>&1
    # ls has exitcode 2 if no files are found
    if [ "$?" = "2" ]; then
        echo "Resource $resource has no example"
        exitcode=1
    fi
done

# Code for data source examples. Currently not used.
#datasources="$(grep "dataSource.*()," netbox/provider.go | sed 's/.*"\(.*\)":.*/\1/')"
#
#for datasource in ${datasources}; do
#    ls -1 examples/data-sources/"$datasource"/*.tf >/dev/null 2>&1
#    if [ "$?" = "2" ]; then
#        echo "Data source $datasource has no example"
#        exitcode=1
#    fi
#done

exit $exitcode
