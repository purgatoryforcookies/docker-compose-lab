# Inits a file to persist fail attempts
# If fail, adds a line to the file
# if linecount > 10, clean the file and kill itself
# on succesfull healthcheck, clears the linecount and keeps it clean

# idea is to debounce the fail attempst and not kill 
# on first fails. Healthcheck interval acts as a sleep here. 


MEM=./mem.txt

[ ! -f $MEM ] && touch $MEM

if [[ ! `curl -f $1` ]]
then
  COUNT=`wc -l < $MEM`

  [ $COUNT -gt 10 ] && (truncate -s 0 $MEM) && kill 1

  echo $COUNT >> $MEM
  exit 1
else
  truncate -s 0 $MEM
  exit 0
fi