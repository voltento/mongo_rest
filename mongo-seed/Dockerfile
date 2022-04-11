FROM mongo

COPY init.json /init.json
CMD mongoimport --host mongodb --db wave --collection numbers --type json --file /init.json --jsonArray