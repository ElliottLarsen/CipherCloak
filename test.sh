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
