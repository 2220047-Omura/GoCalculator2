set -x

TYPE=mpfi
#TYPE=double
#TYPE=count

DIR=datadir-$TYPE
mv -f $DIR $DIR-back
mkdir $DIR

PRG=./cmain-$TYPE

INDATADIR=../mtx2

DATA=10000-02-100
CMD="$PRG $INDATADIR/$DATA.mtx"
echo $CMD > $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat

DATA=10000-02-200
CMD="$PRG $INDATADIR/$DATA.mtx"
echo $CMD > $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat

DATA=10000-02-300
CMD="$PRG $INDATADIR/$DATA.mtx"
echo $CMD > $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat

DATA=10000-02-400
CMD="$PRG $INDATADIR/$DATA.mtx"
echo $CMD > $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat

DATA=10000-02-500
CMD="$PRG $INDATADIR/$DATA.mtx"
echo $CMD > $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat