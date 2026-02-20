set -x

TYPE=mpfi
#TYPE=double
#TYPE=count

DIR=datadir-$TYPE
mv -f $DIR $DIR-back
mkdir $DIR

PRG=./cmain-$TYPE

INDATADIR=../mtx

DATA=bcsstk25
CMD="$PRG $INDATADIR/$DATA.mtx"
echo $CMD > $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat