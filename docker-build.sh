if [[ "$1" = "NO-CACHE" ]]
then
   docker build -f Dockerfile.dev --no-cache --tag atlas-pos:latest .
else
   docker build -f Dockerfile.dev --tag atlas-pos:latest .
fi
