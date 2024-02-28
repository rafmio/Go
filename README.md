# Go
# git@github.com:rafmio/Go.git
# The destiny of this repository is to have access from any host to an unstructured dump of pieces of code in the Go language. There is no structure here. Sometimes it doesn't even make sense

https://www.devdungeon.com/content/go


Если перед любой командой go run поставить GODEBUG=gctrace=1, то Go выводит
аналитические данные о работе сборщика мусора
$ GODEBUG=gctrace=1 go run gColl.go
$ GODEBUG=gctrace=1 go run <filename.go>


html/template
{{.}} {{if}} {{end}} {{range}}
func (t *Template) Funcs(funcMap FuncMap) *Template
variables
define template

http(html?).FormValue()


----------------------------------------------------
INTERFACES

Для работы интерфейса должны быть (без применения указателей):
1.Объявлен тип (type VariableType string/struct/[]float64)
2.Объявлена метод с соответствующим типом и параметрами и возвращ.значениями:
    func(v VariableType) methodName()
3.Объявлен интерфейс с включением методов:
    type InterfaceName interface {
      methodName()
      methodName2(float64) []float64
    }
4.Объявлена переменная типа, который поддерживает интерфейс
  newVar := VariableType("text if underlying type is string") или
  newVar2 := VariableType { field1 : value1, field2 : "value2", field3 : 11.33 }

5.Объявить переменную с типом=названием интерфейс и присвоить ей значение
  ранее объявленной переменной с типом:
  var varForInterface InterfaceName = newVar

6.Через точку обращаться к последней переменной за методами:
  varForInterface.methodName()
  varForInterface.methodName2(3.14)

РАЗОБРАТСЬЯ С УКАЗАТЕЛЯМИ В ИНТЕРФЕЙСАХ
???Все вышеуказанные действия копируют данные в методы. Для того, чтобы передавать
по ссылке нужно
  a) varForInterface присвоить адрес newVar:
    var varForInterface InterfaceName = &newVar
  b) в получателях методов расставить * перед типом:
    func(v *VariableType) methodName


--------------------------------------------------------

//----------------------------------------
Рандомные инты (случайные значения):
1.Подключить  "math/rand" и "time"
2.rand.Seed(time.Now().UnixNano()) - "посеять зерно, от чего будет отталкиваться"
3.rand.Int()
4.инт в промежутке : rand.Intn()

rand.Seed()
