echo=/bin/echo

encryptport=$((RANDOM % 60000 + 1025))
./encrypt_server $encryptport &
sleep 1

${echo}
${echo} "------------------------------"
${echo} "Testing keymaker 10 > key10"
./keymaker 10 > key10
${echo} "key10 must exist."
if [ -s key10 ]
then
  tput bold
  tput setaf 2
  echo "key10 exists!"
  tput sgr0
else
  tput bold
  tput setaf 1
  echo "key10 does not exist."
  tput sgr0
fi

${echo}
${echo} "------------------------------"
${echo} "Testing the number of characters in key10.  It should be 11."
count=$(wc -m< key10)
echo $count key10
if [ $count -eq 11 ]
then
  tput bold
  tput setaf 2
  echo "key10 looks good."
  tput sgr0
else
  tput bold
  tput setaf 1
  echo "key10 does not look good."
  tput sgr0
fi

${echo}
${echo} "------------------------------"
${echo} "Testing creation and the number of characters in key70000.  It should be 70001."
./keymaker 70000 > key70000
count=$(wc -m < key70000)
echo $count key70000
if [ $count -eq 70001 ]
then
  tput bold
  tput setaf 2
  echo "key70000 looks good."
  tput sgr0
else
  tput bold
  tput setaf key10
  echo "key70000 does not look good."
  tput sgr0
fi

${echo}
${echo} "______________________________"
${echo} "Testing ./encrypt_client plaintext1 key70000 $encryptport"
${echo} "This should return encrypted version of plaintext1."
./encrypt_client plaintext1 key70000 $encryptport > ciphertext1
cat ciphertext1

${echo}
${echo} "______________________________"
${echo} "Testing ./encrypt_client plaintext1 key70000 $encryptport > ciphertext1"
${echo} "ciphertext1 should exist."
if [ -s ciphertext1 ]
then 
	tput bold
	tput setaf 2
  echo "ciphertext1 exists." 
	tput sgr0
else 
	tput bold
	tput setaf 1
    echo "ciphertext1 does not exist." 
	tput sgr0
fi 

${echo}
${echo} "______________________________"
${echo} "Testing wc -m plaintext1"
wc -m plaintext1
${echo} "Testing wc -m ciphertext1"
wc -m ciphertext1
${echo} "ciphertext1 must be same number of characters as plaintext1."
if [ $(wc -m < plaintext1) -eq $(wc -m <ciphertext1) ]
then
	tput bold
	tput setaf 2
	echo "Both texts are of the same length."
	tput sgr0
else
	tput bold
	tput setaf 1
	echo "Incorrect length."
	tput sgr0
fi
