set -e

gowsdl -p gen ../../fixtures/test.wsdl

go build ./gen/

