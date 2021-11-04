#!/bin/bash
echo "Let's create two customers, James and Bob."
echo "We'll initialize James with two certificates, and Bob with none."
curl -X POST https://localhost:8080/customers -d @test/createJames.json
curl -X POST https://localhost:8080/customers -d @test/createBob.json
echo "Now let's list out their certificates."
curl -X GET https://localhost:8080/customers/1/certificates
curl -X GET https://localhost:8080/customers/2/certificates
echo "We can see how Bob is lacking any certificates. Let's delete Bob, thank you for your help, Bob."
curl -X DELETE htttps://localhost:8080/customers/2
echo "Let's add a new certificate to James..."
curl -X POST https://localhost:8080/customers/1/certificates -d @test/createCert.json
echo "...and see James' certificates again."
curl -X GET https://localhost:8080/customers/1/certificates
echo "Okay. Deactivate James' certificate with ID 7."
curl -X PUT https://localhost:8080/customers/1/certificates/7 -d @test/deactivate.json
echo "View James' certificates (reminder, this only shows active certificates!)"
curl -X GET https://localhost:8080/customers/1/certificates
echo "Okay. Activate James' certificate with ID 7."
curl -X PUT https://localhost:8080/customers/1/certificates/7 -d @test/activate.json
echo "View James' certificates (reminder, this only shows active certificates!)"
curl -X GET https://localhost:8080/customers/1/certificates