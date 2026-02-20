set -x

#TYPE=mpfi
TYPE=double
#TYPE=count

DIR=datadir-$TYPE
mv -f $DIR $DIR-back
mkdir $DIR

PRG=./cgomain-$TYPE

INDATADIR=../mtx2

DATA=300-02-100
CMD="$PRG $INDATADIR/$DATA.mtx"
echo $CMD > $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat

DATA=400-02-100
CMD="$PRG $INDATADIR/$DATA.mtx"
echo $CMD > $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat

DATA=500-02-100
CMD="$PRG $INDATADIR/$DATA.mtx"
echo $CMD > $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat

DATA=600-02-100
CMD="$PRG $INDATADIR/$DATA.mtx"
echo $CMD > $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat

DATA=700-02-100
CMD="$PRG $INDATADIR/$DATA.mtx"
echo $CMD > $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat

DATA=10000-02-100
CMD="$PRG $INDATADIR/$DATA.mtx"
echo $CMD > $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat
$CMD >> $DIR/$DATA-$TYPE.dat

 