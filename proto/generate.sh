cd proto
buf generate --template buf.gen.gogo.yaml
buf generate --template buf.gen.pulsar.yaml
cd ..

cp -r github.com/noble-assets/authority/* ./
cp -r api/noble/authority/* api/

rm -rf github.com
rm -rf api/noble
rm -rf noble
