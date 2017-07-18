
Please have a look at

    ./values.yaml

and adjust the values if necessary. For fine grained adjustments to the kubernetes config files, please
review the files in './templates' sub folder.
Once finished, make sure kubectl is configured to the desired cluster and execute:

    helm install --name=echo --set image.tag=<docker-tag> --set stage={dev|prod} .

where stage needs to be either 'dev' or 'prod' depending on where you want to create the app.
For deployment, execute:

    helm upgrade --recreate-pods --set image.tag=<docker-tag> --set stage={dev|prod} echo .

If you want helm to just render the templates, please add '--dry-run --debug' to either one of the install or upgrade command.

In order to delete this app from kubernetes, run:

    helm delete --purge echo
