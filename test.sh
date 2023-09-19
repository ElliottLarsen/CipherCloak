echo=/bin/echo

${echo}
${echo} '------------------------------'
${echo} 'Testing keymaker 10 > key10'
./keymaker 10 > key10
${echo} 'key10 must exist.'
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
${echo} '------------------------------'
${echo} 'Testing the number of characters in key10.  It should be 11.'
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
${echo} '------------------------------'
${echo} 'Testing creation and the number of characters in key70000.  It should be 70001.'
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
