#!/bin/bash
echo "Let's create two customers, James and Bob."
echo "We'll initialize James with two certificates, and Bob with none."
curl -X POST localhost:8080/customers -d @test/createJames.json
curl -X POST localhost:8080/customers -d @test/createBob.json
echo "(2 POST requests happened here)"; echo
echo "Now let's list out their certificates."
curl -X GET localhost:8080/customers/1/certificates; echo
curl -X GET localhost:8080/customers/2/certificates; echo
echo
echo "We can see how Bob is lacking certificates. Let's delete Bob, thank you for your help, Bob."
curl -X DELETE localhost:8080/customers/2
echo "(DELETE request happened here)"; echo
echo "Let's add a new certificate to James..."
curl -X POST localhost:8080/customers/1/certificates -d @test/createCert.json
echo "(POST request happened here)"; echo
echo "...and see James' certificates again."
curl -X GET localhost:8080/customers/1/certificates; echo
echo
echo "Okay. Deactivate James' certificate with ID 7."
curl -X PUT localhost:8080/customers/1/certificates/7 -d @test/deactivate.json
echo "(PUT request happened here)"; echo
echo "View James' certificates (reminder, this only shows active certificates!)"
curl -X GET localhost:8080/customers/1/certificates; echo
echo
echo "Okay. Activate James' certificate with ID 7."
curl -X PUT localhost:8080/customers/1/certificates/7 -d @test/activate.json
echo "(PUT request happened here)"; echo
echo "View James' certificates (reminder, this only shows active certificates!)"
curl -X GET localhost:8080/customers/1/certificates; echo