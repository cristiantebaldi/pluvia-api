cd $1
ls
cp config.example.json config.json
sed -i -e 's/$BCRYPT_HASH/'@str@1587'/g' config.json
sed -i -e 's/$DB_USER/'postgres'/g' config.json
sed -i -e 's/$DB_PASS/'postgres'/g' config.json
sed -i -e 's/$DB_HOST/'postgres'/g' config.json
sed -i -e 's/$DB_PORT/'5432'/g' config.json
sed -i -e 's/$DB_NAME/'pluvia'/g' config.json
sed -i -e 's/$SERVER_PORT/'3000'/g' config.json
sed -i -e 's/$SERVER_VERSION/'2025.8.4.0'/g' config.json
