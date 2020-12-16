if [[ "$1" = "NO-CACHE" ]]
then
   docker build --no-cache --tag atlas-pos:latest .
else
   docker build --tag atlas-pos:latest .
fi
