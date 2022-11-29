// func (имя_параметра тип_получателя) имя_метода (параметры) (типы_возвращаемых_результатов){
//     тело_метода
// }

// объявление типа
type Special int 
// объявление метода
func (m Special) GetString() string{
      return fmt.Sprintf("Special: %d", m)      
}
// вызов метода
var m Special = 5
var s string = m.GetString()
