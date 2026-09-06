package orm

import (
	"github.com/beego/beego/v2/client/orm/internal/models"
)

const (
	TypeBooleanField              = models.TypeBooleanField
	TypeVarCharField              = models.TypeVarCharField
	TypeCharField                 = models.TypeCharField
	TypeTextField                 = models.TypeTextField
	TypeTimeField                 = models.TypeTimeField
	TypeDateField                 = models.TypeDateField
	TypeDateTimeField             = models.TypeDateTimeField
	TypeBitField                  = models.TypeBitField
	TypeSmallIntegerField         = models.TypeSmallIntegerField
	TypeIntegerField              = models.TypeIntegerField
	TypeBigIntegerField           = models.TypeBigIntegerField
	TypePositiveBitField          = models.TypePositiveBitField
	TypePositiveSmallIntegerField = models.TypePositiveSmallIntegerField
	TypePositiveIntegerField      = models.TypePositiveIntegerField
	TypePositiveBigIntegerField   = models.TypePositiveBigIntegerField
	TypeFloatField                = models.TypeFloatField
	TypeDecimalField              = models.TypeDecimalField
	TypeJSONField                 = models.TypeJSONField
	TypeJsonbField                = models.TypeJsonbField
	RelForeignKey                 = models.RelForeignKey
	RelOneToOne                   = models.RelOneToOne
	RelManyToMany                 = models.RelManyToMany
	RelReverseOne                 = models.RelReverseOne
	RelReverseMany                = models.RelReverseMany
)

const (
	IsIntegerField         = models.IsIntegerField
	IsPositiveIntegerField = models.IsPositiveIntegerField
	IsRelField             = models.IsRelField
	IsFieldType            = models.IsFieldType
)

type BooleanField = models.BooleanField

var _ Fielder = new(BooleanField)

type CharField = models.CharField

var _ Fielder = new(CharField)

type TimeField = models.TimeField

var _ Fielder = new(TimeField)

type DateField = models.DateField

var _ Fielder = new(DateField)

type DateTimeField = models.DateTimeField

var _ models.Fielder = new(DateTimeField)

type FloatField = models.FloatField

var _ Fielder = new(FloatField)

type SmallIntegerField = models.SmallIntegerField

var _ Fielder = new(SmallIntegerField)

type IntegerField = models.IntegerField

var _ Fielder = new(IntegerField)

type BigIntegerField = models.BigIntegerField

var _ Fielder = new(BigIntegerField)

type PositiveSmallIntegerField = models.PositiveSmallIntegerField

var _ Fielder = new(PositiveSmallIntegerField)

type PositiveIntegerField = models.PositiveIntegerField

var _ Fielder = new(PositiveIntegerField)

type PositiveBigIntegerField = models.PositiveBigIntegerField

var _ Fielder = new(PositiveBigIntegerField)

type TextField = models.TextField

var _ Fielder = new(TextField)

type JSONField = models.JSONField

var _ models.Fielder = new(JSONField)

type JsonbField = models.JsonbField

var _ models.Fielder = new(JsonbField)
